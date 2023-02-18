package handler

import (
	"net/http"

	"github.com/Dev-roll/couponroll/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	PostStoreRequest struct {
		Name        string    `json:"name,omitempty" form:"name"`
		Description string    `json:"description,omitempty" form:"description"`
		IsPublic    bool      `json:"is_public,omitempty" form:"is_public"`
		IconUrl     string    `json:"icon_url,omitempty" form:"icon_url"`
		ImgUrl      string    `json:"img_url,omitempty" form:"img_url"`
		CreatorID   uuid.UUID `json:"creator_id,omitempty" form:"creator_id"`
	}
)

func PostStore(c echo.Context) error {
	var req PostStoreRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	store := model.Store{
		Name:        req.Name,
		Description: req.Description,
		IsPublic:    req.IsPublic,
		IconUrl:     req.IconUrl,
		ImgUrl:      req.ImgUrl,
		CreatorID:   req.CreatorID,
	}

	model.CreateStore(store)

	return c.NoContent(http.StatusCreated)
}

func GetStore(c echo.Context) error {
	storeID := c.Param("store_id")
	store := model.GetStoreFromID(model.Store{ID: uuid.MustParse(storeID)})
	return c.JSON(http.StatusOK, store)
}

func PatchStore(c echo.Context) error {
	storeID := c.Param("store_id")

	var req PostStoreRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	store := model.Store{
		Name:        req.Name,
		Description: req.Description,
		IsPublic:    req.IsPublic,
		IconUrl:     req.IconUrl,
		ImgUrl:      req.ImgUrl,
		CreatorID:   req.CreatorID,
	}

	model.UpdateStore(uuid.MustParse(storeID), store)

	return c.NoContent(http.StatusNoContent)
}

func DeleteStore(c echo.Context) error {
	storeID := c.Param("store_id")
	model.DeleteStore(model.Store{ID: uuid.MustParse(storeID)})
	return c.NoContent(http.StatusNoContent)
}
