package numbers

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

const (
	getAvailableNumbersPath   = "numbers/1/numbers/available"
	listPurchasedNumbersPath  = "numbers/1/numbers"
	purchaseNumberPath        = "numbers/1/numbers"
	getpurchasedNumberPath    = "numbers/1/numbers/%s"
	updatepurchasedNumberPath = "numbers/1/numbers/%s"
	deletepurchasedNumberPath = "numbers/1/numbers/%s"
)

// Numbers provides methods to interact with the Infobip Numbers API.
// Numbers API docs: https://www.infobip.com/docs/api#channels/numbers
type Numbers interface {
	// GetAvailableNumbers returns information about your available numbers.
	GetAvailableNumbers(ctx context.Context, queryParams models.GetAvailableNumbersParams) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	ListPurchasedNumbers(ctx context.Context, queryParams models.ListPurchasedNumbersParam,
	) (resp models.ListPurchasedNumbersResponse, respDetails models.ResponseDetails, err error)

	// Using the number ID or number, this method enables you to buy a new number.
	// For buying a US number, only the number should be provided. For all other purchases, only the numberKey must be provided
	PurchaseNumber(ctx context.Context, request models.PurchaseNumberRequest,
	) (resp models.Number, respDetails models.ResponseDetails, err error)

	GetPurchasedNumber(ctx context.Context, numberKey string,
	) (resp models.Number, respDetails models.ResponseDetails, err error)

	UpdatePurshasedNumbers(ctx context.Context, numberKey string,
		request models.UpdatePurchasedNumberRequest,
	) (resp models.Number, respDetails models.ResponseDetails, err error)

	CancelNumber(ctx context.Context, numberKey string,
	) (respDetails models.ResponseDetails, err error)
}

type Channel struct {
	ReqHandler internal.HTTPHandler
}

func (numbers *Channel) GetAvailableNumbers(
	ctx context.Context,
	queryParams models.GetAvailableNumbersParams,
) (resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "country", Value: queryParams.Country},
		{Name: "state", Value: queryParams.State},
		{Name: "number", Value: queryParams.Number},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}
	if queryParams.Page > 0 {
		params = append(params, internal.QueryParameter{Name: "page", Value: fmt.Sprint(queryParams.Page)})
	}
	if queryParams.NPA > 0 {
		params = append(params, internal.QueryParameter{Name: "npa", Value: fmt.Sprint(queryParams.NPA)})
	}
	if queryParams.Nxx > 0 {
		params = append(params, internal.QueryParameter{Name: "nxx", Value: fmt.Sprint(queryParams.Nxx)})
	}
	for _, capability := range queryParams.Capabilities {
		params = append(params, internal.QueryParameter{Name: "capabilities", Value: capability})
	}

	respDetails, err = numbers.ReqHandler.GetRequest(ctx, &resp, getAvailableNumbersPath, params)

	return resp, respDetails, err
}

func (numbers *Channel) ListPurchasedNumbers(
	ctx context.Context,
	queryParams models.ListPurchasedNumbersParam,
) (resp models.ListPurchasedNumbersResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "number", Value: queryParams.Number},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}
	if queryParams.Page > 0 {
		params = append(params, internal.QueryParameter{Name: "page", Value: fmt.Sprint(queryParams.Page)})
	}
	respDetails, err = numbers.ReqHandler.GetRequest(ctx, &resp, listPurchasedNumbersPath, params)
	return resp, respDetails, err
}

func (numbers *Channel) PurchaseNumber(
	ctx context.Context,
	request models.PurchaseNumberRequest,
) (resp models.Number, respDetails models.ResponseDetails, err error) {
	respDetails, err = numbers.ReqHandler.PostJSONReq(ctx, &request, &resp, purchaseNumberPath)
	return resp, respDetails, err
}

func (numbers *Channel) GetPurchasedNumber(
	ctx context.Context,
	numberKey string,
) (resp models.Number, respDetails models.ResponseDetails, err error) {
	respDetails, err = numbers.ReqHandler.GetRequest(ctx, &resp, fmt.Sprintf(getpurchasedNumberPath, numberKey), nil)
	return resp, respDetails, err
}

func (numbers *Channel) UpdatePurshasedNumbers(
	ctx context.Context,
	numberKey string,
	request models.UpdatePurchasedNumberRequest,
) (resp models.Number, respDetails models.ResponseDetails, err error) {
	respDetails, err = numbers.ReqHandler.PutJSONReq(
		ctx,
		&request,
		&resp,
		fmt.Sprintf(updatepurchasedNumberPath, numberKey),
		nil)
	return resp, respDetails, err
}
func (numbers *Channel) CancelNumber(
	ctx context.Context,
	numberKey string,
) (respDetails models.ResponseDetails, err error) {
	return numbers.ReqHandler.DeleteRequest(ctx, fmt.Sprintf(deletepurchasedNumberPath, numberKey), nil)
}
