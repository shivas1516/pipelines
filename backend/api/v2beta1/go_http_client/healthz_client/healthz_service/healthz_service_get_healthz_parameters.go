// Code generated by go-swagger; DO NOT EDIT.

package healthz_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHealthzServiceGetHealthzParams creates a new HealthzServiceGetHealthzParams object
// with the default values initialized.
func NewHealthzServiceGetHealthzParams() *HealthzServiceGetHealthzParams {

	return &HealthzServiceGetHealthzParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHealthzServiceGetHealthzParamsWithTimeout creates a new HealthzServiceGetHealthzParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHealthzServiceGetHealthzParamsWithTimeout(timeout time.Duration) *HealthzServiceGetHealthzParams {

	return &HealthzServiceGetHealthzParams{

		timeout: timeout,
	}
}

// NewHealthzServiceGetHealthzParamsWithContext creates a new HealthzServiceGetHealthzParams object
// with the default values initialized, and the ability to set a context for a request
func NewHealthzServiceGetHealthzParamsWithContext(ctx context.Context) *HealthzServiceGetHealthzParams {

	return &HealthzServiceGetHealthzParams{

		Context: ctx,
	}
}

// NewHealthzServiceGetHealthzParamsWithHTTPClient creates a new HealthzServiceGetHealthzParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHealthzServiceGetHealthzParamsWithHTTPClient(client *http.Client) *HealthzServiceGetHealthzParams {

	return &HealthzServiceGetHealthzParams{
		HTTPClient: client,
	}
}

/*HealthzServiceGetHealthzParams contains all the parameters to send to the API endpoint
for the healthz service get healthz operation typically these are written to a http.Request
*/
type HealthzServiceGetHealthzParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the healthz service get healthz params
func (o *HealthzServiceGetHealthzParams) WithTimeout(timeout time.Duration) *HealthzServiceGetHealthzParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the healthz service get healthz params
func (o *HealthzServiceGetHealthzParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the healthz service get healthz params
func (o *HealthzServiceGetHealthzParams) WithContext(ctx context.Context) *HealthzServiceGetHealthzParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the healthz service get healthz params
func (o *HealthzServiceGetHealthzParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the healthz service get healthz params
func (o *HealthzServiceGetHealthzParams) WithHTTPClient(client *http.Client) *HealthzServiceGetHealthzParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the healthz service get healthz params
func (o *HealthzServiceGetHealthzParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HealthzServiceGetHealthzParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
