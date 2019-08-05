package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	controller "path/to/service/proto/controller"
)

type controllerKey struct {}

// FromContext retrieves the client from the Context
func ControllerFromContext(ctx context.Context) (controller.ControllerService, bool) {
	c, ok := ctx.Value(controllerKey{}).(controller.ControllerService)
	return c, ok
}

// Client returns a wrapper for the ControllerClient
func ControllerWrapper(service micro.Service) server.HandlerWrapper {
	client := controller.NewControllerService("go.micro.srv.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, controllerKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
