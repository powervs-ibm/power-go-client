// Code generated by go-swagger; DO NOT EDIT.

package internal_power_v_s_locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new internal power v s locations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for internal power v s locations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	InternalV1PowervsLocationsActivatePut(params *InternalV1PowervsLocationsActivatePutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsActivatePutOK, error)

	InternalV1PowervsLocationsTagDelete(params *InternalV1PowervsLocationsTagDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTagDeleteOK, error)

	InternalV1PowervsLocationsTagPost(params *InternalV1PowervsLocationsTagPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTagPostOK, error)

	InternalV1PowervsLocationsTransitgatewayGet(params *InternalV1PowervsLocationsTransitgatewayGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTransitgatewayGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
InternalV1PowervsLocationsActivatePut activates a power v s on prem location
*/
func (a *Client) InternalV1PowervsLocationsActivatePut(params *InternalV1PowervsLocationsActivatePutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsActivatePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1PowervsLocationsActivatePutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.powervs.locations.activate.put",
		Method:             "PUT",
		PathPattern:        "/internal/v1/powervs/locations/activate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1PowervsLocationsActivatePutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InternalV1PowervsLocationsActivatePutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.powervs.locations.activate.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InternalV1PowervsLocationsTagDelete deletes a power satellite tag
*/
func (a *Client) InternalV1PowervsLocationsTagDelete(params *InternalV1PowervsLocationsTagDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTagDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1PowervsLocationsTagDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.powervs.locations.tag.delete",
		Method:             "DELETE",
		PathPattern:        "/internal/v1/powervs/locations/tag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1PowervsLocationsTagDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InternalV1PowervsLocationsTagDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.powervs.locations.tag.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InternalV1PowervsLocationsTagPost adds a power satellite tag
*/
func (a *Client) InternalV1PowervsLocationsTagPost(params *InternalV1PowervsLocationsTagPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTagPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1PowervsLocationsTagPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.powervs.locations.tag.post",
		Method:             "POST",
		PathPattern:        "/internal/v1/powervs/locations/tag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1PowervsLocationsTagPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InternalV1PowervsLocationsTagPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.powervs.locations.tag.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InternalV1PowervsLocationsTransitgatewayGet gets list of p e r enabled power v s locations
*/
func (a *Client) InternalV1PowervsLocationsTransitgatewayGet(params *InternalV1PowervsLocationsTransitgatewayGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTransitgatewayGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1PowervsLocationsTransitgatewayGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.powervs.locations.transitgateway.get",
		Method:             "GET",
		PathPattern:        "/internal/v1/powervs/locations/transit-gateway",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1PowervsLocationsTransitgatewayGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InternalV1PowervsLocationsTransitgatewayGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.powervs.locations.transitgateway.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
