// Code generated by go-swagger; DO NOT EDIT.

package output

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new output API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for output API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetOutputAPIOutputRetrieveJSON gets output for input
*/
func (a *Client) GetOutputAPIOutputRetrieveJSON(params *GetOutputAPIOutputRetrieveJSONParams, authInfo runtime.ClientAuthInfoWriter) (*GetOutputAPIOutputRetrieveJSONOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOutputAPIOutputRetrieveJSONParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOutputAPIOutputRetrieveJSON",
		Method:             "GET",
		PathPattern:        "/output_api/output/retrieve.json",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOutputAPIOutputRetrieveJSONReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOutputAPIOutputRetrieveJSONOK), nil

}

/*
PostOutputAPIPostCustomInputOutputJSON gets ouput for custom input
*/
func (a *Client) PostOutputAPIPostCustomInputOutputJSON(params *PostOutputAPIPostCustomInputOutputJSONParams, authInfo runtime.ClientAuthInfoWriter) (*PostOutputAPIPostCustomInputOutputJSONOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOutputAPIPostCustomInputOutputJSONParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOutputAPIPostCustomInputOutputJSON",
		Method:             "POST",
		PathPattern:        "/output_api/post_custom_input/output.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOutputAPIPostCustomInputOutputJSONReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOutputAPIPostCustomInputOutputJSONOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}