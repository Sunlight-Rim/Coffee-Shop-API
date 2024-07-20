package tools

import (
	"net/http"

	"coffeeshop-api/pkg/errors"

	"github.com/gdexlab/go-render/render"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
)

var failedResponse = []byte(`{"response": null, "error": {"message": "failed encode response", "code": 0}}`)

// General response struct.
// easyjson:json
type Response struct {
	Response any            `json:"response"`
	Error    *ErrorResponse `json:"error"`
}

// General error response struct.
// easyjson:json
type ErrorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// SendResponse sends a response or error to the client.
func SendResponse(c echo.Context, res any, err error) {
	if errRes := c.JSONBlob(fmtResponse(res, err)); errRes != nil {
		logrus.Errorf("Send response: %v. Response was: %v", errRes, render.AsCode(res))
	}
}

// fmtResponse provides response formatting.
func fmtResponse(data any, err error) (int, []byte) {
	var res []byte

	if err != nil {
		// Match error
		status, code, message := errors.GetHTTPErrData(err)

		if res, err = easyjson.Marshal(Response{
			Error: &ErrorResponse{Code: code, Message: message},
		}); err != nil {
			logrus.Errorf("Marshal error response: %v. Error message was: %v", err, message)
			return http.StatusInternalServerError, failedResponse
		}

		return status, res
	}

	if res, err = easyjson.Marshal(Response{
		Response: data,
	}); err != nil {
		logrus.Errorf("Marshal response: %v. Response data was: %#v", err, render.AsCode(data))
		return http.StatusInternalServerError, failedResponse
	}

	return http.StatusOK, res
}
