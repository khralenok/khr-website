package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		data := map[string]interface{}{
			"Title":   "Khralenok Dev",
			"Heading": "Hello from Go",
			"Message": "This page is rendered server-side with Go template",
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Go server running at :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
