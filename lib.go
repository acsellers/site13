package main

import (
	"fmt"
	"net/http"
)

// Put library functions/structs here
func NewCss(name string, sources ...string) {
	http.HandleFunc("/assets/"+name+".css", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		for _, content := range sources {
			fmt.Fprintln(w, content)
		}
	})
}

func NewJS(name string, sources ...string) {
	http.HandleFunc("/assets/"+name+".js", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/javascript")
		for _, content := range sources {
			fmt.Fprintln(w, content)
		}
	})
}

func CloseAssets() {
	http.HandleFunc("/assets/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "")
	})
}

func StaticPage(name, content string) {
	http.HandleFunc(name, func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, content)
	})
}

func DynamicPage(name string, f http.HandlerFunc) {
	http.HandleFunc(name, f)
}
