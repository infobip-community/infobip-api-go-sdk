package account

import (
	"context"

	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

// Account provides methods to interact with the Infobip Account API.
// Account API docs: https://www.infobip.com/docs/api/platform/account-management/
type Account interface {
	// Returns account's credit balance.
	Balance(ctx context.Context) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// Returns account's free messages.
	GetFreeMessagesCount(ctx context.Context) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// Returns account credit balance with currency sign and free message count.
	GetTotalAccountBalance(ctx context.Context) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// Get all accounts
	GetAllAccounts(ctx context.Context, queryParams models.ListPurchasedNumbersParam) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to update an account.
	UpdateAccount(ctx context.Context, accountKey string, request models.RescheduleEmailRequest) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to fetch an API keys by filter. Only users with certain roles can fetch api keys;
	// for example, Account Manager and Integrations Manager roles.
	GetAPIKeysByFilter(ctx context.Context, queryParams models.ListPurchasedNumbersParam) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to create an API key. Only users with certain roles can create api keys;
	// for example, Account Manager and Integrations Manager roles
	CreateAPIKey(ctx context.Context, request models.RescheduleEmailRequest) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to fetch an API key. Only users with certain roles can fetch api key;
	// for example, Account Manager and Integrations Manager roles.
	GetAPIKey(ctx context.Context, APIkeyID string) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to update an API key. Only users with certain roles can create api keys;
	// for example, Account Manager and Integrations Manager roles.
	UpdateAPIKey(ctx context.Context, APIkeyID string, request models.RescheduleEmailRequest) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to create a session (login) which by default will expire after 60 minutes.
	// If you want to create a new token before the session expires, you'll need to destroy it first.
	CreateSession(ctx context.Context, request models.RescheduleEmailRequest) (
		resp models.GetAvailableNumbersResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to destroy a session (login).
	DeleteSession(ctx context.Context) (respDetails models.ResponseDetails, err error)

	// Generate OAuth2 access token that can later on be used to authenticate other Infobip API calls.
	CreateOauth2(ctx context.Context, request models.SMSDestination) (
		token models.Action, respDetails models.ResponseDetails, err error)
}
