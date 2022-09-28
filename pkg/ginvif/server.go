package ginvif

import (
	"github.com/eyetowers/gonvif/pkg/gonvif"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/eyetowers/ginvif/docs"
	"github.com/eyetowers/ginvif/pkg/api/ptz"
)

// @title       Onvif JSON API Adapter
// @version     1.0
// @description HTTP Server translating JSON requests into Onvif (onvif.org) WS/SOAP protocol.

// @contact.name  Ondrej Sery
// @contact.url   https://eyetowers.cz
// @contact.email ondrej.sery@eyetowers.cz

// @license.name MIT
// @license.url  https://github.com/eyetowers/gonvif/blob/master/LICENSE

// @BasePath /api/v1
func Serve(bind string, client gonvif.Client) error {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		ptz.RegisterHandlers(v1, client)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	return r.Run(bind)
}
