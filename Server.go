package main

import (
	"encoding/json"
	"net/http"

	"github.com/Consuelo-Fraga/loja-digiport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8083", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryNome := r.URL.Query().Get("nome")
		if queryNome != "" {
			produtosFiltradosPorNome := produtosPorNome(queryNome)
			json.NewEncoder(w).Encode(produtosFiltradosPorNome)
		} else {
			produtos := criaEstoque()
			json.NewEncoder(w).Encode(produtos)
		}
	} else if r.Method == "POST" {
		var produto model.Produto
		json.NewDecoder(r.Body).Decode(&produto)
		adicionaProduto(produto)

	}
}

// func buscaProdutos(w http.ResponseWriter, r *http.Request) {
// 	queryNome := r.URL.Query().Get("nome")
// 	if queryNome != "" {
// 		produtosFiltradosPorNome := produtosPorNome(queryNome)
// 		json.NewEncoder(w).Encode(produtosFiltradosPorNome)
// 	} else {
// 		produtos := ListaProdutos
// 		json.NewEncoder(w).Encode(produtos)

// 	}
// }
