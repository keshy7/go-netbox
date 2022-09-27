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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsCircuitTypesDeleteReader is a Reader for the CircuitsCircuitTypesDelete structure.
type CircuitsCircuitTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsCircuitTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTypesDeleteNoContent creates a CircuitsCircuitTypesDeleteNoContent with default headers values
func NewCircuitsCircuitTypesDeleteNoContent() *CircuitsCircuitTypesDeleteNoContent {
	return &CircuitsCircuitTypesDeleteNoContent{}
}

/*
CircuitsCircuitTypesDeleteNoContent describes a response with status code 204, with default header values.

CircuitsCircuitTypesDeleteNoContent circuits circuit types delete no content
*/
type CircuitsCircuitTypesDeleteNoContent struct {
}

// IsSuccess returns true when this circuits circuit types delete no content response has a 2xx status code
func (o *CircuitsCircuitTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit types delete no content response has a 3xx status code
func (o *CircuitsCircuitTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types delete no content response has a 4xx status code
func (o *CircuitsCircuitTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit types delete no content response has a 5xx status code
func (o *CircuitsCircuitTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types delete no content response a status code equal to that given
func (o *CircuitsCircuitTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *CircuitsCircuitTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesDeleteNoContent ", 204)
}

func (o *CircuitsCircuitTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesDeleteNoContent ", 204)
}

func (o *CircuitsCircuitTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCircuitsCircuitTypesDeleteDefault creates a CircuitsCircuitTypesDeleteDefault with default headers values
func NewCircuitsCircuitTypesDeleteDefault(code int) *CircuitsCircuitTypesDeleteDefault {
	return &CircuitsCircuitTypesDeleteDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTypesDeleteDefault describes a response with status code -1, with default header values.

CircuitsCircuitTypesDeleteDefault circuits circuit types delete default
*/
type CircuitsCircuitTypesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit types delete default response
func (o *CircuitsCircuitTypesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits circuit types delete default response has a 2xx status code
func (o *CircuitsCircuitTypesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit types delete default response has a 3xx status code
func (o *CircuitsCircuitTypesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit types delete default response has a 4xx status code
func (o *CircuitsCircuitTypesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit types delete default response has a 5xx status code
func (o *CircuitsCircuitTypesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit types delete default response a status code equal to that given
func (o *CircuitsCircuitTypesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsCircuitTypesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-types/{id}/][%d] circuits_circuit-types_delete default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-types/{id}/][%d] circuits_circuit-types_delete default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTypesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
