package api

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/api/request"
	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/toolbelt"
)

func Create(context *gin.Context) {
	var json request.CreateRequest
	if err := context.BindJSON(&json); err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.Validate(); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	go toolbelt.Create(json)

	context.String(http.StatusAccepted, "")
}

func State(context *gin.Context) {
	podName := context.Param("pod")
	if podName == "" {
		context.String(http.StatusBadRequest, "Pod name is not defined!")
		return
	}

	stateJson, err := toolbelt.State(podName)

	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.String(http.StatusOK, stateJson)
}