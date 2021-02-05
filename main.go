package main

import (
	"fmt"

	"github.com/2785/demo/internal/handler"

	"github.com/2785/demo/internal/repository/fake"
	"github.com/gin-gonic/gin"
)

func main() {
	// router shenanigans

	// server shenanigans

	// that we don't care about

	r := gin.Default()

	fakeRepo, err := fake.New("./beers.json")
	if err != nil {
		panic(fmt.Errorf("yo the file path doesn't work: %w", err))
	}

	h := handler.HandleGet(fakeRepo)

	r.GET("/thing/:id", h)

	r.Run()
}
