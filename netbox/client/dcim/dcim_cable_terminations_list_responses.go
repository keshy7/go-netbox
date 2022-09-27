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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimCableTerminationsListReader is a Reader for the DcimCableTerminationsList structure.
type DcimCableTerminationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCableTerminationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCableTerminationsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCableTerminationsListOK creates a DcimCableTerminationsListOK with default headers values
func NewDcimCableTerminationsListOK() *DcimCableTerminationsListOK {
	return &DcimCableTerminationsListOK{}
}

/*
DcimCableTerminationsListOK describes a response with status code 200, with default header values.

DcimCableTerminationsListOK dcim cable terminations list o k
*/
type DcimCableTerminationsListOK struct {
	Payload *DcimCableTerminationsListOKBody
}

// IsSuccess returns true when this dcim cable terminations list o k response has a 2xx status code
func (o *DcimCableTerminationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations list o k response has a 3xx status code
func (o *DcimCableTerminationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations list o k response has a 4xx status code
func (o *DcimCableTerminationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations list o k response has a 5xx status code
func (o *DcimCableTerminationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations list o k response a status code equal to that given
func (o *DcimCableTerminationsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimCableTerminationsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/cable-terminations/][%d] dcimCableTerminationsListOK  %+v", 200, o.Payload)
}

func (o *DcimCableTerminationsListOK) String() string {
	return fmt.Sprintf("[GET /dcim/cable-terminations/][%d] dcimCableTerminationsListOK  %+v", 200, o.Payload)
}

func (o *DcimCableTerminationsListOK) GetPayload() *DcimCableTerminationsListOKBody {
	return o.Payload
}

func (o *DcimCableTerminationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimCableTerminationsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCableTerminationsListDefault creates a DcimCableTerminationsListDefault with default headers values
func NewDcimCableTerminationsListDefault(code int) *DcimCableTerminationsListDefault {
	return &DcimCableTerminationsListDefault{
		_statusCode: code,
	}
}

/*
DcimCableTerminationsListDefault describes a response with status code -1, with default header values.

DcimCableTerminationsListDefault dcim cable terminations list default
*/
type DcimCableTerminationsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cable terminations list default response
func (o *DcimCableTerminationsListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cable terminations list default response has a 2xx status code
func (o *DcimCableTerminationsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cable terminations list default response has a 3xx status code
func (o *DcimCableTerminationsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cable terminations list default response has a 4xx status code
func (o *DcimCableTerminationsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cable terminations list default response has a 5xx status code
func (o *DcimCableTerminationsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cable terminations list default response a status code equal to that given
func (o *DcimCableTerminationsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCableTerminationsListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/cable-terminations/][%d] dcim_cable-terminations_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/cable-terminations/][%d] dcim_cable-terminations_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimCableTerminationsListOKBody dcim cable terminations list o k body
swagger:model DcimCableTerminationsListOKBody
*/
type DcimCableTerminationsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.CableTermination `json:"results"`
}

// Validate validates this dcim cable terminations list o k body
func (o *DcimCableTerminationsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimCableTerminationsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimCableTerminationsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimCableTerminationsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCableTerminationsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCableTerminationsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCableTerminationsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCableTerminationsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimCableTerminationsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim cable terminations list o k body based on the context it is used
func (o *DcimCableTerminationsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimCableTerminationsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCableTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimCableTerminationsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimCableTerminationsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimCableTerminationsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
