package email

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func generateEmailMsg() models.EmailMsg {
	mail := models.EmailMsg{
		From:                    "edcoronag@selfserviceib.com",
		To:                      "edcoronag@gmail.com",
		Cc:                      "",
		Bcc:                     "",
		Subject:                 "Some subject",
		Text:                    "Some text",
		BulkId:                  "",
		MessageId:               "",
		TemplateId:              0,
		Attachment:              nil,
		InlineImage:             nil,
		HTML:                    "",
		ReplyTo:                 "",
		DefaultPlaceholders:     "",
		PreserveRecipients:      false,
		TrackingURL:             "",
		TrackClicks:             false,
		TrackOpens:              false,
		Track:                   false,
		CallbackData:            "",
		IntermediateReport:      false,
		NotifyURL:               "",
		NotifyContentType:       "",
		SendAt:                  "",
		LandingPagePlaceholders: "",
		LandingPageId:           "",
	}

	return mail
}

func TestSendEmailValid(t *testing.T) {
	apiKey := "secret"
	emailMsg := generateEmailMsg()
	rawJSONResp := []byte(`{
    "results": 
[
{

    "bulkId": "string",
    "messageId": "string",
    "to": "string",
    "sentAt": "2022-04-01T17:50:28Z",
    "doneAt": "2022-04-01T17:50:28Z",
    "messageCount": 0,
    "price": 

{

    "pricePerMessage": 0,
    "currency": "string"

},
"status": 
{

    "groupId": 0,
    "groupName": "string",
    "id": 0,
    "name": "string",
    "description": "string",
    "action": "string"

},
"error": 

            {
                "groupId": 0,
                "groupName": "string",
                "id": 0,
                "name": "string",
                "description": "string",
                "permanent": true
            }
        }
    ]

}`)
	var expectedResp models.SendEmailResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.Nil(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendEmailPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedMsg models.EmailMsg
		fmt.Printf("%s\n", parsedBody)
		servErr = json.Unmarshal(parsedBody, &receivedMsg)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedMsg, emailMsg)

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	email := Channel{
		ReqHandler: internal.HTTPHandler{
			HTTPClient: http.Client{},
			BaseURL:    serv.URL,
			APIKey:     apiKey,
		}}

	emailResp, respDetails, err := email.SendFullyFeatured(context.Background(), emailMsg)

	require.Nil(t, err)
	assert.NotEqual(t, models.SendEmailResponse{}, emailResp)
	assert.Equal(t, expectedResp, emailResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestSendEmailValidReq(t *testing.T) {
	apiKey := "secret"
	msg := generateEmailMsg()
	rawJSONResp := []byte(`{
		"bulkId": "1",
		"messages": [
			{
				"to": "41793026727",
				"status": {
					"groupId": 1,
					"groupName": "PENDING",
					"id": 26,
					"name": "PENDING_ACCEPTED",
					"description": "Message accepted, pending for delivery."
				},
				"messageId": "2250be2d4219-3af1-78856-aabe-1362af1edfd2"
			}
		],
		"errorMessage": "string"
	}`)
	var expectedResp models.MMSResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendEmailPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		assert.Regexp(t, regexp.MustCompile(`multipart/form-data; boundary=\w+`), r.Header.Get("Content-Type"))

		require.NoError(t, err)
		assert.Equal(t, string(expectedHead), r.MultipartForm.Value["head"][0])
		assert.Contains(t, tmpFile.Name(), r.MultipartForm.File["media"][0].Filename)
		assert.Equal(t, int64(len(content)), r.MultipartForm.File["media"][0].Size)
		assert.Equal(t, msg.Text, r.MultipartForm.Value["text"][0])
		var expectedExternallyHostedMedia []byte
		expectedExternallyHostedMedia, err = json.Marshal(msg.ExternallyHostedMedia)
		require.NoError(t, err)
		assert.Equal(t, string(expectedExternallyHostedMedia), r.MultipartForm.Value["externallyHostedMedia"][0])
		assert.Equal(t, msg.SMIL, r.MultipartForm.Value["smil"][0])

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	mms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := mms.SendMsg(context.Background(), msg)

	require.NoError(t, err)
	assert.NotEqual(t, models.MMSResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
