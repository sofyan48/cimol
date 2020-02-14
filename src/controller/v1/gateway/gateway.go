package gateway

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	entity "github.com/sofyan48/otp/src/entity/http/v1"
	service "github.com/sofyan48/otp/src/service/v1/gateway"
	"github.com/sofyan48/otp/src/util/helper/rest"
)

// ControllerGateway ...
type ControllerGateway struct {
	ServiceGateway service.GatewayInterface
}

// PostNotification ...
func (ctrl *ControllerGateway) PostNotification(context *gin.Context) {
	// catatan manis untuk infobip jangan lupa untuk mengirim id dari log dynamo ke callback data

	payload := &entity.PostNotificationRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}

	waitgroup := &sync.WaitGroup{}

	ctrl.ServiceGateway.PostNotification(payload, waitgroup)

	result := &entity.PostNotificationResponse{}
	result.ID = payload.UUID
	result.Status = "QUEUE"
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// PostNotificationEmail ...
func (ctrl *ControllerGateway) PostNotificationEmail(context *gin.Context) {
	// catatan manis untuk infobip jangan lupa untuk mengirim id dari log dynamo ke callback data

	payload := &entity.PostNotificationRequestEmail{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	result := &entity.PostNotificationResponse{}
	result.ID = payload.UUID
	result.Status = "QUEUE"
	rest.ResponseData(context, http.StatusOK, payload)
	return
}

// GetHistory ...
func (ctrl *ControllerGateway) GetHistory(context *gin.Context) {
	history, err := ctrl.ServiceGateway.GetHistory(context.Param("receiverAddress"))
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, history)
	return
}

// GetByID ...
func (ctrl *ControllerGateway) GetByID(context *gin.Context) {
	data, err := ctrl.ServiceGateway.GetByID(context.Param("id"))
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, data)
	return
}
