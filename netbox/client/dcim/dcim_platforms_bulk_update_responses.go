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

// DcimPlatformsBulkUpdateReader is a Reader for the DcimPlatformsBulkUpdate structure.
type DcimPlatformsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsBulkUpdateOK creates a DcimPlatformsBulkUpdateOK with default headers values
func NewDcimPlatformsBulkUpdateOK() *DcimPlatformsBulkUpdateOK {
	return &DcimPlatformsBulkUpdateOK{}
}

/*
DcimPlatformsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsBulkUpdateOK dcim platforms bulk update o k
*/
type DcimPlatformsBulkUpdateOK struct {
	Payload *models.Platform
}

// IsSuccess returns true when this dcim platforms bulk update o k response has a 2xx status code
func (o *DcimPlatformsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms bulk update o k response has a 3xx status code
func (o *DcimPlatformsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms bulk update o k response has a 4xx status code
func (o *DcimPlatformsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms bulk update o k response has a 5xx status code
func (o *DcimPlatformsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms bulk update o k response a status code equal to that given
func (o *DcimPlatformsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPlatformsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/platforms/][%d] dcimPlatformsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/platforms/][%d] dcimPlatformsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsBulkUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsBulkUpdateDefault creates a DcimPlatformsBulkUpdateDefault with default headers values
func NewDcimPlatformsBulkUpdateDefault(code int) *DcimPlatformsBulkUpdateDefault {
	return &DcimPlatformsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPlatformsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimPlatformsBulkUpdateDefault dcim platforms bulk update default
*/
type DcimPlatformsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms bulk update default response
func (o *DcimPlatformsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim platforms bulk update default response has a 2xx status code
func (o *DcimPlatformsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim platforms bulk update default response has a 3xx status code
func (o *DcimPlatformsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim platforms bulk update default response has a 4xx status code
func (o *DcimPlatformsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim platforms bulk update default response has a 5xx status code
func (o *DcimPlatformsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim platforms bulk update default response a status code equal to that given
func (o *DcimPlatformsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPlatformsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/platforms/][%d] dcim_platforms_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/platforms/][%d] dcim_platforms_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
