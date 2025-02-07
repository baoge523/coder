package main

import (
	"bytes"
	"context"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"time"
)

func main() {
	clientSend()
}
// clientSend 在发送请求时，如果使用 bufio包装一下reader会导致请求header(Content-Length)不存在
func clientSend() {

	io := bytes.NewReader([]byte("发起一个请求，这是请求数据, client send"))

	http.Post("http://127.0.0.1:8080/abcd","application/json",io)

}

func clientSend2() {
	client := http.Client{
		Transport: &http.Transport{
            DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
                return net.Dial(network, addr)
            },
        },
	}

	req,_ := http.NewRequest("POST", "http://127.0.0.1:8080/abcd",bytes.NewReader([]byte("发起一个请求，这是请求数据")))
	client.Do(req)

}




func clientUseCase() {

	// 创建 SOCKS5 认证
	var auth = &proxy.Auth{
		User:     "user",
		Password: "password",
	}
	// 创建 SOCKS5 代理拨号器
	dialer, err := proxy.SOCKS5("tcp", "ip:port", auth, proxy.Direct)
	if err != nil {
		panic(err)
	}



	// 创建一个客户端
	client := http.Client{
		Transport: &http.Transport{
			// 创建http连接
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				host, port, err := net.SplitHostPort(addr)
				if err != nil {
					return nil, err
				}
				if host == "www.baidu.com" {
					hostname := net.JoinHostPort("127.0.0.1", port)
					return net.Dial(network, hostname)
				}
				// 走代理
				return dialer.Dial(network, addr)
			},
			// 创建https连接
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return nil, nil
			},
			MaxIdleConns: 200,  // 最大空闲连接，所有的hosts，0表示不限制
			MaxIdleConnsPerHost: 2, // 单个hosts的最大空闲连接
			MaxConnsPerHost: 100, // 单个hosts的最大连接数
			IdleConnTimeout: 90 * time.Second, // 空闲连接的超时时间
			WriteBufferSize: 1024, // 写缓冲区大小 默认 4k
			ReadBufferSize: 1024, // 读缓冲区大小 默认 4k

		},
	}

	client.Do(nil)

}