package handler

import(
	"github.com/vedantrajput-privafy/url-shortener/repository"
)

// store the Repository instance in the Handler struct
type Handler struct {
	Repository *repository.URLRepository
}

