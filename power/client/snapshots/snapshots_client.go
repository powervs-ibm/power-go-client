// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new snapshots API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new snapshots API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new snapshots API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for snapshots API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	V1SnapshotsGet(params *V1SnapshotsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1SnapshotsGetOK, error)

	V1SnapshotsGetall(params *V1SnapshotsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1SnapshotsGetallOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
V1SnapshotsGet gets the detail of a snapshot
*/
func (a *Client) V1SnapshotsGet(params *V1SnapshotsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1SnapshotsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1SnapshotsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.snapshots.get",
		Method:             "GET",
		PathPattern:        "/v1/snapshots/{snapshot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1SnapshotsGetReader{formats: a.formats},
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
	success, ok := result.(*V1SnapshotsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.snapshots.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
V1SnapshotsGetall gets a list of all the snapshots on a workspace
*/
func (a *Client) V1SnapshotsGetall(params *V1SnapshotsGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1SnapshotsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1SnapshotsGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.snapshots.getall",
		Method:             "GET",
		PathPattern:        "/v1/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1SnapshotsGetallReader{formats: a.formats},
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
	success, ok := result.(*V1SnapshotsGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.snapshots.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}