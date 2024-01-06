package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ApiJul() {
	urlApi := "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy" //url de l'artiste jul

	httpClient := http.Client{ //pour cancel la request si elle met trop de temps
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil) //création de la request http vers l'api pour récup les info
	if errReq != nil {
		fmt.Println("oups il y a un probleme avec la requete : ", errReq.Error())
	}

	req.Header.Set("jsp quoi écrire", "là / non plus")

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		fmt.Println("oups il y a un probleme avec la requete : ", errRes.Error())
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("oups il y a un probleme avec le corps de la requete : ", errBody.Error())
	}
}
