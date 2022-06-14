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

	"github.com/infobip-community/infobip-api-go-sdk/pkg/infobip/models"
)

type HTTPHandler struct {
	APIKey     string
	BaseURL    string
	HTTPClient http.Client
}

type QueryParameter struct {
	Name  string
	Value string
}

func (h *HTTPHandler) createReq(
	ctx context.Context,
	method string,
	resourcePath string,
	body io.Reader,
	queryParams []QueryParameter,
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", h.BaseURL, resourcePath), body)
	if err != nil {
		return nil, err
	}
	req.Header = h.generateCommonHeaders()
	req.URL.RawQuery = generateQueryParams(queryParams)
	return req, nil
}

func (h *HTTPHandler) executeReq(
	req *http.Request,
) (resp *http.Response, respBody []byte, err error) {
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
	queryParams []QueryParameter,
) (respDetails models.ResponseDetails, err error) {
	req, err := h.createReq(ctx, http.MethodGet, reqPath, nil, queryParams)
	if err != nil {
		return respDetails, err
	}

	resp, parsedBody, err := h.executeReq(req) //nolint: bodyclose // closed in the method itself
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

func (h *HTTPHandler) PostJSONReq(
	ctx context.Context,
	postResource models.Validatable,
	respResource interface{},
	reqPath string,
) (respDetails models.ResponseDetails, err error) {
	err = postResource.Validate()
	if err != nil {
		return respDetails, err
	}
	payload, err := postResource.Marshal()
	if err != nil {
		return respDetails, err
	}
	return h.postRequest(ctx, payload, respResource, reqPath, "application/json")
}

func (h *HTTPHandler) PostNoBodyReq(
	ctx context.Context,
	respResource interface{},
	reqPath string,
) (respDetails models.ResponseDetails, err error) {
	return h.postNoBodyRequest(ctx, respResource, reqPath)
}

func (h *HTTPHandler) PutJSONReq(
	ctx context.Context,
	putResource models.Validatable,
	respResource interface{},
	reqPath string,
	queryParams []QueryParameter,
) (respDetails models.ResponseDetails, err error) {
	err = putResource.Validate()
	if err != nil {
		return respDetails, err
	}
	payload, err := putResource.Marshal()
	if err != nil {
		return respDetails, err
	}
	return h.putRequest(ctx, payload, respResource, reqPath, "application/json", queryParams)
}

func (h *HTTPHandler) PostMultipartReq(
	ctx context.Context,
	postResource models.MultipartValidatable,
	respResource interface{},
	reqPath string,
) (respDetails models.ResponseDetails, err error) {
	err = postResource.Validate()
	if err != nil {
		return respDetails, err
	}
	payload, err := postResource.Marshal()
	if err != nil {
		return respDetails, err
	}
	return h.postRequest(
		ctx,
		payload,
		respResource,
		reqPath,
		fmt.Sprintf("multipart/form-data; boundary=%s", postResource.GetMultipartBoundary()),
	)
}

func (h *HTTPHandler) DeleteRequest(
	ctx context.Context,
	reqPath string,
	queryParams []QueryParameter,
) (respDetails models.ResponseDetails, err error) {
	req, err := h.createReq(ctx, http.MethodDelete, reqPath, nil, queryParams)
	if err != nil {
		return respDetails, err
	}

	resp, parsedBody, err := h.executeReq(req) //nolint: bodyclose // closed in the method itself
	if err != nil {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
		return respDetails, err
	}
	respDetails.HTTPResponse = *resp

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
	}

	return respDetails, err
}

func (h *HTTPHandler) postRequest(
	ctx context.Context,
	payload *bytes.Buffer,
	respResource interface{},
	reqPath string,
	contentType string,
) (respDetails models.ResponseDetails, err error) {
	req, err := h.createReq(ctx, http.MethodPost, reqPath, payload, nil)
	if err != nil {
		return respDetails, err
	}
	req.Header.Set("Content-Type", contentType)

	resp, parsedBody, err := h.executeReq(req) //nolint: bodyclose // closed in the method itself
	if err != nil {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
		return respDetails, err
	}
	respDetails.HTTPResponse = *resp

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(parsedBody, &respResource)
	} else {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
		// MMS 4xx/5xx responses use the same response as 2xx responses
		if _, ok := respResource.(*models.SendMMSResponse); ok {
			_ = json.Unmarshal(parsedBody, &respResource)
		}
	}
	return respDetails, err
}

func (h *HTTPHandler) postNoBodyRequest(
	ctx context.Context,
	respResource interface{},
	reqPath string,
) (respDetails models.ResponseDetails, err error) {
	req, err := h.createReq(ctx, http.MethodPost, reqPath, nil, nil)
	if err != nil {
		return respDetails, err
	}

	resp, parsedBody, err := h.executeReq(req) //nolint: bodyclose // closed in the method itself
	if err != nil {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
		return respDetails, err
	}
	respDetails.HTTPResponse = *resp

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		if len(parsedBody) > 0 {
			err = json.Unmarshal(parsedBody, &respResource)
		}
	} else {
		_ = json.Unmarshal(parsedBody, &respDetails.ErrorResponse)
		// MMS 4xx/5xx responses use the same response as 2xx responses
		if _, ok := respResource.(*models.SendMMSResponse); ok {
			_ = json.Unmarshal(parsedBody, &respResource)
		}
	}
	return respDetails, err
}

func (h *HTTPHandler) putRequest(
	ctx context.Context,
	payload *bytes.Buffer,
	respResource interface{},
	reqPath string,
	contentType string,
	queryParams []QueryParameter,
) (respDetails models.ResponseDetails, err error) {
	req, err := h.createReq(ctx, http.MethodPut, reqPath, payload, queryParams)
	if err != nil {
		return respDetails, err
	}
	req.Header.Set("Content-Type", contentType)

	resp, parsedBody, err := h.executeReq(req) //nolint: bodyclose // closed in the method itself
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

func (h *HTTPHandler) generateCommonHeaders() http.Header {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("App %s", h.APIKey))
	header.Add("Accept", "application/json")
	return header
}

func generateQueryParams(params []QueryParameter) string {
	q := url.Values{}
	for _, param := range params {
		if param.Value != "" {
			q.Add(param.Name, param.Value)
		}
	}

	return q.Encode()
}
