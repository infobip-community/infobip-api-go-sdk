package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type exampleResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestGetReqOK(t *testing.T) {
	rawJSONResp := []byte(`{"id": 1,"name": "John"}`)
	var expectedResp exampleResp
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := exampleResp{}
	respDetails, err := handler.GetRequest(context.Background(), &respResource, "some/path", nil)

	require.NoError(t, err)
	assert.NotEqual(t, exampleResp{}, respResource)
	assert.Equal(t, expectedResp, respResource)
	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestGetReq4xx(t *testing.T) {
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
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusUnauthorized)
		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := exampleResp{}
	respDetails, err := handler.GetRequest(context.Background(), &respResource, "some/path", nil)

	require.NoError(t, err)
	assert.Equal(t, exampleResp{}, respResource)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusUnauthorized, respDetails.HTTPResponse.StatusCode)
	assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
	assert.Equal(t, expectedResp, respDetails.ErrorResponse)
}

func TestGetReqErr(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: "nonexistent"}
	respResource := models.MsgResponse{}
	respDetails, err := handler.GetRequest(context.Background(), &respResource, "some/path", nil)

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.MsgResponse{}, models.MsgResponse{})
}
