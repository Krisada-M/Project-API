package helper

import (
	"Restapi/config"
	"Restapi/models"
	"net/http"
)

// D is pretty res
type D map[string]interface{}

// APIResponse is response RestApi
func (data D) APIResponse() models.Response {
	var resp models.Response

	resp.Response.Datetime = config.Date

	resp.Status.Code = http.StatusOK
	resp.Status.Description = "Success"
	resp.Status.Message = ""

	resp.Data = data

	return resp
}
