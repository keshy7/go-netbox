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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamIPAddressesBulkUpdateReader is a Reader for the IpamIPAddressesBulkUpdate structure.
type IpamIPAddressesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPAddressesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPAddressesBulkUpdateOK creates a IpamIPAddressesBulkUpdateOK with default headers values
func NewIpamIPAddressesBulkUpdateOK() *IpamIPAddressesBulkUpdateOK {
	return &IpamIPAddressesBulkUpdateOK{}
}

/*
IpamIPAddressesBulkUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesBulkUpdateOK ipam Ip addresses bulk update o k
*/
type IpamIPAddressesBulkUpdateOK struct {
	Payload *models.IPAddress
}

// IsSuccess returns true when this ipam Ip addresses bulk update o k response has a 2xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip addresses bulk update o k response has a 3xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip addresses bulk update o k response has a 4xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip addresses bulk update o k response has a 5xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip addresses bulk update o k response a status code equal to that given
func (o *IpamIPAddressesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamIPAddressesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/][%d] ipamIpAddressesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/][%d] ipamIpAddressesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesBulkUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPAddressesBulkUpdateDefault creates a IpamIPAddressesBulkUpdateDefault with default headers values
func NewIpamIPAddressesBulkUpdateDefault(code int) *IpamIPAddressesBulkUpdateDefault {
	return &IpamIPAddressesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamIPAddressesBulkUpdateDefault describes a response with status code -1, with default header values.

IpamIPAddressesBulkUpdateDefault ipam ip addresses bulk update default
*/
type IpamIPAddressesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip addresses bulk update default response
func (o *IpamIPAddressesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam ip addresses bulk update default response has a 2xx status code
func (o *IpamIPAddressesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip addresses bulk update default response has a 3xx status code
func (o *IpamIPAddressesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip addresses bulk update default response has a 4xx status code
func (o *IpamIPAddressesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip addresses bulk update default response has a 5xx status code
func (o *IpamIPAddressesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip addresses bulk update default response a status code equal to that given
func (o *IpamIPAddressesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamIPAddressesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/][%d] ipam_ip-addresses_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/][%d] ipam_ip-addresses_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
