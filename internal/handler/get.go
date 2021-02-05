package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/2785/demo/internal/models"
)

type DataSource interface {
	GetById(ctx context.Context, id string) (*models.Beer, error)
}

func HandleGet(source DataSource) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		data, err := source.GetById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, fmt.Sprintf("beer with id %s not found", id))
			return
		}

		c.JSON(http.StatusOK, data)
		return
	}
}
