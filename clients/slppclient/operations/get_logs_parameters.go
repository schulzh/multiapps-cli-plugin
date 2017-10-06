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

// NewGetLogsParams creates a new GetLogsParams object
// with the default values initialized.
func NewGetLogsParams() *GetLogsParams {

	return &GetLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogsParamsWithTimeout creates a new GetLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogsParamsWithTimeout(timeout time.Duration) *GetLogsParams {

	return &GetLogsParams{

		timeout: timeout,
	}
}

// NewGetLogsParamsWithContext creates a new GetLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogsParamsWithContext(ctx context.Context) *GetLogsParams {

	return &GetLogsParams{

		Context: ctx,
	}
}

// NewGetLogsParamsWithHTTPClient creates a new GetLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogsParamsWithHTTPClient(client *http.Client) *GetLogsParams {

	return &GetLogsParams{
		HTTPClient: client,
	}
}

/*GetLogsParams contains all the parameters to send to the API endpoint
for the get logs operation typically these are written to a http.Request
*/
type GetLogsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get logs params
func (o *GetLogsParams) WithTimeout(timeout time.Duration) *GetLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logs params
func (o *GetLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logs params
func (o *GetLogsParams) WithContext(ctx context.Context) *GetLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logs params
func (o *GetLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) WithHTTPClient(client *http.Client) *GetLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
