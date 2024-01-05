package routeur

import (
	ctrl "TpApiSpotify/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {

	http.HandleFunc("/", ctrl.Home)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)

}
