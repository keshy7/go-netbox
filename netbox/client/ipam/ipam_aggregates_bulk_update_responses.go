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

// IpamAggregatesBulkUpdateReader is a Reader for the IpamAggregatesBulkUpdate structure.
type IpamAggregatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAggregatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAggregatesBulkUpdateOK creates a IpamAggregatesBulkUpdateOK with default headers values
func NewIpamAggregatesBulkUpdateOK() *IpamAggregatesBulkUpdateOK {
	return &IpamAggregatesBulkUpdateOK{}
}

/*
IpamAggregatesBulkUpdateOK describes a response with status code 200, with default header values.

IpamAggregatesBulkUpdateOK ipam aggregates bulk update o k
*/
type IpamAggregatesBulkUpdateOK struct {
	Payload *models.Aggregate
}

// IsSuccess returns true when this ipam aggregates bulk update o k response has a 2xx status code
func (o *IpamAggregatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam aggregates bulk update o k response has a 3xx status code
func (o *IpamAggregatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates bulk update o k response has a 4xx status code
func (o *IpamAggregatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam aggregates bulk update o k response has a 5xx status code
func (o *IpamAggregatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates bulk update o k response a status code equal to that given
func (o *IpamAggregatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamAggregatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/aggregates/][%d] ipamAggregatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/aggregates/][%d] ipamAggregatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesBulkUpdateOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAggregatesBulkUpdateDefault creates a IpamAggregatesBulkUpdateDefault with default headers values
func NewIpamAggregatesBulkUpdateDefault(code int) *IpamAggregatesBulkUpdateDefault {
	return &IpamAggregatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamAggregatesBulkUpdateDefault describes a response with status code -1, with default header values.

IpamAggregatesBulkUpdateDefault ipam aggregates bulk update default
*/
type IpamAggregatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam aggregates bulk update default response
func (o *IpamAggregatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam aggregates bulk update default response has a 2xx status code
func (o *IpamAggregatesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam aggregates bulk update default response has a 3xx status code
func (o *IpamAggregatesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam aggregates bulk update default response has a 4xx status code
func (o *IpamAggregatesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam aggregates bulk update default response has a 5xx status code
func (o *IpamAggregatesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam aggregates bulk update default response a status code equal to that given
func (o *IpamAggregatesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamAggregatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/aggregates/][%d] ipam_aggregates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/aggregates/][%d] ipam_aggregates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAggregatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
