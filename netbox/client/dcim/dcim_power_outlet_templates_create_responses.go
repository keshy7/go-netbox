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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimPowerOutletTemplatesCreateReader is a Reader for the DcimPowerOutletTemplatesCreate structure.
type DcimPowerOutletTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerOutletTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesCreateCreated creates a DcimPowerOutletTemplatesCreateCreated with default headers values
func NewDcimPowerOutletTemplatesCreateCreated() *DcimPowerOutletTemplatesCreateCreated {
	return &DcimPowerOutletTemplatesCreateCreated{}
}

/*
DcimPowerOutletTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimPowerOutletTemplatesCreateCreated dcim power outlet templates create created
*/
type DcimPowerOutletTemplatesCreateCreated struct {
	Payload *models.PowerOutletTemplate
}

// IsSuccess returns true when this dcim power outlet templates create created response has a 2xx status code
func (o *DcimPowerOutletTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlet templates create created response has a 3xx status code
func (o *DcimPowerOutletTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlet templates create created response has a 4xx status code
func (o *DcimPowerOutletTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlet templates create created response has a 5xx status code
func (o *DcimPowerOutletTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlet templates create created response a status code equal to that given
func (o *DcimPowerOutletTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimPowerOutletTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerOutletTemplatesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerOutletTemplatesCreateCreated) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletTemplatesCreateDefault creates a DcimPowerOutletTemplatesCreateDefault with default headers values
func NewDcimPowerOutletTemplatesCreateDefault(code int) *DcimPowerOutletTemplatesCreateDefault {
	return &DcimPowerOutletTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletTemplatesCreateDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesCreateDefault dcim power outlet templates create default
*/
type DcimPowerOutletTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlet templates create default response
func (o *DcimPowerOutletTemplatesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power outlet templates create default response has a 2xx status code
func (o *DcimPowerOutletTemplatesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlet templates create default response has a 3xx status code
func (o *DcimPowerOutletTemplatesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlet templates create default response has a 4xx status code
func (o *DcimPowerOutletTemplatesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlet templates create default response has a 5xx status code
func (o *DcimPowerOutletTemplatesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlet templates create default response a status code equal to that given
func (o *DcimPowerOutletTemplatesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerOutletTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-outlet-templates/][%d] dcim_power-outlet-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/power-outlet-templates/][%d] dcim_power-outlet-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
