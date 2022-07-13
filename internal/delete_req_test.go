package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteReqOK(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		_, servErr := w.Write([]byte(``))
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respDetails, err := handler.DeleteRequest(context.Background(), "some/path", nil)

	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestDeleteReq4xx(t *testing.T) {
	rawJSONResp := []byte(`{
		"requestError": {
			"serviceException": {
				"messageId": "UNAUTHORIZED",
				"text": "Invalid login details"
			}
		}
	}`)
	var expectedResp models.ErrorDetails
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusUnauthorized)
		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respDetails, err := handler.DeleteRequest(context.Background(), "some/path", nil)

	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusUnauthorized, respDetails.HTTPResponse.StatusCode)
	assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
	assert.Equal(t, expectedResp, respDetails.ErrorResponse)
}

func TestDeleteReqErr(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: "nonexistent"}
	respDetails, err := handler.DeleteRequest(context.Background(), "some/path", nil)

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.SendWAMsgResponse{}, models.SendWAMsgResponse{})
}
