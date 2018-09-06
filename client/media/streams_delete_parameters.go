// Code generated by go-swagger; DO NOT EDIT.

package media

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
)

// NewStreamsDeleteParams creates a new StreamsDeleteParams object
// with the default values initialized.
func NewStreamsDeleteParams() *StreamsDeleteParams {
	var ()
	return &StreamsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStreamsDeleteParamsWithTimeout creates a new StreamsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStreamsDeleteParamsWithTimeout(timeout time.Duration) *StreamsDeleteParams {
	var ()
	return &StreamsDeleteParams{

		timeout: timeout,
	}
}

// NewStreamsDeleteParamsWithContext creates a new StreamsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewStreamsDeleteParamsWithContext(ctx context.Context) *StreamsDeleteParams {
	var ()
	return &StreamsDeleteParams{

		Context: ctx,
	}
}

// NewStreamsDeleteParamsWithHTTPClient creates a new StreamsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStreamsDeleteParamsWithHTTPClient(client *http.Client) *StreamsDeleteParams {
	var ()
	return &StreamsDeleteParams{
		HTTPClient: client,
	}
}

/*StreamsDeleteParams contains all the parameters to send to the API endpoint
for the streams delete operation typically these are written to a http.Request
*/
type StreamsDeleteParams struct {

	/*MediaID
	  Mediaを特定するためのIDを指定します

	*/
	MediaID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the streams delete params
func (o *StreamsDeleteParams) WithTimeout(timeout time.Duration) *StreamsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the streams delete params
func (o *StreamsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the streams delete params
func (o *StreamsDeleteParams) WithContext(ctx context.Context) *StreamsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the streams delete params
func (o *StreamsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the streams delete params
func (o *StreamsDeleteParams) WithHTTPClient(client *http.Client) *StreamsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the streams delete params
func (o *StreamsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMediaID adds the mediaID to the streams delete params
func (o *StreamsDeleteParams) WithMediaID(mediaID string) *StreamsDeleteParams {
	o.SetMediaID(mediaID)
	return o
}

// SetMediaID adds the mediaId to the streams delete params
func (o *StreamsDeleteParams) SetMediaID(mediaID string) {
	o.MediaID = mediaID
}

// WriteToRequest writes these params to a swagger request
func (o *StreamsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param media_id
	if err := r.SetPathParam("media_id", o.MediaID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}