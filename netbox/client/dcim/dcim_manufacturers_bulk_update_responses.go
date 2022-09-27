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

// DcimManufacturersBulkUpdateReader is a Reader for the DcimManufacturersBulkUpdate structure.
type DcimManufacturersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimManufacturersBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersBulkUpdateOK creates a DcimManufacturersBulkUpdateOK with default headers values
func NewDcimManufacturersBulkUpdateOK() *DcimManufacturersBulkUpdateOK {
	return &DcimManufacturersBulkUpdateOK{}
}

/*
DcimManufacturersBulkUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersBulkUpdateOK dcim manufacturers bulk update o k
*/
type DcimManufacturersBulkUpdateOK struct {
	Payload *models.Manufacturer
}

// IsSuccess returns true when this dcim manufacturers bulk update o k response has a 2xx status code
func (o *DcimManufacturersBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim manufacturers bulk update o k response has a 3xx status code
func (o *DcimManufacturersBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers bulk update o k response has a 4xx status code
func (o *DcimManufacturersBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim manufacturers bulk update o k response has a 5xx status code
func (o *DcimManufacturersBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers bulk update o k response a status code equal to that given
func (o *DcimManufacturersBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimManufacturersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/manufacturers/][%d] dcimManufacturersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/manufacturers/][%d] dcimManufacturersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersBulkUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersBulkUpdateDefault creates a DcimManufacturersBulkUpdateDefault with default headers values
func NewDcimManufacturersBulkUpdateDefault(code int) *DcimManufacturersBulkUpdateDefault {
	return &DcimManufacturersBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimManufacturersBulkUpdateDefault describes a response with status code -1, with default header values.

DcimManufacturersBulkUpdateDefault dcim manufacturers bulk update default
*/
type DcimManufacturersBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim manufacturers bulk update default response
func (o *DcimManufacturersBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim manufacturers bulk update default response has a 2xx status code
func (o *DcimManufacturersBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim manufacturers bulk update default response has a 3xx status code
func (o *DcimManufacturersBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim manufacturers bulk update default response has a 4xx status code
func (o *DcimManufacturersBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim manufacturers bulk update default response has a 5xx status code
func (o *DcimManufacturersBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim manufacturers bulk update default response a status code equal to that given
func (o *DcimManufacturersBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimManufacturersBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/manufacturers/][%d] dcim_manufacturers_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/manufacturers/][%d] dcim_manufacturers_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
