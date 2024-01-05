package controller

import(
	"net/http"
	initTemp "TpApiSpotify/temps"
)

func Home(w http.ResponseWriter, r *http.Request){
	initTemp.Temp.ExecuteTemplate(w, "home", nil)
}