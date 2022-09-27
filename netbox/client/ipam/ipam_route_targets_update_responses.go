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

// IpamRouteTargetsUpdateReader is a Reader for the IpamRouteTargetsUpdate structure.
type IpamRouteTargetsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRouteTargetsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsUpdateOK creates a IpamRouteTargetsUpdateOK with default headers values
func NewIpamRouteTargetsUpdateOK() *IpamRouteTargetsUpdateOK {
	return &IpamRouteTargetsUpdateOK{}
}

/*
IpamRouteTargetsUpdateOK describes a response with status code 200, with default header values.

IpamRouteTargetsUpdateOK ipam route targets update o k
*/
type IpamRouteTargetsUpdateOK struct {
	Payload *models.RouteTarget
}

// IsSuccess returns true when this ipam route targets update o k response has a 2xx status code
func (o *IpamRouteTargetsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam route targets update o k response has a 3xx status code
func (o *IpamRouteTargetsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets update o k response has a 4xx status code
func (o *IpamRouteTargetsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam route targets update o k response has a 5xx status code
func (o *IpamRouteTargetsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets update o k response a status code equal to that given
func (o *IpamRouteTargetsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamRouteTargetsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipamRouteTargetsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRouteTargetsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipamRouteTargetsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRouteTargetsUpdateOK) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsUpdateDefault creates a IpamRouteTargetsUpdateDefault with default headers values
func NewIpamRouteTargetsUpdateDefault(code int) *IpamRouteTargetsUpdateDefault {
	return &IpamRouteTargetsUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRouteTargetsUpdateDefault describes a response with status code -1, with default header values.

IpamRouteTargetsUpdateDefault ipam route targets update default
*/
type IpamRouteTargetsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets update default response
func (o *IpamRouteTargetsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam route targets update default response has a 2xx status code
func (o *IpamRouteTargetsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam route targets update default response has a 3xx status code
func (o *IpamRouteTargetsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam route targets update default response has a 4xx status code
func (o *IpamRouteTargetsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam route targets update default response has a 5xx status code
func (o *IpamRouteTargetsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam route targets update default response a status code equal to that given
func (o *IpamRouteTargetsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRouteTargetsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipam_route-targets_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipam_route-targets_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
