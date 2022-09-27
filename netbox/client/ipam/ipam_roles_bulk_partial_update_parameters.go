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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewIpamRolesBulkPartialUpdateParams creates a new IpamRolesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRolesBulkPartialUpdateParams() *IpamRolesBulkPartialUpdateParams {
	return &IpamRolesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRolesBulkPartialUpdateParamsWithTimeout creates a new IpamRolesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRolesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamRolesBulkPartialUpdateParams {
	return &IpamRolesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRolesBulkPartialUpdateParamsWithContext creates a new IpamRolesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamRolesBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamRolesBulkPartialUpdateParams {
	return &IpamRolesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamRolesBulkPartialUpdateParamsWithHTTPClient creates a new IpamRolesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRolesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamRolesBulkPartialUpdateParams {
	return &IpamRolesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamRolesBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam roles bulk partial update operation.

	Typically these are written to a http.Request.
*/
type IpamRolesBulkPartialUpdateParams struct {

	// Data.
	Data *models.Role

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam roles bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesBulkPartialUpdateParams) WithDefaults() *IpamRolesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam roles bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamRolesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamRolesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamRolesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) WithData(data *models.Role) *IpamRolesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam roles bulk partial update params
func (o *IpamRolesBulkPartialUpdateParams) SetData(data *models.Role) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRolesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
