package main

import (
	"fmt"
	har "goTest/errhandling/fileList"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"reflect"
)


type appHandler func(writer http.ResponseWriter, request *http.Request) error
func errWrapper (handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover(); if r != nil {
				log.Printf("panic: %v ", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			// 打印日志在终端
			log.Println(" 报错了", err.Error())
			if userError, ok := err.(userError); ok {
				// 展示给用户的
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				code = http.StatusNotFound
				http.Error(writer, http.StatusText(code), code)
			case os.IsPermission(err):
				code = http.StatusPermanentRedirect //没有权限
			default:
				code = http.StatusInternalServerError
			}
		}
		fmt.Println("数据类型", reflect.TypeOf(request.URL.Query()["id"]))
		fmt.Println("请求的参数", request.URL.Query(), request.URL.Query()["id"])
	}
}

type userError interface {
	error
	Message() string
}


/***
	获取文件列表的服务器
 */
func main() {
	http.HandleFunc("/", errWrapper(har.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
