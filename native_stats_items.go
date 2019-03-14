// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NativeStatsItems Stats
//
// Current stats for one object.
// swagger:model nativeStatsItems
type NativeStatsItems struct {

	// backend name
	BackendName string `json:"backend_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// runtime Api
	RuntimeAPI string `json:"runtimeApi,omitempty"`

	// stats
	Stats *NativeStatsItemsStats `json:"stats,omitempty"`

	// type
	// Enum: [backend server frontend]
	Type string `json:"type,omitempty"`
}

// Validate validates this native stats items
func (m *NativeStatsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeStatsItems) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

var nativeStatsItemsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["backend","server","frontend"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatsItemsTypeTypePropEnum = append(nativeStatsItemsTypeTypePropEnum, v)
	}
}

const (

	// NativeStatsItemsTypeBackend captures enum value "backend"
	NativeStatsItemsTypeBackend string = "backend"

	// NativeStatsItemsTypeServer captures enum value "server"
	NativeStatsItemsTypeServer string = "server"

	// NativeStatsItemsTypeFrontend captures enum value "frontend"
	NativeStatsItemsTypeFrontend string = "frontend"
)

// prop value enum
func (m *NativeStatsItems) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nativeStatsItemsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatsItems) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeStatsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeStatsItems) UnmarshalBinary(b []byte) error {
	var res NativeStatsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
