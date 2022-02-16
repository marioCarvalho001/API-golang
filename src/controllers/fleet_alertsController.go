package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarFleet_alert(w http.ResponseWriter, r *http.Request) {

	//responsavel por puxar o id passado na url
	parametros := mux.Vars(r)
	fleet_id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioFleet_alert(db)
	fleet_alert, erro := repositorio.BuscarFleet_alert(fleet_id)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, 200, fleet_alert)

}

//Gabiarra para puxar o id da url

func CriarFleet_alert(w http.ResponseWriter, r *http.Request) {

	//responsavel por puxar o id passado na url e passar para o repositorio fazer a filtragem
	parametros := mux.Vars(r)
	fleet_id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	repositorios.Urlcd(fleet_id)

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	//passanndo o struct de  pra json
	var fleet_alert models.Fleet_alert
	if erro = json.Unmarshal(corpoRequest, &fleet_alert); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = fleet_alert.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioFleet_alert(db)
	fleet_alert.ID, erro = repositorio.CriarFleet_alert(fleet_alert)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, 201, fleet_alert.ID)

}
