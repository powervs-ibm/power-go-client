// Code generated by go-swagger; DO NOT EDIT.

package network_peers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1NetworkPeersListReader is a Reader for the V1NetworkPeersList structure.
type V1NetworkPeersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NetworkPeersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NetworkPeersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1NetworkPeersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1NetworkPeersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1NetworkPeersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1NetworkPeersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1NetworkPeersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/network-peers] v1.networkPeers.list", response, response.Code())
	}
}

// NewV1NetworkPeersListOK creates a V1NetworkPeersListOK with default headers values
func NewV1NetworkPeersListOK() *V1NetworkPeersListOK {
	return &V1NetworkPeersListOK{}
}

/*
V1NetworkPeersListOK describes a response with status code 200, with default header values.

OK
*/
type V1NetworkPeersListOK struct {
	Payload *models.NetworkPeers
}

// IsSuccess returns true when this v1 network peers list o k response has a 2xx status code
func (o *V1NetworkPeersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 network peers list o k response has a 3xx status code
func (o *V1NetworkPeersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network peers list o k response has a 4xx status code
func (o *V1NetworkPeersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network peers list o k response has a 5xx status code
func (o *V1NetworkPeersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network peers list o k response a status code equal to that given
func (o *V1NetworkPeersListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 network peers list o k response
func (o *V1NetworkPeersListOK) Code() int {
	return 200
}

func (o *V1NetworkPeersListOK) Error() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListOK  %+v", 200, o.Payload)
}

func (o *V1NetworkPeersListOK) String() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListOK  %+v", 200, o.Payload)
}

func (o *V1NetworkPeersListOK) GetPayload() *models.NetworkPeers {
	return o.Payload
}

func (o *V1NetworkPeersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPeers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkPeersListBadRequest creates a V1NetworkPeersListBadRequest with default headers values
func NewV1NetworkPeersListBadRequest() *V1NetworkPeersListBadRequest {
	return &V1NetworkPeersListBadRequest{}
}

/*
V1NetworkPeersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1NetworkPeersListBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network peers list bad request response has a 2xx status code
func (o *V1NetworkPeersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network peers list bad request response has a 3xx status code
func (o *V1NetworkPeersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network peers list bad request response has a 4xx status code
func (o *V1NetworkPeersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network peers list bad request response has a 5xx status code
func (o *V1NetworkPeersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network peers list bad request response a status code equal to that given
func (o *V1NetworkPeersListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 network peers list bad request response
func (o *V1NetworkPeersListBadRequest) Code() int {
	return 400
}

func (o *V1NetworkPeersListBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkPeersListBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkPeersListBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkPeersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkPeersListUnauthorized creates a V1NetworkPeersListUnauthorized with default headers values
func NewV1NetworkPeersListUnauthorized() *V1NetworkPeersListUnauthorized {
	return &V1NetworkPeersListUnauthorized{}
}

/*
V1NetworkPeersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1NetworkPeersListUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network peers list unauthorized response has a 2xx status code
func (o *V1NetworkPeersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network peers list unauthorized response has a 3xx status code
func (o *V1NetworkPeersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network peers list unauthorized response has a 4xx status code
func (o *V1NetworkPeersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network peers list unauthorized response has a 5xx status code
func (o *V1NetworkPeersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network peers list unauthorized response a status code equal to that given
func (o *V1NetworkPeersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 network peers list unauthorized response
func (o *V1NetworkPeersListUnauthorized) Code() int {
	return 401
}

func (o *V1NetworkPeersListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkPeersListUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkPeersListUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkPeersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkPeersListForbidden creates a V1NetworkPeersListForbidden with default headers values
func NewV1NetworkPeersListForbidden() *V1NetworkPeersListForbidden {
	return &V1NetworkPeersListForbidden{}
}

/*
V1NetworkPeersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1NetworkPeersListForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network peers list forbidden response has a 2xx status code
func (o *V1NetworkPeersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network peers list forbidden response has a 3xx status code
func (o *V1NetworkPeersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network peers list forbidden response has a 4xx status code
func (o *V1NetworkPeersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network peers list forbidden response has a 5xx status code
func (o *V1NetworkPeersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network peers list forbidden response a status code equal to that given
func (o *V1NetworkPeersListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 network peers list forbidden response
func (o *V1NetworkPeersListForbidden) Code() int {
	return 403
}

func (o *V1NetworkPeersListForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkPeersListForbidden) String() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkPeersListForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkPeersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkPeersListNotFound creates a V1NetworkPeersListNotFound with default headers values
func NewV1NetworkPeersListNotFound() *V1NetworkPeersListNotFound {
	return &V1NetworkPeersListNotFound{}
}

/*
V1NetworkPeersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1NetworkPeersListNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network peers list not found response has a 2xx status code
func (o *V1NetworkPeersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network peers list not found response has a 3xx status code
func (o *V1NetworkPeersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network peers list not found response has a 4xx status code
func (o *V1NetworkPeersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network peers list not found response has a 5xx status code
func (o *V1NetworkPeersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network peers list not found response a status code equal to that given
func (o *V1NetworkPeersListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 network peers list not found response
func (o *V1NetworkPeersListNotFound) Code() int {
	return 404
}

func (o *V1NetworkPeersListNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkPeersListNotFound) String() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkPeersListNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkPeersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkPeersListInternalServerError creates a V1NetworkPeersListInternalServerError with default headers values
func NewV1NetworkPeersListInternalServerError() *V1NetworkPeersListInternalServerError {
	return &V1NetworkPeersListInternalServerError{}
}

/*
V1NetworkPeersListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1NetworkPeersListInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network peers list internal server error response has a 2xx status code
func (o *V1NetworkPeersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network peers list internal server error response has a 3xx status code
func (o *V1NetworkPeersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network peers list internal server error response has a 4xx status code
func (o *V1NetworkPeersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network peers list internal server error response has a 5xx status code
func (o *V1NetworkPeersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 network peers list internal server error response a status code equal to that given
func (o *V1NetworkPeersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 network peers list internal server error response
func (o *V1NetworkPeersListInternalServerError) Code() int {
	return 500
}

func (o *V1NetworkPeersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkPeersListInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/network-peers][%d] v1NetworkPeersListInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkPeersListInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkPeersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
