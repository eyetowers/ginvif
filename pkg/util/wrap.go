package util

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProcessOnvifRequest parses the JSON request body, calls the given Onvif method, and returns the
// response as JSON again.
func ProcessOnvifRequest[REQ any, RESP any, FUNC func(context.Context, *REQ) (*RESP, error)](
	ctx *gin.Context, call FUNC,
) {
	var req REQ
	err := ctx.BindJSON(&req)
	if err != nil {
		FailWithError(ctx, http.StatusBadRequest, err)
		return
	}
	resp, err := call(ctx, &req)
	if err != nil {
		FailWithError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, resp)
}

func FailWithError(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, &Error{
		Message: err.Error(),
		Code:    code,
	})
}
