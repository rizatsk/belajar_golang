package v1

import (
	"golang_restful_api/api/v1/categories/controller"
	"golang_restful_api/helper"
	"golang_restful_api/middleware"

	"github.com/julienschmidt/httprouter"
)

func CategoriesRouter(router *httprouter.Router, controller controller.CategoryController) {
	path_api := helper.GetPathApi()

	router.POST(path_api, middleware.ChainMiddleware(
		controller.Create,
		middleware.LoggingMiddleware,
		middleware.VerifyTokenMiddleware,
	))
	router.PATCH(path_api, middleware.ChainMiddleware(
		controller.Update,
		middleware.LoggingMiddleware,
		middleware.VerifyTokenMiddleware,
	))
	router.DELETE(path_api+"/:categoryId", middleware.ChainMiddleware(
		controller.Delete,
		middleware.LoggingMiddleware,
		middleware.VerifyTokenMiddleware,
	))
	router.GET(path_api+"/:categoryId", middleware.LoggingMiddleware(controller.FindById))
	router.GET(path_api, controller.FindAll)
}
