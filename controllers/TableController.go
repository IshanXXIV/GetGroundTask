package controllers

import (
	"context"
	"net/http"

	"github.com/IshanXXIV/GetGroundTask/requests"

	"github.com/gin-gonic/gin"
)

func AddGuest(c *gin.Context) {
	_ctx, _ := c.Get("context")
	_ = _ctx.(context.Context)

	req := requests.AddGuestRequest{}

	err := c.BindJSON(req)

	print(req)
	if err != nil {
		// apiErr := &utils.ApplicationError{
		// 	Message:    "Enter proper request body",
		// 	StatusCode: http.StatusBadRequest,
		// 	Code:       "bad_request",
		// }

		c.JSON(http.StatusBadRequest, "")
		return
	}

	// name, err := strconv.ParseInt(c.Param("name"), 10, 64)

	if err != nil {
		// apiErr := &utils.ApplicationError{
		// 	Message:    "Enter proper guest name",
		// 	StatusCode: http.StatusBadRequest,
		// 	Code:       "bad_request",
		// }

		c.JSON(http.StatusBadRequest, "")
		return
	}
	return

	// c.JSON(http.StatusOK, response.Success(ctx, url, key, request))
}
