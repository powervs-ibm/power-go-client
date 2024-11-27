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

// V1NetworkAddressGroupsPostReader is a Reader for the V1NetworkAddressGroupsPost structure.
type V1NetworkAddressGroupsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NetworkAddressGroupsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NetworkAddressGroupsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewV1NetworkAddressGroupsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1NetworkAddressGroupsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1NetworkAddressGroupsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1NetworkAddressGroupsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1NetworkAddressGroupsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1NetworkAddressGroupsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewV1NetworkAddressGroupsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1NetworkAddressGroupsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/network-address-groups] v1.networkAddressGroups.post", response, response.Code())
	}
}

// NewV1NetworkAddressGroupsPostOK creates a V1NetworkAddressGroupsPostOK with default headers values
func NewV1NetworkAddressGroupsPostOK() *V1NetworkAddressGroupsPostOK {
	return &V1NetworkAddressGroupsPostOK{}
}

/*
V1NetworkAddressGroupsPostOK describes a response with status code 200, with default header values.

OK
*/
type V1NetworkAddressGroupsPostOK struct {
	Payload *models.NetworkAddressGroup
}

// IsSuccess returns true when this v1 network address groups post o k response has a 2xx status code
func (o *V1NetworkAddressGroupsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 network address groups post o k response has a 3xx status code
func (o *V1NetworkAddressGroupsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post o k response has a 4xx status code
func (o *V1NetworkAddressGroupsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network address groups post o k response has a 5xx status code
func (o *V1NetworkAddressGroupsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post o k response a status code equal to that given
func (o *V1NetworkAddressGroupsPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 network address groups post o k response
func (o *V1NetworkAddressGroupsPostOK) Code() int {
	return 200
}

func (o *V1NetworkAddressGroupsPostOK) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostOK  %+v", 200, o.Payload)
}

func (o *V1NetworkAddressGroupsPostOK) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostOK  %+v", 200, o.Payload)
}

func (o *V1NetworkAddressGroupsPostOK) GetPayload() *models.NetworkAddressGroup {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkAddressGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostCreated creates a V1NetworkAddressGroupsPostCreated with default headers values
func NewV1NetworkAddressGroupsPostCreated() *V1NetworkAddressGroupsPostCreated {
	return &V1NetworkAddressGroupsPostCreated{}
}

/*
V1NetworkAddressGroupsPostCreated describes a response with status code 201, with default header values.

Created
*/
type V1NetworkAddressGroupsPostCreated struct {
	Payload *models.NetworkAddressGroup
}

// IsSuccess returns true when this v1 network address groups post created response has a 2xx status code
func (o *V1NetworkAddressGroupsPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 network address groups post created response has a 3xx status code
func (o *V1NetworkAddressGroupsPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post created response has a 4xx status code
func (o *V1NetworkAddressGroupsPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network address groups post created response has a 5xx status code
func (o *V1NetworkAddressGroupsPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post created response a status code equal to that given
func (o *V1NetworkAddressGroupsPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the v1 network address groups post created response
func (o *V1NetworkAddressGroupsPostCreated) Code() int {
	return 201
}

func (o *V1NetworkAddressGroupsPostCreated) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostCreated  %+v", 201, o.Payload)
}

func (o *V1NetworkAddressGroupsPostCreated) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostCreated  %+v", 201, o.Payload)
}

func (o *V1NetworkAddressGroupsPostCreated) GetPayload() *models.NetworkAddressGroup {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkAddressGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostBadRequest creates a V1NetworkAddressGroupsPostBadRequest with default headers values
func NewV1NetworkAddressGroupsPostBadRequest() *V1NetworkAddressGroupsPostBadRequest {
	return &V1NetworkAddressGroupsPostBadRequest{}
}

/*
V1NetworkAddressGroupsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1NetworkAddressGroupsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups post bad request response has a 2xx status code
func (o *V1NetworkAddressGroupsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups post bad request response has a 3xx status code
func (o *V1NetworkAddressGroupsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post bad request response has a 4xx status code
func (o *V1NetworkAddressGroupsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups post bad request response has a 5xx status code
func (o *V1NetworkAddressGroupsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post bad request response a status code equal to that given
func (o *V1NetworkAddressGroupsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 network address groups post bad request response
func (o *V1NetworkAddressGroupsPostBadRequest) Code() int {
	return 400
}

func (o *V1NetworkAddressGroupsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkAddressGroupsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkAddressGroupsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostUnauthorized creates a V1NetworkAddressGroupsPostUnauthorized with default headers values
func NewV1NetworkAddressGroupsPostUnauthorized() *V1NetworkAddressGroupsPostUnauthorized {
	return &V1NetworkAddressGroupsPostUnauthorized{}
}

/*
V1NetworkAddressGroupsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1NetworkAddressGroupsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups post unauthorized response has a 2xx status code
func (o *V1NetworkAddressGroupsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups post unauthorized response has a 3xx status code
func (o *V1NetworkAddressGroupsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post unauthorized response has a 4xx status code
func (o *V1NetworkAddressGroupsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups post unauthorized response has a 5xx status code
func (o *V1NetworkAddressGroupsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post unauthorized response a status code equal to that given
func (o *V1NetworkAddressGroupsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 network address groups post unauthorized response
func (o *V1NetworkAddressGroupsPostUnauthorized) Code() int {
	return 401
}

func (o *V1NetworkAddressGroupsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkAddressGroupsPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkAddressGroupsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostForbidden creates a V1NetworkAddressGroupsPostForbidden with default headers values
func NewV1NetworkAddressGroupsPostForbidden() *V1NetworkAddressGroupsPostForbidden {
	return &V1NetworkAddressGroupsPostForbidden{}
}

/*
V1NetworkAddressGroupsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1NetworkAddressGroupsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups post forbidden response has a 2xx status code
func (o *V1NetworkAddressGroupsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups post forbidden response has a 3xx status code
func (o *V1NetworkAddressGroupsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post forbidden response has a 4xx status code
func (o *V1NetworkAddressGroupsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups post forbidden response has a 5xx status code
func (o *V1NetworkAddressGroupsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post forbidden response a status code equal to that given
func (o *V1NetworkAddressGroupsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 network address groups post forbidden response
func (o *V1NetworkAddressGroupsPostForbidden) Code() int {
	return 403
}

func (o *V1NetworkAddressGroupsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkAddressGroupsPostForbidden) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkAddressGroupsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostNotFound creates a V1NetworkAddressGroupsPostNotFound with default headers values
func NewV1NetworkAddressGroupsPostNotFound() *V1NetworkAddressGroupsPostNotFound {
	return &V1NetworkAddressGroupsPostNotFound{}
}

/*
V1NetworkAddressGroupsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1NetworkAddressGroupsPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups post not found response has a 2xx status code
func (o *V1NetworkAddressGroupsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups post not found response has a 3xx status code
func (o *V1NetworkAddressGroupsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post not found response has a 4xx status code
func (o *V1NetworkAddressGroupsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups post not found response has a 5xx status code
func (o *V1NetworkAddressGroupsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post not found response a status code equal to that given
func (o *V1NetworkAddressGroupsPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 network address groups post not found response
func (o *V1NetworkAddressGroupsPostNotFound) Code() int {
	return 404
}

func (o *V1NetworkAddressGroupsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkAddressGroupsPostNotFound) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkAddressGroupsPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostConflict creates a V1NetworkAddressGroupsPostConflict with default headers values
func NewV1NetworkAddressGroupsPostConflict() *V1NetworkAddressGroupsPostConflict {
	return &V1NetworkAddressGroupsPostConflict{}
}

/*
V1NetworkAddressGroupsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1NetworkAddressGroupsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups post conflict response has a 2xx status code
func (o *V1NetworkAddressGroupsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups post conflict response has a 3xx status code
func (o *V1NetworkAddressGroupsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post conflict response has a 4xx status code
func (o *V1NetworkAddressGroupsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups post conflict response has a 5xx status code
func (o *V1NetworkAddressGroupsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post conflict response a status code equal to that given
func (o *V1NetworkAddressGroupsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 network address groups post conflict response
func (o *V1NetworkAddressGroupsPostConflict) Code() int {
	return 409
}

func (o *V1NetworkAddressGroupsPostConflict) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostConflict  %+v", 409, o.Payload)
}

func (o *V1NetworkAddressGroupsPostConflict) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostConflict  %+v", 409, o.Payload)
}

func (o *V1NetworkAddressGroupsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostUnprocessableEntity creates a V1NetworkAddressGroupsPostUnprocessableEntity with default headers values
func NewV1NetworkAddressGroupsPostUnprocessableEntity() *V1NetworkAddressGroupsPostUnprocessableEntity {
	return &V1NetworkAddressGroupsPostUnprocessableEntity{}
}

/*
V1NetworkAddressGroupsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type V1NetworkAddressGroupsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups post unprocessable entity response has a 2xx status code
func (o *V1NetworkAddressGroupsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups post unprocessable entity response has a 3xx status code
func (o *V1NetworkAddressGroupsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post unprocessable entity response has a 4xx status code
func (o *V1NetworkAddressGroupsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network address groups post unprocessable entity response has a 5xx status code
func (o *V1NetworkAddressGroupsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network address groups post unprocessable entity response a status code equal to that given
func (o *V1NetworkAddressGroupsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the v1 network address groups post unprocessable entity response
func (o *V1NetworkAddressGroupsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *V1NetworkAddressGroupsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *V1NetworkAddressGroupsPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *V1NetworkAddressGroupsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkAddressGroupsPostInternalServerError creates a V1NetworkAddressGroupsPostInternalServerError with default headers values
func NewV1NetworkAddressGroupsPostInternalServerError() *V1NetworkAddressGroupsPostInternalServerError {
	return &V1NetworkAddressGroupsPostInternalServerError{}
}

/*
V1NetworkAddressGroupsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1NetworkAddressGroupsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network address groups post internal server error response has a 2xx status code
func (o *V1NetworkAddressGroupsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network address groups post internal server error response has a 3xx status code
func (o *V1NetworkAddressGroupsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network address groups post internal server error response has a 4xx status code
func (o *V1NetworkAddressGroupsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network address groups post internal server error response has a 5xx status code
func (o *V1NetworkAddressGroupsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 network address groups post internal server error response a status code equal to that given
func (o *V1NetworkAddressGroupsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 network address groups post internal server error response
func (o *V1NetworkAddressGroupsPostInternalServerError) Code() int {
	return 500
}

func (o *V1NetworkAddressGroupsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkAddressGroupsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/network-address-groups][%d] v1NetworkAddressGroupsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkAddressGroupsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkAddressGroupsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
