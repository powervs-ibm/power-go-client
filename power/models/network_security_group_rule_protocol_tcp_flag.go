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

// NetworkSecurityGroupRuleProtocolTCPFlag network security group rule protocol Tcp flag
//
// swagger:model NetworkSecurityGroupRuleProtocolTcpFlag
type NetworkSecurityGroupRuleProtocolTCPFlag struct {

	// TCP flag
	// Enum: [syn ack fin rst urg psh wnd chk seq]
	Flag string `json:"flag,omitempty"`
}

// Validate validates this network security group rule protocol Tcp flag
func (m *NetworkSecurityGroupRuleProtocolTCPFlag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkSecurityGroupRuleProtocolTcpFlagTypeFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["syn","ack","fin","rst","urg","psh","wnd","chk","seq"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkSecurityGroupRuleProtocolTcpFlagTypeFlagPropEnum = append(networkSecurityGroupRuleProtocolTcpFlagTypeFlagPropEnum, v)
	}
}

const (

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagSyn captures enum value "syn"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagSyn string = "syn"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagAck captures enum value "ack"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagAck string = "ack"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagFin captures enum value "fin"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagFin string = "fin"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagRst captures enum value "rst"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagRst string = "rst"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagUrg captures enum value "urg"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagUrg string = "urg"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagPsh captures enum value "psh"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagPsh string = "psh"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagWnd captures enum value "wnd"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagWnd string = "wnd"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagChk captures enum value "chk"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagChk string = "chk"

	// NetworkSecurityGroupRuleProtocolTCPFlagFlagSeq captures enum value "seq"
	NetworkSecurityGroupRuleProtocolTCPFlagFlagSeq string = "seq"
)

// prop value enum
func (m *NetworkSecurityGroupRuleProtocolTCPFlag) validateFlagEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkSecurityGroupRuleProtocolTcpFlagTypeFlagPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkSecurityGroupRuleProtocolTCPFlag) validateFlag(formats strfmt.Registry) error {
	if swag.IsZero(m.Flag) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlagEnum("flag", "body", m.Flag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network security group rule protocol Tcp flag based on context it is used
func (m *NetworkSecurityGroupRuleProtocolTCPFlag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSecurityGroupRuleProtocolTCPFlag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSecurityGroupRuleProtocolTCPFlag) UnmarshalBinary(b []byte) error {
	var res NetworkSecurityGroupRuleProtocolTCPFlag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
