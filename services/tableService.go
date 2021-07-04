package service

import (
	"net/http"
	"strconv"
)

func addGuest(c *gin.Context) {
	name, err := strconv.ParseInt(c.Param("name"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Enter proper guest name",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

}

func getGuestList(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)

		return
	}

	user, apiErr := services.UsersService.GetUser(userID)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}

func editGuestList(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("name"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Enter proper guest name",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)

		return
	}

	user, apiErr := services.UsersService.GetUser(userID)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}

func deleteGuests(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("name"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Enter proper guest name",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)

		return
	}

	user, apiErr := services.UsersService.GetUser(userID)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
