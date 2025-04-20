package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Hello from /"))
		case "/health":
			w.Write([]byte("Hello from /health"))
		case "/about":
			w.Write([]byte("Hello from /about"))
		case "/test":
			w.Write([]byte("Hello from /test"))
		case "/os":
			w.Write([]byte("Hello from /os"))
		default:
			w.Write([]byte("Hello from default"))
		}

	}
}

func main() {
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}

}
