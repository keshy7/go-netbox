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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimFrontPortsBulkUpdateParams creates a new DcimFrontPortsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsBulkUpdateParams() *DcimFrontPortsBulkUpdateParams {
	return &DcimFrontPortsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsBulkUpdateParamsWithTimeout creates a new DcimFrontPortsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortsBulkUpdateParams {
	return &DcimFrontPortsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsBulkUpdateParamsWithContext creates a new DcimFrontPortsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsBulkUpdateParamsWithContext(ctx context.Context) *DcimFrontPortsBulkUpdateParams {
	return &DcimFrontPortsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsBulkUpdateParamsWithHTTPClient creates a new DcimFrontPortsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortsBulkUpdateParams {
	return &DcimFrontPortsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimFrontPortsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim front ports bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimFrontPortsBulkUpdateParams struct {

	// Data.
	Data *models.WritableFrontPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsBulkUpdateParams) WithDefaults() *DcimFrontPortsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) WithContext(ctx context.Context) *DcimFrontPortsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) WithData(data *models.WritableFrontPort) *DcimFrontPortsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front ports bulk update params
func (o *DcimFrontPortsBulkUpdateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
