package order

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	h.initRoutes(router)
	return router
}

func (h *Handler) initRoutes(router *gin.Engine) {

	rGroup := router.Group("/order")
	rGroup.GET("/health", h.health)
	rGroup.POST("/", h.Create)

}
