package ptz

import (
	ptz "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
	"github.com/gin-gonic/gin"

	"github.com/eyetowers/ginvif/pkg/util"
)

func RegisterHandlers(base *gin.RouterGroup, client gonvif.Client) {
	g := base.Group("/ptz")
	g.POST("/ptz/get-service-capabilities", withPTZ(client, GetServiceCapabilities))
	g.POST("/get-nodes", withPTZ(client, GetNodes))
	g.POST("/get-node", withPTZ(client, GetNode))
	g.POST("/get-configuration", withPTZ(client, GetConfiguration))
	g.POST("/get-configurations", withPTZ(client, GetConfigurations))
	g.POST("/set-configuration", withPTZ(client, SetConfiguration))
	g.POST("/get-configurations-options", withPTZ(client, GetConfigurationOptions))
	g.POST("/send-auxiliary-command", withPTZ(client, SendAuxiliaryCommand))
	g.POST("/get-presets", withPTZ(client, GetPresets))
	g.POST("/set-preset", withPTZ(client, SetPreset))
	g.POST("/remove-preset", withPTZ(client, RemovePreset))
	g.POST("/goto-preset", withPTZ(client, GotoPreset))
	g.POST("/goto-home-position", withPTZ(client, GotoHomePosition))
	g.POST("/set-home-position", withPTZ(client, SetHomePosition))
	g.POST("/continuous-move", withPTZ(client, ContinuousMove))
	g.POST("/relative-move", withPTZ(client, RelativeMove))
	g.POST("/get-status", withPTZ(client, GetStatus))
	g.POST("/absolute-move", withPTZ(client, AbsoluteMove))
	g.POST("/geo-move", withPTZ(client, GeoMove))
	g.POST("/stop", withPTZ(client, Stop))
	g.POST("/get-preset-tours", withPTZ(client, GetPresetTours))
	g.POST("/get-preset-tour", withPTZ(client, GetPresetTour))
	g.POST("/get-preset-tour-options", withPTZ(client, GetPresetTourOptions))
	g.POST("/create-preset-tour", withPTZ(client, CreatePresetTour))
	g.POST("/modify-preset-tour", withPTZ(client, ModifyPresetTour))
	g.POST("/operate-preset-tour", withPTZ(client, OperatePresetTour))
	g.POST("/remove-preset-tour", withPTZ(client, RemovePresetTour))
	g.POST("/get-compatible-configurations", withPTZ(client, GetCompatibleConfigurations))
	g.POST("/move-and-start-tracking", withPTZ(client, MoveAndStartTracking))
}

// @Summary Returns the capabilities of the PTZ service. The result is returned in a typed answer.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.GetServiceCapabilities true "JSON request body"
// @Success 200     {object} ptz.GetServiceCapabilitiesResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/get-service-capabilities [post]
func GetServiceCapabilities(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetServiceCapabilitiesContext)
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

// @Summary Get a specific PTZ Node identified by a reference token or a name.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.GetNode true "JSON request body"
// @Success 200     {object} ptz.GetNodeResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/get-node [post]
func GetNode(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetNodeContext)
}

// @Summary     Get a specific PTZconfiguration from the device, identified by its reference token or name.
// @Description The default Position/Translation/Velocity Spaces are introduced to allow NVCs sending move requests without the need to specify a certain coordinate system. The default Speeds are introduced to control the speed of move requests (absolute, relative, preset), where no explicit speed has been set. The allowed pan and tilt range for Pan/Tilt Limits is defined by a two-dimensional space range that is mapped to a specific Absolute Pan/Tilt Position Space. At least one Pan/Tilt Position Space is required by the PTZNode to support Pan/Tilt limits. The limits apply to all supported absolute, relative and continuous Pan/Tilt movements. The limits shall be checked within the coordinate system for which the limits have been specified. That means that even if movements are specified in a different coordinate system, the requested movements shall be transformed to the coordinate system of the limits where the limits can be checked. When a relative or continuous movements is specified, which would leave the specified limits, the PTZ unit has to move along the specified limits. The Zoom Limits have to be interpreted accordingly.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GetConfiguration true "JSON request body"
// @Success     200     {object} ptz.GetConfigurationResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/get-configuration [post]
func GetConfiguration(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetConfigurationContext)
}

