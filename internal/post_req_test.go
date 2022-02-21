package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pgrubacc/infobip-go-client/pkg/infobip/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostReqOK(t *testing.T) {
	msg := models.TextMsg{
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
	var expectedResp models.MsgResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.Nil(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedMsg models.TextMsg
		servErr = json.Unmarshal(parsedBody, &receivedMsg)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedMsg, msg)

		w.WriteHeader(http.StatusCreated)
		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := models.MsgResponse{}
	respDetails, err := handler.PostRequest(context.Background(), &msg, &respResource, "some/path")

	require.Nil(t, err)
	assert.NotEqual(t, models.MsgResponse{}, respResource)
	assert.Equal(t, expectedResp, respResource)
	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusCreated, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestPostReq4xx(t *testing.T) {
	msg := models.TextMsg{
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
	require.Nil(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	respResource := models.MsgResponse{}
	respDetails, err := handler.PostRequest(context.Background(), &msg, &respResource, "some/path")

	require.Nil(t, err)
	assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
	assert.NotEqual(t, models.ErrorDetails{}, respDetails.ErrorResponse)
	assert.Equal(t, expectedResp, respDetails.ErrorResponse)
	assert.Equal(t, http.StatusUnauthorized, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.MsgResponse{}, respResource)
}

func TestPostReqErr(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: "nonexistent"}
	msg := models.TextMsg{
		MsgCommon: models.MsgCommon{
			From: "+16175551213",
			To:   "+16175551212",
		},
		Content: models.TextContent{Text: "hello world"},
	}
	respResource := models.MsgResponse{}
	respDetails, err := handler.PostRequest(context.Background(), &msg, &respResource, "some/path")

	require.NotNil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, models.MsgResponse{}, models.MsgResponse{})
}
