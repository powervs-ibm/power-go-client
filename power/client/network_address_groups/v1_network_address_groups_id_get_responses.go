// Code generated by go-swagger; DO NOT EDIT.

package network_address_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1NetworkAddressGroupsIDGetReader is a Reader for the V1NetworkAddressGroupsIDGet structure.
type V1NetworkAddressGroupsIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NetworkAddressGroupsIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NetworkAddressGroupsIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1NetworkAddressGroupsIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1NetworkAddressGroupsIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1NetworkAddressGroupsIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1NetworkAddressGroupsIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1NetworkAddressGroupsIDGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/network-address-groups/{network_address_group_id}] v1.networkAddressGroups.id.get", response, response.Code())
	}
}

// NewV1NetworkAddressGroupsIDGetOK creates a V1NetworkAddressGroupsIDGetOK with default headers values
func NewV1NetworkAddressGroupsIDGetOK() *V1NetworkAddressGroupsIDGetOK {
	return &V1NetworkAddressGroupsIDGetOK{}
}

/*
V1NetworkAddressGroupsIDGetOK describes a response with status code 200, with default header values.

OK
*/
type V1NetworkAddressGroupsIDGetOK struct {
	Payload *models.NetworkAddressGroup
}

// IsSuccess returns true when this v1 network address groups Id get o k response has a 2xx status code
func (o *V1NetworkAddressGroupsIDGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 network address groups Id get o k response has a 3xx status code
func (o *V1NetworkAddressGroupsIDGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups Id get o k response has a 4xx status code
func (o *V1NetworkAddressGroupsIDGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network address groups Id get o k response has a 5xx status code
func (o *V1NetworkAddressGroupsIDGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups Id get o k response a status code equal to that given
func (o *V1NetworkAddressGroupsIDGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 network address groups Id get o k response
func (o *V1NetworkAddressGroupsIDGetOK) Code() int {
	return 200
}

func (o *V1NetworkAddressGroupsIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetOK  %+v", 200, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetOK) String() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetOK  %+v", 200, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetOK) GetPayload() *models.NetworkAddressGroup {
	return o.Payload
}

func (o *V1NetworkAddressGroupsIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkAddressGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsIDGetBadRequest creates a V1NetworkAddressGroupsIDGetBadRequest with default headers values
func NewV1NetworkAddressGroupsIDGetBadRequest() *V1NetworkAddressGroupsIDGetBadRequest {
	return &V1NetworkAddressGroupsIDGetBadRequest{}
}

/*
V1NetworkAddressGroupsIDGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1NetworkAddressGroupsIDGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups Id get bad request response has a 2xx status code
func (o *V1NetworkAddressGroupsIDGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups Id get bad request response has a 3xx status code
func (o *V1NetworkAddressGroupsIDGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups Id get bad request response has a 4xx status code
func (o *V1NetworkAddressGroupsIDGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups Id get bad request response has a 5xx status code
func (o *V1NetworkAddressGroupsIDGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups Id get bad request response a status code equal to that given
func (o *V1NetworkAddressGroupsIDGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 network address groups Id get bad request response
func (o *V1NetworkAddressGroupsIDGetBadRequest) Code() int {
	return 400
}

func (o *V1NetworkAddressGroupsIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsIDGetUnauthorized creates a V1NetworkAddressGroupsIDGetUnauthorized with default headers values
func NewV1NetworkAddressGroupsIDGetUnauthorized() *V1NetworkAddressGroupsIDGetUnauthorized {
	return &V1NetworkAddressGroupsIDGetUnauthorized{}
}

/*
V1NetworkAddressGroupsIDGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1NetworkAddressGroupsIDGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups Id get unauthorized response has a 2xx status code
func (o *V1NetworkAddressGroupsIDGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups Id get unauthorized response has a 3xx status code
func (o *V1NetworkAddressGroupsIDGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups Id get unauthorized response has a 4xx status code
func (o *V1NetworkAddressGroupsIDGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups Id get unauthorized response has a 5xx status code
func (o *V1NetworkAddressGroupsIDGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups Id get unauthorized response a status code equal to that given
func (o *V1NetworkAddressGroupsIDGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 network address groups Id get unauthorized response
func (o *V1NetworkAddressGroupsIDGetUnauthorized) Code() int {
	return 401
}

func (o *V1NetworkAddressGroupsIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsIDGetForbidden creates a V1NetworkAddressGroupsIDGetForbidden with default headers values
func NewV1NetworkAddressGroupsIDGetForbidden() *V1NetworkAddressGroupsIDGetForbidden {
	return &V1NetworkAddressGroupsIDGetForbidden{}
}

/*
V1NetworkAddressGroupsIDGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1NetworkAddressGroupsIDGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups Id get forbidden response has a 2xx status code
func (o *V1NetworkAddressGroupsIDGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups Id get forbidden response has a 3xx status code
func (o *V1NetworkAddressGroupsIDGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups Id get forbidden response has a 4xx status code
func (o *V1NetworkAddressGroupsIDGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups Id get forbidden response has a 5xx status code
func (o *V1NetworkAddressGroupsIDGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups Id get forbidden response a status code equal to that given
func (o *V1NetworkAddressGroupsIDGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 network address groups Id get forbidden response
func (o *V1NetworkAddressGroupsIDGetForbidden) Code() int {
	return 403
}

func (o *V1NetworkAddressGroupsIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetForbidden) String() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsIDGetNotFound creates a V1NetworkAddressGroupsIDGetNotFound with default headers values
func NewV1NetworkAddressGroupsIDGetNotFound() *V1NetworkAddressGroupsIDGetNotFound {
	return &V1NetworkAddressGroupsIDGetNotFound{}
}

/*
V1NetworkAddressGroupsIDGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1NetworkAddressGroupsIDGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups Id get not found response has a 2xx status code
func (o *V1NetworkAddressGroupsIDGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups Id get not found response has a 3xx status code
func (o *V1NetworkAddressGroupsIDGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups Id get not found response has a 4xx status code
func (o *V1NetworkAddressGroupsIDGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups Id get not found response has a 5xx status code
func (o *V1NetworkAddressGroupsIDGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups Id get not found response a status code equal to that given
func (o *V1NetworkAddressGroupsIDGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 network address groups Id get not found response
func (o *V1NetworkAddressGroupsIDGetNotFound) Code() int {
	return 404
}

func (o *V1NetworkAddressGroupsIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetNotFound) String() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsIDGetInternalServerError creates a V1NetworkAddressGroupsIDGetInternalServerError with default headers values
func NewV1NetworkAddressGroupsIDGetInternalServerError() *V1NetworkAddressGroupsIDGetInternalServerError {
	return &V1NetworkAddressGroupsIDGetInternalServerError{}
}

/*
V1NetworkAddressGroupsIDGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1NetworkAddressGroupsIDGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups Id get internal server error response has a 2xx status code
func (o *V1NetworkAddressGroupsIDGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups Id get internal server error response has a 3xx status code
func (o *V1NetworkAddressGroupsIDGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups Id get internal server error response has a 4xx status code
func (o *V1NetworkAddressGroupsIDGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network address groups Id get internal server error response has a 5xx status code
func (o *V1NetworkAddressGroupsIDGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 network address groups Id get internal server error response a status code equal to that given
func (o *V1NetworkAddressGroupsIDGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 network address groups Id get internal server error response
func (o *V1NetworkAddressGroupsIDGetInternalServerError) Code() int {
	return 500
}

func (o *V1NetworkAddressGroupsIDGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/network-address-groups/{network_address_group_id}][%d] v1NetworkAddressGroupsIdGetInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkAddressGroupsIDGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsIDGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
