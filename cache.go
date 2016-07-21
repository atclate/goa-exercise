package main

import (
	"github.com/atclate/goa-exercise/app"
	"github.com/atclate/goa-exercise/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"fmt"
)

// CacheController implements the cache resource.
type CacheController struct {
	*goa.Controller
}

// NewCacheController creates a cache controller.
func NewCacheController(service *goa.Service) *CacheController {
	return &CacheController{Controller: service.NewController("CacheController")}
}

// Create runs the create action.
func (c *CacheController) Create(ctx *app.CreateCacheContext) error {
	// CacheController_Create: start_implement

	// Put your logic here
	s := models.CacheFromCreateCachePayload(ctx.Payload)
	err := cdb.Add(ctx, s)
	if err != nil {
		return ctx.BadRequest(&goa.Error{Code: "internal_error", Status: 500, Detail: err.Error()})
	}

	// CacheController_Create: end_implement
	return ctx.Created(s.CacheToCache())
}

// Show runs the show action.
func (c *CacheController) Show(ctx *app.ShowCacheContext) error {
	// CacheController_Show: start_implement

	// Put your logic here
	cache, err := cdb.Get(ctx.Context, ctx.ID)
	if err == gorm.ErrRecordNotFound || cache == nil {
		return ctx.NotFound()
	} else if err != nil {
		return c.Service.Send(ctx, 500, err.Error)
	}

	source, err := sdb.Get(ctx.Context, cache.SourceID)
	if err == nil && source != nil {
		cache.Source = *source
	}

	cacheExtended := cache.CacheToCacheExtended()
	cacheExtended.Source = &app.SourceMediaLink{
		ID: source.ID,
		Name: source.Name,
		Href: fmt.Sprintf("/api/sources/%v", source.ID),
	}

	// CacheController_Show: end_implement
	return ctx.OKExtended(cacheExtended)
}

// Update runs the update action.
func (c *CacheController) Update(ctx *app.UpdateCacheContext) error {
	// CacheController_Update: start_implement

	// Put your logic here

	// CacheController_Update: end_implement
	res := &app.Cache{}
	return ctx.OK(res)
}