// @Summary     Get all the existing PTZConfigurations from the device.
// @Description The default Position/Translation/Velocity Spaces are introduced to allow NVCs sending move requests without the need to specify a certain coordinate system. The default Speeds are introduced to control the speed of move requests (absolute, relative, preset), where no explicit speed has been set. The allowed pan and tilt range for Pan/Tilt Limits is defined by a two-dimensional space range that is mapped to a specific Absolute Pan/Tilt Position Space. At least one Pan/Tilt Position Space is required by the PTZNode to support Pan/Tilt limits. The limits apply to all supported absolute, relative and continuous Pan/Tilt movements. The limits shall be checked within the coordinate system for which the limits have been specified. That means that even if movements are specified in a different coordinate system, the requested movements shall be transformed to the coordinate system of the limits where the limits can be checked. When a relative or continuous movements is specified, which would leave the specified limits, the PTZ unit has to move along the specified limits. The Zoom Limits have to be interpreted accordingly.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GetConfigurations true "JSON request body"
// @Success     200     {object} ptz.GetConfigurationsResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/get-configurations [post]
func GetConfigurations(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetConfigurationsContext)
}

// @Summary Set/update a existing PTZConfiguration on the device.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.SetConfiguration true "JSON request body"
// @Success 200     {object} ptz.SetConfigurationResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/set-configuration [post]
func SetConfiguration(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.SetConfigurationContext)
}

// @Summary     List supported coordinate systems including their range limitations.
// @Description Therefore, the options MAY differ depending on whether the PTZ Configuration is assigned to a Profile containing a Video Source Configuration. In that case, the options may additionally contain coordinate systems referring to the image coordinate system described by the Video Source Configuration. If the PTZ Node supports continuous movements, it shall return a Timeout Range within which Timeouts are accepted by the PTZ Node.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GetConfigurationOptions true "JSON request body"
// @Success     200     {object} ptz.GetConfigurationOptionsResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/get-configurations-options [post]
func GetConfigurationOptions(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetConfigurationOptionsContext)
}

// @Summary     Operation to send auxiliary commands to the PTZ device mapped by the PTZNode in the selected profile.
// @Description The operation is supported if the AuxiliarySupported element of the PTZNode is true
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.SendAuxiliaryCommand true "JSON request body"
// @Success     200     {object} ptz.SendAuxiliaryCommandResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/send-auxiliary-command [post]
func SendAuxiliaryCommand(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.SendAuxiliaryCommandContext)
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

// @Summary     The SetPreset command saves the current device position parameters so that the device can move to the saved preset position through the GotoPreset operation.
// @Description In order to create a new preset, the SetPresetRequest contains no PresetToken. If creation is successful, the Response contains the PresetToken which uniquely identifies the Preset. An existing Preset can be overwritten by specifying the PresetToken of the corresponding Preset. In both cases (overwriting or creation) an optional PresetName can be specified. The operation fails if the PTZ device is moving during the SetPreset operation. The device MAY internally save additional states such as imaging properties in the PTZ Preset which then should be recalled in the GotoPreset operation. The operation is supported if there is support for at least on PTZ preset by the PTZNode.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.SetPreset true "JSON request body"
// @Success     200     {object} ptz.SetPresetResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/set-preset [post]
func SetPreset(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.SetPresetContext)
}

// @Summary     Operation to remove a PTZ preset for the Node in the selected profile.
// @Description The operation is supported if the PresetPosition capability exists for teh Node in the selected profile.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.RemovePreset true "JSON request body"
// @Success     200     {object} ptz.RemovePresetResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/remove-preset [post]
func RemovePreset(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.RemovePresetContext)
}

