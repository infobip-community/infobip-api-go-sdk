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

func TestGetDomainsValidReq(t *testing.T) {
	apiKey := "apiKey"
	rawJSONResp := []byte(`
		{
		  "paging": {
			"page": 0,
			"size": 0,
			"totalPages": 0,
			"totalResults": 0
		  },
		  "results": [
			{
			  "domainId": 1,
			  "domainName": "newDomain.com",
			  "active": false,
			  "tracking": {
				"clicks": true,
				"opens": true,
				"unsubscribe": true
			  },
			  "dnsRecords": [
				{
				  "recordType": "string",
				  "name": "string",
				  "expectedValue": "string",
				  "verified": true
				}
			  ],
			  "blocked": false,
			  "createdAt": "2022-05-05T17:32:28.777+01:00"
			}
		  ]
		}
	`)

	var expectedResp models.GetEmailDomainsResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		assert.True(t, strings.HasSuffix(r.URL.Path, getDomainsPath))
		assert.Equal(t, fmt.Sprint("App ", apiKey), r.Header.Get("Authorization"))

		_, servErr := w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	queryParams := models.GetEmailDomainsParams{
		Size: 10,
		Page: 0,
	}

	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	resp, respDetails, err := email.GetDomains(context.Background(), queryParams)

	require.NoError(t, err)
	assert.NotEqual(t, models.GetEmailDomainsResponse{}, resp)
	assert.Equal(t, expectedResp, resp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
