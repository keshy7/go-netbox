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

// DcimPowerFeedsCreateReader is a Reader for the DcimPowerFeedsCreate structure.
type DcimPowerFeedsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerFeedsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsCreateCreated creates a DcimPowerFeedsCreateCreated with default headers values
func NewDcimPowerFeedsCreateCreated() *DcimPowerFeedsCreateCreated {
	return &DcimPowerFeedsCreateCreated{}
}

/*
DcimPowerFeedsCreateCreated describes a response with status code 201, with default header values.

DcimPowerFeedsCreateCreated dcim power feeds create created
*/
type DcimPowerFeedsCreateCreated struct {
	Payload *models.PowerFeed
}

// IsSuccess returns true when this dcim power feeds create created response has a 2xx status code
func (o *DcimPowerFeedsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds create created response has a 3xx status code
func (o *DcimPowerFeedsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds create created response has a 4xx status code
func (o *DcimPowerFeedsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds create created response has a 5xx status code
func (o *DcimPowerFeedsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds create created response a status code equal to that given
func (o *DcimPowerFeedsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimPowerFeedsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-feeds/][%d] dcimPowerFeedsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerFeedsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/power-feeds/][%d] dcimPowerFeedsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerFeedsCreateCreated) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsCreateDefault creates a DcimPowerFeedsCreateDefault with default headers values
func NewDcimPowerFeedsCreateDefault(code int) *DcimPowerFeedsCreateDefault {
	return &DcimPowerFeedsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerFeedsCreateDefault describes a response with status code -1, with default header values.

DcimPowerFeedsCreateDefault dcim power feeds create default
*/
type DcimPowerFeedsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds create default response
func (o *DcimPowerFeedsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power feeds create default response has a 2xx status code
func (o *DcimPowerFeedsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power feeds create default response has a 3xx status code
func (o *DcimPowerFeedsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power feeds create default response has a 4xx status code
func (o *DcimPowerFeedsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power feeds create default response has a 5xx status code
func (o *DcimPowerFeedsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power feeds create default response a status code equal to that given
func (o *DcimPowerFeedsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerFeedsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-feeds/][%d] dcim_power-feeds_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/power-feeds/][%d] dcim_power-feeds_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
