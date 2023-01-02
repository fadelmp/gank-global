package controller

import (
	"net/http"
	"product/config"
	"product/dto"
	"product/service"
	"strconv"

	"github.com/labstack/echo"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func ProviderCategoryController(c service.CategoryService) CategoryController {
	return CategoryController{
		CategoryService: c,
	}
}

func (c *CategoryController) GetAll(e echo.Context) error {

	categories := c.CategoryService.GetAll()

	if len(categories) == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CategoryNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, categories)
}

func (c *CategoryController) GetByID(e echo.Context) error {

	id, _ := strconv.ParseUint(e.Param("id"), 10, 32)

	category := c.CategoryService.GetByID(uint(id))

	if category.ID == 0 {
		return config.SuccessResponse(e, http.StatusNoContent, config.CategoryNotFound)
	}

	return config.SuccessResponse(e, http.StatusOK, category)
}

func (c *CategoryController) Create(e echo.Context) error {

	var category dto.Category

	if e.Bind(&category) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := c.CategoryService.Create(category)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (c *CategoryController) Update(e echo.Context) error {

	var category dto.Category

	if e.Bind(&category) != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	res, err := c.CategoryService.Update(category)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, res)
}

func (a *CategoryController) Delete(e echo.Context) error {

	id, err := strconv.ParseUint(e.Param("id"), 10, 64)

	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, config.BadRequest)
	}

	err = a.CategoryService.Delete(uint(id))
	if err != nil {
		return config.ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return config.SuccessResponse(e, http.StatusOK, "Delete Success")
}
