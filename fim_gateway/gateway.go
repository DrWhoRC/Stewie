package main

import (
	"bytes"
	"encoding/json"
	"fim/common/etcd"
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	Addr      string
	Etcd      string
	WhiteList []string
}

var config Config

//用户直接用login登录的话，会返回jwt，但网关目前的设计是所有服务都先通过auth再进行转发，
//所以需要想一种模式，满足的逻辑为：访问login的时候，不需要auth，直接返回jwt，再走auth
//如果没有login，本来就在login的话，就不需要login，还是auth->service，
//如果没有login，最好是auth->login->auth->service

func gateway(res http.ResponseWriter, req *http.Request) {

	//匹配请求前缀 /api/user/xxx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrlist := regex.FindStringSubmatch(req.URL.Path)
	fmt.Println(addrlist)
	if len(addrlist) != 2 {
		res.Write([]byte("err"))
	}

	service := addrlist[1]
	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		fmt.Println("service not found", service)
		res.Write([]byte("err"))
		return
	}
	fmt.Println("addr:", addr)

	remoteAddr := strings.Split(req.RemoteAddr, ":")

	var counter int = 0
	fmt.Println("config.whitelist:", config.WhiteList)
	for i := 0; i < len(config.WhiteList); i++ {
		if req.URL.String() == config.WhiteList[i] {
			counter++
		}
	}
	if counter == 0 {
		Auth(res, req, remoteAddr)
	}

	url := fmt.Sprintf("http://%s/%s", addr, req.URL.String())

	//先将请求体读取到body中，此时req.Body已经被读取过一次，所以要重新设置
	body, err := io.ReadAll(req.Body)
	//随后使用NopCloser方法再次设置req.Body，req.body来设置proxybody，body来获取contentlength
	req.Body = io.NopCloser(bytes.NewReader(body))

	proxyReq, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		fmt.Println(err)
		res.Write([]byte("认证服务异常2"))
		return
	}

	proxyReq.ContentLength = int64(len(body))
	//罪魁祸首！！！！！！！！！！！！！
	//就是上边那一行代码！！！！！！！！
	//header中的contentlength确实是56，这也就解释了为什么打印的header信息中没有错误
	//但是！！！proxyReq.ContentLength和header中的contentlength不是一个东西，
	//所以要手动设置proxyReq.ContentLength的长度！

	proxyReq.Header.Set("Content-Type", "application/json")

	for name, values := range req.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	response, err := http.DefaultClient.Do(proxyReq)

	if err != nil {
		fmt.Println(err)
		res.Write([]byte("服务异常"))
		return
	}
	io.Copy(res, response.Body)
}

func Auth(res http.ResponseWriter, req *http.Request, remoteAddr []string) {
	// 请求认证服务的地址
	authAddr := etcd.GetServiceAddr((config.Etcd), "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
	fmt.Println("authUrl:", authUrl)
	body01, err := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewReader(body01))

	authReq, _ := http.NewRequest("POST", authUrl, req.Body)
	req.Body = io.NopCloser(bytes.NewReader(body01))
	authReq.Header.Set("Content-Type", "application/json")
	authReq.ContentLength = int64(len(body01))
	authReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	for name, values := range req.Header {
		for _, value := range values {
			authReq.Header.Add(name, value)
		}
	}

	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		fmt.Println("authres:", err)
		res.Write([]byte("服务异常"))
		return
	}
	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	var authResponse Response
	byteData, _ := io.ReadAll(authRes.Body)
	fmt.Println(string(byteData))
	authErr := json.Unmarshal(byteData, &authResponse)
	if authErr != nil {
		fmt.Println(&authResponse, ":", authErr.Error())
	}

	if err != nil {
		logx.Error(authErr)
		res.Write([]byte("认证服务异常1"))
		return
	}
	// authentication failed
	if authResponse.Code == 1 {
		res.Write([]byte("请重新登录"))
	}
}

var configFile = flag.String("f", "settings.yaml", "The Config File")

func main() {

	flag.Parse()

	conf.MustLoad(*configFile, &config)

	http.HandleFunc("/", gateway)
	fmt.Printf("gateway running %s \n", config.Addr)

	http.ListenAndServe(config.Addr, nil)
}
