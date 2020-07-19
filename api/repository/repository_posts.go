package repository

import "gotut/api/models"

type PostRepository interface {
	FindAll() ([]models.Post, error)
}
