package http

import (
	"github.com/gofiber/fiber/v2"
)

type blogHandler struct {
	uc BlogUsecase
}

func (h *blogHandler) showPosts(c *fiber.Ctx) error {
	posts, _ := h.uc.GetPosts(c.UserContext())
	return c.Render("templates/index.html", posts)
}

func (h *blogHandler) showPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, _ := h.uc.GetPost(c.UserContext(), id)
	return c.Render("templates/post.html", post)
}
