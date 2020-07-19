package crud

import "database/sql"

type repositoryPostsCRUD struct {
	db *sql.DB
}

func (r *repositoryPostsCRUD) NewRepositoryPostsCRUD(db *sql.DB) *repositoryPostsCRUD {
	return &repositoryPostsCRUD{db}
}
