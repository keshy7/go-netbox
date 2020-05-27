// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableFrontPort writable front port
//
// swagger:model WritableFrontPort
type WritableFrontPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Rear port
	// Required: true
	RearPort *int64 `json:"rear_port"`

	// Rear port position
	// Maximum: 64
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Required: true
	// Enum: [8p8c 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st]
	Type *string `json:"type"`
}

// Validate validates this writable front port
func (m *WritableFrontPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *WritableFrontPort) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableFrontPort) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateRearPortPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.RearPortPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("rear_port_position", "body", int64(m.RearPortPosition), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rear_port_position", "body", int64(m.RearPortPosition), 64, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPort) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

var writableFrontPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableFrontPortTypeTypePropEnum = append(writableFrontPortTypeTypePropEnum, v)
	}
}

const (

	// WritableFrontPortTypeNr8p8c captures enum value "8p8c"
	WritableFrontPortTypeNr8p8c string = "8p8c"

	// WritableFrontPortTypeNr110Punch captures enum value "110-punch"
	WritableFrontPortTypeNr110Punch string = "110-punch"

	// WritableFrontPortTypeBnc captures enum value "bnc"
	WritableFrontPortTypeBnc string = "bnc"

	// WritableFrontPortTypeMrj21 captures enum value "mrj21"
	WritableFrontPortTypeMrj21 string = "mrj21"

	// WritableFrontPortTypeFc captures enum value "fc"
	WritableFrontPortTypeFc string = "fc"

	// WritableFrontPortTypeLc captures enum value "lc"
	WritableFrontPortTypeLc string = "lc"

	// WritableFrontPortTypeLcApc captures enum value "lc-apc"
	WritableFrontPortTypeLcApc string = "lc-apc"

	// WritableFrontPortTypeLsh captures enum value "lsh"
	WritableFrontPortTypeLsh string = "lsh"

	// WritableFrontPortTypeLshApc captures enum value "lsh-apc"
	WritableFrontPortTypeLshApc string = "lsh-apc"

	// WritableFrontPortTypeMpo captures enum value "mpo"
	WritableFrontPortTypeMpo string = "mpo"

	// WritableFrontPortTypeMtrj captures enum value "mtrj"
	WritableFrontPortTypeMtrj string = "mtrj"

	// WritableFrontPortTypeSc captures enum value "sc"
	WritableFrontPortTypeSc string = "sc"

	// WritableFrontPortTypeScApc captures enum value "sc-apc"
	WritableFrontPortTypeScApc string = "sc-apc"

	// WritableFrontPortTypeSt captures enum value "st"
	WritableFrontPortTypeSt string = "st"
)

// prop value enum
func (m *WritableFrontPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableFrontPortTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableFrontPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableFrontPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableFrontPort) UnmarshalBinary(b []byte) error {
	var res WritableFrontPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
