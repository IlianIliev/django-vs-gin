package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilianiliev/django-vs-gin/gin/pkg/services"
	"net/http"

	"github.com/ilianiliev/django-vs-gin/gin/pkg/api"
)

func home(c *gin.Context) {
	c.String(http.StatusOK, "Gin home page")
}

func main() {
	groupService := services.NewInMemoryGroupService()

	router := gin.Default()
	router.GET("/", home)

	routerGroup := router.Group("/group")

	routerGroup.GET("/", api.ListGroupMembers(groupService))
	routerGroup.GET("/:name", api.GetGroupMember(groupService))

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
