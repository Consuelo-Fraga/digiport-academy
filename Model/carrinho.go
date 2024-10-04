package model

type Carrinho struct {
	ID         string
	Produtos   QuantidadeProduto
	ValorTotal float64
}
type QuantidadeProduto map[string]int
