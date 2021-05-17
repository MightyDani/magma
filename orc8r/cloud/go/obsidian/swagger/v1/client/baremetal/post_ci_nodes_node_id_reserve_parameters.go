// Code generated by go-swagger; DO NOT EDIT.

package baremetal

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

// NewPostCiNodesNodeIDReserveParams creates a new PostCiNodesNodeIDReserveParams object
// with the default values initialized.
func NewPostCiNodesNodeIDReserveParams() *PostCiNodesNodeIDReserveParams {
	var ()
	return &PostCiNodesNodeIDReserveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCiNodesNodeIDReserveParamsWithTimeout creates a new PostCiNodesNodeIDReserveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCiNodesNodeIDReserveParamsWithTimeout(timeout time.Duration) *PostCiNodesNodeIDReserveParams {
	var ()
	return &PostCiNodesNodeIDReserveParams{

		timeout: timeout,
	}
}

// NewPostCiNodesNodeIDReserveParamsWithContext creates a new PostCiNodesNodeIDReserveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCiNodesNodeIDReserveParamsWithContext(ctx context.Context) *PostCiNodesNodeIDReserveParams {
	var ()
	return &PostCiNodesNodeIDReserveParams{

		Context: ctx,
	}
}

// NewPostCiNodesNodeIDReserveParamsWithHTTPClient creates a new PostCiNodesNodeIDReserveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCiNodesNodeIDReserveParamsWithHTTPClient(client *http.Client) *PostCiNodesNodeIDReserveParams {
	var ()
	return &PostCiNodesNodeIDReserveParams{
		HTTPClient: client,
	}
}

/*PostCiNodesNodeIDReserveParams contains all the parameters to send to the API endpoint
for the post ci nodes node ID reserve operation typically these are written to a http.Request
*/
type PostCiNodesNodeIDReserveParams struct {

	/*NodeID
	  CI node ID

	*/
	NodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) WithTimeout(timeout time.Duration) *PostCiNodesNodeIDReserveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) WithContext(ctx context.Context) *PostCiNodesNodeIDReserveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) WithHTTPClient(client *http.Client) *PostCiNodesNodeIDReserveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeID adds the nodeID to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) WithNodeID(nodeID string) *PostCiNodesNodeIDReserveParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the post ci nodes node ID reserve params
func (o *PostCiNodesNodeIDReserveParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCiNodesNodeIDReserveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param node_id
	if err := r.SetPathParam("node_id", o.NodeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
