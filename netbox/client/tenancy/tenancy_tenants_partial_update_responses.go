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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// TenancyTenantsPartialUpdateReader is a Reader for the TenancyTenantsPartialUpdate structure.
type TenancyTenantsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantsPartialUpdateOK creates a TenancyTenantsPartialUpdateOK with default headers values
func NewTenancyTenantsPartialUpdateOK() *TenancyTenantsPartialUpdateOK {
	return &TenancyTenantsPartialUpdateOK{}
}

/*
TenancyTenantsPartialUpdateOK describes a response with status code 200, with default header values.

TenancyTenantsPartialUpdateOK tenancy tenants partial update o k
*/
type TenancyTenantsPartialUpdateOK struct {
	Payload *models.Tenant
}

// IsSuccess returns true when this tenancy tenants partial update o k response has a 2xx status code
func (o *TenancyTenantsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenants partial update o k response has a 3xx status code
func (o *TenancyTenantsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenants partial update o k response has a 4xx status code
func (o *TenancyTenantsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenants partial update o k response has a 5xx status code
func (o *TenancyTenantsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenants partial update o k response a status code equal to that given
func (o *TenancyTenantsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyTenantsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenants/{id}/][%d] tenancyTenantsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /tenancy/tenants/{id}/][%d] tenancyTenantsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantsPartialUpdateOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantsPartialUpdateDefault creates a TenancyTenantsPartialUpdateDefault with default headers values
func NewTenancyTenantsPartialUpdateDefault(code int) *TenancyTenantsPartialUpdateDefault {
	return &TenancyTenantsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyTenantsPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyTenantsPartialUpdateDefault tenancy tenants partial update default
*/
type TenancyTenantsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenants partial update default response
func (o *TenancyTenantsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy tenants partial update default response has a 2xx status code
func (o *TenancyTenantsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy tenants partial update default response has a 3xx status code
func (o *TenancyTenantsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy tenants partial update default response has a 4xx status code
func (o *TenancyTenantsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy tenants partial update default response has a 5xx status code
func (o *TenancyTenantsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy tenants partial update default response a status code equal to that given
func (o *TenancyTenantsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyTenantsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/tenants/{id}/][%d] tenancy_tenants_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /tenancy/tenants/{id}/][%d] tenancy_tenants_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
