// Code generated by go-swagger; DO NOT EDIT.

package datacenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1DatacentersPrivateGetallReader is a Reader for the V1DatacentersPrivateGetall structure.
type V1DatacentersPrivateGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DatacentersPrivateGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DatacentersPrivateGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1DatacentersPrivateGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1DatacentersPrivateGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1DatacentersPrivateGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1DatacentersPrivateGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/datacenters/private] v1.datacentersPrivate.getall", response, response.Code())
	}
}

// NewV1DatacentersPrivateGetallOK creates a V1DatacentersPrivateGetallOK with default headers values
func NewV1DatacentersPrivateGetallOK() *V1DatacentersPrivateGetallOK {
	return &V1DatacentersPrivateGetallOK{}
}

/*
V1DatacentersPrivateGetallOK describes a response with status code 200, with default header values.

OK
*/
type V1DatacentersPrivateGetallOK struct {
	Payload *models.Datacenters
}

// IsSuccess returns true when this v1 datacenters private getall o k response has a 2xx status code
func (o *V1DatacentersPrivateGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 datacenters private getall o k response has a 3xx status code
func (o *V1DatacentersPrivateGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters private getall o k response has a 4xx status code
func (o *V1DatacentersPrivateGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 datacenters private getall o k response has a 5xx status code
func (o *V1DatacentersPrivateGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters private getall o k response a status code equal to that given
func (o *V1DatacentersPrivateGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 datacenters private getall o k response
func (o *V1DatacentersPrivateGetallOK) Code() int {
	return 200
}

func (o *V1DatacentersPrivateGetallOK) Error() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallOK  %+v", 200, o.Payload)
}

func (o *V1DatacentersPrivateGetallOK) String() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallOK  %+v", 200, o.Payload)
}

func (o *V1DatacentersPrivateGetallOK) GetPayload() *models.Datacenters {
	return o.Payload
}

func (o *V1DatacentersPrivateGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenters)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersPrivateGetallBadRequest creates a V1DatacentersPrivateGetallBadRequest with default headers values
func NewV1DatacentersPrivateGetallBadRequest() *V1DatacentersPrivateGetallBadRequest {
	return &V1DatacentersPrivateGetallBadRequest{}
}

/*
V1DatacentersPrivateGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1DatacentersPrivateGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters private getall bad request response has a 2xx status code
func (o *V1DatacentersPrivateGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters private getall bad request response has a 3xx status code
func (o *V1DatacentersPrivateGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters private getall bad request response has a 4xx status code
func (o *V1DatacentersPrivateGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 datacenters private getall bad request response has a 5xx status code
func (o *V1DatacentersPrivateGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters private getall bad request response a status code equal to that given
func (o *V1DatacentersPrivateGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 datacenters private getall bad request response
func (o *V1DatacentersPrivateGetallBadRequest) Code() int {
	return 400
}

func (o *V1DatacentersPrivateGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallBadRequest  %+v", 400, o.Payload)
}

func (o *V1DatacentersPrivateGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallBadRequest  %+v", 400, o.Payload)
}

func (o *V1DatacentersPrivateGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersPrivateGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersPrivateGetallUnauthorized creates a V1DatacentersPrivateGetallUnauthorized with default headers values
func NewV1DatacentersPrivateGetallUnauthorized() *V1DatacentersPrivateGetallUnauthorized {
	return &V1DatacentersPrivateGetallUnauthorized{}
}

/*
V1DatacentersPrivateGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1DatacentersPrivateGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters private getall unauthorized response has a 2xx status code
func (o *V1DatacentersPrivateGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters private getall unauthorized response has a 3xx status code
func (o *V1DatacentersPrivateGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters private getall unauthorized response has a 4xx status code
func (o *V1DatacentersPrivateGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 datacenters private getall unauthorized response has a 5xx status code
func (o *V1DatacentersPrivateGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters private getall unauthorized response a status code equal to that given
func (o *V1DatacentersPrivateGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 datacenters private getall unauthorized response
func (o *V1DatacentersPrivateGetallUnauthorized) Code() int {
	return 401
}

func (o *V1DatacentersPrivateGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *V1DatacentersPrivateGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *V1DatacentersPrivateGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersPrivateGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersPrivateGetallForbidden creates a V1DatacentersPrivateGetallForbidden with default headers values
func NewV1DatacentersPrivateGetallForbidden() *V1DatacentersPrivateGetallForbidden {
	return &V1DatacentersPrivateGetallForbidden{}
}

/*
V1DatacentersPrivateGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1DatacentersPrivateGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters private getall forbidden response has a 2xx status code
func (o *V1DatacentersPrivateGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters private getall forbidden response has a 3xx status code
func (o *V1DatacentersPrivateGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters private getall forbidden response has a 4xx status code
func (o *V1DatacentersPrivateGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 datacenters private getall forbidden response has a 5xx status code
func (o *V1DatacentersPrivateGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters private getall forbidden response a status code equal to that given
func (o *V1DatacentersPrivateGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 datacenters private getall forbidden response
func (o *V1DatacentersPrivateGetallForbidden) Code() int {
	return 403
}

func (o *V1DatacentersPrivateGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallForbidden  %+v", 403, o.Payload)
}

func (o *V1DatacentersPrivateGetallForbidden) String() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallForbidden  %+v", 403, o.Payload)
}

func (o *V1DatacentersPrivateGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersPrivateGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersPrivateGetallInternalServerError creates a V1DatacentersPrivateGetallInternalServerError with default headers values
func NewV1DatacentersPrivateGetallInternalServerError() *V1DatacentersPrivateGetallInternalServerError {
	return &V1DatacentersPrivateGetallInternalServerError{}
}

/*
V1DatacentersPrivateGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1DatacentersPrivateGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters private getall internal server error response has a 2xx status code
func (o *V1DatacentersPrivateGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters private getall internal server error response has a 3xx status code
func (o *V1DatacentersPrivateGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters private getall internal server error response has a 4xx status code
func (o *V1DatacentersPrivateGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 datacenters private getall internal server error response has a 5xx status code
func (o *V1DatacentersPrivateGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 datacenters private getall internal server error response a status code equal to that given
func (o *V1DatacentersPrivateGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 datacenters private getall internal server error response
func (o *V1DatacentersPrivateGetallInternalServerError) Code() int {
	return 500
}

func (o *V1DatacentersPrivateGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *V1DatacentersPrivateGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/datacenters/private][%d] v1DatacentersPrivateGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *V1DatacentersPrivateGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersPrivateGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
