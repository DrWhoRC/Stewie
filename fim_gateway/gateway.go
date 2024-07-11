package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
)

var serviceMap = map[string]string{
	"auth": "http://localhost:8888",
	"user": "http://localhost:8080",
}

func gateway(res http.ResponseWriter, req *http.Request) {

	//匹配请求前缀 /api/user/xxx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrlist := regex.FindStringSubmatch(req.URL.Path)
	fmt.Println(addrlist)
	if len(addrlist) != 2 {
		res.Write([]byte("err"))
	}

	service := addrlist[1]
	addr, ok := serviceMap[service]
	if !ok {
		fmt.Println("service not found", service)
		res.Write([]byte("err"))
		return
	}

	url := fmt.Sprintf("%s/%s", addr, req.URL.String())
	proxyReq, _ := http.NewRequest(req.Method, url, req.Body)
	fmt.Println(req.Header, "\n", req.Body)

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

type Config struct {
	Addr string
}

func main() {

	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)

	http.HandleFunc("/", gateway)
	fmt.Printf("gateway running %s \n", c.Addr)

	http.ListenAndServe(c.Addr, nil)
}
