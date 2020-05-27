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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimPowerOutletsCreateReader is a Reader for the DcimPowerOutletsCreate structure.
type DcimPowerOutletsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerOutletsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletsCreateCreated creates a DcimPowerOutletsCreateCreated with default headers values
func NewDcimPowerOutletsCreateCreated() *DcimPowerOutletsCreateCreated {
	return &DcimPowerOutletsCreateCreated{}
}

/*DcimPowerOutletsCreateCreated handles this case with default header values.

DcimPowerOutletsCreateCreated dcim power outlets create created
*/
type DcimPowerOutletsCreateCreated struct {
	Payload *models.PowerOutlet
}

func (o *DcimPowerOutletsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-outlets/][%d] dcimPowerOutletsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerOutletsCreateCreated) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
