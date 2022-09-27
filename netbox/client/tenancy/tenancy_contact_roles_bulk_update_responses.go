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

// TenancyContactRolesBulkUpdateReader is a Reader for the TenancyContactRolesBulkUpdate structure.
type TenancyContactRolesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactRolesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactRolesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesBulkUpdateOK creates a TenancyContactRolesBulkUpdateOK with default headers values
func NewTenancyContactRolesBulkUpdateOK() *TenancyContactRolesBulkUpdateOK {
	return &TenancyContactRolesBulkUpdateOK{}
}

/*
TenancyContactRolesBulkUpdateOK describes a response with status code 200, with default header values.

TenancyContactRolesBulkUpdateOK tenancy contact roles bulk update o k
*/
type TenancyContactRolesBulkUpdateOK struct {
	Payload *models.ContactRole
}

// IsSuccess returns true when this tenancy contact roles bulk update o k response has a 2xx status code
func (o *TenancyContactRolesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles bulk update o k response has a 3xx status code
func (o *TenancyContactRolesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles bulk update o k response has a 4xx status code
func (o *TenancyContactRolesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles bulk update o k response has a 5xx status code
func (o *TenancyContactRolesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles bulk update o k response a status code equal to that given
func (o *TenancyContactRolesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactRolesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/][%d] tenancyContactRolesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/][%d] tenancyContactRolesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesBulkUpdateOK) GetPayload() *models.ContactRole {
	return o.Payload
}

func (o *TenancyContactRolesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactRolesBulkUpdateDefault creates a TenancyContactRolesBulkUpdateDefault with default headers values
func NewTenancyContactRolesBulkUpdateDefault(code int) *TenancyContactRolesBulkUpdateDefault {
	return &TenancyContactRolesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactRolesBulkUpdateDefault describes a response with status code -1, with default header values.

TenancyContactRolesBulkUpdateDefault tenancy contact roles bulk update default
*/
type TenancyContactRolesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact roles bulk update default response
func (o *TenancyContactRolesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact roles bulk update default response has a 2xx status code
func (o *TenancyContactRolesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact roles bulk update default response has a 3xx status code
func (o *TenancyContactRolesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact roles bulk update default response has a 4xx status code
func (o *TenancyContactRolesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact roles bulk update default response has a 5xx status code
func (o *TenancyContactRolesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact roles bulk update default response a status code equal to that given
func (o *TenancyContactRolesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactRolesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/][%d] tenancy_contact-roles_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/][%d] tenancy_contact-roles_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
