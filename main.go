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

	rr := renderer.NewRenderer("./views", ".go.html", "./views/layouts", ".go.html", false)

	router.HandleFunc("GET /", handleHome(rr))

	fmt.Println("Server running on port: 8080")
	log.Fatal(server.ListenAndServe())
}

func handleHome(rr *renderer.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handleHome func")
		fmt.Println(len(rr.TemplateCache))
		for k, t := range rr.TemplateCache {
			fmt.Println(k, t.Name())
		}
		rr.Render(w, r, "home.go.html", make(map[string]interface{}))
	}
}
