package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestRespBody struct {
	StringField string  `json:"testField,omitempty"`
	FloatField  float64 `json:"floatField,omitempty"`
}

func (t *TestRespBody) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func TestReqValidInput(t *testing.T) {
	type test struct {
		handler     HTTPHandler
		method      string
		path        string
		body        *TestRespBody
		servResp    string
		httpClient  http.Client
		queryParams map[string]string
	}

	path := "some/path"
	queryParams := map[string]string{"key": "value", "key2": "value2"}
	servResp := `{"data": "test"}`
	apiKey := "secret"

	tests := []test{
		{
			path:        path,
			queryParams: queryParams,
			method:      http.MethodGet,
			servResp:    servResp,
			body:        nil,
		},
		{
			path:        path,
			queryParams: queryParams,
			method:      http.MethodPost,
			servResp:    servResp,
			body:        &TestRespBody{StringField: "test"},
		},
		{
			path:        fmt.Sprintf("%s/1", path),
			queryParams: queryParams,
			method:      http.MethodPatch,
			servResp:    servResp,
			body:        &TestRespBody{StringField: "test"},
		},
		{
			path:        fmt.Sprintf("%s/1", path),
			queryParams: queryParams,
			method:      http.MethodPut,
			servResp:    servResp,
			body:        &TestRespBody{StringField: "test"},
		},
		{
			path:        fmt.Sprintf("%s/1", path),
			queryParams: queryParams,
			method:      http.MethodDelete,
			servResp:    servResp,
			body:        nil,
		},
		{
			path:   path,
			method: http.MethodHead,
		},
	}

	for _, tc := range tests {
		t.Run(tc.method, func(t *testing.T) {
			serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tc.path, r.URL.Path[1:])
				assert.Equal(t, tc.method, r.Method)
				assert.Equal(t, fmt.Sprintf("App %s", apiKey), r.Header.Get("Authorization"))
				expectedBody := []byte{}
				var err error
				if tc.body != nil {
					expectedBody, err = json.Marshal(tc.body)
					require.Nil(t, err)
				}
				parsedBody, err := ioutil.ReadAll(r.Body)
				require.Nil(t, err)
				assert.Equal(t, expectedBody, parsedBody)

				w.WriteHeader(http.StatusOK)
				_, err = w.Write([]byte(tc.servResp))
				require.Nil(t, err)
			}))
			defer serv.Close()

			host := serv.URL
			tc.handler.BaseURL = host
			tc.handler.HTTPClient = tc.httpClient
			tc.handler.APIKey = apiKey

			var payloadBuf io.Reader
			if tc.body != nil {
				payload, err := json.Marshal(tc.body)
				require.Nil(t, err)
				payloadBuf = bytes.NewBuffer(payload)
			}

			resp, body, err := tc.handler.request(
				context.Background(),
				tc.method,
				tc.path,
				payloadBuf,
				tc.queryParams,
			)

			require.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, []byte(tc.servResp), body)
		})
	}
}

func TestReqContext(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		cancel()
	}()
	resp, _, err := handler.request(ctx, http.MethodGet, "some/path", nil, nil)

	require.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestReqInvalidMethod(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	resp, _, err := handler.request(context.Background(), "ČĆŽŽ", "some/path", nil, nil)

	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid method")
	assert.Nil(t, resp)
}

func TestReqInvalidResBody(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	resp, _, err := handler.request(context.Background(), http.MethodGet, "some/path", nil, nil)

	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected EOF")
	assert.NotNil(t, resp)
}

func TestReqInvalidHost(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: "nonexistent"}
	resp, _, err := handler.request(context.Background(), http.MethodGet, "some/path", nil, nil)

	require.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGenerateHeaders(t *testing.T) {
	type test struct {
		method              string
		expectedContentType string
	}

	apiKey := "secret"
	handler := HTTPHandler{APIKey: apiKey}
	tests := []test{{method: "GET"}, {method: "POST", expectedContentType: "application/json"}}
	for _, tc := range tests {
		t.Run(tc.method, func(t *testing.T) {
			headers := handler.generateHeaders(tc.method)
			assert.NotNil(t, headers)
			assert.Equal(t, fmt.Sprintf("App %s", apiKey), headers.Get("Authorization"))
			assert.Equal(t, tc.expectedContentType, headers.Get("Content-Type"))
		})
	}
}

func TestGenerateQueryParams(t *testing.T) {
	type test struct {
		scenario string
		params   map[string]string
		expected string
	}

	tests := []test{
		{
			scenario: "params passed",
			params:   map[string]string{"key1": "value1", "key2": "value2"},
			expected: "key1=value1&key2=value2",
		},
		{
			scenario: "empty params",
		},
	}

	for _, tc := range tests {
		t.Run(tc.scenario, func(t *testing.T) {
			queryParams := generateQueryParams(tc.params)
			assert.NotNil(t, queryParams)
			assert.Equal(t, tc.expected, queryParams)
		})
	}
}
