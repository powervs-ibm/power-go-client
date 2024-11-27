// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1NetworksNetworkInterfacesGetReader is a Reader for the V1NetworksNetworkInterfacesGet structure.
type V1NetworksNetworkInterfacesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NetworksNetworkInterfacesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NetworksNetworkInterfacesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1NetworksNetworkInterfacesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1NetworksNetworkInterfacesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1NetworksNetworkInterfacesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1NetworksNetworkInterfacesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1NetworksNetworkInterfacesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}] v1.networks.network-interfaces.get", response, response.Code())
	}
}

// NewV1NetworksNetworkInterfacesGetOK creates a V1NetworksNetworkInterfacesGetOK with default headers values
func NewV1NetworksNetworkInterfacesGetOK() *V1NetworksNetworkInterfacesGetOK {
	return &V1NetworksNetworkInterfacesGetOK{}
}

/*
V1NetworksNetworkInterfacesGetOK describes a response with status code 200, with default header values.

OK
*/
type V1NetworksNetworkInterfacesGetOK struct {
	Payload *models.NetworkInterface
}

// IsSuccess returns true when this v1 networks network interfaces get o k response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 networks network interfaces get o k response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces get o k response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 networks network interfaces get o k response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces get o k response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 networks network interfaces get o k response
func (o *V1NetworksNetworkInterfacesGetOK) Code() int {
	return 200
}

func (o *V1NetworksNetworkInterfacesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetOK  %+v", 200, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetOK) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetOK  %+v", 200, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetOK) GetPayload() *models.NetworkInterface {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetBadRequest creates a V1NetworksNetworkInterfacesGetBadRequest with default headers values
func NewV1NetworksNetworkInterfacesGetBadRequest() *V1NetworksNetworkInterfacesGetBadRequest {
	return &V1NetworksNetworkInterfacesGetBadRequest{}
}

/*
V1NetworksNetworkInterfacesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1NetworksNetworkInterfacesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces get bad request response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces get bad request response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces get bad request response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces get bad request response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces get bad request response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 networks network interfaces get bad request response
func (o *V1NetworksNetworkInterfacesGetBadRequest) Code() int {
	return 400
}

func (o *V1NetworksNetworkInterfacesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetUnauthorized creates a V1NetworksNetworkInterfacesGetUnauthorized with default headers values
func NewV1NetworksNetworkInterfacesGetUnauthorized() *V1NetworksNetworkInterfacesGetUnauthorized {
	return &V1NetworksNetworkInterfacesGetUnauthorized{}
}

/*
V1NetworksNetworkInterfacesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1NetworksNetworkInterfacesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces get unauthorized response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces get unauthorized response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces get unauthorized response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces get unauthorized response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces get unauthorized response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 networks network interfaces get unauthorized response
func (o *V1NetworksNetworkInterfacesGetUnauthorized) Code() int {
	return 401
}

func (o *V1NetworksNetworkInterfacesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetForbidden creates a V1NetworksNetworkInterfacesGetForbidden with default headers values
func NewV1NetworksNetworkInterfacesGetForbidden() *V1NetworksNetworkInterfacesGetForbidden {
	return &V1NetworksNetworkInterfacesGetForbidden{}
}

/*
V1NetworksNetworkInterfacesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1NetworksNetworkInterfacesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces get forbidden response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces get forbidden response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces get forbidden response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces get forbidden response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces get forbidden response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 networks network interfaces get forbidden response
func (o *V1NetworksNetworkInterfacesGetForbidden) Code() int {
	return 403
}

func (o *V1NetworksNetworkInterfacesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetForbidden) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetNotFound creates a V1NetworksNetworkInterfacesGetNotFound with default headers values
func NewV1NetworksNetworkInterfacesGetNotFound() *V1NetworksNetworkInterfacesGetNotFound {
	return &V1NetworksNetworkInterfacesGetNotFound{}
}

/*
V1NetworksNetworkInterfacesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1NetworksNetworkInterfacesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces get not found response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces get not found response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces get not found response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces get not found response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces get not found response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 networks network interfaces get not found response
func (o *V1NetworksNetworkInterfacesGetNotFound) Code() int {
	return 404
}

func (o *V1NetworksNetworkInterfacesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetNotFound) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetInternalServerError creates a V1NetworksNetworkInterfacesGetInternalServerError with default headers values
func NewV1NetworksNetworkInterfacesGetInternalServerError() *V1NetworksNetworkInterfacesGetInternalServerError {
	return &V1NetworksNetworkInterfacesGetInternalServerError{}
}

/*
V1NetworksNetworkInterfacesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1NetworksNetworkInterfacesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces get internal server error response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces get internal server error response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces get internal server error response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 networks network interfaces get internal server error response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 networks network interfaces get internal server error response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 networks network interfaces get internal server error response
func (o *V1NetworksNetworkInterfacesGetInternalServerError) Code() int {
	return 500
}

func (o *V1NetworksNetworkInterfacesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces/{network_interface_id}][%d] v1NetworksNetworkInterfacesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
