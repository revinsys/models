package main

var (
ModelTemplate = `package models

// $UPPER_NAME
type $UPPER_NAME struct {

}
`

RepositoryTemplate = `package store

import (
	"github.com/jmoiron/sqlx"
	"$GIT_PATH/internal/app/models"
)

type $UPPER_NAMERepository struct {
	db *sqlx.DB
}
`

StoreTemplate = `// $UPPER_NAME ...
func (s *Store) $UPPER_NAME() *$UPPER_NAMERepository {
	if s.$MODEL_NAMERepository == nil {
		s.$MODEL_NAMERepository = &$UPPER_NAMERepository{db: s.db}
	}

	return s.$MODEL_NAMERepository
}
`)
