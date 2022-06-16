package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v2/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutReqOK(t *testing.T) {
	req := models.UpdateScheduledSMSStatusRequest{
		Status: "PAUSED",
	}
	rawJSONResp := []byte(`
		{
			"bulkId": "test-bulk-73",
			"status": "PAUSED"
		}
	`)
	var expectedResp models.UpdateScheduledSMSStatusResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.UpdateScheduledSMSStatusRequest
		servErr = json.Unmarshal(parsedBody, &receivedReq)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedReq, req)

		w.WriteHeader(http.StatusOK)
		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := models.UpdateScheduledSMSStatusResponse{}
	respDetails, err := handler.PutJSONReq(
		context.Background(), &req, &respResource, "some/path", []QueryParameter{})

	require.NoError(t, err)
	assert.NotEqual(t, models.UpdateScheduledSMSStatusResponse{}, respResource)
	assert.Equal(t, expectedResp, respResource)
	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestPutReq4xx(t *testing.T) {
	req := models.RescheduleSMSRequest{
		SendAt: "2020-01-01T00:00:00Z",
	}
	rawJSONResp := []byte(`
		{
		  "requestError": {
			"serviceException": {
			  "messageId": "string",
			  "text": "string"
			}
		  }
		}
	`)
	var expectedResp models.ErrorDetails
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := models.RescheduleSMSResponse{}
	respDetails, err := handler.PutJSONReq(
		context.Background(), &req, &respResource, "some/path", []QueryParameter{})

	require.NoError(t, err)
	assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
	assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
	assert.Equal(t, expectedResp, respDetails.ErrorResponse)
	assert.Equal(t, http.StatusBadRequest, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.RescheduleSMSResponse{}, respResource)
}

func TestPutReqErr(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: "nonexistent"}
	req := models.RescheduleSMSRequest{}
	respResource := models.RescheduleSMSResponse{}
	respDetails, err := handler.PutJSONReq(
		context.Background(), &req, &respResource, "some/path", []QueryParameter{})

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.RescheduleSMSResponse{}, respResource)
}

type InvalidTestMsg struct {
	FloatField float64 `json:"floatField"`
}

func (t *InvalidTestMsg) Validate() error {
	return nil
}

func (t *InvalidTestMsg) Marshal() (*bytes.Buffer, error) {
	payload, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(payload), nil
}

func TestPutInvalidPayload(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	msg := InvalidTestMsg{FloatField: math.Inf(1)}
	respResource := models.RescheduleSMSResponse{}
	respDetails, err := handler.PutJSONReq(
		context.Background(), &msg, &respResource, "some/path", []QueryParameter{})

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.RescheduleSMSResponse{}, respResource)
}
