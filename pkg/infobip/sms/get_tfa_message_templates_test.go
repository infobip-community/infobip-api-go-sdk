package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTFAMessageTemplatesValidReq(t *testing.T) {
	rawJSONResp := []byte(`
		[
		  {
			"messageId": "9C815F8AF3328",
			"applicationId": "HJ675435E3A6EA43432G5F37A635KJ8B",
			"pinPlaceholder": "{{pin}}",
			"messageText": "Your PIN is {{pin}}.",
			"pinLength": 4,
			"pinType": "NUMERIC",
			"language": "en",
			"repeatDTMF": "1#",
			"speechRate": 1
		  },
		  {
			"messageId": "8F0792F86035A",
			"applicationId": "HJ675435E3A6EA43432G5F37A635KJ8B",
			"pinPlaceholder": "{{pin}}",
			"messageText": "Your PIN is {{pin}}.",
			"pinLength": 6,
			"pinType": "HEX",
			"repeatDTMF": "1#",
			"speechRate": 1.5
		  }
		]
	`)

	var expectedResp models.GetTFAMessageTemplatesResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "some-api-key"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.Path, getTFAMessageTemplatesPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()
	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := sms.GetTFAMessageTemplates(context.Background(),
		"HJ675435E3A6EA43432G5F37A635KJ8B")

	require.NoError(t, err)
	assert.NotEqual(t, models.GetTFAMessageTemplatesResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
