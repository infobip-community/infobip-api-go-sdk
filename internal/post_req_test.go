package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostReqOK(t *testing.T) {
	msg := models.WATextMsg{
		MsgCommon: models.GenerateTestMsgCommon(),
		Content:   models.TextContent{Text: "hello world"},
	}
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
		assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedMsg models.WATextMsg
		servErr = json.Unmarshal(parsedBody, &receivedMsg)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedMsg, msg)

		w.WriteHeader(http.StatusCreated)
		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := models.SendWAMsgResponse{}
	respDetails, err := handler.PostJSONReq(context.Background(), &msg, &respResource, "some/path")

	require.NoError(t, err)
	assert.NotEqual(t, models.SendWAMsgResponse{}, respResource)
	assert.Equal(t, expectedResp, respResource)
	require.NoError(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusCreated, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestPostReq4xx(t *testing.T) {
	msg := models.WATextMsg{
		MsgCommon: models.GenerateTestMsgCommon(),
		Content:   models.TextContent{Text: "hello world"},
	}
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
	respDetails, err := handler.PostJSONReq(context.Background(), &msg, &respResource, "some/path")

	require.NoError(t, err)
	assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
	assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
	assert.Equal(t, expectedResp, respDetails.ErrorResponse)
	assert.Equal(t, http.StatusUnauthorized, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.SendWAMsgResponse{}, respResource)
}

func TestPostReqErr(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: "nonexistent"}
	msg := models.WATextMsg{
		MsgCommon: models.MsgCommon{
			From: "+16175551213",
			To:   "+16175551212",
		},
		Content: models.TextContent{Text: "hello world"},
	}
	respResource := models.SendWAMsgResponse{}
	respDetails, err := handler.PostJSONReq(context.Background(), &msg, &respResource, "some/path")

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.SendWAMsgResponse{}, models.SendWAMsgResponse{})
}

func TestPostInvalidPayload(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	msg := InvalidTestMsg{FloatField: math.Inf(1)}
	respResource := models.SendWAMsgResponse{}
	respDetails, err := handler.PostJSONReq(context.Background(), &msg, &respResource, "some/path")

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.SendWAMsgResponse{}, models.SendWAMsgResponse{})
}
