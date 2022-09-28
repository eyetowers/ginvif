package ptz

import (
	"net/http"

	ptz "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
	"github.com/gin-gonic/gin"

	"github.com/eyetowers/ginvif/pkg/util"
)

func withPTZ(
	client gonvif.Client, handler func(*gin.Context, ptz.PTZ),
) func(*gin.Context) {
	return func(ctx *gin.Context) {
		p, err := client.PTZ()
		if err != nil {
			util.FailWithError(ctx, http.StatusNotImplemented, err)
			return
		}
		handler(ctx, p)
	}
}
