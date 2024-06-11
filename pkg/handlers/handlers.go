package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplaate(w,"home.page.html")
}


func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplaate(w, "about.page.html")
}

