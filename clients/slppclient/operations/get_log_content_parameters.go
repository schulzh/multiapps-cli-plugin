// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetLogContentParams creates a new GetLogContentParams object
// with the default values initialized.
func NewGetLogContentParams() *GetLogContentParams {
	var ()
	return &GetLogContentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogContentParamsWithTimeout creates a new GetLogContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogContentParamsWithTimeout(timeout time.Duration) *GetLogContentParams {
	var ()
	return &GetLogContentParams{

		timeout: timeout,
	}
}

// NewGetLogContentParamsWithContext creates a new GetLogContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogContentParamsWithContext(ctx context.Context) *GetLogContentParams {
	var ()
	return &GetLogContentParams{

		Context: ctx,
	}
}

// NewGetLogContentParamsWithHTTPClient creates a new GetLogContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogContentParamsWithHTTPClient(client *http.Client) *GetLogContentParams {
	var ()
	return &GetLogContentParams{
		HTTPClient: client,
	}
}

/*GetLogContentParams contains all the parameters to send to the API endpoint
for the get log content operation typically these are written to a http.Request
*/
type GetLogContentParams struct {

	/*LogID*/
	LogID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get log content params
func (o *GetLogContentParams) WithTimeout(timeout time.Duration) *GetLogContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log content params
func (o *GetLogContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log content params
func (o *GetLogContentParams) WithContext(ctx context.Context) *GetLogContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log content params
func (o *GetLogContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log content params
func (o *GetLogContentParams) WithHTTPClient(client *http.Client) *GetLogContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log content params
func (o *GetLogContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogID adds the logID to the get log content params
func (o *GetLogContentParams) WithLogID(logID string) *GetLogContentParams {
	o.SetLogID(logID)
	return o
}

// SetLogID adds the logId to the get log content params
func (o *GetLogContentParams) SetLogID(logID string) {
	o.LogID = logID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param logId
	if err := r.SetPathParam("logId", o.LogID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
