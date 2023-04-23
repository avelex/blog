package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type blogHandler struct {
	uc BlogUsecase
}

func (h *blogHandler) showPosts(c *fiber.Ctx) error {
	posts, err := h.uc.GetPosts(c.UserContext())
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return c.Render("index", posts)
}

func (h *blogHandler) showPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := h.uc.GetPost(c.UserContext(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return c.Render("post", post)
}

func (h *blogHandler) createPost(c *fiber.Ctx) error {
	if err := h.uc.CreatePost(c.UserContext()); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
