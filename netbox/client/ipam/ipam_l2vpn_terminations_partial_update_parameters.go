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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewIpamL2vpnTerminationsPartialUpdateParams creates a new IpamL2vpnTerminationsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamL2vpnTerminationsPartialUpdateParams() *IpamL2vpnTerminationsPartialUpdateParams {
	return &IpamL2vpnTerminationsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamL2vpnTerminationsPartialUpdateParamsWithTimeout creates a new IpamL2vpnTerminationsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamL2vpnTerminationsPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamL2vpnTerminationsPartialUpdateParams {
	return &IpamL2vpnTerminationsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamL2vpnTerminationsPartialUpdateParamsWithContext creates a new IpamL2vpnTerminationsPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamL2vpnTerminationsPartialUpdateParamsWithContext(ctx context.Context) *IpamL2vpnTerminationsPartialUpdateParams {
	return &IpamL2vpnTerminationsPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamL2vpnTerminationsPartialUpdateParamsWithHTTPClient creates a new IpamL2vpnTerminationsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamL2vpnTerminationsPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamL2vpnTerminationsPartialUpdateParams {
	return &IpamL2vpnTerminationsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamL2vpnTerminationsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam l2vpn terminations partial update operation.

	Typically these are written to a http.Request.
*/
type IpamL2vpnTerminationsPartialUpdateParams struct {

	// Data.
	Data *models.WritableL2VPNTermination

	/* ID.

	   A unique integer value identifying this L2VPN termination.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam l2vpn terminations partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsPartialUpdateParams) WithDefaults() *IpamL2vpnTerminationsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam l2vpn terminations partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamL2vpnTerminationsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) WithContext(ctx context.Context) *IpamL2vpnTerminationsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamL2vpnTerminationsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) WithData(data *models.WritableL2VPNTermination) *IpamL2vpnTerminationsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) SetData(data *models.WritableL2VPNTermination) {
	o.Data = data
}

// WithID adds the id to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) WithID(id int64) *IpamL2vpnTerminationsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam l2vpn terminations partial update params
func (o *IpamL2vpnTerminationsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamL2vpnTerminationsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
