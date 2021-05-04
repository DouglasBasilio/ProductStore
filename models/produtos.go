package models

import "github.com/ProductStore/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBancoDeDados()

	selectDosProdutos, err := db.Query("select * from produtos order by 1")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDeDados() //conecta com banco de dados

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}

//Função para trazer as informações do produto
func EditaProduto(id string) Produto {
	db := db.ConectaBancoDeDados()

	//Trazer o ID do produto
	produtoDoBanco, err := db.Query("select * from produtos where id = $1", id)

	//Se erro não for igual a null, exibe o erro
	if err != nil {
		panic(err.Error())
	}

	//Preparar para armazenar as informações oriundas do BD
	//Crio uma variável que é igual a uma instância de um Produto{}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//Colocar o ID que trouxe do banco dentro da memória da variável, colocando o &
		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		//Caso não tenha erro e consegui armazenar dentro das variáveis,
		//pego o produto e falo qual é a variável
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	//Depois que tudo executou, fecho a conexão com o banco
	defer db.Close()
	//Retorno o produto
	return produtoParaAtualizar

}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
