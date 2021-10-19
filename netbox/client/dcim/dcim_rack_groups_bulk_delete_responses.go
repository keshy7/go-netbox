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
)

// DcimRackGroupsBulkDeleteReader is a Reader for the DcimRackGroupsBulkDelete structure.
type DcimRackGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackGroupsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackGroupsBulkDeleteNoContent creates a DcimRackGroupsBulkDeleteNoContent with default headers values
func NewDcimRackGroupsBulkDeleteNoContent() *DcimRackGroupsBulkDeleteNoContent {
	return &DcimRackGroupsBulkDeleteNoContent{}
}

/* DcimRackGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRackGroupsBulkDeleteNoContent dcim rack groups bulk delete no content
*/
type DcimRackGroupsBulkDeleteNoContent struct {
}

func (o *DcimRackGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-groups/][%d] dcimRackGroupsBulkDeleteNoContent ", 204)
}

func (o *DcimRackGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRackGroupsBulkDeleteDefault creates a DcimRackGroupsBulkDeleteDefault with default headers values
func NewDcimRackGroupsBulkDeleteDefault(code int) *DcimRackGroupsBulkDeleteDefault {
	return &DcimRackGroupsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* DcimRackGroupsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimRackGroupsBulkDeleteDefault dcim rack groups bulk delete default
*/
type DcimRackGroupsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack groups bulk delete default response
func (o *DcimRackGroupsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackGroupsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-groups/][%d] dcim_rack-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRackGroupsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackGroupsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}