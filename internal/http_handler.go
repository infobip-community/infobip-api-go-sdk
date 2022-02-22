package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pgrubacc/infobip-go-client/pkg/infobip/models"
)

type HTTPHandler struct {
	APIKey     string
	BaseURL    string
	HTTPClient http.Client
}

// request is a wrapper around the net/http request method, while
// also appending mandatory headers, formatting query parameters
// along with handling and parsing the response status and body.
//
// The body is immediately parsed and closed.
func (h *HTTPHandler) request(
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

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", h.BaseURL, resourcePath), buf)
	if err != nil {
		return nil, nil, err
	}

	req.Header = h.generateHeaders(method)
	req.URL.RawQuery = generateQueryParams(params)

	resp, err = h.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	parsedBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		parsedBody = nil
	}

	return resp, parsedBody, err
}

func (h *HTTPHandler) GetRequest(
	ctx context.Context,
	respResource interface{},
	reqPath string,
) (respDetails models.ResponseDetails, err error) {
	resp, parsedBody, err := h.request( //nolint: bodyclose // closed in the method below
		ctx,
		http.MethodGet,
		reqPath,
		nil,
		nil,
	)
	if err != nil {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
		return respDetails, err
	}
	respDetails.HTTPResponse = *resp

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(parsedBody, &respResource)
	} else {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
	}
	return respDetails, err
}

func (h *HTTPHandler) PostRequest(
	ctx context.Context,
	postResource models.Validatable,
	respResource interface{},
	reqPath string,
) (respDetails models.ResponseDetails, err error) {
	err = postResource.Validate()
	if err != nil {
		return respDetails, err
	}
	resp, parsedBody, err := h.request( //nolint: bodyclose // closed in the method below
		ctx,
		http.MethodPost,
		reqPath,
		postResource,
		nil,
	)
	if err != nil {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
		return respDetails, err
	}
	respDetails.HTTPResponse = *resp

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(parsedBody, &respResource)
	} else {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
	}
	return respDetails, err
}

func (h *HTTPHandler) generateHeaders(method string) http.Header {
	header := http.Header{}

	header.Add("Authorization", fmt.Sprintf("App %s", h.APIKey))
	header.Add("Accept", "application/json")

	if method == http.MethodPost {
		header.Add("Content-Type", "application/json")
	}

	return header
}

func generateQueryParams(params map[string]string) string {
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}

	return q.Encode()
}