// @Summary     Operation to go to a saved preset position for the PTZNode in the selected profile.
// @Description The operation is supported if there is support for at least on PTZ preset by the PTZNode.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GotoPreset true "JSON request body"
// @Success     200     {object} ptz.GotoPresetResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/goto-preset [post]
func GotoPreset(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GotoPresetContext)
}

// @Summary     Operation to move the PTZ device to it's "home" position.
// @Description The operation is supported if the HomeSupported element in the PTZNode is true.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GotoHomePosition true "JSON request body"
// @Success     200     {object} ptz.GotoHomePositionResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/goto-home-position [post]
func GotoHomePosition(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GotoHomePositionContext)
}

// @Summary     Operation to save current position as the home position.
// @Description The SetHomePosition command returns with a failure if the “home” position is fixed and cannot be overwritten. If the SetHomePosition is successful, it is possible to recall the Home Position with the GotoHomePosition command.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.SetHomePosition true "JSON request body"
// @Success     200     {object} ptz.SetHomePositionResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/set-home-position [post]
func SetHomePosition(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.SetHomePositionContext)
}

// @Summary     Operation for continuous Pan/Tilt and Zoom movements.
// @Description The operation is supported if the PTZNode supports at least one continuous Pan/Tilt or Zoom space. If the space argument is omitted, the default space set by the PTZConfiguration will be used.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.ContinuousMove true "JSON request body"
// @Success     200     {object} ptz.ContinuousMoveResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/continuous-move [post]
func ContinuousMove(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.ContinuousMoveContext)
}

// @Summary     Operation for Relative Pan/Tilt and Zoom Move.
// @Description The operation is supported if the PTZNode supports at least one relative Pan/Tilt or Zoom space. The speed argument is optional. If an x/y speed value is given it is up to the device to either use the x value as absolute resoluting speed vector or to map x and y to the component speed. If the speed argument is omitted, the default speed set by the PTZConfiguration will be used.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.RelativeMove true "JSON request body"
// @Success     200     {object} ptz.RelativeMoveResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/relative-move [post]
func RelativeMove(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.RelativeMoveContext)
}

// @Summary Operation to request PTZ status for the Node in the selected profile.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.GetStatus true "JSON request body"
// @Success 200     {object} ptz.GetStatusResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/get-status [post]
func GetStatus(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetStatusContext)
}

// @Summary     Operation to move pan,tilt or zoom to a absolute destination.
// @Description The speed argument is optional. If an x/y speed value is given it is up to the device to either use the x value as absolute resoluting speed vector or to map x and y to the component speed. If the speed argument is omitted, the default speed set by the PTZConfiguration will be used.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.AbsoluteMove true "JSON request body"
// @Success     200     {object} ptz.AbsoluteMoveResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/absolute-move [post]
func AbsoluteMove(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.AbsoluteMoveContext)
}

// @Summary     Operation to move pan,tilt or zoom to point to a destination based on the geolocation of the target.
// @Description The speed argument is optional. If an x/y speed value is given it is up to the device to either use the x value as absolute resoluting speed vector or to map x and y to the component speed. If the speed argument is omitted, the default speed set by the PTZConfiguration will be used. The area height and area dwidth parameters are optional, they can be used independently and may be used by the device to automatically determine the best zoom level to show the target.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GeoMove true "JSON request body"
// @Success     200     {object} ptz.GeoMoveResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/geo-move [post]
func GeoMove(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GeoMoveContext)
}

// @Summary     Operation to stop ongoing pan, tilt and zoom movements of absolute relative and continuous type.
// @Description If no stop argument for pan, tilt or zoom is set, the device will stop all ongoing pan, tilt and zoom movements.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.Stop true "JSON request body"
// @Success     200     {object} ptz.StopResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/stop [post]
func Stop(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.StopContext)
}

