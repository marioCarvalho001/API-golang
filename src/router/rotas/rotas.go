package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Representa todas as rotas da api
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

//Colocar todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := ExtensoesRotas

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
