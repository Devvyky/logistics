package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var orderResp = make(map[int]int)

type createOrderParams struct {
	ProductLine string `json:"product_line" binding:"required"`
	Quantity    int64  `json:"quantity" binding:"required,min=1"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	orderResp[250] = 1
	orderResp[500] = 2
	ctx.JSON(http.StatusOK, orderResp)
}
