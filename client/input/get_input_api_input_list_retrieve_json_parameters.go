// Code generated by go-swagger; DO NOT EDIT.

package input

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

// NewGetInputAPIInputListRetrieveJSONParams creates a new GetInputAPIInputListRetrieveJSONParams object
// with the default values initialized.
func NewGetInputAPIInputListRetrieveJSONParams() *GetInputAPIInputListRetrieveJSONParams {
	var ()
	return &GetInputAPIInputListRetrieveJSONParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInputAPIInputListRetrieveJSONParamsWithTimeout creates a new GetInputAPIInputListRetrieveJSONParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInputAPIInputListRetrieveJSONParamsWithTimeout(timeout time.Duration) *GetInputAPIInputListRetrieveJSONParams {
	var ()
	return &GetInputAPIInputListRetrieveJSONParams{

		timeout: timeout,
	}
}

// NewGetInputAPIInputListRetrieveJSONParamsWithContext creates a new GetInputAPIInputListRetrieveJSONParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInputAPIInputListRetrieveJSONParamsWithContext(ctx context.Context) *GetInputAPIInputListRetrieveJSONParams {
	var ()
	return &GetInputAPIInputListRetrieveJSONParams{

		Context: ctx,
	}
}

// NewGetInputAPIInputListRetrieveJSONParamsWithHTTPClient creates a new GetInputAPIInputListRetrieveJSONParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInputAPIInputListRetrieveJSONParamsWithHTTPClient(client *http.Client) *GetInputAPIInputListRetrieveJSONParams {
	var ()
	return &GetInputAPIInputListRetrieveJSONParams{
		HTTPClient: client,
	}
}

/*GetInputAPIInputListRetrieveJSONParams contains all the parameters to send to the API endpoint
for the get input API input list retrieve JSON operation typically these are written to a http.Request
*/
type GetInputAPIInputListRetrieveJSONParams struct {

	/*JudgeAlias
	  name of judge(ex. UVa)

	*/
	JudgeAlias string
	/*ProblemID
	  id of problem(ex. 100)

	*/
	ProblemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) WithTimeout(timeout time.Duration) *GetInputAPIInputListRetrieveJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) WithContext(ctx context.Context) *GetInputAPIInputListRetrieveJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) WithHTTPClient(client *http.Client) *GetInputAPIInputListRetrieveJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJudgeAlias adds the judgeAlias to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) WithJudgeAlias(judgeAlias string) *GetInputAPIInputListRetrieveJSONParams {
	o.SetJudgeAlias(judgeAlias)
	return o
}

// SetJudgeAlias adds the judgeAlias to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) SetJudgeAlias(judgeAlias string) {
	o.JudgeAlias = judgeAlias
}

// WithProblemID adds the problemID to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) WithProblemID(problemID string) *GetInputAPIInputListRetrieveJSONParams {
	o.SetProblemID(problemID)
	return o
}

// SetProblemID adds the problemId to the get input API input list retrieve JSON params
func (o *GetInputAPIInputListRetrieveJSONParams) SetProblemID(problemID string) {
	o.ProblemID = problemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInputAPIInputListRetrieveJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param judge_alias
	qrJudgeAlias := o.JudgeAlias
	qJudgeAlias := qrJudgeAlias
	if qJudgeAlias != "" {
		if err := r.SetQueryParam("judge_alias", qJudgeAlias); err != nil {
			return err
		}
	}

	// query param problem_id
	qrProblemID := o.ProblemID
	qProblemID := qrProblemID
	if qProblemID != "" {
		if err := r.SetQueryParam("problem_id", qProblemID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
