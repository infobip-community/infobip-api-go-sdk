package infobip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReqValidInput(t *testing.T) {
	type test struct {
		handler     httpHandler
		method      string
		path        string
		body        interface{}
		servResp    string
		httpClient  http.Client
		queryParams map[string]string
	}

	type TestRespBody struct {
		TestField string `json:"testField,omitempty"`
	}

	path := "some/path"
	queryParams := map[string]string{"key": "value", "key2": "value2"}
	servResp := `{"data": "test"}`
	apiKey := "jlf3cdef5b20acc82019482a2ce463cc-9b3d6g39-8af8-4e9b-9206-e76340dc43e5"

	tests := []test{
		{
			path:        path,
			queryParams: queryParams,
			method:      http.MethodGet,
			servResp:    servResp,
		},
		{
			path:        path,
			queryParams: queryParams,
			method:      http.MethodPost,
			servResp:    servResp,
			body:        TestRespBody{TestField: "test"},
		},
		{
			path:        fmt.Sprintf("%s/1", path),
			queryParams: queryParams,
			method:      http.MethodPatch,
			servResp:    servResp,
			body:        TestRespBody{TestField: "test"},
		},
		{
			path:        fmt.Sprintf("%s/1", path),
			queryParams: queryParams,
			method:      http.MethodPut,
			servResp:    servResp,
			body:        TestRespBody{TestField: "test"},
		},
		{
			path:        fmt.Sprintf("%s/1", path),
			queryParams: queryParams,
			method:      http.MethodDelete,
			servResp:    servResp,
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
				assert.Equal(t, apiKey, r.Header.Get("Authorization"))
				expectedBody := []byte{}
				var err error
				if tc.body != nil {
					expectedBody, err = json.Marshal(tc.body)
					assert.NotNil(t, err)
				}
				parsedBody, err := ioutil.ReadAll(r.Body)
				assert.Nil(t, err)
				assert.Equal(t, expectedBody, parsedBody)

				w.WriteHeader(http.StatusOK)
				_, err = w.Write([]byte(tc.servResp))
				assert.Nil(t, err)
			}))
			defer serv.Close()

			host := serv.URL
			tc.handler.baseURL = host
			tc.handler.httpClient = tc.httpClient
			tc.handler.apiKey = apiKey

			resp, body, err := tc.handler.request(tc.method, tc.path, tc.body, tc.queryParams)

			assert.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, []byte(tc.servResp), body)
		})
	}
}

func TestReqInvalidMethod(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer serv.Close()

	handler := httpHandler{httpClient: http.Client{}, baseURL: serv.URL}
	resp, _, err := handler.request("ČĆŽŽ", "some/path", nil, nil)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid method")
	assert.Nil(t, resp)
}

func TestReqInvalidResBody(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer serv.Close()

	handler := httpHandler{httpClient: http.Client{}, baseURL: serv.URL}
	resp, _, err := handler.request(http.MethodGet, "some/path", nil, nil)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected EOF")
	assert.NotNil(t, resp)
}

func TestReqInvalidPayload(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer serv.Close()

	handler := httpHandler{httpClient: http.Client{}, baseURL: serv.URL}
	resp, _, err := handler.request(http.MethodPost, "some/path", math.Inf(1), nil)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "json: unsupported value")
	assert.Nil(t, resp)
}

func TestReqInvalidHost(t *testing.T) {
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer serv.Close()

	handler := httpHandler{httpClient: http.Client{}, baseURL: "nonexistent"}
	resp, _, err := handler.request(http.MethodGet, "some/path", nil, nil)

	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGenerateHeaders(t *testing.T) {
	type test struct {
		method              string
		expectedContentType string
	}

	apiKey := "jlf3cdef5b20acc82019482a2ce463cc-9b3d6g39-8af8-4e9b-9206-e76340dc43e5"
	handler := httpHandler{apiKey: apiKey}
	tests := []test{{method: "GET"}, {method: "POST", expectedContentType: "application/json"}}
	for _, tc := range tests {
		t.Run(tc.method, func(t *testing.T) {
			headers := handler.generateHeaders(tc.method)
			assert.NotNil(t, headers)
			assert.Equal(t, apiKey, headers.Get("Authorization"))
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
