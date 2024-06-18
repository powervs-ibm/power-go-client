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

// StorageTier Storage tier detail
//
// swagger:model StorageTier
type StorageTier struct {

	// Description, storage tier label
	Description string `json:"description,omitempty"`

	// Name of the storage tier
	Name string `json:"name,omitempty"`

	// State of the storage tier (active or inactive)
	// Enum: ["active","inactive"]
	State *string `json:"state,omitempty"`
}

// Validate validates this storage tier
func (m *StorageTier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storageTierTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageTierTypeStatePropEnum = append(storageTierTypeStatePropEnum, v)
	}
}

const (

	// StorageTierStateActive captures enum value "active"
	StorageTierStateActive string = "active"

	// StorageTierStateInactive captures enum value "inactive"
	StorageTierStateInactive string = "inactive"
)

// prop value enum
func (m *StorageTier) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, storageTierTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StorageTier) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this storage tier based on context it is used
func (m *StorageTier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageTier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageTier) UnmarshalBinary(b []byte) error {
	var res StorageTier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
