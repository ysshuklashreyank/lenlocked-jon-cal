package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ysshuklashreyank/lenslocked/controllers"
	"github.com/ysshuklashreyank/lenslocked/templates"
	"github.com/ysshuklashreyank/lenslocked/views"
)

func galleriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Hare Krishna haribol</h1> ")
	id := chi.URLParam(r, "id")
	fmt.Fprint(w, id)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.Get("/signup", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))))
	r.Get("/galleries/{id}", galleriesHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3001...")
	http.ListenAndServe(":3001", r)
}
