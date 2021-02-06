package main

import (
	data2 "go_test/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for i:=0; i>3;i++ {
			log.Println(data2.Add("https://github.com/EDDYCJY"))
		}
	}()
	http.ListenAndServe(":6060", nil)
}