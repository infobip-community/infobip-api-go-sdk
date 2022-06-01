package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRescheduleMessagesValidReq(t *testing.T) {
	rawJSONResp := []byte(`
		{
			"bulkId": "test-bulk-73",
			"sendAt": "2021-11-25T16:00:00.000+0000"
		}
	`)

	var expectedResp models.RescheduleSMSResponse

	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	apiKey := "apiKey"
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, rescheduleSMSPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	sms := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	req := models.RescheduleSMSRequest{
		SendAt: "2022-04-12T17:56:07Z",
	}
	queryParams := models.RescheduleSMSParams{}

	resp, respDetails, err := sms.RescheduleMessages(context.Background(), req, queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.RescheduleEmailResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
