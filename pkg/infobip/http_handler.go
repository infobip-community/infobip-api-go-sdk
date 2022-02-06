package infobip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// httpHandler provides methods for handling http requests.
type httpHandler struct {
	apiKey     string
	baseURL    string
	httpClient http.Client
}

// request is a wrapper around the net/http request method, while
// also appending mandatory headers, formatting query parameters
// along with handling and parsing the response status and body.
//
// The body is immediately parsed and closed and
// the output is returned to the caller.
func (h *httpHandler) request(
	ctx context.Context,
	method string,
	resourcePath string,
	body interface{},
	params map[string]string,
) (resp *http.Response, respBody []byte, err error) {
	var buf io.Reader

	if body != nil {
		var parsedBody []byte
		parsedBody, err = json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
		buf = bytes.NewBuffer(parsedBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", h.baseURL, resourcePath), buf)
	if err != nil {
		return nil, nil, err
	}

	req.Header = h.generateHeaders(method)
	req.URL.RawQuery = generateQueryParams(params)

	resp, err = h.httpClient.Do(req)
	if err != nil {
		return resp, nil, err
	}
	defer resp.Body.Close()

	parsedBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		parsedBody = nil
	}

	return resp, parsedBody, err
}

// generateHeaders returns a http.Header object depending on the passed method.
// The headers follow the specification in the API docs.
// Common headers that http.Client automatically generates, e.g. "Host", are omitted.

// https://api-docs.form3.tech/api.html#introduction-and-api-conventions-headers
func (h *httpHandler) generateHeaders(method string) http.Header {
	header := http.Header{}

	header.Add("Authorization", h.apiKey)

	if method == http.MethodPost {
		header.Add("Content-Type", "application/json")
	}

	return header
}

// generateQueryParams parses the map of query parameters and returns
// them in the URL encoded format.
func generateQueryParams(params map[string]string) string {
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}

	return q.Encode()
}
