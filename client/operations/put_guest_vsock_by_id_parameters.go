// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	client_models "github.com/firecracker-microvm/go-firecracker/client/models"
)

// NewPutGuestVsockByIDParams creates a new PutGuestVsockByIDParams object
// with the default values initialized.
func NewPutGuestVsockByIDParams() *PutGuestVsockByIDParams {
	var ()
	return &PutGuestVsockByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutGuestVsockByIDParamsWithTimeout creates a new PutGuestVsockByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutGuestVsockByIDParamsWithTimeout(timeout time.Duration) *PutGuestVsockByIDParams {
	var ()
	return &PutGuestVsockByIDParams{

		timeout: timeout,
	}
}

// NewPutGuestVsockByIDParamsWithContext creates a new PutGuestVsockByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutGuestVsockByIDParamsWithContext(ctx context.Context) *PutGuestVsockByIDParams {
	var ()
	return &PutGuestVsockByIDParams{

		Context: ctx,
	}
}

// NewPutGuestVsockByIDParamsWithHTTPClient creates a new PutGuestVsockByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutGuestVsockByIDParamsWithHTTPClient(client *http.Client) *PutGuestVsockByIDParams {
	var ()
	return &PutGuestVsockByIDParams{
		HTTPClient: client,
	}
}

/*PutGuestVsockByIDParams contains all the parameters to send to the API endpoint
for the put guest vsock by ID operation typically these are written to a http.Request
*/
type PutGuestVsockByIDParams struct {

	/*Body
	  Guest vsock properties

	*/
	Body *client_models.Vsock
	/*ID
	  The id of the vsock device

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) WithTimeout(timeout time.Duration) *PutGuestVsockByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) WithContext(ctx context.Context) *PutGuestVsockByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) WithHTTPClient(client *http.Client) *PutGuestVsockByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) WithBody(body *client_models.Vsock) *PutGuestVsockByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) SetBody(body *client_models.Vsock) {
	o.Body = body
}

// WithID adds the id to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) WithID(id string) *PutGuestVsockByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put guest vsock by ID params
func (o *PutGuestVsockByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutGuestVsockByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}