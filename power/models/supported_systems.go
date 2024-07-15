// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SupportedSystems supported systems
//
// swagger:model SupportedSystems
type SupportedSystems struct {

	// List of all available dedicated host types
	// Required: true
	Dedicated []string `json:"dedicated"`

	// List of all available host types
	// Required: true
	General []string `json:"general"`
}

// Validate validates this supported systems
func (m *SupportedSystems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDedicated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneral(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportedSystems) validateDedicated(formats strfmt.Registry) error {

	if err := validate.Required("dedicated", "body", m.Dedicated); err != nil {
		return err
	}

	return nil
}

func (m *SupportedSystems) validateGeneral(formats strfmt.Registry) error {

	if err := validate.Required("general", "body", m.General); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this supported systems based on context it is used
func (m *SupportedSystems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupportedSystems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedSystems) UnmarshalBinary(b []byte) error {
	var res SupportedSystems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
