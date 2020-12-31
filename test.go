package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	fmt.Println(GetIP(r))
	http.Redirect(w, r, "http://www.google.com", 301)
	elapsed := time.Since(start)
	log.Printf("redirect took %s", elapsed)
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
