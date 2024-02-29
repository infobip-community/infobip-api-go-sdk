package account

import (
	"context"
	"fmt"

	"github.com/infobip-community/infobip-api-go-sdk/v3/internal"
	"github.com/infobip-community/infobip-api-go-sdk/v3/pkg/infobip/models"
)

// Account provides methods to interact with the Infobip Account API.
// Account API docs: https://www.infobip.com/docs/api/platform/account-management/
type Account interface {
	// Returns account's credit balance.
	Balance(ctx context.Context) (
		resp models.AccountBalance, respDetails models.ResponseDetails, err error)

	// Returns account's free messages.
	GetFreeMessagesCount(ctx context.Context) (
		resp models.FreeMessagesCount, respDetails models.ResponseDetails, err error)

	// Returns account credit balance with currency sign and free message count.
	GetTotalAccountBalance(ctx context.Context) (
		resp models.TotalAccountBalance, respDetails models.ResponseDetails, err error)

	// Get all accounts
	GetAllAccounts(ctx context.Context, queryParams models.GetAllAccountsParams) (
		resp models.GetAllAccountsResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to update an account.
	UpdateAccount(ctx context.Context, accountKey string, request models.UpdateAccountRequest) (
		resp models.UpdateAccountResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to fetch an API keys by filter. Only users with certain roles can fetch api keys;
	// for example, Account Manager and Integrations Manager roles.
	GetAPIKeysByFilter(ctx context.Context, queryParams models.GetAPIKeybyFilterParam) (
		resp models.GetAPIKeybyFilterResponse, respDetails models.ResponseDetails, err error)

	// This method allows you to create an API key. Only users with certain roles can create api keys;
	// for example, Account Manager and Integrations Manager roles.
	CreateAPIKey(ctx context.Context, request models.APIKey) (
		resp models.APIKey, respDetails models.ResponseDetails, err error)

	// This method allows you to fetch an API key. Only users with certain roles can fetch api key;
	// for example, Account Manager and Integrations Manager roles.
	GetAPIKey(ctx context.Context, apiKeyID string) (
		resp models.APIKey, respDetails models.ResponseDetails, err error)

	// This method allows you to update an API key. Only users with certain roles can create api keys;
	// for example, Account Manager and Integrations Manager roles.
	UpdateAPIKey(ctx context.Context, apiKeyID string, request models.UpdateAPIKeyRequest) (
		resp models.APIKey, respDetails models.ResponseDetails, err error)

	// This method allows you to create a session (login) which by default will expire after 60 minutes.
	// If you want to create a new token before the session expires, you'll need to destroy it first.
	CreateSession(ctx context.Context, request models.CreateSessionRequest) (
		resp models.Token, respDetails models.ResponseDetails, err error)

	// This method allows you to destroy a session (login).
	DeleteSession(ctx context.Context) (respDetails models.ResponseDetails, err error)

	// Generate OAuth2 access token that can later on be used to authenticate other Infobip API calls.
	CreateOauth2(ctx context.Context, request models.CreateOauth2TokenRequest) (
		token models.CreateOauth2TokenResponse, respDetails models.ResponseDetails, err error)
}

type Platform struct {
	ReqHandler internal.HTTPHandler
}

const (
	getAccountBalancePath      = "account/1/balance"
	getFreeMessagesCountPath   = "account/1/free-messages"
	getTotalAccountBalancePath = "account/1/total-balance"
	getAllAccountsPath         = "settings/1/accounts"
	updateAccountPath          = "settings/1/accounts/%s"
	getAPIKeysByFilterPath     = "settings/2/api-keys"    //nolint:gosec // Just a path, not a secret
	createAPIKeyPath           = "settings/2/api-keys"    //nolint:gosec // Just a path, not a secret
	getAPIKeyPath              = "settings/2/api-keys/%s" //nolint:gosec // Just a path, not a secret
	updateAPIKeyPath           = "settings/2/api-keys/%s" //nolint:gosec // Just a path, not a secret
	createSessionPath          = "auth/1/session"
	deleteSessionPath          = "auth/1/session"
	createOAuth2Path           = "auth/1/oauth2/token"
)

// Returns account's credit balance.
func (platform *Platform) Balance(ctx context.Context) (
	resp models.AccountBalance, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.GetRequest(ctx, &resp, getAccountBalancePath, nil)
	return resp, respDetails, err
}

// Returns account's free messages.
func (platform *Platform) GetFreeMessagesCount(ctx context.Context) (
	resp models.FreeMessagesCount, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.GetRequest(ctx, &resp, getFreeMessagesCountPath, nil)
	return resp, respDetails, err
}

// Returns account credit balance with currency sign and free message count.
func (platform *Platform) GetTotalAccountBalance(ctx context.Context) (
	resp models.TotalAccountBalance, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.GetRequest(ctx, &resp, getTotalAccountBalancePath, nil)
	return resp, respDetails, err
}

// Get all accounts.
func (platform *Platform) GetAllAccounts(ctx context.Context, queryParams models.GetAllAccountsParams) (
	resp models.GetAllAccountsResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "name", Value: queryParams.Name},
	}
	if queryParams.Limit > 0 {
		params = append(params, internal.QueryParameter{Name: "limit", Value: fmt.Sprint(queryParams.Limit)})
	}
	if queryParams.Enable != nil {
		params = append(params, internal.QueryParameter{Name: "enable", Value: fmt.Sprint(queryParams.Enable)})
	}
	respDetails, err = platform.ReqHandler.GetRequest(ctx, &resp, getAllAccountsPath, params)
	return resp, respDetails, err
}

