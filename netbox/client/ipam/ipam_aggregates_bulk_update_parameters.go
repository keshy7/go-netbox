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

// NewIpamAggregatesBulkUpdateParams creates a new IpamAggregatesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesBulkUpdateParams() *IpamAggregatesBulkUpdateParams {
	return &IpamAggregatesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesBulkUpdateParamsWithTimeout creates a new IpamAggregatesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamAggregatesBulkUpdateParams {
	return &IpamAggregatesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesBulkUpdateParamsWithContext creates a new IpamAggregatesBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamAggregatesBulkUpdateParamsWithContext(ctx context.Context) *IpamAggregatesBulkUpdateParams {
	return &IpamAggregatesBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamAggregatesBulkUpdateParamsWithHTTPClient creates a new IpamAggregatesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamAggregatesBulkUpdateParams {
	return &IpamAggregatesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamAggregatesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the ipam aggregates bulk update operation.

	Typically these are written to a http.Request.
*/
type IpamAggregatesBulkUpdateParams struct {

	// Data.
	Data *models.WritableAggregate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesBulkUpdateParams) WithDefaults() *IpamAggregatesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamAggregatesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) WithContext(ctx context.Context) *IpamAggregatesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamAggregatesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) WithData(data *models.WritableAggregate) *IpamAggregatesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam aggregates bulk update params
func (o *IpamAggregatesBulkUpdateParams) SetData(data *models.WritableAggregate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