// @Summary Operation to request PTZ preset tours in the selected media profiles.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.GetPresetTours true "JSON request body"
// @Success 200     {object} ptz.GetPresetToursResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/get-preset-tours [post]
func GetPresetTours(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetPresetToursContext)
}

// @Summary Operation to request a specific PTZ preset tour in the selected media profile.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.GetPresetTour true "JSON request body"
// @Success 200     {object} ptz.GetPresetTourResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/get-preset-tour [post]
func GetPresetTour(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetPresetTourContext)
}

// @Summary Operation to request available options to configure PTZ preset tour.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.GetPresetTourOptions true "JSON request body"
// @Success 200     {object} ptz.GetPresetTourOptionsResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/get-preset-tour-options [post]
func GetPresetTourOptions(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetPresetTourOptionsContext)
}

// @Summary Operation to create a preset tour for the selected media profile.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.CreatePresetTour true "JSON request body"
// @Success 200     {object} ptz.CreatePresetTourResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/create-preset-tour [post]
func CreatePresetTour(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.CreatePresetTourContext)
}

// @Summary Operation to modify a preset tour for the selected media profile.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.ModifyPresetTour true "JSON request body"
// @Success 200     {object} ptz.ModifyPresetTourResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/modify-preset-tour [post]
func ModifyPresetTour(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.ModifyPresetTourContext)
}

// @Summary Operation to perform specific operation on the preset tour in selected media profile.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.OperatePresetTour true "JSON request body"
// @Success 200     {object} ptz.OperatePresetTourResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/operate-preset-tour [post]
func OperatePresetTour(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.OperatePresetTourContext)
}

// @Summary Operation to delete a specific preset tour from the media profile.
// @Tags    ptz
// @Accept  json
// @Produce json
// @Param   request body     ptz.RemovePresetTour true "JSON request body"
// @Success 200     {object} ptz.RemovePresetTourResponse
// @Failure 400     {object} util.Error
// @Failure 500     {object} util.Error
// @Failure 501     {object} util.Error
// @Router  /ptz/remove-preset-tour [post]
func RemovePresetTour(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.RemovePresetTourContext)
}

// @Summary     Operation to get all available PTZConfigurations that can be added to the referenced media profile.
// @Description A device providing more than one PTZConfiguration or more than one VideoSourceConfiguration or which has any other resource interdependency between PTZConfiguration entities and other resources listable in a media profile should implement this operation. PTZConfiguration entities returned by this operation shall not fail on adding them to the referenced media profile.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.GetCompatibleConfigurations true "JSON request body"
// @Success     200     {object} ptz.GetCompatibleConfigurationsResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/get-compatible-configurations [post]
func GetCompatibleConfigurations(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.GetCompatibleConfigurationsContext)
}

// @Summary     Operation to send an an atomic command to the device: move the camera to a wanted position and then delegate the PTZ control to the tracking algorithm.
// @Description An existing Speed argument overrides DefaultSpeed of the corresponding PTZ configuration during movement to the requested position. If spaces are referenced within the Speed argument, they shall be speed spaces supported by the PTZ node. If the detection and the tracking are done in the same device, an ObjectID reference can be passed as an argument, in order to specify which object should be tracked. The operation shall fail if the requested absolute position is not reachable.
// @Tags        ptz
// @Accept      json
// @Produce     json
// @Param       request body     ptz.MoveAndStartTracking true "JSON request body"
// @Success     200     {object} ptz.MoveAndStartTrackingResponse
// @Failure     400     {object} util.Error
// @Failure     500     {object} util.Error
// @Failure     501     {object} util.Error
// @Router      /ptz/move-and-start-tracking [post]
func MoveAndStartTracking(ctx *gin.Context, client ptz.PTZ) {
	util.ProcessOnvifRequest(ctx, client.MoveAndStartTrackingContext)
}
