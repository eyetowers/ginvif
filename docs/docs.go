// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ondrej Sery",
            "url": "https://eyetowers.cz",
            "email": "ondrej.sery@eyetowers.cz"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/eyetowers/gonvif/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ptz/get-nodes": {
            "post": {
                "description": "A PTZ-capable device may have multiple PTZ Nodes. The PTZ Nodes may represent mechanical PTZ drivers, uploaded PTZ drivers or digital PTZ drivers. PTZ Nodes are the lowest level entities in the PTZ control API and reflect the supported PTZ capabilities. The PTZ Node is referenced either by its name or by its reference token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ptz"
                ],
                "summary": "Get the descriptions of the available PTZ Nodes.",
                "parameters": [
                    {
                        "description": "JSON request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wsdl.GetNodes"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wsdl.GetNodesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            }
        },
        "/ptz/get-presets": {
            "post": {
                "description": "The operation is supported if there is support for at least on PTZ preset by the PTZNode.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ptz"
                ],
                "summary": "Operation to request all PTZ presets for the PTZNode in the selected profile.",
                "parameters": [
                    {
                        "description": "JSON request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wsdl.GetPresets"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wsdl.GetPresetsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/util.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.FloatRange": {
            "type": "object",
            "properties": {
                "Max": {
                    "type": "number"
                },
                "Min": {
                    "type": "number"
                }
            }
        },
        "schema.PTZNode": {
            "type": "object",
            "properties": {
                "AuxiliaryCommands": {
                    "description": "A list of supported Auxiliary commands. If the list is not empty, the Auxiliary Operations MUST be available for this PTZ Node.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "Extension": {
                    "$ref": "#/definitions/schema.PTZNodeExtension"
                },
                "FixedHomePosition": {
                    "type": "boolean"
                },
                "GeoMove": {
                    "type": "boolean"
                },
                "HomeSupported": {
                    "description": "A boolean operator specifying the availability of a home position. If set to true, the Home Position Operations MUST be available for this PTZ Node.",
                    "type": "boolean"
                },
                "MaximumNumberOfPresets": {
                    "description": "All preset operations MUST be available for this PTZ Node if one preset is supported.",
                    "type": "integer"
                },
                "Name": {
                    "description": "A unique identifier that is used to reference PTZ Nodes.",
                    "type": "string"
                },
                "SupportedPTZSpaces": {
                    "description": "A list of Coordinate Systems available for the PTZ Node. For each Coordinate System, the PTZ Node MUST specify its allowed range.",
                    "$ref": "#/definitions/schema.PTZSpaces"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "schema.PTZNodeExtension": {
            "type": "object",
            "properties": {
                "Extension": {
                    "$ref": "#/definitions/schema.PTZNodeExtension2"
                },
                "SupportedPresetTour": {
                    "description": "Detail of supported Preset Tour feature.",
                    "$ref": "#/definitions/schema.PTZPresetTourSupported"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.PTZNodeExtension2": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.PTZPreset": {
            "type": "object",
            "properties": {
                "Name": {
                    "description": "A list of preset position name.",
                    "type": "string"
                },
                "PTZPosition": {
                    "description": "A list of preset position.",
                    "$ref": "#/definitions/schema.PTZVector"
                },
                "token": {
                    "description": "A reference to the MediaProfile where the operation should take place.",
                    "type": "string"
                }
            }
        },
        "schema.PTZPresetTourSupported": {
            "type": "object",
            "properties": {
                "Extension": {
                    "$ref": "#/definitions/schema.PTZPresetTourSupportedExtension"
                },
                "MaximumNumberOfPresetTours": {
                    "description": "Indicates number of preset tours that can be created. Required preset tour operations shall be available for this PTZ Node if one or more preset tour is supported.",
                    "type": "integer"
                },
                "PTZPresetTourOperation": {
                    "description": "Indicates which preset tour operations are available for this PTZ Node.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.PTZPresetTourSupportedExtension": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.PTZSpaces": {
            "type": "object",
            "properties": {
                "AbsolutePanTiltPositionSpace": {
                    "description": "The Generic Pan/Tilt Position space is provided by every PTZ node that supports absolute Pan/Tilt, since it does not relate to a specific physical range.\nInstead, the range should be defined as the full range of the PTZ unit normalized to the range -1 to 1 resulting in the following space description.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space2DDescription"
                    }
                },
                "AbsoluteZoomPositionSpace": {
                    "description": "The Generic Zoom Position Space is provided by every PTZ node that supports absolute Zoom, since it does not relate to a specific physical range.\nInstead, the range should be defined as the full range of the Zoom normalized to the range 0 (wide) to 1 (tele).\nThere is no assumption about how the generic zoom range is mapped to magnification, FOV or other physical zoom dimension.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space1DDescription"
                    }
                },
                "ContinuousPanTiltVelocitySpace": {
                    "description": "The generic Pan/Tilt velocity space shall be provided by every PTZ node, since it does not relate to a specific physical range.\nInstead, the range should be defined as a range of the PTZ unit’s speed normalized to the range -1 to 1, where a positive velocity would map to clockwise\nrotation or movement in the right/up direction. A signed speed can be independently specified for the pan and tilt component resulting in the following space description.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space2DDescription"
                    }
                },
                "ContinuousZoomVelocitySpace": {
                    "description": "The generic zoom velocity space specifies a zoom factor velocity without knowing the underlying physical model. The range should be normalized from -1 to 1,\nwhere a positive velocity would map to TELE direction. A generic zoom velocity space description resembles the following.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space1DDescription"
                    }
                },
                "Extension": {
                    "$ref": "#/definitions/schema.PTZSpacesExtension"
                },
                "PanTiltSpeedSpace": {
                    "description": "The speed space specifies the speed for a Pan/Tilt movement when moving to an absolute position or to a relative translation.\nIn contrast to the velocity spaces, speed spaces do not contain any directional information. The speed of a combined Pan/Tilt\nmovement is represented by a single non-negative scalar value.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space1DDescription"
                    }
                },
                "RelativePanTiltTranslationSpace": {
                    "description": "The Generic Pan/Tilt translation space is provided by every PTZ node that supports relative Pan/Tilt, since it does not relate to a specific physical range.\nInstead, the range should be defined as the full positive and negative translation range of the PTZ unit normalized to the range -1 to 1,\nwhere positive translation would mean clockwise rotation or movement in right/up direction resulting in the following space description.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space2DDescription"
                    }
                },
                "RelativeZoomTranslationSpace": {
                    "description": "The Generic Zoom Translation Space is provided by every PTZ node that supports relative Zoom, since it does not relate to a specific physical range.\nInstead, the corresponding absolute range should be defined as the full positive and negative translation range of the Zoom normalized to the range -1 to1,\nwhere a positive translation maps to a movement in TELE direction. The translation is signed to indicate direction (negative is to wide, positive is to tele).\nThere is no assumption about how the generic zoom range is mapped to magnification, FOV or other physical zoom dimension. This results in the following space description.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space1DDescription"
                    }
                },
                "ZoomSpeedSpace": {
                    "description": "The speed space specifies the speed for a Zoom movement when moving to an absolute position or to a relative translation.\nIn contrast to the velocity spaces, speed spaces do not contain any directional information.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Space1DDescription"
                    }
                }
            }
        },
        "schema.PTZSpacesExtension": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schema.PTZVector": {
            "type": "object",
            "properties": {
                "PanTilt": {
                    "description": "Pan and tilt position. The x component corresponds to pan and the y component to tilt.",
                    "$ref": "#/definitions/schema.Vector2D"
                },
                "Zoom": {
                    "description": "A zoom position.",
                    "$ref": "#/definitions/schema.Vector1D"
                }
            }
        },
        "schema.Space1DDescription": {
            "type": "object",
            "properties": {
                "URI": {
                    "description": "A URI of coordinate systems.",
                    "type": "string"
                },
                "XRange": {
                    "description": "A range of x-axis.",
                    "$ref": "#/definitions/schema.FloatRange"
                }
            }
        },
        "schema.Space2DDescription": {
            "type": "object",
            "properties": {
                "URI": {
                    "description": "A URI of coordinate systems.",
                    "type": "string"
                },
                "XRange": {
                    "description": "A range of x-axis.",
                    "$ref": "#/definitions/schema.FloatRange"
                },
                "YRange": {
                    "description": "A range of y-axis.",
                    "$ref": "#/definitions/schema.FloatRange"
                }
            }
        },
        "schema.Vector1D": {
            "type": "object",
            "properties": {
                "space": {
                    "type": "string"
                },
                "x": {
                    "type": "number"
                }
            }
        },
        "schema.Vector2D": {
            "type": "object",
            "properties": {
                "space": {
                    "type": "string"
                },
                "x": {
                    "type": "number"
                },
                "y": {
                    "type": "number"
                }
            }
        },
        "util.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "wsdl.GetNodes": {
            "type": "object"
        },
        "wsdl.GetNodesResponse": {
            "type": "object",
            "properties": {
                "PTZNode": {
                    "description": "A list of the existing PTZ Nodes on the device.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.PTZNode"
                    }
                }
            }
        },
        "wsdl.GetPresets": {
            "type": "object",
            "properties": {
                "ProfileToken": {
                    "description": "A reference to the MediaProfile where the operation should take place.",
                    "type": "string"
                }
            }
        },
        "wsdl.GetPresetsResponse": {
            "type": "object",
            "properties": {
                "Preset": {
                    "description": "A list of presets which are available for the requested MediaProfile.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.PTZPreset"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Onvif JSON API Adapter",
	Description:      "HTTP Server translating JSON requests into Onvif (onvif.org) WS/SOAP protocol.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
