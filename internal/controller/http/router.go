package http

import (
	"context"

	"github.com/avelex/blog/internal/blog/entity"

	"github.com/gofiber/fiber/v2"
)

type BlogUsecase interface {
	GetPosts(ctx context.Context) ([]entity.Post, error)
	GetPost(ctx context.Context, id string) (entity.Post, error)
}

func NewRouter(router fiber.Router, blogUsecase BlogUsecase) {
	router.Get("/status", func(c *fiber.Ctx) error { return c.SendString("OK") })

	h := blogHandler{uc: blogUsecase}

	router.Get("/", h.showPosts)
	router.Get("/post/:id", h.showPost)
}