// This method allows you to update an account.
func (platform *Platform) UpdateAccount(ctx context.Context, accountKey string, request models.UpdateAccountRequest) (
	resp models.UpdateAccountResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.PutJSONReq(ctx, &request, &resp, fmt.Sprintf(updateAccountPath, accountKey), nil)
	return resp, respDetails, err
}

// This method allows you to fetch an API keys by filter. Only users with certain roles can fetch api keys;
// for example, Account Manager and Integrations Manager roles.
func (platform *Platform) GetAPIKeysByFilter(ctx context.Context, queryParams models.GetAPIKeybyFilterParam) (
	resp models.GetAPIKeybyFilterResponse, respDetails models.ResponseDetails, err error) {
	params := []internal.QueryParameter{
		{Name: "accountId", Value: queryParams.AccountID},
		{Name: "name", Value: queryParams.Name},
		{Name: "orderBy", Value: queryParams.OrderBy},
		{Name: "orderDirection", Value: queryParams.OrderDirection},
		{Name: "apiKeySecret", Value: queryParams.APIKeystring},
	}
	if queryParams.Page > 0 {
		params = append(params, internal.QueryParameter{Name: "page", Value: fmt.Sprint(queryParams.Page)})
	}
	if queryParams.Size > 0 {
		params = append(params, internal.QueryParameter{Name: "size", Value: fmt.Sprint(queryParams.Size)})
	}
	if queryParams.Enable != nil {
		params = append(params, internal.QueryParameter{Name: "enable", Value: fmt.Sprint(*queryParams.Enable)})
	}
	respDetails, err = platform.ReqHandler.GetRequest(ctx, &resp, getAPIKeysByFilterPath, params)
	return resp, respDetails, err
}

// This method allows you to create an API key. Only users with certain roles can create api keys;
// for example, Account Manager and Integrations Manager roles.
func (platform *Platform) CreateAPIKey(ctx context.Context, request models.APIKey) (
	resp models.APIKey, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.PostJSONReq(ctx, &request, &resp, createAPIKeyPath)
	return resp, respDetails, err
}

// This method allows you to fetch an API key. Only users with certain roles can fetch api key;
// for example, Account Manager and Integrations Manager roles.
func (platform *Platform) GetAPIKey(ctx context.Context, apiKeyID string) (
	resp models.APIKey, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.GetRequest(ctx, &resp, fmt.Sprintf(getAPIKeyPath, apiKeyID), nil)
	return resp, respDetails, err
}

// This method allows you to update an API key. Only users with certain roles can create api keys;
// for example, Account Manager and Integrations Manager roles.
func (platform *Platform) UpdateAPIKey(ctx context.Context, apiKeyID string, request models.UpdateAPIKeyRequest) (
	resp models.APIKey, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.PutJSONReq(ctx, &request, &resp, fmt.Sprintf(updateAPIKeyPath, apiKeyID), nil)
	return resp, respDetails, err
}

// This method allows you to create a session (login) which by default will expire after 60 minutes.
// If you want to create a new token before the session expires, you'll need to destroy it first.
func (platform *Platform) CreateSession(ctx context.Context, request models.CreateSessionRequest) (
	resp models.Token, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.PostJSONReq(ctx, &request, &resp, createSessionPath)
	return resp, respDetails, err
}

// This method allows you to destroy a session (login).
func (platform *Platform) DeleteSession(ctx context.Context) (respDetails models.ResponseDetails, err error) {
	return platform.ReqHandler.DeleteRequest(ctx, deleteSessionPath, nil)
}

// Generate OAuth2 access token that can later on be used to authenticate other Infobip API calls.
func (platform *Platform) CreateOauth2(ctx context.Context, request models.CreateOauth2TokenRequest) (
	token models.CreateOauth2TokenResponse, respDetails models.ResponseDetails, err error) {
	respDetails, err = platform.ReqHandler.PostJSONReq(ctx, &request, &token, createAPIKeyPath)
	return token, respDetails, err
}
