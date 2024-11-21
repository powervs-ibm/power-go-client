// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_network_interfaces

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
)

// NewInternalV1OperationsNetworkinterfacesDeleteParams creates a new InternalV1OperationsNetworkinterfacesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1OperationsNetworkinterfacesDeleteParams() *InternalV1OperationsNetworkinterfacesDeleteParams {
	return &InternalV1OperationsNetworkinterfacesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1OperationsNetworkinterfacesDeleteParamsWithTimeout creates a new InternalV1OperationsNetworkinterfacesDeleteParams object
// with the ability to set a timeout on a request.
func NewInternalV1OperationsNetworkinterfacesDeleteParamsWithTimeout(timeout time.Duration) *InternalV1OperationsNetworkinterfacesDeleteParams {
	return &InternalV1OperationsNetworkinterfacesDeleteParams{
		timeout: timeout,
	}
}

// NewInternalV1OperationsNetworkinterfacesDeleteParamsWithContext creates a new InternalV1OperationsNetworkinterfacesDeleteParams object
// with the ability to set a context for a request.
func NewInternalV1OperationsNetworkinterfacesDeleteParamsWithContext(ctx context.Context) *InternalV1OperationsNetworkinterfacesDeleteParams {
	return &InternalV1OperationsNetworkinterfacesDeleteParams{
		Context: ctx,
	}
}

// NewInternalV1OperationsNetworkinterfacesDeleteParamsWithHTTPClient creates a new InternalV1OperationsNetworkinterfacesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1OperationsNetworkinterfacesDeleteParamsWithHTTPClient(client *http.Client) *InternalV1OperationsNetworkinterfacesDeleteParams {
	return &InternalV1OperationsNetworkinterfacesDeleteParams{
		HTTPClient: client,
	}
}

/*
InternalV1OperationsNetworkinterfacesDeleteParams contains all the parameters to send to the API endpoint

	for the internal v1 operations networkinterfaces delete operation.

	Typically these are written to a http.Request.
*/
type InternalV1OperationsNetworkinterfacesDeleteParams struct {

	/* ResourceCrn.

	   Encoded resource CRN, "/" to be encoded into "%2F", example 'crn:v1:staging:public:power-iaas:satloc_dal_clp2joc20ppo19876n50:a%2Fc7e6bd2517ad44eabbd61fcc25cf68d5:79bffc73-0035-4e7b-b34a-15da38927424:network:d8d51d44-053b-4df3-90b6-31fbe72ba600'
	*/
	ResourceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 operations networkinterfaces delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) WithDefaults() *InternalV1OperationsNetworkinterfacesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 operations networkinterfaces delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) WithTimeout(timeout time.Duration) *InternalV1OperationsNetworkinterfacesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) WithContext(ctx context.Context) *InternalV1OperationsNetworkinterfacesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) WithHTTPClient(client *http.Client) *InternalV1OperationsNetworkinterfacesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceCrn adds the resourceCrn to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) WithResourceCrn(resourceCrn string) *InternalV1OperationsNetworkinterfacesDeleteParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the internal v1 operations networkinterfaces delete params
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1OperationsNetworkinterfacesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_crn
	if err := r.SetPathParam("resource_crn", o.ResourceCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}