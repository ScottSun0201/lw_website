package client

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ClientUserApi
	ClientSEOApi
	ClientAboutApi
}

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
