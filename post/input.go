package post

type InputStorePost struct {
    Title       string `json:"title" validate:"required,min=20"`
	Content    string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status string `json:"status" validate:"required,oneof=draft publish"`
}

type InputUpdatePost struct {
    Title       string `json:"title" validate:"required,min=20"`
	Content    string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status string `json:"status" validate:"omitempty,oneof=draft publish trash"`
}

type InputPostDetail struct {
	ID int `uri:"id" binding:"required"`
}