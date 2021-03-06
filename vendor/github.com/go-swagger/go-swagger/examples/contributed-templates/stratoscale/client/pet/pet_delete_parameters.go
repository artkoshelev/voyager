// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPetDeleteParams creates a new PetDeleteParams object
// with the default values initialized.
func NewPetDeleteParams() *PetDeleteParams {
	var ()
	return &PetDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPetDeleteParamsWithTimeout creates a new PetDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPetDeleteParamsWithTimeout(timeout time.Duration) *PetDeleteParams {
	var ()
	return &PetDeleteParams{

		timeout: timeout,
	}
}

// NewPetDeleteParamsWithContext creates a new PetDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPetDeleteParamsWithContext(ctx context.Context) *PetDeleteParams {
	var ()
	return &PetDeleteParams{

		Context: ctx,
	}
}

// NewPetDeleteParamsWithHTTPClient creates a new PetDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPetDeleteParamsWithHTTPClient(client *http.Client) *PetDeleteParams {
	var ()
	return &PetDeleteParams{
		HTTPClient: client,
	}
}

/*PetDeleteParams contains all the parameters to send to the API endpoint
for the pet delete operation typically these are written to a http.Request
*/
type PetDeleteParams struct {

	/*APIKey*/
	APIKey *string
	/*PetID
	  Pet id to delete

	*/
	PetID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pet delete params
func (o *PetDeleteParams) WithTimeout(timeout time.Duration) *PetDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pet delete params
func (o *PetDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pet delete params
func (o *PetDeleteParams) WithContext(ctx context.Context) *PetDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pet delete params
func (o *PetDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pet delete params
func (o *PetDeleteParams) WithHTTPClient(client *http.Client) *PetDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pet delete params
func (o *PetDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the pet delete params
func (o *PetDeleteParams) WithAPIKey(aPIKey *string) *PetDeleteParams {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the pet delete params
func (o *PetDeleteParams) SetAPIKey(aPIKey *string) {
	o.APIKey = aPIKey
}

// WithPetID adds the petID to the pet delete params
func (o *PetDeleteParams) WithPetID(petID int64) *PetDeleteParams {
	o.SetPetID(petID)
	return o
}

// SetPetID adds the petId to the pet delete params
func (o *PetDeleteParams) SetPetID(petID int64) {
	o.PetID = petID
}

// WriteToRequest writes these params to a swagger request
func (o *PetDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIKey != nil {

		// header param api_key
		if err := r.SetHeaderParam("api_key", *o.APIKey); err != nil {
			return err
		}

	}

	// path param petId
	if err := r.SetPathParam("petId", swag.FormatInt64(o.PetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
