package repository

import "testing/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
