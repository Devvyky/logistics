package api

import (
	"database/sql"
	"net/http"

	db "github.com/devvyky/logistics/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createPackSizeParams struct {
	ProductLine string `json:"product_line" binding:"required"`
	PackSize    int64  `json:"pack_size" binding:"required"`
}

func (server *Server) createPackSize(ctx *gin.Context) {
	var req createPackSizeParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePackSizeParams{
		ProductLine: req.ProductLine,
		PackSize:    req.PackSize,
	}

	packSize, err := server.store.CreatePackSize(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, packSize)
}

type getPackSizeParams struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getPackSize(ctx *gin.Context) {
	var req getPackSizeParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	packSize, err := server.store.GetPackSize(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, packSize)
}

type listPackSizeByProductLinesParams struct {
	ProductLine string `form:"product_line" binding:"required"`
}

func (server *Server) listPackSizeByProductLines(ctx *gin.Context) {
	var req listPackSizeByProductLinesParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	packSizes, err := listPackSizeByProductLines(ctx, req, server)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, packSizes)
}

type updatePackSizeParams struct {
	ProductLine string `json:"product_line" binding:"required"`
	PackSize    int64  `json:"pack_size" binding:"required"`
}

func (server *Server) updatePackSize(ctx *gin.Context) {
	var body updatePackSizeParams
	var param getPackSizeParams

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := uuid.Parse(param.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePackSizesParams{
		ID:          id,
		ProductLine: body.ProductLine,
		PackSize:    body.PackSize,
	}
	packSize, err := server.store.UpdatePackSizes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, packSize)
}

func (server *Server) deletePackSize(ctx *gin.Context) {
	var req getPackSizeParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeletePackSize(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (server *Server) listProductLines(ctx *gin.Context) {
	productLines, err := server.store.ListProductLines(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, productLines)
}

func listPackSizeByProductLines(
	ctx *gin.Context,
	req listPackSizeByProductLinesParams,
	server *Server) ([]db.ProductPackSize, error) {
	packSizes, err := server.store.ListPackSizesByProductLine(ctx, req.ProductLine)
	return packSizes, err
}
