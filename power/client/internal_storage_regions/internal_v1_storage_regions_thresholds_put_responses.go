// Code generated by go-swagger; DO NOT EDIT.

package internal_storage_regions

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

// InternalV1StorageRegionsThresholdsPutReader is a Reader for the InternalV1StorageRegionsThresholdsPut structure.
type InternalV1StorageRegionsThresholdsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1StorageRegionsThresholdsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewInternalV1StorageRegionsThresholdsPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalV1StorageRegionsThresholdsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInternalV1StorageRegionsThresholdsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1StorageRegionsThresholdsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewInternalV1StorageRegionsThresholdsPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewInternalV1StorageRegionsThresholdsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1StorageRegionsThresholdsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds] internal.v1.storage.regions.thresholds.put", response, response.Code())
	}
}

// NewInternalV1StorageRegionsThresholdsPutAccepted creates a InternalV1StorageRegionsThresholdsPutAccepted with default headers values
func NewInternalV1StorageRegionsThresholdsPutAccepted() *InternalV1StorageRegionsThresholdsPutAccepted {
	return &InternalV1StorageRegionsThresholdsPutAccepted{}
}

/*
InternalV1StorageRegionsThresholdsPutAccepted describes a response with status code 202, with default header values.

OK, region-zone default threshold settings update
*/
type InternalV1StorageRegionsThresholdsPutAccepted struct {
	Payload *models.Thresholds
}

// IsSuccess returns true when this internal v1 storage regions thresholds put accepted response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsPutAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 storage regions thresholds put accepted response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsPutAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds put accepted response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsPutAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 storage regions thresholds put accepted response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsPutAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds put accepted response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsPutAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the internal v1 storage regions thresholds put accepted response
func (o *InternalV1StorageRegionsThresholdsPutAccepted) Code() int {
	return 202
}

func (o *InternalV1StorageRegionsThresholdsPutAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutAccepted %s", 202, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutAccepted %s", 202, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutAccepted) GetPayload() *models.Thresholds {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thresholds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsPutBadRequest creates a InternalV1StorageRegionsThresholdsPutBadRequest with default headers values
func NewInternalV1StorageRegionsThresholdsPutBadRequest() *InternalV1StorageRegionsThresholdsPutBadRequest {
	return &InternalV1StorageRegionsThresholdsPutBadRequest{}
}

/*
InternalV1StorageRegionsThresholdsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalV1StorageRegionsThresholdsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds put bad request response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds put bad request response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds put bad request response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds put bad request response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds put bad request response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal v1 storage regions thresholds put bad request response
func (o *InternalV1StorageRegionsThresholdsPutBadRequest) Code() int {
	return 400
}

func (o *InternalV1StorageRegionsThresholdsPutBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutBadRequest %s", 400, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutBadRequest %s", 400, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsPutUnauthorized creates a InternalV1StorageRegionsThresholdsPutUnauthorized with default headers values
func NewInternalV1StorageRegionsThresholdsPutUnauthorized() *InternalV1StorageRegionsThresholdsPutUnauthorized {
	return &InternalV1StorageRegionsThresholdsPutUnauthorized{}
}

/*
InternalV1StorageRegionsThresholdsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1StorageRegionsThresholdsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds put unauthorized response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds put unauthorized response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds put unauthorized response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds put unauthorized response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds put unauthorized response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 storage regions thresholds put unauthorized response
func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) Code() int {
	return 401
}

func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutUnauthorized %s", 401, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutUnauthorized %s", 401, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsPutForbidden creates a InternalV1StorageRegionsThresholdsPutForbidden with default headers values
func NewInternalV1StorageRegionsThresholdsPutForbidden() *InternalV1StorageRegionsThresholdsPutForbidden {
	return &InternalV1StorageRegionsThresholdsPutForbidden{}
}

/*
InternalV1StorageRegionsThresholdsPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1StorageRegionsThresholdsPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds put forbidden response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds put forbidden response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds put forbidden response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds put forbidden response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds put forbidden response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 storage regions thresholds put forbidden response
func (o *InternalV1StorageRegionsThresholdsPutForbidden) Code() int {
	return 403
}

func (o *InternalV1StorageRegionsThresholdsPutForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutForbidden %s", 403, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutForbidden %s", 403, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsPutConflict creates a InternalV1StorageRegionsThresholdsPutConflict with default headers values
func NewInternalV1StorageRegionsThresholdsPutConflict() *InternalV1StorageRegionsThresholdsPutConflict {
	return &InternalV1StorageRegionsThresholdsPutConflict{}
}

/*
InternalV1StorageRegionsThresholdsPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type InternalV1StorageRegionsThresholdsPutConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds put conflict response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds put conflict response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds put conflict response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds put conflict response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds put conflict response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the internal v1 storage regions thresholds put conflict response
func (o *InternalV1StorageRegionsThresholdsPutConflict) Code() int {
	return 409
}

func (o *InternalV1StorageRegionsThresholdsPutConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutConflict %s", 409, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutConflict %s", 409, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsPutUnprocessableEntity creates a InternalV1StorageRegionsThresholdsPutUnprocessableEntity with default headers values
func NewInternalV1StorageRegionsThresholdsPutUnprocessableEntity() *InternalV1StorageRegionsThresholdsPutUnprocessableEntity {
	return &InternalV1StorageRegionsThresholdsPutUnprocessableEntity{}
}

/*
InternalV1StorageRegionsThresholdsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type InternalV1StorageRegionsThresholdsPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds put unprocessable entity response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds put unprocessable entity response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds put unprocessable entity response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds put unprocessable entity response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds put unprocessable entity response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the internal v1 storage regions thresholds put unprocessable entity response
func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) Code() int {
	return 422
}

func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutUnprocessableEntity %s", 422, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutUnprocessableEntity %s", 422, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsPutInternalServerError creates a InternalV1StorageRegionsThresholdsPutInternalServerError with default headers values
func NewInternalV1StorageRegionsThresholdsPutInternalServerError() *InternalV1StorageRegionsThresholdsPutInternalServerError {
	return &InternalV1StorageRegionsThresholdsPutInternalServerError{}
}

/*
InternalV1StorageRegionsThresholdsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1StorageRegionsThresholdsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds put internal server error response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds put internal server error response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds put internal server error response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 storage regions thresholds put internal server error response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 storage regions thresholds put internal server error response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 storage regions thresholds put internal server error response
func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) Code() int {
	return 500
}

func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutInternalServerError %s", 500, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsPutInternalServerError %s", 500, payload)
}

func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
