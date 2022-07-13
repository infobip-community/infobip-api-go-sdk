package email

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

func TestUpdateScheduledMessagesStatusValidReq(t *testing.T) {
	apiKey := "apiKey"
	rawJSONResp := []byte(`
		{
			"bulkId": "test-bulk-525",
			"status": "CANCELED"
		}
	`)

	var expectedResp models.UpdateScheduledStatusResponse

	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, updateScheduledMessagesStatusPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	req := models.UpdateScheduledEmailStatusRequest{
		Status: "CANCELED",
	}
	queryParams := models.UpdateScheduledEmailStatusParams{}

	resp, respDetails, err := email.UpdateScheduledMessagesStatus(context.Background(), req, queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.UpdateScheduledStatusResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
