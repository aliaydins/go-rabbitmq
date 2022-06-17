package order

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) health(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"message": "Order service is up"})
}

func (h *Handler) Create(g *gin.Context) {

	req := new(CreateOrderRequest)
	if err := g.BindJSON(req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := h.service.CreateOrder(req)
	if err != nil {
		fmt.Println(err.Error())
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}
	g.JSON(http.StatusCreated, order)

}
