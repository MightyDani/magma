// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLTENetworkIDCellularParams creates a new GetLTENetworkIDCellularParams object
// with the default values initialized.
func NewGetLTENetworkIDCellularParams() *GetLTENetworkIDCellularParams {
	var ()
	return &GetLTENetworkIDCellularParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDCellularParamsWithTimeout creates a new GetLTENetworkIDCellularParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDCellularParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDCellularParams {
	var ()
	return &GetLTENetworkIDCellularParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDCellularParamsWithContext creates a new GetLTENetworkIDCellularParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDCellularParamsWithContext(ctx context.Context) *GetLTENetworkIDCellularParams {
	var ()
	return &GetLTENetworkIDCellularParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDCellularParamsWithHTTPClient creates a new GetLTENetworkIDCellularParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDCellularParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDCellularParams {
	var ()
	return &GetLTENetworkIDCellularParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDCellularParams contains all the parameters to send to the API endpoint
for the get LTE network ID cellular operation typically these are written to a http.Request
*/
type GetLTENetworkIDCellularParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDCellularParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) WithContext(ctx context.Context) *GetLTENetworkIDCellularParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDCellularParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) WithNetworkID(networkID string) *GetLTENetworkIDCellularParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID cellular params
func (o *GetLTENetworkIDCellularParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDCellularParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
