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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewExtrasCustomLinksBulkUpdateParams creates a new ExtrasCustomLinksBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomLinksBulkUpdateParams() *ExtrasCustomLinksBulkUpdateParams {
	return &ExtrasCustomLinksBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomLinksBulkUpdateParamsWithTimeout creates a new ExtrasCustomLinksBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomLinksBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasCustomLinksBulkUpdateParams {
	return &ExtrasCustomLinksBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomLinksBulkUpdateParamsWithContext creates a new ExtrasCustomLinksBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasCustomLinksBulkUpdateParamsWithContext(ctx context.Context) *ExtrasCustomLinksBulkUpdateParams {
	return &ExtrasCustomLinksBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasCustomLinksBulkUpdateParamsWithHTTPClient creates a new ExtrasCustomLinksBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomLinksBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasCustomLinksBulkUpdateParams {
	return &ExtrasCustomLinksBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasCustomLinksBulkUpdateParams contains all the parameters to send to the API endpoint

	for the extras custom links bulk update operation.

	Typically these are written to a http.Request.
*/
type ExtrasCustomLinksBulkUpdateParams struct {

	// Data.
	Data *models.CustomLink

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom links bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksBulkUpdateParams) WithDefaults() *ExtrasCustomLinksBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom links bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasCustomLinksBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) WithContext(ctx context.Context) *ExtrasCustomLinksBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasCustomLinksBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) WithData(data *models.CustomLink) *ExtrasCustomLinksBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom links bulk update params
func (o *ExtrasCustomLinksBulkUpdateParams) SetData(data *models.CustomLink) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomLinksBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
