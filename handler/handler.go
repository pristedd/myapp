package handler

import (
	"github.com/labstack/echo/v4"
	"myapp/entity"
	"myapp/logic"
	"net/http"
	"strconv"
)

type PersonHandler struct {
	PersonLogic logic.PersonLogic
}

func (h *PersonHandler) GetPersons(c echo.Context) error {
	persons, err := h.PersonLogic.GetPersons()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) GetPerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	person, err := h.PersonLogic.GetPerson(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) CreatePerson(c echo.Context) error {
	person := new(entity.Person)
	if err := c.Bind(person); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdPerson, err := h.PersonLogic.CreatePerson(person)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, createdPerson)
}

func (h *PersonHandler) UpdatePerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	person := new(entity.Person)
	if err := c.Bind(person); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedPerson, err := h.PersonLogic.UpdatePerson(id, person)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedPerson)
}

func (h *PersonHandler) DeletePerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.PersonLogic.DeletePerson(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
