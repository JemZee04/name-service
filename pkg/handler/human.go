package handler

import (
	"github.com/gin-gonic/gin"
	"name-service/model"
	"net/http"
	"strconv"
)

const (
	pageSize = 10
)

func (h *Handler) createHuman(c *gin.Context) {
	var input model.Human
	var err error

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Age, err = h.getAge(input.Name)
	input.Gender, err = h.getGender(input.Name)
	input.Nationality, err = h.getNationality(input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Human.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK, map[string]interface{}{
			"id": id,
		},
	)
}

type getAllHumansResponse struct {
	Data []model.Human `json:"data"`
}

func (h *Handler) getAllHumans(c *gin.Context) {
	var input model.FilterHuman
	var err error
	input.Name = c.DefaultQuery("name", "")
	input.Surname = c.DefaultQuery("surname", "")
	input.Patronymic = c.DefaultQuery("patronymic", "")
	input.MinAge, err = strconv.Atoi(c.DefaultQuery("minAge", "0"))
	input.MaxAge, err = strconv.Atoi(c.DefaultQuery("maxAge", "1000"))
	input.Gender = c.DefaultQuery("gender", "")
	input.Nationality = c.DefaultQuery("nationality", "")
	input.Page, err = strconv.Atoi(c.DefaultQuery("maxAge", "1"))

	humans, err := h.services.Human.GetAll(input, pageSize)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK, getAllHumansResponse{
			Data: humans,
		},
	)
}

func (h *Handler) updateHuman(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateHumanInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})

}

func (h *Handler) deleteHuman(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Human.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK, statusResponse{
			Status: "ok",
		},
	)
}
