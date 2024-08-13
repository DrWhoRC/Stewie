package main

import (
	"bytes"
	"encoding/json"
	"fim/common/etcd"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	Addr      string
	Etcd      string
	WhiteList []string
	Log       logx.LogConf
}

var config Config

//用户直接用login登录的话，会返回jwt，但网关目前的设计是所有服务都先通过auth再进行转发，
//所以需要想一种模式，满足的逻辑为：访问login的时候，不需要auth，直接返回jwt，再走auth
//如果没有login，本来就在login的话，就不需要login，还是auth->service，
//如果没有login，最好是auth->login->auth->service

type Proxy struct {
}

func (Proxy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//匹配请求前缀 /api/user/xxx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrlist := regex.FindStringSubmatch(req.URL.Path)
	fmt.Println(addrlist)
	if len(addrlist) != 2 {
		res.Write([]byte("err"))
	}

	service := addrlist[1]
	fmt.Println(service)
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

	ProxyUrl := fmt.Sprintf("http://%s%s", addr, req.URL.String())
	logx.Infof("%s, %s", remoteAddr[0], ProxyUrl)

	//将格式化后的 URL 字符串解析为一个 url.URL 对象。这个对象包含了 URL 的各个部分（如协议、主机、路径等），便于后续操作。
	remote, _ := url.Parse(fmt.Sprintf("http://%s", addr))     //目标服务器地址
	reverseProxy := httputil.NewSingleHostReverseProxy(remote) //httputil.ReverseProxy 类型的对象。它负责将请求转发到目标服务器，并将目标服务器的响应返回给客户端。
	reverseProxy.ServeHTTP(res, req)                           //这个ServeHTTP是httputil.ReverseProxy的方法，用于处理请求和响应。并不是Proxy的ServeHTTP方法
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
	contentType := req.Header.Get("Content-Type")
	authReq.Header.Set("Content-Type", contentType)

	// 然后复制其他header
	for name, values := range req.Header {
		// 跳过Content-Type，因为我们已经设置过了
		if name != "Content-Type" {
			for _, value := range values {
				authReq.Header.Add(name, value)
			}
		}
	}
	authReq.ContentLength = int64(len(body01))
	authReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		fmt.Println("authres:", err)
		res.Write([]byte("auth服务异常"))
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
	logx.SetUp(config.Log)
	proxy := Proxy{}
	fmt.Printf("gateway running %s \n", config.Addr)

	http.ListenAndServe(config.Addr, proxy)
}
