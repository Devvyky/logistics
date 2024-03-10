package api

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

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

	arg := listPackSizeByProductLinesParams{
		ProductLine: req.ProductLine,
	}
	res, err := listPackSizeByProductLines(ctx, arg, server)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var packSizes []int
	for _, p := range res {
		packSizes = append(packSizes, int(p.PackSize))
	}

	resp := fulfilOrder(int(req.Quantity), packSizes)
	ctx.JSON(http.StatusOK, resp)
}

func fulfilOrder(qty int, packSizes []int) map[int]int {
	var packSizeQtyMap = make(map[int]int)
	left, right, remaining := 0, len(packSizes)-1, qty

	if qty <= packSizes[0] {
		packSizeQtyMap[packSizes[0]] = 1
		return packSizeQtyMap
	}

	for left <= right && remaining > 0 {
		var currPackSize = packSizes[right]

		if right == left && remaining < currPackSize {
			packSizeQtyMap[currPackSize] = packSizeQtyMap[currPackSize] + 1
		}

		if remaining >= currPackSize {
			packSizeQtyMap[currPackSize] = packSizeQtyMap[currPackSize] + 1
			remaining = remaining - currPackSize
			if remaining >= currPackSize {
				continue
			} else {
				right--
			}
		} else {
			right--
		}
	}

	if remaining > 0 && len(packSizeQtyMap) == 1 {
		for packSize := range packSizeQtyMap {
			packSizeQtyMap[packSize]++
		}
	}

	// fulfill by pack size
	for packSize, count := range packSizeQtyMap {
		if count == 2 {
			multipleOfKey := packSize * 2

			if slices.Contains(packSizes, multipleOfKey) {
				packSizeQtyMap[multipleOfKey] = 1
				delete(packSizeQtyMap, packSize)
			}
		}
	}
	return packSizeQtyMap
}
