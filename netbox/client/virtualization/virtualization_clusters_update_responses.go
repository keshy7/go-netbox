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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualizationClustersUpdateReader is a Reader for the VirtualizationClustersUpdate structure.
type VirtualizationClustersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClustersUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersUpdateOK creates a VirtualizationClustersUpdateOK with default headers values
func NewVirtualizationClustersUpdateOK() *VirtualizationClustersUpdateOK {
	return &VirtualizationClustersUpdateOK{}
}

/*
VirtualizationClustersUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersUpdateOK virtualization clusters update o k
*/
type VirtualizationClustersUpdateOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters update o k response has a 2xx status code
func (o *VirtualizationClustersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters update o k response has a 3xx status code
func (o *VirtualizationClustersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters update o k response has a 4xx status code
func (o *VirtualizationClustersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters update o k response has a 5xx status code
func (o *VirtualizationClustersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters update o k response a status code equal to that given
func (o *VirtualizationClustersUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClustersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/{id}/][%d] virtualizationClustersUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/{id}/][%d] virtualizationClustersUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersUpdateDefault creates a VirtualizationClustersUpdateDefault with default headers values
func NewVirtualizationClustersUpdateDefault(code int) *VirtualizationClustersUpdateDefault {
	return &VirtualizationClustersUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClustersUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClustersUpdateDefault virtualization clusters update default
*/
type VirtualizationClustersUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization clusters update default response
func (o *VirtualizationClustersUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization clusters update default response has a 2xx status code
func (o *VirtualizationClustersUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization clusters update default response has a 3xx status code
func (o *VirtualizationClustersUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization clusters update default response has a 4xx status code
func (o *VirtualizationClustersUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization clusters update default response has a 5xx status code
func (o *VirtualizationClustersUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization clusters update default response a status code equal to that given
func (o *VirtualizationClustersUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClustersUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/{id}/][%d] virtualization_clusters_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/{id}/][%d] virtualization_clusters_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
