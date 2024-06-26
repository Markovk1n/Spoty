package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/models"
)

type singInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) singUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "You registred",
	})

}

func (h *Handler) singIn(c *gin.Context) {
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:  "session_token",
		Value: token,
		Path:  "/",
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) singOut(c *gin.Context) {

	_, err := c.Request.Cookie("session_token")
	if err != nil {

		c.Redirect(http.StatusFound, "/")
		return
	}

	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(c.Writer, cookie)

	c.Redirect(http.StatusFound, "/")

}

func (h *Handler) authPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth.html", nil)
}
