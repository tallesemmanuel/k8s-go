package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1> Brincando com k8s + Golang webserver - Version 2.0 </h1>"))
	})
	http.ListenAndServe(":8080", nil)
}