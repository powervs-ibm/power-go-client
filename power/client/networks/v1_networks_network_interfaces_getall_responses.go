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

// V1NetworksNetworkInterfacesGetallReader is a Reader for the V1NetworksNetworkInterfacesGetall structure.
type V1NetworksNetworkInterfacesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NetworksNetworkInterfacesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NetworksNetworkInterfacesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1NetworksNetworkInterfacesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1NetworksNetworkInterfacesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1NetworksNetworkInterfacesGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1NetworksNetworkInterfacesGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1NetworksNetworkInterfacesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/networks/{network_id}/network-interfaces] v1.networks.network-interfaces.getall", response, response.Code())
	}
}

// NewV1NetworksNetworkInterfacesGetallOK creates a V1NetworksNetworkInterfacesGetallOK with default headers values
func NewV1NetworksNetworkInterfacesGetallOK() *V1NetworksNetworkInterfacesGetallOK {
	return &V1NetworksNetworkInterfacesGetallOK{}
}

/*
V1NetworksNetworkInterfacesGetallOK describes a response with status code 200, with default header values.

OK
*/
type V1NetworksNetworkInterfacesGetallOK struct {
	Payload *models.NetworkInterfaces
}

// IsSuccess returns true when this v1 networks network interfaces getall o k response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 networks network interfaces getall o k response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces getall o k response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 networks network interfaces getall o k response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces getall o k response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 networks network interfaces getall o k response
func (o *V1NetworksNetworkInterfacesGetallOK) Code() int {
	return 200
}

func (o *V1NetworksNetworkInterfacesGetallOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallOK  %+v", 200, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallOK) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallOK  %+v", 200, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallOK) GetPayload() *models.NetworkInterfaces {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkInterfaces)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetallBadRequest creates a V1NetworksNetworkInterfacesGetallBadRequest with default headers values
func NewV1NetworksNetworkInterfacesGetallBadRequest() *V1NetworksNetworkInterfacesGetallBadRequest {
	return &V1NetworksNetworkInterfacesGetallBadRequest{}
}

/*
V1NetworksNetworkInterfacesGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1NetworksNetworkInterfacesGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces getall bad request response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces getall bad request response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces getall bad request response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces getall bad request response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces getall bad request response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 networks network interfaces getall bad request response
func (o *V1NetworksNetworkInterfacesGetallBadRequest) Code() int {
	return 400
}

func (o *V1NetworksNetworkInterfacesGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetallUnauthorized creates a V1NetworksNetworkInterfacesGetallUnauthorized with default headers values
func NewV1NetworksNetworkInterfacesGetallUnauthorized() *V1NetworksNetworkInterfacesGetallUnauthorized {
	return &V1NetworksNetworkInterfacesGetallUnauthorized{}
}

/*
V1NetworksNetworkInterfacesGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1NetworksNetworkInterfacesGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces getall unauthorized response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces getall unauthorized response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces getall unauthorized response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces getall unauthorized response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces getall unauthorized response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 networks network interfaces getall unauthorized response
func (o *V1NetworksNetworkInterfacesGetallUnauthorized) Code() int {
	return 401
}

func (o *V1NetworksNetworkInterfacesGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetallForbidden creates a V1NetworksNetworkInterfacesGetallForbidden with default headers values
func NewV1NetworksNetworkInterfacesGetallForbidden() *V1NetworksNetworkInterfacesGetallForbidden {
	return &V1NetworksNetworkInterfacesGetallForbidden{}
}

/*
V1NetworksNetworkInterfacesGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1NetworksNetworkInterfacesGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces getall forbidden response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces getall forbidden response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces getall forbidden response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces getall forbidden response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces getall forbidden response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 networks network interfaces getall forbidden response
func (o *V1NetworksNetworkInterfacesGetallForbidden) Code() int {
	return 403
}

func (o *V1NetworksNetworkInterfacesGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallForbidden) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetallNotFound creates a V1NetworksNetworkInterfacesGetallNotFound with default headers values
func NewV1NetworksNetworkInterfacesGetallNotFound() *V1NetworksNetworkInterfacesGetallNotFound {
	return &V1NetworksNetworkInterfacesGetallNotFound{}
}

/*
V1NetworksNetworkInterfacesGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1NetworksNetworkInterfacesGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces getall not found response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces getall not found response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces getall not found response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 networks network interfaces getall not found response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 networks network interfaces getall not found response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 networks network interfaces getall not found response
func (o *V1NetworksNetworkInterfacesGetallNotFound) Code() int {
	return 404
}

func (o *V1NetworksNetworkInterfacesGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallNotFound) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworksNetworkInterfacesGetallInternalServerError creates a V1NetworksNetworkInterfacesGetallInternalServerError with default headers values
func NewV1NetworksNetworkInterfacesGetallInternalServerError() *V1NetworksNetworkInterfacesGetallInternalServerError {
	return &V1NetworksNetworkInterfacesGetallInternalServerError{}
}

/*
V1NetworksNetworkInterfacesGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1NetworksNetworkInterfacesGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 networks network interfaces getall internal server error response has a 2xx status code
func (o *V1NetworksNetworkInterfacesGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 networks network interfaces getall internal server error response has a 3xx status code
func (o *V1NetworksNetworkInterfacesGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 networks network interfaces getall internal server error response has a 4xx status code
func (o *V1NetworksNetworkInterfacesGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 networks network interfaces getall internal server error response has a 5xx status code
func (o *V1NetworksNetworkInterfacesGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 networks network interfaces getall internal server error response a status code equal to that given
func (o *V1NetworksNetworkInterfacesGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 networks network interfaces getall internal server error response
func (o *V1NetworksNetworkInterfacesGetallInternalServerError) Code() int {
	return 500
}

func (o *V1NetworksNetworkInterfacesGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/networks/{network_id}/network-interfaces][%d] v1NetworksNetworkInterfacesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworksNetworkInterfacesGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworksNetworkInterfacesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
