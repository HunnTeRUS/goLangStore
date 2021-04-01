package controller

import (
	"fmt"
	"net/http"
	"stock/model"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", model.GetAllProducts())
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	templates.ExecuteTemplate(w, "Edit", model.GetProductById(id))
}

func NewProducts(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func EditFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	if r.Method == "POST" {
		id := r.URL.Query().Get("id")

		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		quantity := r.FormValue("quantidade")
		price := r.FormValue("preco")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic("Price convertion error: " + err.Error())
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			panic("Quantity convertion error: " + err.Error())
		}

		model.EditProduct(id, name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		quantity := r.FormValue("quantidade")
		price := r.FormValue("preco")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic("Price convertion error: " + err.Error())
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			panic("Quantity convertion error: " + err.Error())
		}

		model.CreateProduct(name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	model.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}
