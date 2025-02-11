package handler

import (
	"be-post/helpers/validator"
	"be-post/post"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type postHandler struct {
	postService post.PostService
}

func NewPostHandler(postService post.PostService) *postHandler {
    return &postHandler{postService: postService}
}

func (h *postHandler) Index(ctx *gin.Context) {
	// Set default value
	defaultLimit := 10
	defaultOffset := 0

	// Get the limit from query parameters, with a default value
	limitStr := ctx.DefaultQuery("limit", strconv.Itoa(defaultLimit))
	offsetStr := ctx.DefaultQuery("offset", strconv.Itoa(defaultOffset))

	// Convert to integer
	limit, err := strconv.Atoi(limitStr)
	
	if err != nil {
		// If conversion fails, use default limit
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(offsetStr)
	
	if err != nil {
		// If conversion fails, use default offset
		offset = defaultOffset
	}

	posts, err := h.postService.Index(limit, offset)

	if err!= nil {
		response := Response{
            Success: false,
            Message: "Something went wrong",
            Data:    err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
	}

	response := Response{
        Success: true,
        Message: "Posts retrieved successfully",
        Data:    posts,
    }

	ctx.JSON(http.StatusOK, response)
}

func (h *postHandler) Store(ctx *gin.Context) {
	var input post.InputStorePost

	err := ctx.ShouldBindJSON(&input)
	if err!= nil {
        response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
        return 
    }

	// form validation
	errorsMap, err := validator.ValidateStruct(input)
	if err!= nil {
		response := Response{
            Success: false,
            Message: "Validation failed",
            Data:    errorsMap,
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }
	
	// save to database
	newPost, err := h.postService.Store(input)
	if err!= nil {
		response := Response{
            Success: false,
            Message: "Something went wrong",
            Data:    err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
	}

	response := Response{
		Success: true,
		Message: "Post has been stored successfully",
		Data: newPost,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *postHandler) Show(ctx *gin.Context) {
	var input post.InputPostDetail

	err := ctx.ShouldBindUri(&input)
	if err!= nil {
        response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
        return 
    }

	newPost, err := h.postService.SelectById(input)

	if err!= nil {
		response := Response{
            Success: false,
            Message: "Something went wrong",
            Data:    err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
	}

	response := Response{
		Success: true,
		Message: "Get Post By ID",
		Data: newPost,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *postHandler) Update(ctx *gin.Context) {
	var inputDetail post.InputPostDetail
	var input post.InputUpdatePost

	err := ctx.ShouldBindUri(&inputDetail)
	if err!= nil {
        response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
        return 
    }

	err = ctx.ShouldBindJSON(&input)

	if err!= nil {
		response := Response{
            Success: false,
            Message: "Something went wrong",
            Data:    err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

	// form validation
	errorsMap, err := validator.ValidateStruct(input)
	if err!= nil {
		response := Response{
            Success: false,
            Message: "Validation failed",
            Data:    errorsMap,
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

	newPost, err := h.postService.Update(inputDetail, input)
	if err!= nil {
		response := Response{
            Success: false,
            Message: "Something went wrong",
            Data:    err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

	response := Response{
		Success: true,
		Message: "Post updated successfully",
		Data: newPost,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *postHandler) Destroy(ctx *gin.Context) {
	var input post.InputPostDetail

	err := ctx.ShouldBindUri(&input)
	if err!= nil {
        response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
        return 
    }

	newPost, err := h.postService.Destroy(input)

	if err!= nil {
		response := Response{
            Success: false,
            Message: "Something went wrong",
            Data:    err.Error(),
        }
        ctx.JSON(http.StatusBadRequest, response)
        return
	}

	response := Response{
		Success: true,
		Message: "Post has been deleted successfully",
		Data: newPost,
	}

	ctx.JSON(http.StatusOK, response)
}