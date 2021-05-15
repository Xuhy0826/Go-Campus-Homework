// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

package main

import (
	"log"
	"net/http"
	"week03/application"
	h "week03/transport/http"
)

func main(){
	h1 := h.NewServer(h.Address("127.0.0.1:8080"))
	h1.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("h1: hello"))
	})

	h2 := h.NewServer(h.Address("127.0.0.1:8081"))
	h2.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("h2: hello"))
	})

	app := application.New(
		application.Name("week03"),
		application.Version("v1.0.0"),
		application.Server(h1, h2),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

