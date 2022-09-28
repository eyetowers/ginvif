package ptz

import (
	ptz "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
	"github.com/gin-gonic/gin"

	"github.com/eyetowers/ginvif/pkg/util"
)

func RegisterHandlers(base *gin.RouterGroup, client gonvif.Client) {
	g := base.Group("/ptz")
	g.POST("/get-nodes", withPTZ(client, GetNodes))
	g.POST("/get-presets", withPTZ(client, GetPresets))
}

// @Summary     Get the descriptions of the available PTZ Nodes.
// @Description A PTZ-capable device may have multiple PTZ Nodes. The PTZ Nodes may represent mechanical PTZ drivers, uploaded PTZ drivers or digital PTZ drivers. PTZ Nodes are the lowest level entities in the PTZ control API and reflect the supported PTZ capabilities. The PTZ Node is referenced either by its name or by its reference token.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GetNodes true "JSON request body"
// @Success     200     {object} ptz.GetNodesResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/get-nodes [post]
func GetNodes(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetNodesContext)
}

// @Summary     Operation to request all PTZ presets for the PTZNode in the selected profile.
// @Description The operation is supported if there is support for at least on PTZ preset by the PTZNode.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GetPresets true "JSON request body"
// @Success     200     {object} ptz.GetPresetsResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/get-presets [post]
func GetPresets(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetPresetsContext)
}
