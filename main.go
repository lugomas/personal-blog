package main

import (
	"log"
	"net/http"
	"roadmaps/projects/personal-blog/internal/auth"
	"roadmaps/projects/personal-blog/internal/handler"
)

func main() {

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/view/", handler.ViewHandler)
	http.HandleFunc("/logout", handler.LogoutHandler)

	http.HandleFunc("/admin/", auth.BasicAuthMiddleware(handler.AdminHandler))
	http.HandleFunc("/save/", auth.BasicAuthMiddleware(handler.SaveHandler))
	http.HandleFunc("/edit/", auth.BasicAuthMiddleware(handler.EditHandler))
	http.HandleFunc("/new/", auth.BasicAuthMiddleware(handler.NewHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
