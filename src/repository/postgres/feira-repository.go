package postgres

import (
	"api-unico/dto"
	"api-unico/errors"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type feiraRepositorySQL struct {
	db *sqlx.DB
}

func NewfeiraRepositorySQL(db *sqlx.DB) feiraRepositorySQL {
	return feiraRepositorySQL{
		db: db,
	}
}

func (h *feiraRepositorySQL) Get(id int64) (dto.Feira, error) {
	var feirasNullable []feiraDtoNullable
	err := h.db.Select(&feirasNullable,
		`SELECT id, long, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, 
		regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia 
		FROM deinfo 
		WHERE id=$1 AND ativo=1`, id)
	if err != nil {
		return dto.Feira{}, err
	}

	if len(feirasNullable) == 0 {
		return dto.Feira{}, errors.ErrResourceNotFound
	}

	feiras := NullableToDto(feirasNullable)
	return feiras[0], nil
}

func (h *feiraRepositorySQL) Create(feira dto.Feira) (int64, error) {
	var id int64
	err := h.db.QueryRow(
		`INSERT INTO deinfo(
			long, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, 
			regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		RETURNING "id"`,
		feira.Long, feira.Lat, feira.SetCens, feira.Areap, feira.CodDist, feira.Distrito,
		feira.CodSubPref, feira.SubPrefe, feira.Regiao5, feira.Regiao8, feira.NomeFeira,
		feira.Registro, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (h *feiraRepositorySQL) Update(feira dto.Feira) error {
	res, err := h.db.Exec(
		`UPDATE deinfo 
		 SET long=$1, lat=$2, setcens=$3, areap=$4, coddist=$5, distrito=$6, codsubpref=$7, 
		 subprefe=$8, regiao5=$9, regiao8=$10, nome_feira=$11, registro=$12, logradouro=$13, 
		 numero=$14, bairro=$15, referencia=$16
		WHERE id=$17`,
		feira.Long, feira.Lat, feira.SetCens, feira.Areap, feira.CodDist, feira.Distrito,
		feira.CodSubPref, feira.SubPrefe, feira.Regiao5, feira.Regiao8, feira.NomeFeira,
		feira.Registro, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia, feira.Id,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.ErrResourceNotFound
	}

	return nil
}

func (h *feiraRepositorySQL) Delete(id int64) error {
	res, err := h.db.Exec(`UPDATE deinfo SET ativo=0 WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.ErrResourceNotFound
	}

	return nil
}

func (h *feiraRepositorySQL) FindBy(column string, query string) ([]dto.Feira, error) {
	var feirasNullable []feiraDtoNullable
	err := h.db.Select(&feirasNullable,
		fmt.Sprintf(`SELECT id, long, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, 
		regiao5, regiao8, nome_feira, registro, logradouro, numero, bairro, referencia 
		FROM deinfo 
		WHERE ativo=1 AND %s LIKE $1`, column), "%"+query+"%")
	if err != nil {
		return []dto.Feira{}, err
	}

	if len(feirasNullable) == 0 {
		return []dto.Feira{}, errors.ErrResourceNotFound
	}

	feiras := NullableToDto(feirasNullable)
	return feiras, nil
}

type feiraDtoNullable struct {
	CodDist    sql.NullString `json:"cod_dist" db:"coddist"`
	Distrito   sql.NullString `json:"distrito" db:"distrito"`
	SubPrefe   sql.NullString `json:"sub_prefe" db:"subprefe"`
	Regiao5    sql.NullString `json:"regiao5" db:"regiao5"`
	Regiao8    sql.NullString `json:"regiao8" db:"regiao8"`
	NomeFeira  sql.NullString `json:"nome_feira" db:"nome_feira"`
	Registro   sql.NullString `json:"registro" db:"registro"`
	Logradouro sql.NullString `json:"logradouro" db:"logradouro"`
	Numero     sql.NullString `json:"numero" db:"numero"`
	Bairro     sql.NullString `json:"bairro" db:"bairro"`
	Referencia sql.NullString `json:"referencia" db:"referencia"`
	SetCens    sql.NullInt64  `json:"set_cens" db:"setcens"`
	Id         sql.NullInt64  `json:"id" db:"id"`
	Long       sql.NullInt64  `json:"long" db:"long"`
	Lat        sql.NullInt64  `json:"lat" db:"lat"`
	Areap      sql.NullInt64  `json:"areap" db:"areap"`
	CodSubPref sql.NullInt64  `json:"cod_sub_pref" db:"codsubpref"`
}

func NullableToDto(input []feiraDtoNullable) (output []dto.Feira) {
	for _, v := range input {
		f := dto.Feira{
			CodDist:    v.CodDist.String,
			Distrito:   v.Distrito.String,
			SubPrefe:   v.SubPrefe.String,
			Regiao5:    v.Regiao5.String,
			Regiao8:    v.Regiao8.String,
			NomeFeira:  v.NomeFeira.String,
			Registro:   v.Registro.String,
			Logradouro: v.Logradouro.String,
			Numero:     v.Numero.String,
			Bairro:     v.Bairro.String,
			Referencia: v.Referencia.String,
			SetCens:    v.SetCens.Int64,
			Id:         v.Id.Int64,
			Long:       v.Long.Int64,
			Lat:        v.Lat.Int64,
			Areap:      v.Areap.Int64,
			CodSubPref: v.CodSubPref.Int64,
		}
		output = append(output, f)
	}
	return output
}
