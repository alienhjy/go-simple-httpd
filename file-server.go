package main

import (
	"os"
	"net/http"
	"log"
	"flag"
	"strings"
)

var (
	serv_http string
	serv_dir StringFlags
	handle_path []string
)

func main() {
	flag.Var(&serv_dir, "dir",
		"Dirs for http file server. (e.g. 'list1:/dir1')")
	flag.StringVar(&serv_http, "http", ":8880",
		"HTTP service address.")
	flag.Parse()

	if len(serv_dir) == 0 {
		serv_dir = append(serv_dir, "/:.")
	}
	for _, dir := range serv_dir{
		handle_path = strings.SplitN(dir, ":", 2)
		if ! strings.HasPrefix(handle_path[0], "/") {
			handle_path[0] = "/" + handle_path[0]
		}
		if ! strings.HasSuffix(handle_path[0], "/") {
			handle_path[0] = handle_path[0] + "/"
		}
		http.Handle(handle_path[0], http.StripPrefix(handle_path[0],
			http.FileServer(http.Dir(handle_path[1]))))
	}
	log.Println("Http server is started at", serv_http)
	log.Println(serv_dir)
	err := http.ListenAndServe(serv_http, nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
		os.Exit(1)
	}

	return
}
