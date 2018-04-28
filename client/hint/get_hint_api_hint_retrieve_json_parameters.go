// Code generated by go-swagger; DO NOT EDIT.

package hint

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

// NewGetHintAPIHintRetrieveJSONParams creates a new GetHintAPIHintRetrieveJSONParams object
// with the default values initialized.
func NewGetHintAPIHintRetrieveJSONParams() *GetHintAPIHintRetrieveJSONParams {
	var ()
	return &GetHintAPIHintRetrieveJSONParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHintAPIHintRetrieveJSONParamsWithTimeout creates a new GetHintAPIHintRetrieveJSONParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHintAPIHintRetrieveJSONParamsWithTimeout(timeout time.Duration) *GetHintAPIHintRetrieveJSONParams {
	var ()
	return &GetHintAPIHintRetrieveJSONParams{

		timeout: timeout,
	}
}

// NewGetHintAPIHintRetrieveJSONParamsWithContext creates a new GetHintAPIHintRetrieveJSONParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHintAPIHintRetrieveJSONParamsWithContext(ctx context.Context) *GetHintAPIHintRetrieveJSONParams {
	var ()
	return &GetHintAPIHintRetrieveJSONParams{

		Context: ctx,
	}
}

// NewGetHintAPIHintRetrieveJSONParamsWithHTTPClient creates a new GetHintAPIHintRetrieveJSONParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHintAPIHintRetrieveJSONParamsWithHTTPClient(client *http.Client) *GetHintAPIHintRetrieveJSONParams {
	var ()
	return &GetHintAPIHintRetrieveJSONParams{
		HTTPClient: client,
	}
}

/*GetHintAPIHintRetrieveJSONParams contains all the parameters to send to the API endpoint
for the get hint API hint retrieve JSON operation typically these are written to a http.Request
*/
type GetHintAPIHintRetrieveJSONParams struct {

	/*HintID
	  id of hint(ex. 813889)

	*/
	HintID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) WithTimeout(timeout time.Duration) *GetHintAPIHintRetrieveJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) WithContext(ctx context.Context) *GetHintAPIHintRetrieveJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) WithHTTPClient(client *http.Client) *GetHintAPIHintRetrieveJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHintID adds the hintID to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) WithHintID(hintID string) *GetHintAPIHintRetrieveJSONParams {
	o.SetHintID(hintID)
	return o
}

// SetHintID adds the hintId to the get hint API hint retrieve JSON params
func (o *GetHintAPIHintRetrieveJSONParams) SetHintID(hintID string) {
	o.HintID = hintID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHintAPIHintRetrieveJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param hint_id
	qrHintID := o.HintID
	qHintID := qrHintID
	if qHintID != "" {
		if err := r.SetQueryParam("hint_id", qHintID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}