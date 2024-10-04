package main

import "github.com/Consuelo-Fraga/loja-digiport-backend/model"

var produtos []model.Produto = []model.Produto{}

func criaEstoque() {
	produtos = []model.Produto{
		{
			ID:                  "1",
			Nome:                "Cojunto Dino",
			Preco:               39.90,
			Descricao:           "Conjunto Menino Dino",
			Imagem:              "roupa1.jpg",
			QuantidadeEmEstoque: 15,
		},

		{
			ID:                  "2",
			Nome:                "Cojunto Jacare",
			Preco:               39.90,
			Descricao:           "Conjunto Menino Jacare",
			Imagem:              "roupa2.jpg",
			QuantidadeEmEstoque: 20,
		},

		{
			ID:                  "3",
			Nome:                "Cojunto Raposa",
			Preco:               39.90,
			Descricao:           "Conjunto Menino Raposa",
			Imagem:              "roupa3.jpg",
			QuantidadeEmEstoque: 15,
		},

		{
			ID:                  "4",
			Nome:                "Cojunto Tigre",
			Preco:               39.90,
			Descricao:           "Conjunto Menino Tigre",
			Imagem:              "roupa4.jpg",
			QuantidadeEmEstoque: 15,
		},
	}
	//ListaProdutos = produtos

}

func produtosPorNome(nome string) []model.Produto {
	resultado := []model.Produto{}

	for _, produto := range produtos {
		if produto.Nome == nome {
			resultado = append(resultado, produto)
		}
	}
	return resultado
}

func adicionaProduto(produto model.Produto) {
	produtos = append(produtos, produto)
}
