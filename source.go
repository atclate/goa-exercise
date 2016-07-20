package main

import (
	"github.com/atclate/goa-exercise/app"
	"github.com/atclate/goa-exercise/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// SourceController implements the source resource.
type SourceController struct {
	*goa.Controller
}

// NewSourceController creates a source controller.
func NewSourceController(service *goa.Service) *SourceController {
	return &SourceController{Controller: service.NewController("SourceController")}
}

// Caches runs the caches action.
func (c *SourceController) Caches(ctx *app.CachesSourceContext) error {
	// SourceController_Caches: start_implement

	// Put your logic here
	source, err := sdb.OneSourceMedia(ctx.Context, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return c.Service.Send(ctx, 500, err.Error)
	}

	// SourceController_Show: end_implement
	return ctx.OK(source.Caches)
}

// Create runs the create action.
func (c *SourceController) Create(ctx *app.CreateSourceContext) error {
	// SourceController_Create: start_implement

	// Put your logic here
	s := models.SourceFromCreateSourcePayload(ctx.Payload)
	sdb.Add(ctx, s)

	// SourceController_Create: end_implement
	return ctx.Created(s.SourceToSourceMedia())
}

// Delete runs the delete action.
func (c *SourceController) Delete(ctx *app.DeleteSourceContext) error {
	// SourceController_Delete: start_implement

	// Put your logic here
	err := sdb.Delete(ctx, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return c.Service.Send(ctx, 500, err.Error)
	}

	// SourceController_Delete: end_implement
	return ctx.OK(nil)
}

// List runs the list action.
func (c *SourceController) List(ctx *app.ListSourceContext) error {
	// SourceController_List: start_implement

	// Put your logic here
	source := sdb.ListSourceMedia(ctx.Context)

	// SourceController_Show: end_implement
	return ctx.OK(source)
}

// Show runs the show action.
func (c *SourceController) Show(ctx *app.ShowSourceContext) error {
	// SourceController_Show: start_implement

	// Put your logic here
	source, err := sdb.OneSourceMedia(ctx.Context, ctx.ID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return c.Service.Send(ctx, 500, err.Error)
	}

	// SourceController_Show: end_implement
	return ctx.OK(source)
}

// Update runs the update action.
func (c *SourceController) Update(ctx *app.UpdateSourceContext) error {
	// SourceController_Update: start_implement

	// Put your logic here

	// SourceController_Update: end_implement
	res := &app.SourceMedia{}
	return ctx.OK(res)
}
