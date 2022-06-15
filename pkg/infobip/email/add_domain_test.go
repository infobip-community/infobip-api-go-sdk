package email

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/internal"
	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDomainValidReq(t *testing.T) {
	apiKey := "some-key"
	req := models.AddEmailDomainRequest{DomainName: "some-domain.com"}
	rawJSONResp := []byte(`
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
	`)

	var expectedResp models.AddEmailDomainResponse
	err := json.Unmarshal(rawJSONResp, &expectedResp)
	require.NoError(t, err)

	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.HasSuffix(r.URL.Path, addDomainPath))
		assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
		parsedBody, servErr := ioutil.ReadAll(r.Body)
		assert.Nil(t, servErr)

		var receivedReq models.AddEmailDomainRequest
		servErr = json.Unmarshal(parsedBody, &receivedReq)
		assert.Nil(t, servErr)
		assert.Equal(t, receivedReq, req)

		_, servErr = w.Write(rawJSONResp)
		assert.Nil(t, servErr)
	}))
	defer serv.Close()

	email := Channel{ReqHandler: internal.HTTPHandler{
		HTTPClient: http.Client{},
		BaseURL:    serv.URL,
		APIKey:     apiKey,
	}}

	msgResp, respDetails, err := email.AddDomain(context.Background(), req)

	require.NoError(t, err)
	assert.NotEqual(t, models.AddEmailDomainResponse{}, msgResp)
	assert.Equal(t, expectedResp, msgResp)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
	assert.Equal(t, models.ErrorDetails{}, respDetails.ErrorResponse)
}
