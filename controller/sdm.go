package controller

import(
	"net/http"
	initTemp "TpApiSpotify/temps"
)

func Sdm(w http.ResponseWriter, r *http.Request){
	initTemp.Temp.ExecuteTemplate(w, "sdm", nil)
}