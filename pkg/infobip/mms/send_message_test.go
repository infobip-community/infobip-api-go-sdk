package mms

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendMsgValidReq(t *testing.T) {
	content := []byte("temporary file's content")
	tmpFile, err := ioutil.TempFile("", "example")
	require.NoError(t, err)
	_, err = tmpFile.Write(content)
	require.NoError(t, err)
	_, err = tmpFile.Seek(0, 0)
	require.NoError(t, err)

	apiKey := "secret"
	msg := models.MMSMsg{
		Head: models.MMSHead{
			From:                  "16175551213",
			To:                    "16175551212",
			ID:                    "26",
			Subject:               "This is a sample message",
			ValidityPeriodMinutes: 10,
			CallbackData:          "data",
			NotifyURL:             "https://www.google.com",
			SendAt:                "2006-01-02T15:04:05.123Z",
			IntermediateReport:    true,
			DeliveryTimeWindow: &models.DeliveryTimeWindow{
				Days: []string{"MONDAY", "TUESDAY"},
				From: &models.MMSTime{Minute: 1, Hour: 1},
				To:   &models.MMSTime{Minute: 1, Hour: 2},
			},
		},
		Media: tmpFile,
		Text:  "Some text",
		ExternallyHostedMedia: []models.ExternallyHostedMedia{
			{
				ContentType: "image/jpeg",
				ContentID:   "1",
				ContentURL:  "https://myurl.com/asd.jpg",
			},
		},
		SMIL: "<smil></smil>",
	}
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
	err = json.Unmarshal(rawJSONResp, &expectedResp)
	require.Nil(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, sendMessagePath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		assert.Regexp(t, regexp.MustCompile(`multipart/form-data; boundary=\w+`), r.Header.Get("Content-Type"))

		err = r.ParseMultipartForm(int64(len(content)))
		require.NoError(t, err)
		var expectedHead []byte
		expectedHead, err = json.Marshal(msg.Head)
		require.Nil(t, err)
		assert.Equal(t, string(expectedHead), r.MultipartForm.Value["head"][0])
		assert.Contains(t, tmpFile.Name(), r.MultipartForm.File["media"][0].Filename)
		assert.Equal(t, int64(len(content)), r.MultipartForm.File["media"][0].Size)
		assert.Equal(t, msg.Text, r.MultipartForm.Value["text"][0])
		var expectedExternallyHostedMedia []byte
		expectedExternallyHostedMedia, err = json.Marshal(msg.ExternallyHostedMedia)
		require.Nil(t, err)
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

	require.Nil(t, err)
	assert.NotEqual(t, models.MMSResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}

func TestInvalidMsg(t *testing.T) {
	mms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    "https://something.api.infobip.com",
		APIKey:     "secret",
	}}
	msgResp, respDetails, err := mms.SendMsg(context.Background(), models.MMSMsg{Text: "Some text"})

	require.NotNil(t, err)
	assert.IsType(t, err, validator.ValidationErrors{})
	assert.Equal(t, models.MMSResponse{}, msgResp)
	assert.Equal(t, models.ResponseDetails{}, respDetails)
}

func TestSendMsgErrors(t *testing.T) {
	tests := []struct {
		rawJSONResp []byte
		statusCode  int
	}{
		{
			rawJSONResp: []byte(`{
				"bulkId": "",
				"messages": [ ],
				"errorMessage": "Head part is mandatory. Check API documentation"
			}`),
			statusCode: http.StatusBadRequest,
		},
		{
			rawJSONResp: []byte(`{
				"bulkId": "",
				"messages": [ ],
				"errorMessage": "Internal error"
			}`),
			statusCode: http.StatusInternalServerError,
		},
	}
	msg := models.MMSMsg{
		Head: models.MMSHead{From: "16175551213", To: "16175551212"},
		Text: "Some text",
	}

	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.statusCode), func(t *testing.T) {
			var expectedResp models.MMSResponse
			err := json.Unmarshal(tc.rawJSONResp, &expectedResp)
			require.Nil(t, err)
			serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				_, servErr := w.Write(tc.rawJSONResp)
				assert.Nil(t, servErr)
			}))
			mms := Channel{ReqHandler: internal.HTTPHandler{
				HTTPClient: http.Client{},
				BaseURL:    serv.URL,
				APIKey:     "secret",
			}}

			msgResp, respDetails, err := mms.SendMsg(context.Background(), msg)
			serv.Close()

			require.Nil(t, err)
			assert.NotEqual(t, http.Response{}, respDetails.HTTPResponse)
			assert.Equal(t, tc.statusCode, respDetails.HTTPResponse.StatusCode)
			assert.Equal(t, expectedResp, msgResp)
		})
	}
}
