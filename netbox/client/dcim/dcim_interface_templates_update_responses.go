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

// DcimInterfaceTemplatesUpdateReader is a Reader for the DcimInterfaceTemplatesUpdate structure.
type DcimInterfaceTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfaceTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfaceTemplatesUpdateOK creates a DcimInterfaceTemplatesUpdateOK with default headers values
func NewDcimInterfaceTemplatesUpdateOK() *DcimInterfaceTemplatesUpdateOK {
	return &DcimInterfaceTemplatesUpdateOK{}
}

/*
DcimInterfaceTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesUpdateOK dcim interface templates update o k
*/
type DcimInterfaceTemplatesUpdateOK struct {
	Payload *models.InterfaceTemplate
}

// IsSuccess returns true when this dcim interface templates update o k response has a 2xx status code
func (o *DcimInterfaceTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates update o k response has a 3xx status code
func (o *DcimInterfaceTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates update o k response has a 4xx status code
func (o *DcimInterfaceTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates update o k response has a 5xx status code
func (o *DcimInterfaceTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates update o k response a status code equal to that given
func (o *DcimInterfaceTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfaceTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfaceTemplatesUpdateDefault creates a DcimInterfaceTemplatesUpdateDefault with default headers values
func NewDcimInterfaceTemplatesUpdateDefault(code int) *DcimInterfaceTemplatesUpdateDefault {
	return &DcimInterfaceTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInterfaceTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimInterfaceTemplatesUpdateDefault dcim interface templates update default
*/
type DcimInterfaceTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interface templates update default response
func (o *DcimInterfaceTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interface templates update default response has a 2xx status code
func (o *DcimInterfaceTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interface templates update default response has a 3xx status code
func (o *DcimInterfaceTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interface templates update default response has a 4xx status code
func (o *DcimInterfaceTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interface templates update default response has a 5xx status code
func (o *DcimInterfaceTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interface templates update default response a status code equal to that given
func (o *DcimInterfaceTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfaceTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/{id}/][%d] dcim_interface-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/{id}/][%d] dcim_interface-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
