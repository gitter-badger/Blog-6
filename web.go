package main

import (
	"fmt"

	"github.com/captncraig/blog/designBrowser"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/resume", resume)
	http.HandleFunc("/colors", designSeeds)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	if bind == ":" {
		bind = ":3322"
	}
	fmt.Printf("listening on %s...\n", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func resume(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/resume.html")
}

func designSeeds(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.ServeFile(w, r, "designBrowser/index.html")
	} else {
		designBrowser.GetMore(w, r, q)
	}
}
