package httpresputils

import (
	"mini-crm-billing-api/source/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	Status     string    `json:"status"`
	AppName    string    `json:"app_name"`
	Message    *string   `json:"message,omitempty"`
	Time       time.Time `json:"timestamp"`
	AppVersion string    `json:"app_version"`
	Opt        any       `json:"options,omitempty"`
	Data       any       `json:"data,omitempty"`
}

var (
	cfg = config.Load()
)

func httpResp(c *gin.Context, statusCode int, data, options any, msg *string) {
	resp := response{
		Status:     http.StatusText(statusCode),
		Time:       time.Now(),
		AppVersion: cfg.AppVersion,
		Message:    msg,
		AppName:    cfg.AppName,
		Data:       data,
		Opt:        options,
	}

	c.JSON(statusCode, resp)
}

// HTTP Success Responses
func HttpRespOK(c *gin.Context, data, options any, msg *string) {
	httpResp(c, http.StatusOK, data, options, msg)
}

func HttpRespCreated(c *gin.Context, data any, msg *string) {
	httpResp(c, http.StatusCreated, data, nil, msg)
}

func HttpRespNoContent(c *gin.Context, msg *string) {
	httpResp(c, http.StatusNoContent, nil, nil, msg)
}

func HttpRespAccepted(c *gin.Context, msg *string) {
	httpResp(c, http.StatusAccepted, nil, nil, msg)
}

// HTTP Error Responses
func HttpRespNotFound(c *gin.Context, msg *string) {
	httpResp(c, http.StatusNotFound, nil, nil, msg)
}

func HttpRespBadRequest(c *gin.Context, msg *string) {
	httpResp(c, http.StatusBadRequest, nil, nil, msg)
}

func HttpRespBadGateway(c *gin.Context, msg *string) {
	httpResp(c, http.StatusBadGateway, nil, nil, msg)
}

func HttpResponseUnAuth(c *gin.Context, msg *string) {
	httpResp(c, http.StatusUnauthorized, nil, nil, msg)
}

func HttpResponseForbidden(c *gin.Context, msg *string) {
	httpResp(c, http.StatusForbidden, nil, nil, msg)
}
