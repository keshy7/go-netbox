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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasObjectChangesReadReader is a Reader for the ExtrasObjectChangesRead structure.
type ExtrasObjectChangesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasObjectChangesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasObjectChangesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasObjectChangesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasObjectChangesReadOK creates a ExtrasObjectChangesReadOK with default headers values
func NewExtrasObjectChangesReadOK() *ExtrasObjectChangesReadOK {
	return &ExtrasObjectChangesReadOK{}
}

/*
ExtrasObjectChangesReadOK describes a response with status code 200, with default header values.

ExtrasObjectChangesReadOK extras object changes read o k
*/
type ExtrasObjectChangesReadOK struct {
	Payload *models.ObjectChange
}

// IsSuccess returns true when this extras object changes read o k response has a 2xx status code
func (o *ExtrasObjectChangesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras object changes read o k response has a 3xx status code
func (o *ExtrasObjectChangesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras object changes read o k response has a 4xx status code
func (o *ExtrasObjectChangesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras object changes read o k response has a 5xx status code
func (o *ExtrasObjectChangesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras object changes read o k response a status code equal to that given
func (o *ExtrasObjectChangesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasObjectChangesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasObjectChangesReadOK) String() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasObjectChangesReadOK) GetPayload() *models.ObjectChange {
	return o.Payload
}

func (o *ExtrasObjectChangesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectChange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasObjectChangesReadDefault creates a ExtrasObjectChangesReadDefault with default headers values
func NewExtrasObjectChangesReadDefault(code int) *ExtrasObjectChangesReadDefault {
	return &ExtrasObjectChangesReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasObjectChangesReadDefault describes a response with status code -1, with default header values.

ExtrasObjectChangesReadDefault extras object changes read default
*/
type ExtrasObjectChangesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras object changes read default response
func (o *ExtrasObjectChangesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras object changes read default response has a 2xx status code
func (o *ExtrasObjectChangesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras object changes read default response has a 3xx status code
func (o *ExtrasObjectChangesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras object changes read default response has a 4xx status code
func (o *ExtrasObjectChangesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras object changes read default response has a 5xx status code
func (o *ExtrasObjectChangesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras object changes read default response a status code equal to that given
func (o *ExtrasObjectChangesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasObjectChangesReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extras_object-changes_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasObjectChangesReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extras_object-changes_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasObjectChangesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasObjectChangesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
