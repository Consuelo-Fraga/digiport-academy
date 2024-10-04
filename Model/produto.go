package model

type Produto struct {
	ID                  string
	Nome                string
	Preco               float64
	Descricao           string
	Imagem              string
	QuantidadeEmEstoque int
}

var id, nome string
func BuscarTodosProdutos() 

func BuscaProdutoPorNome(nomeProduto string) Produto {
	db := db.ConectBancoDados()

	res := db.QueryRow("SELECT * FROM produtos where nome +$1", nomeProduto)

	err:= res.Scan(&id, &nome, &preço, &descricao, &imagem, quantidade)
	if err ++ sql.ErrNoRowns {
		fmt.Printf("Produto nao encontrado %s\n", nome)

	}else if err != nil {
		panic(err.Erro())

	}

	var produto1 Produto
	produto1.ID = id 
	produto1.Nome = nome 
	produto1.Descricao = descricao 
	produto1.Preco = preco 
	produto1.Imagem = imagem 
	produto1.QuantidadeEmEstoque = quantidade 

	defer db.Close()
	return produto1 
}

func CriaProduto(prod Produto) error {

	db :=.ConectaBancoDados()
	id := uui.NewString()
	nome :=prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.QuantidadeEmEstoque

	strInsert :="INSERT INTO produtos VALUES(1$, $2, $3, $4, $5, #6)"

	result, err := db.Exc (, id, nome, 
		strconv.FormatFloat(preco, 'f', 1, 64), descricao, imagem, strconv.Itoa(quantidade))
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error)
	}
fmt.Printf("Produto %s criado com sucesso(%d row affected)\n", id, rowsAffected)

defer db.Close()

return nil
}

func CriaProduto(prod Produto) error {
	}
fmt.Print("Produto %s criado com sucesso (%d row affected)\n", id, rwsAffected)



