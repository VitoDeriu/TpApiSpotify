package controller

import(
	"net/http"
	initTemp "TpApiSpotify/temps"
	callApi "TpApiSpotify/api"
)

func Jul(w http.ResponseWriter, r *http.Request){

	callApi.ApiJul()
	initTemp.Temp.ExecuteTemplate(w, "jul", nil)
}