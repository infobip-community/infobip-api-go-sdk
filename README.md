# Infobip API Go SDK

![Workflow](https://github.com/infobip-community/infobip-api-go-sdk/actions/workflows/checks.yml/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/infobip-community/infobip-api-go-sdk)
[![Licence](https://img.shields.io/github/license/infobip-community/infobip-api-go-sdk)](LICENSE)
![Lines](https://img.shields.io/tokei/lines/github/infobip-community/infobip-api-go-sdk)

Go client SDK for Infobip's API Channels

---

## 📡 Supported Channels

- [SMS](https://www.infobip.com/docs/api#channels/sms)
- [WhatsApp](https://www.infobip.com/docs/api#channels/whatsapp)
- [Email](https://www.infobip.com/docs/api#channels/email)
- [MMS](https://www.infobip.com/docs/api#channels/mms)
- [RCS](https://www.infobip.com/docs/api#channels/rcs)
- [WebRTC](https://www.infobip.com/docs/api#channels/webrtc)

More channels to be added in the near future.

## 🔐 Authentication

Currently, infobip-api-go-sdk only supports API Key authentication, and the key needs to be passed during client creation.
This will most likely change with future versions, once more authentication methods are included. You can get your base URL and API key by logging into Portal. Follow the instructions [here](https://www.infobip.com/docs/api).

## 📦 Installation

Currently, infobip-api-go-sdk requires Go version 1.13 or greater.
We'll do our best not to break older versions of Go unless it's absolutely necessary, but due to tooling constraints,
we don't always test older versions.

infobip-api-go-sdk is compatible with modern Go modules. With Go installed, running the following:

```bash
go get "github.com/infobip-community/infobip-api-go-sdk"
```

will add the client to the current module, along with all of its dependencies.

## 🚀 Usage

```go
import "github.com/infobip-community/infobip-api-go-sdk/pkg/infobip"
```

Construct a new Infobip client:

```go
client, err := infobip.NewClient("https://myinfobipurl.com", "secret")
```

Or, passing your own HTTP client:

```go
client, err := NewClient(baseURL, apiKey, WithHTTPClient(http.Client{Timeout: 3 * time.Second}))
```

Afterwards, use the various services on the client to
access different channels of the Infobip API. For example:

```go
client, err := infobip.NewClient(baseURL, apiKey)

// Send a WhatsApp text message
message := models.TextMsg{
    MsgCommon: models.MsgCommon{
        From: "111111111111",
        To:   "222222222222",
    },
    Content: models.TextContent{
		Text: "This message was sent from the Infobip API using the Go API client."
	},
}
msgResp, respDetails, err := client.WhatsApp.SendTextMsg(context.Background(), message)
```

Requests return the resource returned by the server (if applicable), response details and an error.
Response details contain the raw http.Response object along with ErrorDetails which will be populated for cases
where the server does not return a successful HTTP response code.

An error will only be returned if the underlying HTTP request failed (a network issue, failure reading the body, etc.).
In other words, 4xx/5xx responses do **not** return an error, and the user should instead check for them
by inspecting the ResponseDetails.HTTPResponse.StatusCode value. Note that for requests which require a payload (e.g. POST, PATCH),
the object representing the payload will be validated before it is sent.

The channels of the client divide the API into multiple parts, corresponding to the Infobip Channels documented at
https://www.infobip.com/docs/api#channels.

NOTE: Using the [context](https://godoc.org/context) package, the user can pass cancellation signals and deadlines
to the underlying requests that the client makes. If you don't want to use this feature, then using `context.Background()`
should be sufficient.


## 👀 Examples

The best way to learn how to use the library is to check the examples. The [examples](https://github.com/infobip-community/infobip-api-go-sdk/tree/main/examples) directory
contains tests intended to be used for live testing all endpoints of available channels. The prerequisite for running an individual test is changing
the apiKey and baseURL variables, along with certain message fields, depending on the endpoint (e.g. From/To for WhatsApp).

## ⚖️ License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
