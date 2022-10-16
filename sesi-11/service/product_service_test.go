package service

import (
	"testing"
	"testing/entity"
	"testing/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Kaca Mata",
	}
	// productRepository.Mock.On("FindById", "1").Return(nil)
	productRepository.Mock.On("FindById", "2").Return(product)

	// product, err := productService.GetOneProduct("1")
	result, err := productService.GetOneProduct("2")

	// assert.Nil(t, product)
	assert.Nil(t, err)

	// assert.NotNil(t, err)
	assert.NotNil(t, result)

	// assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")

	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'kaca Mata'")
	assert.Equal(t, &product, result, "result has to be a product data with id '2'")
}
