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

// NewGetHintAPIHintListRetrieveJSONParams creates a new GetHintAPIHintListRetrieveJSONParams object
// with the default values initialized.
func NewGetHintAPIHintListRetrieveJSONParams() *GetHintAPIHintListRetrieveJSONParams {
	var ()
	return &GetHintAPIHintListRetrieveJSONParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHintAPIHintListRetrieveJSONParamsWithTimeout creates a new GetHintAPIHintListRetrieveJSONParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHintAPIHintListRetrieveJSONParamsWithTimeout(timeout time.Duration) *GetHintAPIHintListRetrieveJSONParams {
	var ()
	return &GetHintAPIHintListRetrieveJSONParams{

		timeout: timeout,
	}
}

// NewGetHintAPIHintListRetrieveJSONParamsWithContext creates a new GetHintAPIHintListRetrieveJSONParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHintAPIHintListRetrieveJSONParamsWithContext(ctx context.Context) *GetHintAPIHintListRetrieveJSONParams {
	var ()
	return &GetHintAPIHintListRetrieveJSONParams{

		Context: ctx,
	}
}

// NewGetHintAPIHintListRetrieveJSONParamsWithHTTPClient creates a new GetHintAPIHintListRetrieveJSONParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHintAPIHintListRetrieveJSONParamsWithHTTPClient(client *http.Client) *GetHintAPIHintListRetrieveJSONParams {
	var ()
	return &GetHintAPIHintListRetrieveJSONParams{
		HTTPClient: client,
	}
}

/*GetHintAPIHintListRetrieveJSONParams contains all the parameters to send to the API endpoint
for the get hint API hint list retrieve JSON operation typically these are written to a http.Request
*/
type GetHintAPIHintListRetrieveJSONParams struct {

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

// WithTimeout adds the timeout to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) WithTimeout(timeout time.Duration) *GetHintAPIHintListRetrieveJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) WithContext(ctx context.Context) *GetHintAPIHintListRetrieveJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) WithHTTPClient(client *http.Client) *GetHintAPIHintListRetrieveJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJudgeAlias adds the judgeAlias to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) WithJudgeAlias(judgeAlias string) *GetHintAPIHintListRetrieveJSONParams {
	o.SetJudgeAlias(judgeAlias)
	return o
}

// SetJudgeAlias adds the judgeAlias to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) SetJudgeAlias(judgeAlias string) {
	o.JudgeAlias = judgeAlias
}

// WithProblemID adds the problemID to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) WithProblemID(problemID string) *GetHintAPIHintListRetrieveJSONParams {
	o.SetProblemID(problemID)
	return o
}

// SetProblemID adds the problemId to the get hint API hint list retrieve JSON params
func (o *GetHintAPIHintListRetrieveJSONParams) SetProblemID(problemID string) {
	o.ProblemID = problemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHintAPIHintListRetrieveJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
