package temps

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplate() {
	temp, errTemp := template.ParseGlob("./temps/*.html")
	if errTemp != nil {
		fmt.Printf("Oups une erreur lié au Templates : %v", errTemp.Error())
		os.Exit(1)
	}
	Temp = temp
}
