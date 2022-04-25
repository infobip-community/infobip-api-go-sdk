package email

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
	apiKey := "apiKey"
	rawJSONResp := []byte(`
		{
			"bulkId": "test-bulk-78",
			"sendAt": 1649872567000
		}
	`)

	var expectedResp models.RescheduleMessagesResponse

	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, rescheduleMessagesPath))
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

	req := models.RescheduleMessagesRequest{
		SendAt: "2022-04-12T17:56:07Z",
	}
	queryParams := models.RescheduleMessagesOpts{}

	resp, respDetails, err := email.RescheduleMessages(context.Background(), req, queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.RescheduleMessagesResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
