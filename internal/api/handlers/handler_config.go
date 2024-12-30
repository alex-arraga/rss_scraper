package handlers

import (
	"github.com/alex-arraga/rss_project/internal/di"
)

type HandlerConfig struct {
	Container *di.Container
}

func NewHandlerConfig(c *di.Container) *HandlerConfig {
	return &HandlerConfig{
		Container: c,
	}
}