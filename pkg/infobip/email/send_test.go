package email

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendEmailValidReq(t *testing.T) {
	apiKey := "secret"
	msg := models.GenerateEmailMsg()

	content := []byte("temporary file's content")
	attachment, err := ioutil.TempFile("", "example")
	require.NoError(t, err)
	_, err = attachment.Write(content)
	require.NoError(t, err)
	_, err = attachment.Seek(0, 0)
	require.NoError(t, err)

	image, err := os.Open("testdata/image.png")
	require.NoError(t, err)

	msg.Attachment = attachment
	msg.InlineImage = image

	defer image.Close()
	defer attachment.Close()

	rawJSONResp := []byte(`{
	  "bulkId": "esy82u725261jz8e6pi3",
	  "messages": [
		{
		  "to": "joan.doe0@example.com",
		  "messageCount": 1,
		  "messageId": "somexternalMessageId0",
		  "status": {
			"groupId": 1,
			"groupName": "PENDING",
			"id": 7,
			"name": "PENDING_ENROUTE",
			"description": "Message sent to next instance"
		  }
		},
		{
		  "to": "joan.doe1@example.com",
		  "messageCount": 1,
		  "messageId": "somexternalMessageId1",
		  "status": {
			"groupId": 1,
			"groupName": "PENDING",
			"id": 7,
			"name": "PENDING_ENROUTE",
			"description": "Message sent to next instance"
		  }
		},
		{
		  "to": "joan.doe2@example.com",
		  "messageCount": 1,
		  "messageId": "somexternalMessageId2",
		  "status": {
			"groupId": 1,
			"groupName": "PENDING",
			"id": 7,
			"name": "PENDING_ENROUTE",
			"description": "Message sent to next instance"
		  }
		}
	  ]
	}`)

	var expectedResp models.SendEmailResponse
	err = json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendEmailPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		assert.Regexp(t, regexp.MustCompile(`multipart/form-data; boundary=\w+`), r.Header.Get("Content-Type"))

		err = r.ParseMultipartForm(int64(len(content) + 10240))
		require.NoError(t, err)
		assert.Equal(t, msg.From, r.MultipartForm.Value["from"][0])
		assert.Equal(t, msg.To, r.MultipartForm.Value["to"][0])
		assert.Equal(t, msg.Cc, r.MultipartForm.Value["cc"][0])
		assert.Equal(t, msg.Bcc, r.MultipartForm.Value["bcc"][0])
		assert.Equal(t, msg.Subject, r.MultipartForm.Value["subject"][0])
		assert.Equal(t, msg.Text, r.MultipartForm.Value["text"][0])
		assert.Equal(t, msg.BulkID, r.MultipartForm.Value["bulkId"][0])
		assert.Equal(t, msg.MessageID, r.MultipartForm.Value["messageId"][0])
		assert.Equal(t, fmt.Sprintf("%d", msg.TemplateID), r.MultipartForm.Value["templateid"][0])
		assert.Contains(t, attachment.Name(), r.MultipartForm.File["attachment"][0].Filename)
		assert.Equal(t, int64(len(content)), r.MultipartForm.File["attachment"][0].Size)
		assert.Contains(t, image.Name(), r.MultipartForm.File["inlineImage"][0].Filename)
		// TODO: missing size check.
		assert.Equal(t, msg.HTML, r.MultipartForm.Value["HTML"][0])
		assert.Equal(t, msg.ReplyTo, r.MultipartForm.Value["replyto"][0])
		assert.Equal(t, msg.DefaultPlaceholders, r.MultipartForm.Value["defaultplaceholders"][0])
		assert.Equal(t, fmt.Sprintf("%t", msg.PreserveRecipients), r.MultipartForm.Value["preserverecipients"][0])
		assert.Equal(t, msg.TrackingURL, r.MultipartForm.Value["trackingUrl"][0])
		assert.Equal(t, fmt.Sprintf("%t", msg.TrackClicks), r.MultipartForm.Value["trackclicks"][0])
		assert.Equal(t, fmt.Sprintf("%t", msg.TrackOpens), r.MultipartForm.Value["trackopens"][0])
		assert.Equal(t, fmt.Sprintf("%t", msg.Track), r.MultipartForm.Value["track"][0])
		assert.Equal(t, msg.CallbackData, r.MultipartForm.Value["callbackData"][0])
		assert.Equal(t, fmt.Sprintf("%t", msg.IntermediateReport), r.MultipartForm.Value["intermediateReport"][0])
		assert.Equal(t, msg.NotifyURL, r.MultipartForm.Value["notifyUrl"][0])
		assert.Equal(t, msg.NotifyContentType, r.MultipartForm.Value["notifyContentType"][0])
		assert.Equal(t, msg.SendAt, r.MultipartForm.Value["sendAt"][0])
		assert.Equal(t, msg.LandingPagePlaceholders, r.MultipartForm.Value["landingPagePlaceholders"][0])
		assert.Equal(t, msg.LandingPageID, r.MultipartForm.Value["landingPageId"][0])

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := email.Send(context.Background(), msg)

	require.NoError(t, err)
	assert.NotEqual(t, models.SendEmailResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestInvalidEmailMsg(t *testing.T) {
	msg := models.GenerateEmailMsg()
	msg.To = ""

	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    "http://something.com",
		APIKey:     "secret",
	}}

	msgResp, respDetails, err := email.Send(context.Background(), msg)

	require.NotNil(t, err)
	assert.IsType(t, err, validator.ValidationErrors{})
	assert.Equal(t, models.SendEmailResponse{}, msgResp)
	assert.Equal(t, models.ResponseDetails{}, respDetails)
}

func TestSendEmailErrors(t *testing.T) {
	tests := []struct {
		rawJSONResp []byte
		statusCode  int
	}{
		{
			rawJSONResp: []byte(`{
				  "requestError": {
					 "serviceException": {
					   "messageId": "string",
					   "text": "string"
					 }
				  }
			}`),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		var expectedResp models.SendEmailResponse
		err := json.Unmarshal(test.rawJSONResp, &expectedResp)
		require.NoError(t, err)

		serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(test.statusCode)
			_, servErr := w.Write(test.rawJSONResp)
			assert.Nil(t, servErr)
		}))

		email := Channel{ReqHandler: internal.HTTPHandler{
			HTTPClient: http.Client{},
			BaseURL:    serv.URL,
			APIKey:     "secret",
		}}

		msgResp, respDetails, err := email.Send(context.Background(), models.GenerateEmailMsg())
		serv.Close()

		require.NoError(t, err)
		assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
		assert.Equal(t, test.statusCode, respDetails.HTTPResponse.StatusCode)
		assert.Equal(t, expectedResp, msgResp)
	}
}
