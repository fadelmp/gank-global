package request

import (
	"encoding/json"
	entity "order/entity"
	"os"
	"strconv"
)

type ProductRequestContract interface {
	GetAllProduct() ([]entity.Product, error)
	GetProductById(uint) (entity.Product, error)
}

type ProductRequest struct{}

func ProviderProductRequest() ProductRequest {
	return ProductRequest{}
}

func (p *ProductRequest) GetAllProduct() ([]entity.Product, error) {

	var product []entity.Product
	var product_response entity.ProductAllResponse

	var uri = os.Getenv("PRODUCT_SERVICE_URI")
	uri += "product"

	result, err := GetRequest(uri)
	if err == nil {
		json.Unmarshal([]byte(result), &product_response)
		product = product_response.Data
	}

	return product, err

}

func (p *ProductRequest) GetProductById(product_id uint) (entity.Product, error) {

	var product entity.Product
	var product_response entity.ProductResponse

	var uri = os.Getenv("PRODUCT_SERVICE_URI")
	uri += "product/"
	uri += strconv.FormatUint(uint64(product_id), 10)

	result, err := GetRequest(uri)
	if err == nil {
		json.Unmarshal([]byte(result), &product_response)
		product = product_response.Data
	}

	return product, err

}
