// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPlacementgroupsGetReader is a Reader for the PcloudPlacementgroupsGet structure.
type PcloudPlacementgroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPlacementgroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPlacementgroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPlacementgroupsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPlacementgroupsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPlacementgroupsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPlacementgroupsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPlacementgroupsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}] pcloud.placementgroups.get", response, response.Code())
	}
}

// NewPcloudPlacementgroupsGetOK creates a PcloudPlacementgroupsGetOK with default headers values
func NewPcloudPlacementgroupsGetOK() *PcloudPlacementgroupsGetOK {
	return &PcloudPlacementgroupsGetOK{}
}

/*
PcloudPlacementgroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPlacementgroupsGetOK struct {
	Payload *models.PlacementGroup
}

// IsSuccess returns true when this pcloud placementgroups get o k response has a 2xx status code
func (o *PcloudPlacementgroupsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud placementgroups get o k response has a 3xx status code
func (o *PcloudPlacementgroupsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups get o k response has a 4xx status code
func (o *PcloudPlacementgroupsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups get o k response has a 5xx status code
func (o *PcloudPlacementgroupsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups get o k response a status code equal to that given
func (o *PcloudPlacementgroupsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud placementgroups get o k response
func (o *PcloudPlacementgroupsGetOK) Code() int {
	return 200
}

func (o *PcloudPlacementgroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPlacementgroupsGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPlacementgroupsGetOK) GetPayload() *models.PlacementGroup {
	return o.Payload
}

func (o *PcloudPlacementgroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsGetBadRequest creates a PcloudPlacementgroupsGetBadRequest with default headers values
func NewPcloudPlacementgroupsGetBadRequest() *PcloudPlacementgroupsGetBadRequest {
	return &PcloudPlacementgroupsGetBadRequest{}
}

/*
PcloudPlacementgroupsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPlacementgroupsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups get bad request response has a 2xx status code
func (o *PcloudPlacementgroupsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups get bad request response has a 3xx status code
func (o *PcloudPlacementgroupsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups get bad request response has a 4xx status code
func (o *PcloudPlacementgroupsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups get bad request response has a 5xx status code
func (o *PcloudPlacementgroupsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups get bad request response a status code equal to that given
func (o *PcloudPlacementgroupsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud placementgroups get bad request response
func (o *PcloudPlacementgroupsGetBadRequest) Code() int {
	return 400
}

func (o *PcloudPlacementgroupsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPlacementgroupsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPlacementgroupsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsGetUnauthorized creates a PcloudPlacementgroupsGetUnauthorized with default headers values
func NewPcloudPlacementgroupsGetUnauthorized() *PcloudPlacementgroupsGetUnauthorized {
	return &PcloudPlacementgroupsGetUnauthorized{}
}

/*
PcloudPlacementgroupsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPlacementgroupsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups get unauthorized response has a 2xx status code
func (o *PcloudPlacementgroupsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups get unauthorized response has a 3xx status code
func (o *PcloudPlacementgroupsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups get unauthorized response has a 4xx status code
func (o *PcloudPlacementgroupsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups get unauthorized response has a 5xx status code
func (o *PcloudPlacementgroupsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups get unauthorized response a status code equal to that given
func (o *PcloudPlacementgroupsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud placementgroups get unauthorized response
func (o *PcloudPlacementgroupsGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudPlacementgroupsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPlacementgroupsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPlacementgroupsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsGetForbidden creates a PcloudPlacementgroupsGetForbidden with default headers values
func NewPcloudPlacementgroupsGetForbidden() *PcloudPlacementgroupsGetForbidden {
	return &PcloudPlacementgroupsGetForbidden{}
}

/*
PcloudPlacementgroupsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPlacementgroupsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups get forbidden response has a 2xx status code
func (o *PcloudPlacementgroupsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups get forbidden response has a 3xx status code
func (o *PcloudPlacementgroupsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups get forbidden response has a 4xx status code
func (o *PcloudPlacementgroupsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups get forbidden response has a 5xx status code
func (o *PcloudPlacementgroupsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups get forbidden response a status code equal to that given
func (o *PcloudPlacementgroupsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud placementgroups get forbidden response
func (o *PcloudPlacementgroupsGetForbidden) Code() int {
	return 403
}

func (o *PcloudPlacementgroupsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPlacementgroupsGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPlacementgroupsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsGetNotFound creates a PcloudPlacementgroupsGetNotFound with default headers values
func NewPcloudPlacementgroupsGetNotFound() *PcloudPlacementgroupsGetNotFound {
	return &PcloudPlacementgroupsGetNotFound{}
}

/*
PcloudPlacementgroupsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPlacementgroupsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups get not found response has a 2xx status code
func (o *PcloudPlacementgroupsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups get not found response has a 3xx status code
func (o *PcloudPlacementgroupsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups get not found response has a 4xx status code
func (o *PcloudPlacementgroupsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups get not found response has a 5xx status code
func (o *PcloudPlacementgroupsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups get not found response a status code equal to that given
func (o *PcloudPlacementgroupsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud placementgroups get not found response
func (o *PcloudPlacementgroupsGetNotFound) Code() int {
	return 404
}

func (o *PcloudPlacementgroupsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPlacementgroupsGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPlacementgroupsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsGetInternalServerError creates a PcloudPlacementgroupsGetInternalServerError with default headers values
func NewPcloudPlacementgroupsGetInternalServerError() *PcloudPlacementgroupsGetInternalServerError {
	return &PcloudPlacementgroupsGetInternalServerError{}
}

/*
PcloudPlacementgroupsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPlacementgroupsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups get internal server error response has a 2xx status code
func (o *PcloudPlacementgroupsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups get internal server error response has a 3xx status code
func (o *PcloudPlacementgroupsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups get internal server error response has a 4xx status code
func (o *PcloudPlacementgroupsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups get internal server error response has a 5xx status code
func (o *PcloudPlacementgroupsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud placementgroups get internal server error response a status code equal to that given
func (o *PcloudPlacementgroupsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud placementgroups get internal server error response
func (o *PcloudPlacementgroupsGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudPlacementgroupsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPlacementgroupsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPlacementgroupsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
