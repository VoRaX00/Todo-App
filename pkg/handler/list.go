package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoApp/Entity"
)

func (h *Handler) createList(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		return
	}

	var input Entity.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.service.TodoList.Create(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getAllLists(c *gin.Context) {}

func (h *Handler) getListById(c *gin.Context) {}

func (h *Handler) updateList(c *gin.Context) {}

func (h *Handler) deleteList(c *gin.Context) {}
