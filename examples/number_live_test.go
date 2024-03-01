package examples

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The following examples can also be used to test the client against a real environment.
// Replace the apiKey and baseURL params along with the From, To and Content fields of the message, then run the test.

func TestGetAvailableNumbersExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)

	paramNumber := models.GetAvailableNumbersParams{
		Limit: 1,
	}

	resp, respDetails, err := client.Numbers.GetAvailableNumbers(context.Background(), paramNumber)

	fmt.Printf("%+v \n", resp)
	fmt.Printf("%+v \n", respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.GetAvailableNumbersResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestListPurchasedNumbersExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)

	paramNumber := models.ListPurchasedNumbersParam{
		Limit: 1,
	}

	resp, respDetails, err := client.Numbers.ListPurchasedNumbers(context.Background(), paramNumber)

	fmt.Printf("%+v \n", resp)
	fmt.Printf("%+v \n", respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.GetAvailableNumbersResponse{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestPurchaseNumberExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)

	paramNumber := models.PurchaseNumberRequest{
		Number: "56465456", // Only for US number
	}

	resp, respDetails, err := client.Numbers.PurchaseNumber(context.Background(), paramNumber)

	fmt.Printf("%+v \n", resp)
	fmt.Printf("%+v \n", respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.Number{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestGetPurchasedNumberExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)

	resp, respDetails, err := client.Numbers.GetPurchasedNumber(context.Background(), sender)

	fmt.Printf("%+v \n", resp)
	fmt.Printf("%+v \n", respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.Number{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestUpdatePurchasedNumberExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)

	request := models.UpdatePurchasedNumberRequest{
		ApplicationID: "appId",
	}

	resp, respDetails, err := client.Numbers.UpdatePurshasedNumbers(
		context.Background(), sender, request)

	fmt.Printf("%+v \n", resp)
	fmt.Printf("%+v \n", respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.NotEqual(t, models.Number{}, resp)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}

func TestCalcelPurchasedNumberExample(t *testing.T) {
	client, err := infobip.NewClient(baseURL, apiKey)
	require.NoError(t, err)

	respDetails, err := client.Numbers.CancelNumber(
		context.Background(), sender)

	fmt.Printf("%+v \n", respDetails)

	require.Nil(t, err)
	assert.NotNil(t, respDetails)
	assert.Equal(t, http.StatusOK, respDetails.HTTPResponse.StatusCode)
}
