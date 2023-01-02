package controller

import (
	"net/http"
	"product/config"
	"product/dto"
	"product/service"
	"strconv"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type ProductController struct {
	ProductService service.ProductService
}

func ProviderProductController(p service.ProductService) ProductController {
	return ProductController{ProductService: p}
}

func (p *ProductController) GetAll(e echo.Context) error {

	products := p.ProductService.GetAll()

	if len(products) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.ProductNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, products)
}

func (p *ProductController) GetByID(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	product := p.ProductService.GetByID(uint(id))
	if product.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.ProductNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, product)
}

func (p *ProductController) GetByCategoryID(e echo.Context) error {

	category_id, err := strconv.ParseUint(e.Param("category_id"), 10, 64)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	products := p.ProductService.GetByCategoryID(uint(category_id))
	if len(products) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, products)
	}

	return config.SuccessResponse(e, http.StatusOK, products)
}

func (p *ProductController) Create(e echo.Context) error {

	var product dto.Product
	if e.Bind(&product) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := p.ProductService.Create(product)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (p *ProductController) Update(e echo.Context) error {

	var product dto.Product
	if e.Bind(&product) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := p.ProductService.Update(product)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (p *ProductController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 32)
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err = p.ProductService.Delete(uint(id))
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, "Delete Success")
}
