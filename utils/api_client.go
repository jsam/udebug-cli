package utils

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jsam/udebug-cli/client"
	"os"
)

var username, password string

func init() {
	username = os.Getenv("UDEBUG_USERNAME")
	password = os.Getenv("UDEBUG_PASSWORD")
}

type APIClient struct {
	Client *client.UDebug
	Auth   runtime.ClientAuthInfoWriter
}

func NewAPIClient() *APIClient {
	client := client.NewHTTPClientWithConfig(strfmt.Default, client.DefaultTransportConfig())
	auth := httptransport.BasicAuth(username, password)
	return &APIClient{
		client,
		auth,
	}
}
