package errors

import "errors"

var (
	ErrEmptyNomeFeira   = errors.New("empty nome_feira")
	ErrEmptyRegiao5     = errors.New("empty região5")
	ErrEmptyBairro      = errors.New("empty bairro")
	ErrEmptyDistrito    = errors.New("empty distrito")
	ErrEmptyLogradouro  = errors.New("empty logradouro")
	ErrResourceNotFound = errors.New("resource not found")
)
