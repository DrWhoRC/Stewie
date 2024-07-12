package main

import (
	"bytes"
	"fim/common/etcd"
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Addr string
	Etcd string
}

var config Config

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

	url := fmt.Sprintf("%s%s/%s", req.URL.Scheme, addr, req.URL.String())

	//先将请求体读取到body中，此时req.Body已经被读取过一次，所以要重新设置
	body, err := io.ReadAll(req.Body)
	//随后使用NopCloser方法再次设置req.Body，req.body来设置proxybody，body来获取contentlength
	req.Body = io.NopCloser(bytes.NewReader(body))

	proxyReq, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		fmt.Println(err)
		res.Write([]byte("服务异常"))
		return
	}
	//fmt.Println(req.Header, "\n", proxyReq.Body)

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

	remoteAddr := strings.Split(req.RemoteAddr, ":")
	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	response, err := http.DefaultClient.Do(proxyReq)

	if err != nil {
		fmt.Println(err)
		res.Write([]byte("服务异常"))
		return
	}
	io.Copy(res, response.Body)
}

var configFile = flag.String("f", "settings.yaml", "The Config File")

func main() {

	flag.Parse()

	conf.MustLoad(*configFile, &config)

	http.HandleFunc("/", gateway)
	fmt.Printf("gateway running %s \n", config.Addr)

	http.ListenAndServe(config.Addr, nil)
}
