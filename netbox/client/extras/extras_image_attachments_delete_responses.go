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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasImageAttachmentsDeleteReader is a Reader for the ExtrasImageAttachmentsDelete structure.
type ExtrasImageAttachmentsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasImageAttachmentsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsDeleteNoContent creates a ExtrasImageAttachmentsDeleteNoContent with default headers values
func NewExtrasImageAttachmentsDeleteNoContent() *ExtrasImageAttachmentsDeleteNoContent {
	return &ExtrasImageAttachmentsDeleteNoContent{}
}

/*
ExtrasImageAttachmentsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasImageAttachmentsDeleteNoContent extras image attachments delete no content
*/
type ExtrasImageAttachmentsDeleteNoContent struct {
}

// IsSuccess returns true when this extras image attachments delete no content response has a 2xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments delete no content response has a 3xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments delete no content response has a 4xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments delete no content response has a 5xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments delete no content response a status code equal to that given
func (o *ExtrasImageAttachmentsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasImageAttachmentsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extrasImageAttachmentsDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extrasImageAttachmentsDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasImageAttachmentsDeleteDefault creates a ExtrasImageAttachmentsDeleteDefault with default headers values
func NewExtrasImageAttachmentsDeleteDefault(code int) *ExtrasImageAttachmentsDeleteDefault {
	return &ExtrasImageAttachmentsDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsDeleteDefault describes a response with status code -1, with default header values.

ExtrasImageAttachmentsDeleteDefault extras image attachments delete default
*/
type ExtrasImageAttachmentsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras image attachments delete default response
func (o *ExtrasImageAttachmentsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras image attachments delete default response has a 2xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments delete default response has a 3xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments delete default response has a 4xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments delete default response has a 5xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments delete default response a status code equal to that given
func (o *ExtrasImageAttachmentsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasImageAttachmentsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extras_image-attachments_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extras_image-attachments_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
