package main

import (
	"fmt"
	"log"
	"net/http"

	renderer "github.com/IDOMATH/CheetahRender/Render"
)

func main() {
	fmt.Println("Hello world!")
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	rr := renderer.NewRenderer("./views", ".html", "./views", ".html", false)

	router.HandleFunc("GET /", handleHome(rr))

	log.Fatal(server.ListenAndServe())
}

func handleHome(rr *renderer.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rr.Render(w, r, "home.html", make(map[string]interface{}))
	}
}
