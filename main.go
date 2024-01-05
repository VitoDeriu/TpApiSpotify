package main

import (
	t "TpApiSpotify/temps"
	r "TpApiSpotify/routeur"
)

func main () {
	t.InitTemplate()
	r.InitServe()
}