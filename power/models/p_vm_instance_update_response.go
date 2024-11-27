// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceUpdateResponse p VM instance update response
//
// swagger:model PVMInstanceUpdateResponse
type PVMInstanceUpdateResponse struct {

	// The VTL license repository capacity TB value
	LicenseRepositoryCapacity int64 `json:"licenseRepositoryCapacity,omitempty"`

	// Amount of memory allocated (in GB)
	Memory float64 `json:"memory,omitempty"`

	// pin policy
	PinPolicy PinPolicy `json:"pinPolicy,omitempty"`

	// Processor type (dedicated, shared, capped)
	// Enum: [dedicated shared capped]
	ProcType string `json:"procType,omitempty"`

	// Number of processors allocated
	Processors float64 `json:"processors,omitempty"`

	// Name of the server
	ServerName string `json:"serverName,omitempty"`

	// URL to check for status of the operation (for now, just the URL for the GET on the server, which has status information from powervc)
	StatusURL string `json:"statusUrl,omitempty"`

	// Indicates if all volumes attached to the server must reside in the same storage pool; If set to false then volumes from any storage type and pool can be attached to the PVMInstance; Impacts PVMInstance snapshot, capture, and clone, for capture and clone - only data volumes that are of the same storage type and in the same storage pool of the PVMInstance's boot volume can be included; for snapshot - all data volumes to be included in the snapshot must reside in the same storage type and pool. Once set to false, cannot be set back to true unless all volumes attached reside in the same storage type and pool.
	StoragePoolAffinity *bool `json:"storagePoolAffinity,omitempty"`

	// The pvm instance virtual CPU information
	VirtualCores *VirtualCores `json:"virtualCores,omitempty"`
}

// Validate validates this p VM instance update response
func (m *PVMInstanceUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePinPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualCores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceUpdateResponse) validatePinPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.PinPolicy) { // not required
		return nil
	}

	if err := m.PinPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pinPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("pinPolicy")
		}
		return err
	}

	return nil
}

var pVmInstanceUpdateResponseTypeProcTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dedicated","shared","capped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceUpdateResponseTypeProcTypePropEnum = append(pVmInstanceUpdateResponseTypeProcTypePropEnum, v)
	}
}

const (

	// PVMInstanceUpdateResponseProcTypeDedicated captures enum value "dedicated"
	PVMInstanceUpdateResponseProcTypeDedicated string = "dedicated"

	// PVMInstanceUpdateResponseProcTypeShared captures enum value "shared"
	PVMInstanceUpdateResponseProcTypeShared string = "shared"

	// PVMInstanceUpdateResponseProcTypeCapped captures enum value "capped"
	PVMInstanceUpdateResponseProcTypeCapped string = "capped"
)

// prop value enum
func (m *PVMInstanceUpdateResponse) validateProcTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceUpdateResponseTypeProcTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceUpdateResponse) validateProcType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProcTypeEnum("procType", "body", m.ProcType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceUpdateResponse) validateVirtualCores(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualCores) { // not required
		return nil
	}

	if m.VirtualCores != nil {
		if err := m.VirtualCores.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualCores")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualCores")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this p VM instance update response based on the context it is used
func (m *PVMInstanceUpdateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePinPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualCores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceUpdateResponse) contextValidatePinPolicy(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PinPolicy) { // not required
		return nil
	}

	if err := m.PinPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pinPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("pinPolicy")
		}
		return err
	}

	return nil
}

func (m *PVMInstanceUpdateResponse) contextValidateVirtualCores(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualCores != nil {

		if swag.IsZero(m.VirtualCores) { // not required
			return nil
		}

		if err := m.VirtualCores.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualCores")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualCores")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceUpdateResponse) UnmarshalBinary(b []byte) error {
	var res PVMInstanceUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
