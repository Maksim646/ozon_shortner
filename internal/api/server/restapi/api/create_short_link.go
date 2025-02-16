// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateShortLinkHandlerFunc turns a function with the right signature into a create short link handler
type CreateShortLinkHandlerFunc func(CreateShortLinkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateShortLinkHandlerFunc) Handle(params CreateShortLinkParams) middleware.Responder {
	return fn(params)
}

// CreateShortLinkHandler interface for that can handle valid create short link params
type CreateShortLinkHandler interface {
	Handle(CreateShortLinkParams) middleware.Responder
}

// NewCreateShortLink creates a new http.Handler for the create short link operation
func NewCreateShortLink(ctx *middleware.Context, handler CreateShortLinkHandler) *CreateShortLink {
	return &CreateShortLink{Context: ctx, Handler: handler}
}

/*
	CreateShortLink swagger:route POST /shortner_link Link createShortLink

Create Short Link
*/
type CreateShortLink struct {
	Context *middleware.Context
	Handler CreateShortLinkHandler
}

func (o *CreateShortLink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateShortLinkParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
