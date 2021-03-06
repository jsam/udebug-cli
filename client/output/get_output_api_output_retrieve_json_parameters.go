// Code generated by go-swagger; DO NOT EDIT.

package output

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOutputAPIOutputRetrieveJSONParams creates a new GetOutputAPIOutputRetrieveJSONParams object
// with the default values initialized.
func NewGetOutputAPIOutputRetrieveJSONParams() *GetOutputAPIOutputRetrieveJSONParams {
	var ()
	return &GetOutputAPIOutputRetrieveJSONParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutputAPIOutputRetrieveJSONParamsWithTimeout creates a new GetOutputAPIOutputRetrieveJSONParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutputAPIOutputRetrieveJSONParamsWithTimeout(timeout time.Duration) *GetOutputAPIOutputRetrieveJSONParams {
	var ()
	return &GetOutputAPIOutputRetrieveJSONParams{

		timeout: timeout,
	}
}

// NewGetOutputAPIOutputRetrieveJSONParamsWithContext creates a new GetOutputAPIOutputRetrieveJSONParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutputAPIOutputRetrieveJSONParamsWithContext(ctx context.Context) *GetOutputAPIOutputRetrieveJSONParams {
	var ()
	return &GetOutputAPIOutputRetrieveJSONParams{

		Context: ctx,
	}
}

// NewGetOutputAPIOutputRetrieveJSONParamsWithHTTPClient creates a new GetOutputAPIOutputRetrieveJSONParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutputAPIOutputRetrieveJSONParamsWithHTTPClient(client *http.Client) *GetOutputAPIOutputRetrieveJSONParams {
	var ()
	return &GetOutputAPIOutputRetrieveJSONParams{
		HTTPClient: client,
	}
}

/*GetOutputAPIOutputRetrieveJSONParams contains all the parameters to send to the API endpoint
for the get output API output retrieve JSON operation typically these are written to a http.Request
*/
type GetOutputAPIOutputRetrieveJSONParams struct {

	/*InputID
	  id of input(ex. 809768)

	*/
	InputID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) WithTimeout(timeout time.Duration) *GetOutputAPIOutputRetrieveJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) WithContext(ctx context.Context) *GetOutputAPIOutputRetrieveJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) WithHTTPClient(client *http.Client) *GetOutputAPIOutputRetrieveJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInputID adds the inputID to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) WithInputID(inputID string) *GetOutputAPIOutputRetrieveJSONParams {
	o.SetInputID(inputID)
	return o
}

// SetInputID adds the inputId to the get output API output retrieve JSON params
func (o *GetOutputAPIOutputRetrieveJSONParams) SetInputID(inputID string) {
	o.InputID = inputID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutputAPIOutputRetrieveJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param input_id
	qrInputID := o.InputID
	qInputID := qrInputID
	if qInputID != "" {
		if err := r.SetQueryParam("input_id", qInputID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
