package tools

import (
	"net/http"

	"coffeeshop-api/pkg/errors"

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
	if errResp := c.JSONBlob(mapResponse(res, err)); errResp != nil {
		logrus.Errorf("Send response err: %v. Response was: %v", err, res)
	}
}

// mapResponse provides response mapping/matching.
func mapResponse(data any, err error) (int, []byte) {
	var res []byte

	if err != nil {
		// Match error
		errStatus, code, message := errors.GetHTTPErrData(err)

		if res, err = easyjson.Marshal(Response{Error: &ErrorResponse{Code: code, Message: message}}); err != nil {
			return http.StatusInternalServerError, failedResponse
		}

		return errStatus, res
	}

	if res, err = easyjson.Marshal(Response{Response: data}); err != nil {
		return http.StatusInternalServerError, failedResponse
	}

	return http.StatusOK, res
}
