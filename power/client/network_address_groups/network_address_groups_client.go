// Code generated by go-swagger; DO NOT EDIT.

package network_address_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new network address groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network address groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	V1NetworkAddressGroupsGet(params *V1NetworkAddressGroupsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsGetOK, error)

	V1NetworkAddressGroupsIDDelete(params *V1NetworkAddressGroupsIDDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsIDDeleteOK, error)

	V1NetworkAddressGroupsIDGet(params *V1NetworkAddressGroupsIDGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsIDGetOK, error)

	V1NetworkAddressGroupsIDPut(params *V1NetworkAddressGroupsIDPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsIDPutOK, error)

	V1NetworkAddressGroupsMembersDelete(params *V1NetworkAddressGroupsMembersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsMembersDeleteOK, error)

	V1NetworkAddressGroupsMembersPost(params *V1NetworkAddressGroupsMembersPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsMembersPostOK, error)

	V1NetworkAddressGroupsPost(params *V1NetworkAddressGroupsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsPostOK, *V1NetworkAddressGroupsPostCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
V1NetworkAddressGroupsGet gets the list of network address groups for a workspace
*/
func (a *Client) V1NetworkAddressGroupsGet(params *V1NetworkAddressGroupsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1NetworkAddressGroupsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.networkAddressGroups.get",
		Method:             "GET",
		PathPattern:        "/v1/network-address-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1NetworkAddressGroupsGetReader{formats: a.formats},
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
	success, ok := result.(*V1NetworkAddressGroupsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.networkAddressGroups.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1NetworkAddressGroupsIDDelete deletes a network address group from a workspace
*/
func (a *Client) V1NetworkAddressGroupsIDDelete(params *V1NetworkAddressGroupsIDDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1NetworkAddressGroupsIDDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.networkAddressGroups.id.delete",
		Method:             "DELETE",
		PathPattern:        "/v1/network-address-groups/{network_address_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1NetworkAddressGroupsIDDeleteReader{formats: a.formats},
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
	success, ok := result.(*V1NetworkAddressGroupsIDDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.networkAddressGroups.id.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1NetworkAddressGroupsIDGet gets the detail of a network address group
*/
func (a *Client) V1NetworkAddressGroupsIDGet(params *V1NetworkAddressGroupsIDGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1NetworkAddressGroupsIDGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.networkAddressGroups.id.get",
		Method:             "GET",
		PathPattern:        "/v1/network-address-groups/{network_address_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1NetworkAddressGroupsIDGetReader{formats: a.formats},
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
	success, ok := result.(*V1NetworkAddressGroupsIDGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.networkAddressGroups.id.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1NetworkAddressGroupsIDPut updates a network address group
*/
func (a *Client) V1NetworkAddressGroupsIDPut(params *V1NetworkAddressGroupsIDPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1NetworkAddressGroupsIDPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.networkAddressGroups.id.put",
		Method:             "PUT",
		PathPattern:        "/v1/network-address-groups/{network_address_group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1NetworkAddressGroupsIDPutReader{formats: a.formats},
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
	success, ok := result.(*V1NetworkAddressGroupsIDPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.networkAddressGroups.id.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1NetworkAddressGroupsMembersDelete deletes the member from a network address group
*/
func (a *Client) V1NetworkAddressGroupsMembersDelete(params *V1NetworkAddressGroupsMembersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsMembersDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1NetworkAddressGroupsMembersDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.networkAddressGroups.members.delete",
		Method:             "DELETE",
		PathPattern:        "/v1/network-address-groups/{network_address_group_id}/members/{network_address_group_member_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1NetworkAddressGroupsMembersDeleteReader{formats: a.formats},
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
	success, ok := result.(*V1NetworkAddressGroupsMembersDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.networkAddressGroups.members.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1NetworkAddressGroupsMembersPost adds a member to a network address group
*/
func (a *Client) V1NetworkAddressGroupsMembersPost(params *V1NetworkAddressGroupsMembersPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsMembersPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1NetworkAddressGroupsMembersPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.networkAddressGroups.members.post",
		Method:             "POST",
		PathPattern:        "/v1/network-address-groups/{network_address_group_id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1NetworkAddressGroupsMembersPostReader{formats: a.formats},
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
	success, ok := result.(*V1NetworkAddressGroupsMembersPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.networkAddressGroups.members.post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1NetworkAddressGroupsPost creates a new network address group
*/
func (a *Client) V1NetworkAddressGroupsPost(params *V1NetworkAddressGroupsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1NetworkAddressGroupsPostOK, *V1NetworkAddressGroupsPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1NetworkAddressGroupsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.networkAddressGroups.post",
		Method:             "POST",
		PathPattern:        "/v1/network-address-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1NetworkAddressGroupsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *V1NetworkAddressGroupsPostOK:
		return value, nil, nil
	case *V1NetworkAddressGroupsPostCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for network_address_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
