package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/css/").Handler(http.FileServer(http.FS(ffs)))
	r.PathPrefix("/js/").Handler(http.FileServer(http.FS(ffs)))
	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/img/{id:.*\\/?}", Home).Methods(http.MethodGet)
	r.HandleFunc("/{blog:blog\\/?}", Blog).Methods(http.MethodGet)
	r.HandleFunc("/blog/{entryName:.*\\/?}", BlogEntry).Methods(http.MethodGet)
	r.HandleFunc("/{music:music\\/?}", Music).Methods(http.MethodGet)
	r.HandleFunc("/{video:video\\/?}", Video).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/{multivideo:multivideo\\/?}", MultiVideo).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/video/{connections:connections\\/?}", VideoConnections).Methods(http.MethodGet)
	r.HandleFunc("/multivideo/{connections:connections\\/?}", MultiVideoConnections).Methods(http.MethodGet)
	return r
}
