// Code generated by go-swagger; DO NOT EDIT.

package data

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

	models "github.com/atotto/skyway-webrtc-gateway-go-client/models"
)

// NewDataConnectionPutParams creates a new DataConnectionPutParams object
// with the default values initialized.
func NewDataConnectionPutParams() *DataConnectionPutParams {
	var ()
	return &DataConnectionPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataConnectionPutParamsWithTimeout creates a new DataConnectionPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataConnectionPutParamsWithTimeout(timeout time.Duration) *DataConnectionPutParams {
	var ()
	return &DataConnectionPutParams{

		timeout: timeout,
	}
}

// NewDataConnectionPutParamsWithContext creates a new DataConnectionPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataConnectionPutParamsWithContext(ctx context.Context) *DataConnectionPutParams {
	var ()
	return &DataConnectionPutParams{

		Context: ctx,
	}
}

// NewDataConnectionPutParamsWithHTTPClient creates a new DataConnectionPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataConnectionPutParamsWithHTTPClient(client *http.Client) *DataConnectionPutParams {
	var ()
	return &DataConnectionPutParams{
		HTTPClient: client,
	}
}

/*DataConnectionPutParams contains all the parameters to send to the API endpoint
for the data connection put operation typically these are written to a http.Request
*/
type DataConnectionPutParams struct {

	/*Body
	  受信したDataの転送先と、相手側へ送信するDataのIDを指定します

	*/
	Body *models.DataConnectionPutOptions
	/*DataConnectionID
	  DataConnectionを特定するためのidを指定します

	*/
	DataConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data connection put params
func (o *DataConnectionPutParams) WithTimeout(timeout time.Duration) *DataConnectionPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data connection put params
func (o *DataConnectionPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data connection put params
func (o *DataConnectionPutParams) WithContext(ctx context.Context) *DataConnectionPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data connection put params
func (o *DataConnectionPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data connection put params
func (o *DataConnectionPutParams) WithHTTPClient(client *http.Client) *DataConnectionPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data connection put params
func (o *DataConnectionPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the data connection put params
func (o *DataConnectionPutParams) WithBody(body *models.DataConnectionPutOptions) *DataConnectionPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data connection put params
func (o *DataConnectionPutParams) SetBody(body *models.DataConnectionPutOptions) {
	o.Body = body
}

// WithDataConnectionID adds the dataConnectionID to the data connection put params
func (o *DataConnectionPutParams) WithDataConnectionID(dataConnectionID string) *DataConnectionPutParams {
	o.SetDataConnectionID(dataConnectionID)
	return o
}

// SetDataConnectionID adds the dataConnectionId to the data connection put params
func (o *DataConnectionPutParams) SetDataConnectionID(dataConnectionID string) {
	o.DataConnectionID = dataConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *DataConnectionPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param data_connection_id
	if err := r.SetPathParam("data_connection_id", o.DataConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}