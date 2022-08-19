package dto

type Feira struct {
	SetCens    string `json:"set_cens"`
	CodDist    string `json:"cod_dist"`
	Distrito   string `json:"distrito"`
	SubPrefe   string `json:"sub_prefe"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	NomeFeira  string `json:"nome_feira"`
	Registro   string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
	Id         int64  `json:"id"`
	Long       int64  `json:"long"`
	Lat        int64  `json:"lat"`
	Areap      int64  `json:"areap"`
	CodSubPref int64  `json:"cod_sub_pref"`
}
