// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"work6/microservice_test/go-zero_test/rest/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search/:name",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: CreateUserHandler(serverCtx),
			},
		},
	)
}