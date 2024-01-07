package ApiData

type ApiAlbum struct {
 	Result []struct {
		album_group string `json:"album_group"`
		album_type string `json:"album_type"`
		artists []string `json:"artists"`
	}

}

type ApiArtist struct {
	

}