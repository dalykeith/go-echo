// A simple web application example.
// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io :: 2017-09-13

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "r.Method:           ",  r.Method           )
	fmt.Fprintln(w, "r.URL:              ",  r.URL              )
	fmt.Fprintln(w, "r.Proto:            ",  r.Proto            )
	fmt.Fprintln(w, "r.Header:           ",  r.Header           )
	fmt.Fprintln(w, "r.ContentLength:    ",  r.ContentLength    )
	fmt.Fprintln(w, "r.TransferEncoding: ",  r.TransferEncoding )
	fmt.Fprintln(w, "r.Close:            ",  r.Close            )
	fmt.Fprintln(w, "r.Host:             ",  r.Host             )
	fmt.Fprintln(w, "r.Form:             ",  r.Form             )
	fmt.Fprintln(w, "r.PostForm:         ",  r.PostForm         )
	fmt.Fprintln(w, "r.RemoteAddr:       ",  r.RemoteAddr       )
	fmt.Fprintln(w, "r.RequestURI:       ",  r.RequestURI       )
	//fmt.Fprintln(w, "r.Body:             ",  w(bytes.Buffer).ReadFrom(r.Body).String())

	fmt.Fprintln(w, "r.URL.Opaque:       ", r.URL.Opaque        )
	fmt.Fprintln(w, "r.URL.Scheme:       ", r.URL.Scheme        )
	fmt.Fprintln(w, "r.URL.Host:         ", r.URL.Host          )
	fmt.Fprintln(w, "r.URL.Path:         ", r.URL.Path          )
	fmt.Fprintln(w, "r.URL.RawPath:      ", r.URL.RawPath       )
	fmt.Fprintln(w, "r.URL.RawQuery:     ", r.URL.RawQuery      )
	fmt.Fprintln(w, "r.URL.Fragment:     ", r.URL.Fragment      )
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}