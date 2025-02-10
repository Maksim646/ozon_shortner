// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetOriginalLinkHandlerFunc turns a function with the right signature into a get original link handler
type GetOriginalLinkHandlerFunc func(GetOriginalLinkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOriginalLinkHandlerFunc) Handle(params GetOriginalLinkParams) middleware.Responder {
	return fn(params)
}

// GetOriginalLinkHandler interface for that can handle valid get original link params
type GetOriginalLinkHandler interface {
	Handle(GetOriginalLinkParams) middleware.Responder
}

// NewGetOriginalLink creates a new http.Handler for the get original link operation
func NewGetOriginalLink(ctx *middleware.Context, handler GetOriginalLinkHandler) *GetOriginalLink {
	return &GetOriginalLink{Context: ctx, Handler: handler}
}

/*
	GetOriginalLink swagger:route GET /original_link Link getOriginalLink

Get Original Link
*/
type GetOriginalLink struct {
	Context *middleware.Context
	Handler GetOriginalLinkHandler
}

func (o *GetOriginalLink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetOriginalLinkParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
