package api

import (
	"database/sql"
	"net/http"

	db "github.com/devvyky/logistics/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
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

type listPackSizeParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPackSizes(ctx *gin.Context) {
	var req listPackSizeParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPackSizesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * (req.PageSize),
	}

	packSizes, err := server.store.ListPackSizes(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
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
