// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewInternalV1OperationsSnapshotsPostParams creates a new InternalV1OperationsSnapshotsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1OperationsSnapshotsPostParams() *InternalV1OperationsSnapshotsPostParams {
	return &InternalV1OperationsSnapshotsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1OperationsSnapshotsPostParamsWithTimeout creates a new InternalV1OperationsSnapshotsPostParams object
// with the ability to set a timeout on a request.
func NewInternalV1OperationsSnapshotsPostParamsWithTimeout(timeout time.Duration) *InternalV1OperationsSnapshotsPostParams {
	return &InternalV1OperationsSnapshotsPostParams{
		timeout: timeout,
	}
}

// NewInternalV1OperationsSnapshotsPostParamsWithContext creates a new InternalV1OperationsSnapshotsPostParams object
// with the ability to set a context for a request.
func NewInternalV1OperationsSnapshotsPostParamsWithContext(ctx context.Context) *InternalV1OperationsSnapshotsPostParams {
	return &InternalV1OperationsSnapshotsPostParams{
		Context: ctx,
	}
}

// NewInternalV1OperationsSnapshotsPostParamsWithHTTPClient creates a new InternalV1OperationsSnapshotsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1OperationsSnapshotsPostParamsWithHTTPClient(client *http.Client) *InternalV1OperationsSnapshotsPostParams {
	return &InternalV1OperationsSnapshotsPostParams{
		HTTPClient: client,
	}
}

/*
InternalV1OperationsSnapshotsPostParams contains all the parameters to send to the API endpoint

	for the internal v1 operations snapshots post operation.

	Typically these are written to a http.Request.
*/
type InternalV1OperationsSnapshotsPostParams struct {

	/* Body.

	   Parameters for creating snapshot CRN
	*/
	Body *models.InternalOperationsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 operations snapshots post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsSnapshotsPostParams) WithDefaults() *InternalV1OperationsSnapshotsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 operations snapshots post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsSnapshotsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) WithTimeout(timeout time.Duration) *InternalV1OperationsSnapshotsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) WithContext(ctx context.Context) *InternalV1OperationsSnapshotsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) WithHTTPClient(client *http.Client) *InternalV1OperationsSnapshotsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) WithBody(body *models.InternalOperationsRequest) *InternalV1OperationsSnapshotsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the internal v1 operations snapshots post params
func (o *InternalV1OperationsSnapshotsPostParams) SetBody(body *models.InternalOperationsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1OperationsSnapshotsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}