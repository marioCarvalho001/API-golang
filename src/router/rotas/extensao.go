package rotas

import (
	"api/src/controllers"
	"net/http"
)

var ExtensoesRotas = []Rota{

	//lista todas as frotas
	{
		URI:    "/fleets",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarFleets,
	},
	//inseri uma frota
	{URI: "/fleets",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarFleets,
	},
	{
		URI:    "/fleets/{id}/alerts",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarFleet_alert,
	},

	//inseri uma frota
	{URI: "/fleets/{id}/alerts",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarFleet_alert,
	},
	{
		URI:    "/vehicles",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarVehicle,
	},

	//inseri uma frota
	{URI: "/vehicles",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarVehicle,
	},
	{
		URI:    "/vehicles/{id}/positions",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPosition,
	},

	//inseri uma frota
	{URI: "/vehicles/{id}/positions",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPosition,
	},
	{URI: "/database",
		Metodo: http.MethodDelete,
		Funcao: controllers.Deletadb,
	},
}
