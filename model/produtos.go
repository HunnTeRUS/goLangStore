package model

import (
	"stock/db"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Product {
	db := db.DatabaseConnect()

	productsSelect, err := db.Query("select * from alura_loja.produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	produtos := []Product{}

	for productsSelect.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = productsSelect.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func GetProductById(productId string) Product {
	db := db.DatabaseConnect()

	productsSelect, err := db.Query("select * from alura_loja.produtos where id = " + productId)

	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	productsSelect.Next()

	var id, quantidade int
	var nome, descricao string
	var preco float64

	err = productsSelect.Scan(&id, &nome, &descricao, &preco, &quantidade)

	if err != nil {
		panic(err.Error())
	}

	p.Id = id
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Quantidade = quantidade

	defer db.Close()

	return p
}

func CreateProduct(name string, description string, price float64, quantity int) {
	db := db.DatabaseConnect()

	insert, err := db.Prepare("INSERT INTO alura_loja.produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic("Invalid prepare update: " + err.Error())
	}

	insert.Exec(name, description, price, quantity)

	defer db.Close()
}

func EditProduct(id string, name string, description string, price float64, quantity int) {
	db := db.DatabaseConnect()

	edit, err := db.Prepare("UPDATE alura_loja.produtos SET (nome, descricao, preco, quantidade) = ($1, $2, $3, $4) WHERE id = $5")

	if err != nil {
		panic("Invalid prepare edit: " + err.Error())
	}

	edit.Exec(name, description, price, quantity, id)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DatabaseConnect()

	remove, err := db.Prepare("DELETE FROM alura_loja.produtos WHERE id = $1")

	if err != nil {
		panic("Invalid prepare remove: " + err.Error())
	}

	remove.Exec(id)

	defer db.Close()
}
