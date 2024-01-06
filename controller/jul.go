package controller

import(
	"net/http"
	initTemp "TpApiSpotify/temps"
)

func Jul(w http.ResponseWriter, r *http.Request){
	initTemp.Temp.ExecuteTemplate(w, "jul", nil)
}