// Code generated by go-swagger; DO NOT EDIT.

package network_security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// V1NetworkSecurityGroupsRulesDeleteReader is a Reader for the V1NetworkSecurityGroupsRulesDelete structure.
type V1NetworkSecurityGroupsRulesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NetworkSecurityGroupsRulesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NetworkSecurityGroupsRulesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1NetworkSecurityGroupsRulesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1NetworkSecurityGroupsRulesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1NetworkSecurityGroupsRulesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1NetworkSecurityGroupsRulesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV1NetworkSecurityGroupsRulesDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1NetworkSecurityGroupsRulesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}] v1.networkSecurityGroups.rules.delete", response, response.Code())
	}
}

// NewV1NetworkSecurityGroupsRulesDeleteOK creates a V1NetworkSecurityGroupsRulesDeleteOK with default headers values
func NewV1NetworkSecurityGroupsRulesDeleteOK() *V1NetworkSecurityGroupsRulesDeleteOK {
	return &V1NetworkSecurityGroupsRulesDeleteOK{}
}

/*
V1NetworkSecurityGroupsRulesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type V1NetworkSecurityGroupsRulesDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this v1 network security groups rules delete o k response has a 2xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 network security groups rules delete o k response has a 3xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups rules delete o k response has a 4xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network security groups rules delete o k response has a 5xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups rules delete o k response a status code equal to that given
func (o *V1NetworkSecurityGroupsRulesDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 network security groups rules delete o k response
func (o *V1NetworkSecurityGroupsRulesDeleteOK) Code() int {
	return 200
}

func (o *V1NetworkSecurityGroupsRulesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteOK  %+v", 200, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteOK  %+v", 200, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsRulesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsRulesDeleteBadRequest creates a V1NetworkSecurityGroupsRulesDeleteBadRequest with default headers values
func NewV1NetworkSecurityGroupsRulesDeleteBadRequest() *V1NetworkSecurityGroupsRulesDeleteBadRequest {
	return &V1NetworkSecurityGroupsRulesDeleteBadRequest{}
}

/*
V1NetworkSecurityGroupsRulesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1NetworkSecurityGroupsRulesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups rules delete bad request response has a 2xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups rules delete bad request response has a 3xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups rules delete bad request response has a 4xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups rules delete bad request response has a 5xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups rules delete bad request response a status code equal to that given
func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 network security groups rules delete bad request response
func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) Code() int {
	return 400
}

func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsRulesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsRulesDeleteUnauthorized creates a V1NetworkSecurityGroupsRulesDeleteUnauthorized with default headers values
func NewV1NetworkSecurityGroupsRulesDeleteUnauthorized() *V1NetworkSecurityGroupsRulesDeleteUnauthorized {
	return &V1NetworkSecurityGroupsRulesDeleteUnauthorized{}
}

/*
V1NetworkSecurityGroupsRulesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1NetworkSecurityGroupsRulesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups rules delete unauthorized response has a 2xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups rules delete unauthorized response has a 3xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups rules delete unauthorized response has a 4xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups rules delete unauthorized response has a 5xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups rules delete unauthorized response a status code equal to that given
func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 network security groups rules delete unauthorized response
func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) Code() int {
	return 401
}

func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsRulesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsRulesDeleteForbidden creates a V1NetworkSecurityGroupsRulesDeleteForbidden with default headers values
func NewV1NetworkSecurityGroupsRulesDeleteForbidden() *V1NetworkSecurityGroupsRulesDeleteForbidden {
	return &V1NetworkSecurityGroupsRulesDeleteForbidden{}
}

/*
V1NetworkSecurityGroupsRulesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1NetworkSecurityGroupsRulesDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups rules delete forbidden response has a 2xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups rules delete forbidden response has a 3xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups rules delete forbidden response has a 4xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups rules delete forbidden response has a 5xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups rules delete forbidden response a status code equal to that given
func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 network security groups rules delete forbidden response
func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) Code() int {
	return 403
}

func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsRulesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsRulesDeleteNotFound creates a V1NetworkSecurityGroupsRulesDeleteNotFound with default headers values
func NewV1NetworkSecurityGroupsRulesDeleteNotFound() *V1NetworkSecurityGroupsRulesDeleteNotFound {
	return &V1NetworkSecurityGroupsRulesDeleteNotFound{}
}

/*
V1NetworkSecurityGroupsRulesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1NetworkSecurityGroupsRulesDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups rules delete not found response has a 2xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups rules delete not found response has a 3xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups rules delete not found response has a 4xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups rules delete not found response has a 5xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups rules delete not found response a status code equal to that given
func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 network security groups rules delete not found response
func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) Code() int {
	return 404
}

func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsRulesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsRulesDeleteConflict creates a V1NetworkSecurityGroupsRulesDeleteConflict with default headers values
func NewV1NetworkSecurityGroupsRulesDeleteConflict() *V1NetworkSecurityGroupsRulesDeleteConflict {
	return &V1NetworkSecurityGroupsRulesDeleteConflict{}
}

/*
V1NetworkSecurityGroupsRulesDeleteConflict describes a response with status code 409, with default header values.

Conflict
*/
type V1NetworkSecurityGroupsRulesDeleteConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups rules delete conflict response has a 2xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups rules delete conflict response has a 3xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups rules delete conflict response has a 4xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 network security groups rules delete conflict response has a 5xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 network security groups rules delete conflict response a status code equal to that given
func (o *V1NetworkSecurityGroupsRulesDeleteConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the v1 network security groups rules delete conflict response
func (o *V1NetworkSecurityGroupsRulesDeleteConflict) Code() int {
	return 409
}

func (o *V1NetworkSecurityGroupsRulesDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteConflict  %+v", 409, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteConflict) String() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteConflict  %+v", 409, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsRulesDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1NetworkSecurityGroupsRulesDeleteInternalServerError creates a V1NetworkSecurityGroupsRulesDeleteInternalServerError with default headers values
func NewV1NetworkSecurityGroupsRulesDeleteInternalServerError() *V1NetworkSecurityGroupsRulesDeleteInternalServerError {
	return &V1NetworkSecurityGroupsRulesDeleteInternalServerError{}
}

/*
V1NetworkSecurityGroupsRulesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1NetworkSecurityGroupsRulesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 network security groups rules delete internal server error response has a 2xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 network security groups rules delete internal server error response has a 3xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 network security groups rules delete internal server error response has a 4xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 network security groups rules delete internal server error response has a 5xx status code
func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 network security groups rules delete internal server error response a status code equal to that given
func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 network security groups rules delete internal server error response
func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) Code() int {
	return 500
}

func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/network-security-groups/{network_security_group_id}/rules/{network_security_group_rule_id}][%d] v1NetworkSecurityGroupsRulesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1NetworkSecurityGroupsRulesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
