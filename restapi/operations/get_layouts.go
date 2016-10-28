package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLayoutsHandlerFunc turns a function with the right signature into a get layouts handler
type GetLayoutsHandlerFunc func(context.Context, GetLayoutsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLayoutsHandlerFunc) Handle(ctx context.Context, params GetLayoutsParams) middleware.Responder {
	return fn(ctx, params)
}

// GetLayoutsHandler interface for that can handle valid get layouts params
type GetLayoutsHandler interface {
	Handle(context.Context, GetLayoutsParams) middleware.Responder
}

// NewGetLayouts creates a new http.Handler for the get layouts operation
func NewGetLayouts(ctx *middleware.Context, handler GetLayoutsHandler) *GetLayouts {
	return &GetLayouts{Context: ctx, Handler: handler}
}

/*GetLayouts swagger:route GET /layouts getLayouts

Pre-configured layouts

Layouts are a collection of `Cells` that visualize time-series data.


*/
type GetLayouts struct {
	Context *middleware.Context
	Handler GetLayoutsHandler
}

func (o *GetLayouts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetLayoutsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
