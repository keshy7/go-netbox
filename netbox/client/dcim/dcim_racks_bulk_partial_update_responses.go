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

// DcimRacksBulkPartialUpdateReader is a Reader for the DcimRacksBulkPartialUpdate structure.
type DcimRacksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksBulkPartialUpdateOK creates a DcimRacksBulkPartialUpdateOK with default headers values
func NewDcimRacksBulkPartialUpdateOK() *DcimRacksBulkPartialUpdateOK {
	return &DcimRacksBulkPartialUpdateOK{}
}

/*
DcimRacksBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRacksBulkPartialUpdateOK dcim racks bulk partial update o k
*/
type DcimRacksBulkPartialUpdateOK struct {
	Payload *models.Rack
}

// IsSuccess returns true when this dcim racks bulk partial update o k response has a 2xx status code
func (o *DcimRacksBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks bulk partial update o k response has a 3xx status code
func (o *DcimRacksBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks bulk partial update o k response has a 4xx status code
func (o *DcimRacksBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks bulk partial update o k response has a 5xx status code
func (o *DcimRacksBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks bulk partial update o k response a status code equal to that given
func (o *DcimRacksBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRacksBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/racks/][%d] dcimRacksBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRacksBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/racks/][%d] dcimRacksBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRacksBulkPartialUpdateOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksBulkPartialUpdateDefault creates a DcimRacksBulkPartialUpdateDefault with default headers values
func NewDcimRacksBulkPartialUpdateDefault(code int) *DcimRacksBulkPartialUpdateDefault {
	return &DcimRacksBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRacksBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimRacksBulkPartialUpdateDefault dcim racks bulk partial update default
*/
type DcimRacksBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim racks bulk partial update default response
func (o *DcimRacksBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim racks bulk partial update default response has a 2xx status code
func (o *DcimRacksBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks bulk partial update default response has a 3xx status code
func (o *DcimRacksBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks bulk partial update default response has a 4xx status code
func (o *DcimRacksBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks bulk partial update default response has a 5xx status code
func (o *DcimRacksBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks bulk partial update default response a status code equal to that given
func (o *DcimRacksBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRacksBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/racks/][%d] dcim_racks_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/racks/][%d] dcim_racks_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
