// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

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

// NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParams creates a new GetWifiNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized.
func NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParams() *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDDescriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout creates a new GetWifiNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParamsWithTimeout(timeout time.Duration) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDDescriptionParams{

		timeout: timeout,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParamsWithContext creates a new GetWifiNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParamsWithContext(ctx context.Context) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDDescriptionParams{

		Context: ctx,
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient creates a new GetWifiNetworkIDGatewaysGatewayIDDescriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWifiNetworkIDGatewaysGatewayIDDescriptionParamsWithHTTPClient(client *http.Client) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	var ()
	return &GetWifiNetworkIDGatewaysGatewayIDDescriptionParams{
		HTTPClient: client,
	}
}

/*GetWifiNetworkIDGatewaysGatewayIDDescriptionParams contains all the parameters to send to the API endpoint
for the get wifi network ID gateways gateway ID description operation typically these are written to a http.Request
*/
type GetWifiNetworkIDGatewaysGatewayIDDescriptionParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) WithTimeout(timeout time.Duration) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) WithContext(ctx context.Context) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) WithHTTPClient(client *http.Client) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) WithGatewayID(gatewayID string) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) WithNetworkID(networkID string) *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get wifi network ID gateways gateway ID description params
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
