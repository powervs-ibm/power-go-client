// Code generated by go-swagger; DO NOT EDIT.

package network_security_groups

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

// NewV1NetworkSecurityGroupsIDPutParams creates a new V1NetworkSecurityGroupsIDPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1NetworkSecurityGroupsIDPutParams() *V1NetworkSecurityGroupsIDPutParams {
	return &V1NetworkSecurityGroupsIDPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1NetworkSecurityGroupsIDPutParamsWithTimeout creates a new V1NetworkSecurityGroupsIDPutParams object
// with the ability to set a timeout on a request.
func NewV1NetworkSecurityGroupsIDPutParamsWithTimeout(timeout time.Duration) *V1NetworkSecurityGroupsIDPutParams {
	return &V1NetworkSecurityGroupsIDPutParams{
		timeout: timeout,
	}
}

// NewV1NetworkSecurityGroupsIDPutParamsWithContext creates a new V1NetworkSecurityGroupsIDPutParams object
// with the ability to set a context for a request.
func NewV1NetworkSecurityGroupsIDPutParamsWithContext(ctx context.Context) *V1NetworkSecurityGroupsIDPutParams {
	return &V1NetworkSecurityGroupsIDPutParams{
		Context: ctx,
	}
}

// NewV1NetworkSecurityGroupsIDPutParamsWithHTTPClient creates a new V1NetworkSecurityGroupsIDPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1NetworkSecurityGroupsIDPutParamsWithHTTPClient(client *http.Client) *V1NetworkSecurityGroupsIDPutParams {
	return &V1NetworkSecurityGroupsIDPutParams{
		HTTPClient: client,
	}
}

/*
V1NetworkSecurityGroupsIDPutParams contains all the parameters to send to the API endpoint

	for the v1 network security groups id put operation.

	Typically these are written to a http.Request.
*/
type V1NetworkSecurityGroupsIDPutParams struct {

	/* Body.

	   Parameters for the update of a Network Security Group
	*/
	Body *models.NetworkSecurityGroupUpdate

	/* NetworkSecurityGroupID.

	   Network Security Group ID
	*/
	NetworkSecurityGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 network security groups id put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworkSecurityGroupsIDPutParams) WithDefaults() *V1NetworkSecurityGroupsIDPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 network security groups id put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1NetworkSecurityGroupsIDPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) WithTimeout(timeout time.Duration) *V1NetworkSecurityGroupsIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) WithContext(ctx context.Context) *V1NetworkSecurityGroupsIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) WithHTTPClient(client *http.Client) *V1NetworkSecurityGroupsIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) WithBody(body *models.NetworkSecurityGroupUpdate) *V1NetworkSecurityGroupsIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) SetBody(body *models.NetworkSecurityGroupUpdate) {
	o.Body = body
}

// WithNetworkSecurityGroupID adds the networkSecurityGroupID to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) WithNetworkSecurityGroupID(networkSecurityGroupID string) *V1NetworkSecurityGroupsIDPutParams {
	o.SetNetworkSecurityGroupID(networkSecurityGroupID)
	return o
}

// SetNetworkSecurityGroupID adds the networkSecurityGroupId to the v1 network security groups id put params
func (o *V1NetworkSecurityGroupsIDPutParams) SetNetworkSecurityGroupID(networkSecurityGroupID string) {
	o.NetworkSecurityGroupID = networkSecurityGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *V1NetworkSecurityGroupsIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param network_security_group_id
	if err := r.SetPathParam("network_security_group_id", o.NetworkSecurityGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}