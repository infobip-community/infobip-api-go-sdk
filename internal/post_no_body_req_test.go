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

func TestPostNoBodyReqOK(t *testing.T) {
	rawJSONResp := []byte(`{
		"to": "441134960001",
		"messageCount": 1,
		"messageId": "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
		"status": {
			"groupId": 1,
			"groupName": "PENDING",
			"id": 7,
			"name": "PENDING_ENROUTE",
			"description": "Message sent to next instance"
		}
	}`)
	var expectedResp models.SendWAMsgResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.ContentLength, int64(0))
		w.WriteHeader(http.StatusAccepted)
		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := models.SendWAMsgResponse{}
	respDetails, err := handler.PostNoBodyReq(context.Background(), &respResource, "some/path")

	require.NoError(t, err)
	assert.NotEqual(t, models.SendWAMsgResponse{}, respResource)
	assert.Equal(t, expectedResp, respResource)
	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusAccepted, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestNoBodyPostReq4xx(t *testing.T) {
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
		w.WriteHeader(http.StatusUnauthorized)
		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := models.SendWAMsgResponse{}
	respDetails, err := handler.PostNoBodyReq(context.Background(), &respResource, "some/path")

	require.NoError(t, err)
	assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
	assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
	assert.Equal(t, expectedResp, respDetails.ErrorResponse)
	assert.Equal(t, http.StatusUnauthorized, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.SendWAMsgResponse{}, respResource)
}

func TestPostNoBodyReqErr(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: "nonexistent"}
	respResource := models.SendWAMsgResponse{}
	respDetails, err := handler.PostNoBodyReq(context.Background(), &respResource, "some/path")

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.SendWAMsgResponse{}, models.SendWAMsgResponse{})
}
