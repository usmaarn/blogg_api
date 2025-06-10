package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/usmaarn/blogg/models"
	"github.com/usmaarn/blogg/utils"
)

func (h *Handler) PostsHandler(w http.ResponseWriter, r *http.Request) {
	var data models.ResponseData
	posts, err := h.Store.GetPosts()
	if err != nil {
		data.Errors = []string{err.Error()}
	} else {
		data.Data = posts
	}

	fmt.Println(data)

	utils.RenderTemplate(w, []string{"base", "index"}, data)
}

func (h *Handler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, []string{"base", "posts/create"}, nil)
}

func (h *Handler) StorePostHandler(w http.ResponseWriter, r *http.Request) {
	var body = models.CreatePost{
		Title:   r.FormValue("title"),
		Content: template.HTML(r.FormValue("content")),
	}

	err := h.Validate.Struct(body)

	var data = models.ResponseData{
		Data: body,
	}

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			data.Errors = utils.FormatValidationErrors(errs)
		} else {
			data.Errors = []string{err.Error()}
		}
		utils.RenderTemplate(w, []string{"base", "posts/create"}, data)
		return
	}

	var post = models.Post{
		UserID:  1,
		Title:   body.Title,
		Content: body.Content,
	}

	err = h.Store.SavePost(&post)
	if err != nil {
		data.Errors = []string{err.Error()}
		utils.RenderTemplate(w, []string{"base", "posts/create"}, data)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
