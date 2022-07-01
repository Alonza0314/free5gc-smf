/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pdusession

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/smf/internal/logger"
	"github.com/free5gc/smf/internal/sbi/producer"
	"github.com/free5gc/util/httpwrapper"
)

// HTTPReleaseSmContext - Release SM Context
func HTTPReleaseSmContext(c *gin.Context) {
	logger.PduSessLog.Info("Receive Release SM Context Request")
	var request models.ReleaseSmContextRequest
	request.JsonData = new(models.SmContextReleaseData)

	s := strings.Split(c.GetHeader("Content-Type"), ";")
	var err error
	switch s[0] {
	case "application/json":
		err = c.ShouldBindJSON(request.JsonData)
	case "multipart/related":
		err = c.ShouldBindWith(&request, openapi.MultipartRelatedBinding{})
	}
	if err != nil {
		log.Print(err)
		return
	}

	req := httpwrapper.NewRequest(c.Request, request)
	req.Params["smContextRef"] = c.Params.ByName("smContextRef")

	smContextRef := req.Params["smContextRef"]
	HTTPResponse := producer.HandlePDUSessionSMContextRelease(
		smContextRef, req.Body.(models.ReleaseSmContextRequest))

	if HTTPResponse.Status < 300 {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	}
}

// RetrieveSmContext - Retrieve SM Context
func RetrieveSmContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HTTPUpdateSmContext - Update SM Context
func HTTPUpdateSmContext(c *gin.Context) {
	logger.PduSessLog.Info("Receive Update SM Context Request")
	var request models.UpdateSmContextRequest
	request.JsonData = new(models.SmContextUpdateData)

	s := strings.Split(c.GetHeader("Content-Type"), ";")
	var err error
	switch s[0] {
	case "application/json":
		err = c.ShouldBindJSON(request.JsonData)
	case "multipart/related":
		err = c.ShouldBindWith(&request, openapi.MultipartRelatedBinding{})
	}
	if err != nil {
		log.Print(err)
		return
	}

	req := httpwrapper.NewRequest(c.Request, request)
	req.Params["smContextRef"] = c.Params.ByName("smContextRef")

	smContextRef := req.Params["smContextRef"]
	HTTPResponse := producer.HandlePDUSessionSMContextUpdate(
		smContextRef, req.Body.(models.UpdateSmContextRequest))

	if HTTPResponse.Status < 300 {
		c.Render(HTTPResponse.Status, openapi.MultipartRelatedRender{Data: HTTPResponse.Body})
	} else {
		c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	}
}
