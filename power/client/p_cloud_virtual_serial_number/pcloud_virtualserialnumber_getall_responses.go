// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_virtual_serial_number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVirtualserialnumberGetallReader is a Reader for the PcloudVirtualserialnumberGetall structure.
type PcloudVirtualserialnumberGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVirtualserialnumberGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVirtualserialnumberGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVirtualserialnumberGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVirtualserialnumberGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVirtualserialnumberGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVirtualserialnumberGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVirtualserialnumberGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/virtual-serial-number] pcloud.virtualserialnumber.getall", response, response.Code())
	}
}

// NewPcloudVirtualserialnumberGetallOK creates a PcloudVirtualserialnumberGetallOK with default headers values
func NewPcloudVirtualserialnumberGetallOK() *PcloudVirtualserialnumberGetallOK {
	return &PcloudVirtualserialnumberGetallOK{}
}

/*
PcloudVirtualserialnumberGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVirtualserialnumberGetallOK struct {
	Payload models.UnifiedVirtualSerialNumberList
}

// IsSuccess returns true when this pcloud virtualserialnumber getall o k response has a 2xx status code
func (o *PcloudVirtualserialnumberGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud virtualserialnumber getall o k response has a 3xx status code
func (o *PcloudVirtualserialnumberGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud virtualserialnumber getall o k response has a 4xx status code
func (o *PcloudVirtualserialnumberGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud virtualserialnumber getall o k response has a 5xx status code
func (o *PcloudVirtualserialnumberGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud virtualserialnumber getall o k response a status code equal to that given
func (o *PcloudVirtualserialnumberGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud virtualserialnumber getall o k response
func (o *PcloudVirtualserialnumberGetallOK) Code() int {
	return 200
}

func (o *PcloudVirtualserialnumberGetallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallOK %s", 200, payload)
}

func (o *PcloudVirtualserialnumberGetallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallOK %s", 200, payload)
}

func (o *PcloudVirtualserialnumberGetallOK) GetPayload() models.UnifiedVirtualSerialNumberList {
	return o.Payload
}

func (o *PcloudVirtualserialnumberGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVirtualserialnumberGetallBadRequest creates a PcloudVirtualserialnumberGetallBadRequest with default headers values
func NewPcloudVirtualserialnumberGetallBadRequest() *PcloudVirtualserialnumberGetallBadRequest {
	return &PcloudVirtualserialnumberGetallBadRequest{}
}

/*
PcloudVirtualserialnumberGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVirtualserialnumberGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud virtualserialnumber getall bad request response has a 2xx status code
func (o *PcloudVirtualserialnumberGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud virtualserialnumber getall bad request response has a 3xx status code
func (o *PcloudVirtualserialnumberGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud virtualserialnumber getall bad request response has a 4xx status code
func (o *PcloudVirtualserialnumberGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud virtualserialnumber getall bad request response has a 5xx status code
func (o *PcloudVirtualserialnumberGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud virtualserialnumber getall bad request response a status code equal to that given
func (o *PcloudVirtualserialnumberGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud virtualserialnumber getall bad request response
func (o *PcloudVirtualserialnumberGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudVirtualserialnumberGetallBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallBadRequest %s", 400, payload)
}

func (o *PcloudVirtualserialnumberGetallBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallBadRequest %s", 400, payload)
}

func (o *PcloudVirtualserialnumberGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVirtualserialnumberGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVirtualserialnumberGetallUnauthorized creates a PcloudVirtualserialnumberGetallUnauthorized with default headers values
func NewPcloudVirtualserialnumberGetallUnauthorized() *PcloudVirtualserialnumberGetallUnauthorized {
	return &PcloudVirtualserialnumberGetallUnauthorized{}
}

/*
PcloudVirtualserialnumberGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVirtualserialnumberGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud virtualserialnumber getall unauthorized response has a 2xx status code
func (o *PcloudVirtualserialnumberGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud virtualserialnumber getall unauthorized response has a 3xx status code
func (o *PcloudVirtualserialnumberGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud virtualserialnumber getall unauthorized response has a 4xx status code
func (o *PcloudVirtualserialnumberGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud virtualserialnumber getall unauthorized response has a 5xx status code
func (o *PcloudVirtualserialnumberGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud virtualserialnumber getall unauthorized response a status code equal to that given
func (o *PcloudVirtualserialnumberGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud virtualserialnumber getall unauthorized response
func (o *PcloudVirtualserialnumberGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudVirtualserialnumberGetallUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallUnauthorized %s", 401, payload)
}

func (o *PcloudVirtualserialnumberGetallUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallUnauthorized %s", 401, payload)
}

func (o *PcloudVirtualserialnumberGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVirtualserialnumberGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVirtualserialnumberGetallForbidden creates a PcloudVirtualserialnumberGetallForbidden with default headers values
func NewPcloudVirtualserialnumberGetallForbidden() *PcloudVirtualserialnumberGetallForbidden {
	return &PcloudVirtualserialnumberGetallForbidden{}
}

/*
PcloudVirtualserialnumberGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVirtualserialnumberGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud virtualserialnumber getall forbidden response has a 2xx status code
func (o *PcloudVirtualserialnumberGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud virtualserialnumber getall forbidden response has a 3xx status code
func (o *PcloudVirtualserialnumberGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud virtualserialnumber getall forbidden response has a 4xx status code
func (o *PcloudVirtualserialnumberGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud virtualserialnumber getall forbidden response has a 5xx status code
func (o *PcloudVirtualserialnumberGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud virtualserialnumber getall forbidden response a status code equal to that given
func (o *PcloudVirtualserialnumberGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud virtualserialnumber getall forbidden response
func (o *PcloudVirtualserialnumberGetallForbidden) Code() int {
	return 403
}

func (o *PcloudVirtualserialnumberGetallForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallForbidden %s", 403, payload)
}

func (o *PcloudVirtualserialnumberGetallForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallForbidden %s", 403, payload)
}

func (o *PcloudVirtualserialnumberGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVirtualserialnumberGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVirtualserialnumberGetallNotFound creates a PcloudVirtualserialnumberGetallNotFound with default headers values
func NewPcloudVirtualserialnumberGetallNotFound() *PcloudVirtualserialnumberGetallNotFound {
	return &PcloudVirtualserialnumberGetallNotFound{}
}

/*
PcloudVirtualserialnumberGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVirtualserialnumberGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud virtualserialnumber getall not found response has a 2xx status code
func (o *PcloudVirtualserialnumberGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud virtualserialnumber getall not found response has a 3xx status code
func (o *PcloudVirtualserialnumberGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud virtualserialnumber getall not found response has a 4xx status code
func (o *PcloudVirtualserialnumberGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud virtualserialnumber getall not found response has a 5xx status code
func (o *PcloudVirtualserialnumberGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud virtualserialnumber getall not found response a status code equal to that given
func (o *PcloudVirtualserialnumberGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud virtualserialnumber getall not found response
func (o *PcloudVirtualserialnumberGetallNotFound) Code() int {
	return 404
}

func (o *PcloudVirtualserialnumberGetallNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallNotFound %s", 404, payload)
}

func (o *PcloudVirtualserialnumberGetallNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallNotFound %s", 404, payload)
}

func (o *PcloudVirtualserialnumberGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVirtualserialnumberGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVirtualserialnumberGetallInternalServerError creates a PcloudVirtualserialnumberGetallInternalServerError with default headers values
func NewPcloudVirtualserialnumberGetallInternalServerError() *PcloudVirtualserialnumberGetallInternalServerError {
	return &PcloudVirtualserialnumberGetallInternalServerError{}
}

/*
PcloudVirtualserialnumberGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVirtualserialnumberGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud virtualserialnumber getall internal server error response has a 2xx status code
func (o *PcloudVirtualserialnumberGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud virtualserialnumber getall internal server error response has a 3xx status code
func (o *PcloudVirtualserialnumberGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud virtualserialnumber getall internal server error response has a 4xx status code
func (o *PcloudVirtualserialnumberGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud virtualserialnumber getall internal server error response has a 5xx status code
func (o *PcloudVirtualserialnumberGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud virtualserialnumber getall internal server error response a status code equal to that given
func (o *PcloudVirtualserialnumberGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud virtualserialnumber getall internal server error response
func (o *PcloudVirtualserialnumberGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudVirtualserialnumberGetallInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallInternalServerError %s", 500, payload)
}

func (o *PcloudVirtualserialnumberGetallInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/virtual-serial-number][%d] pcloudVirtualserialnumberGetallInternalServerError %s", 500, payload)
}

func (o *PcloudVirtualserialnumberGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVirtualserialnumberGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
