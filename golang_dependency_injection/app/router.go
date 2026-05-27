package app

import (
	v1 "golang_restful_api/api/v1/categories"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bootstrap *Bootstrap) *httprouter.Router {
	router := httprouter.New()

	// Register router per modul
	v1.CategoriesRouter(router, bootstrap.CategoryController)

	return router
}
