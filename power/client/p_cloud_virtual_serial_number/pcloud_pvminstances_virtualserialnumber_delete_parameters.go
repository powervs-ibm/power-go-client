// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_virtual_serial_number

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

// NewPcloudPvminstancesVirtualserialnumberDeleteParams creates a new PcloudPvminstancesVirtualserialnumberDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesVirtualserialnumberDeleteParams() *PcloudPvminstancesVirtualserialnumberDeleteParams {
	return &PcloudPvminstancesVirtualserialnumberDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesVirtualserialnumberDeleteParamsWithTimeout creates a new PcloudPvminstancesVirtualserialnumberDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesVirtualserialnumberDeleteParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	return &PcloudPvminstancesVirtualserialnumberDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesVirtualserialnumberDeleteParamsWithContext creates a new PcloudPvminstancesVirtualserialnumberDeleteParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesVirtualserialnumberDeleteParamsWithContext(ctx context.Context) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	return &PcloudPvminstancesVirtualserialnumberDeleteParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesVirtualserialnumberDeleteParamsWithHTTPClient creates a new PcloudPvminstancesVirtualserialnumberDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesVirtualserialnumberDeleteParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	return &PcloudPvminstancesVirtualserialnumberDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesVirtualserialnumberDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances virtualserialnumber delete operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesVirtualserialnumberDeleteParams struct {

	/* Body.

	   Parameters to unassign a Virtual Serial Number attached to a PVM Instance
	*/
	Body *models.DeleteServerVirtualSerialNumber

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances virtualserialnumber delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WithDefaults() *PcloudPvminstancesVirtualserialnumberDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances virtualserialnumber delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WithContext(ctx context.Context) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WithBody(body *models.DeleteServerVirtualSerialNumber) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) SetBody(body *models.DeleteServerVirtualSerialNumber) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesVirtualserialnumberDeleteParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances virtualserialnumber delete params
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesVirtualserialnumberDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
