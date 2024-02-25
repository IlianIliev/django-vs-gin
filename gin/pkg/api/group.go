package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ilianiliev/django-vs-gin/gin/pkg/entities"
	"github.com/ilianiliev/django-vs-gin/gin/pkg/services"
	"net/http"
)

type Member struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type getMemberURIParams struct {
	Name string `uri:"name" binding:"required,alpha"`
}

func ListGroupMembers(groupService services.Group) gin.HandlerFunc {
	return func(c *gin.Context) {

		members, err := groupService.ListMembers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		apiMembers := make([]Member, len(members))
		for i, member := range members {
			fmt.Println("member", member, i)
			apiMembers[i] = mapMemberEntityToMember(member)
		}

		c.JSON(http.StatusOK, apiMembers)
	}

}

func mapMemberEntityToMember(member entities.Member) Member {
	return Member{
		Name: member.Name,
		Role: member.Role,
	}
}

func GetGroupMember(groupService services.Group) gin.HandlerFunc {
	return func(c *gin.Context) {
		var uriParams getMemberURIParams

		if err := c.ShouldBindUri(&uriParams); err != nil {
			var validationErr validator.ValidationErrors

			if errors.As(err, &validationErr) {
				details := make([]string, len(validationErr))
				for i, fieldErr := range validationErr {
					details[i] = fmt.Sprintf(
						"Validation for '%s' failed on '%s'",
						fieldErr.Field(),
						fieldErr.Tag(),
					)
				}

				c.JSON(http.StatusBadRequest, details)
				return
			}

			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		member, err := groupService.GetMember(uriParams.Name)
		if err != nil {
			if errors.Is(err, services.ErrMemberNotFound) {
				c.Status(http.StatusNotFound)
				return
			}
		}

		c.JSON(http.StatusOK, mapMemberEntityToMember(member))
	}
}
