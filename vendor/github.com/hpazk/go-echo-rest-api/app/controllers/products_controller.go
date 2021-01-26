package controllers

import (
	"fmt"
	"net/http"

	"github.com/hpazk/go-echo-rest-api/app/database"
	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/hpazk/go-echo-rest-api/app/middlewares"
	ProductModel "github.com/hpazk/go-echo-rest-api/app/models/products"
	"github.com/labstack/echo/v4"
)

type (
	ProductsController struct {
	}

	AddProductRequest struct {
		Name          string `json:"name" validate:"required"`
		Description   string `json:"description" validate:"required"`
		Price         int    `json:"price" validate:"required"`
		ProductRating int    `json:"product_rating"`
		PicturePath   string `json:"picture_path" validate:"required"`
		CategoryID    int    `json:"category_id" validate:"required"`
	}
)

func (controller ProductsController) Routes() []helpers.Route {
	return []helpers.Route{
		{
			Method:     echo.GET,
			Path:       "/products",
			Handler:    controller.GetProducts,
			Middleware: []echo.MiddlewareFunc{middlewares.JWTMiddleWare()},
		},
		{
			Method:     echo.GET,
			Path:       "/products/:productId",
			Handler:    controller.GetProduct,
			Middleware: []echo.MiddlewareFunc{middlewares.JWTMiddleWare()},
		},
		{
			Method:     echo.POST,
			Path:       "/products",
			Handler:    controller.AddProduct,
			Middleware: []echo.MiddlewareFunc{middlewares.JWTMiddleWare()},
		},
		// {
		// 	Method: echo.GET,
		// 	Path:   "/blogs",
		// 	// Handler: controller.GetBlogs,
		// },
		// {
		// 	Method: echo.GET,
		// 	Path:   "/blog/:blogId",
		// 	// Handler: controller.GetBlog,
		// },
	}
}

func (controller ProductsController) GetProduct(c echo.Context) error {
	productId := c.Param("productId")
	db := database.GetInstance()
	var product ProductModel.Product
	err := db.First(&product, "id = ?", productId).Error
	if err != nil {
		response := helpers.ResponseFormatter{
			Code:    404,
			Status:  "error",
			Message: fmt.Sprintf("product id %s not found", productId),
		}

		return c.JSON(http.StatusNotFound, response)
	}

	response := helpers.ResponseFormatter{
		Code:    200,
		Status:  "success",
		Message: "get product successfully",
		Data:    product,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller ProductsController) GetProducts(c echo.Context) error {
	db := database.GetInstance()
	var products []ProductModel.Product
	db.Preload("Category").Find(&products)
	fmt.Println(products)

	response := helpers.ResponseFormatter{
		Code:    200,
		Status:  "success",
		Message: "get products successfully fetched",
		Data:    products,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller ProductsController) AddProduct(c echo.Context) error {
	params := new(AddProductRequest)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := database.GetInstance()
	var product ProductModel.Product
	product.Name = params.Name
	product.Description = params.Description
	product.Price = params.Price
	product.ProductRating = params.ProductRating
	product.PicturePath = params.PicturePath
	product.CategoryID = params.CategoryID
	if err := db.Create(&product).Error; err != nil {
		err_response := helpers.ResponseFormatter{
			Code:    500,
			Status:  "error",
			Message: "something went wrong",
		}
		return c.JSON(http.StatusInternalServerError, err_response)
	}
	response := helpers.ResponseFormatter{
		Code:    201,
		Status:  "success",
		Message: "products saved",
		Data:    product,
	}
	return c.JSON(http.StatusOK, response)
}
