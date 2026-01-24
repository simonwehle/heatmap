package cmd

import (
	"log"
	"net/http"
	"text/template"

	"running-heatmap/internal/files"
	"running-heatmap/internal/style"
)

func Execute() {
	mapStyle := style.GetMapStyle()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		tmpl, _ := template.ParseFiles("./web/index.html")
		tmpl.Execute(w, map[string]string{"MapStyle": mapStyle})
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/api/gpx", files.ListGPXFiles)
	addr := ":8080"
	log.Println("Server started at port" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}