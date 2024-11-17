package web

type CategoryUpdateRequset struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=200,min=1"`
}
