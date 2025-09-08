package api

type CategoryCreateRequest struct {
	Name string `validate:"required,min=3,max=200"`
}

type CategoryUpdateRequest struct {
	Id   string `validate:"required"`
	Name string `validate:"required,min=3,max=200"`
}
