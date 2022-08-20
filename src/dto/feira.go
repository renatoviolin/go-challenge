package dto

type Feira struct {
	CodDist    string `json:"cod_dist" db:"coddist"`
	Distrito   string `json:"distrito" db:"distrito"`
	SubPrefe   string `json:"sub_prefe" db:"subprefe"`
	Regiao5    string `json:"regiao5" db:"regiao5"`
	Regiao8    string `json:"regiao8" db:"regiao8"`
	NomeFeira  string `json:"nome_feira" db:"nome_feira"`
	Registro   string `json:"registro" db:"registro"`
	Logradouro string `json:"logradouro" db:"logradouro"`
	Numero     string `json:"numero" db:"numero"`
	Bairro     string `json:"bairro" db:"bairro"`
	Referencia string `json:"referencia" db:"referencia"`
	SetCens    int64  `json:"set_cens" db:"setcens"`
	Id         int64  `json:"id" db:"id"`
	Long       int64  `json:"long" db:"long"`
	Lat        int64  `json:"lat" db:"lat"`
	Areap      int64  `json:"areap" db:"areap"`
	CodSubPref int64  `json:"cod_sub_pref" db:"codsubpref"`
}
