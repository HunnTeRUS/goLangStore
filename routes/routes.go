package routes

import (
	"net/http"
	"stock/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.NewProducts)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Remove)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.EditFunc)
}
