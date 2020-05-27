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

package secrets

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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewSecretsSecretRolesCreateParams creates a new SecretsSecretRolesCreateParams object
// with the default values initialized.
func NewSecretsSecretRolesCreateParams() *SecretsSecretRolesCreateParams {
	var ()
	return &SecretsSecretRolesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretRolesCreateParamsWithTimeout creates a new SecretsSecretRolesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretRolesCreateParamsWithTimeout(timeout time.Duration) *SecretsSecretRolesCreateParams {
	var ()
	return &SecretsSecretRolesCreateParams{

		timeout: timeout,
	}
}

// NewSecretsSecretRolesCreateParamsWithContext creates a new SecretsSecretRolesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretRolesCreateParamsWithContext(ctx context.Context) *SecretsSecretRolesCreateParams {
	var ()
	return &SecretsSecretRolesCreateParams{

		Context: ctx,
	}
}

// NewSecretsSecretRolesCreateParamsWithHTTPClient creates a new SecretsSecretRolesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretRolesCreateParamsWithHTTPClient(client *http.Client) *SecretsSecretRolesCreateParams {
	var ()
	return &SecretsSecretRolesCreateParams{
		HTTPClient: client,
	}
}

/*SecretsSecretRolesCreateParams contains all the parameters to send to the API endpoint
for the secrets secret roles create operation typically these are written to a http.Request
*/
type SecretsSecretRolesCreateParams struct {

	/*Data*/
	Data *models.SecretRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) WithTimeout(timeout time.Duration) *SecretsSecretRolesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) WithContext(ctx context.Context) *SecretsSecretRolesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) WithHTTPClient(client *http.Client) *SecretsSecretRolesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) WithData(data *models.SecretRole) *SecretsSecretRolesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the secrets secret roles create params
func (o *SecretsSecretRolesCreateParams) SetData(data *models.SecretRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretRolesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
