package model

type Usuario struct {
	IdUsuario      string
	NomeUsuario    string
	Email          string
	Senha          string
	ListaDeTarefas []string
}
