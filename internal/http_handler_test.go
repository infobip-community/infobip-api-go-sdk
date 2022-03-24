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
	path := "some/path"
	servResp := `{"data": "test"}`
	apiKey := "secret"

	tests := []struct {
		handler    HTTPHandler
		method     string
		path       string
		body       *TestRespBody
		servResp   string
		httpClient http.Client
	}{
		{
			path:     path,
			method:   http.MethodGet,
			servResp: servResp,
			body:     nil,
		},
		{
			path:     path,
			method:   http.MethodPost,
			servResp: servResp,
			body:     &TestRespBody{StringField: "test"},
		},
		{
			path:     fmt.Sprintf("%s/1", path),
			method:   http.MethodPatch,
			servResp: servResp,
			body:     &TestRespBody{StringField: "test"},
		},
		{
			path:     fmt.Sprintf("%s/1", path),
			method:   http.MethodPut,
			servResp: servResp,
			body:     &TestRespBody{StringField: "test"},
		},
		{
			path:     fmt.Sprintf("%s/1", path),
			method:   http.MethodDelete,
			servResp: servResp,
			body:     nil,
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

			req, err := tc.handler.createReq(context.Background(), tc.method, tc.path, payloadBuf)
			require.NoError(t, err)

			resp, body, err := tc.handler.executeReq(req)
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
	req, err := handler.createReq(ctx, http.MethodGet, "some/path", nil)
	require.NoError(t, err)

	resp, _, err := handler.executeReq(req)
	require.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestReqInvalidMethod(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	_, err := handler.createReq(context.Background(), "ČĆŽŽ", "some/path", nil)
	require.Error(t, err)
}

func TestReqInvalidResBody(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer serv.Close()

	handler := HTTPHandler{HTTPClient: http.Client{}, BaseURL: serv.URL}
	req, err := handler.createReq(context.Background(), http.MethodGet, "some/path", nil)
	require.NoError(t, err)
	resp, _, err := handler.executeReq(req)

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
	req, err := handler.createReq(context.Background(), http.MethodGet, "some/path", nil)
	require.NoError(t, err)

	resp, _, err := handler.executeReq(req)
	require.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGenerateQueryParams(t *testing.T) {
	tests := []struct {
		scenario string
		params   map[string]string
		expected string
	}{
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
