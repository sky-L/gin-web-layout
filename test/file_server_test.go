package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	port := ":8080" // 服务器端
	dir := "./"     // 文件目录

	// 设置静态文件服务器
	fs := http.FileServer(http.Dir(dir))

	// http.Handle("/", fs)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println(request.URL.Path)

		// writer.Header().Set("Content-Disposition", "attachment")
		fs.ServeHTTP(writer, request)
	})

	// 启动服务器
	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
