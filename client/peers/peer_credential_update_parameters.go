// Code generated by go-swagger; DO NOT EDIT.

package peers

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

// NewPeerCredentialUpdateParams creates a new PeerCredentialUpdateParams object
// with the default values initialized.
func NewPeerCredentialUpdateParams() *PeerCredentialUpdateParams {
	var ()
	return &PeerCredentialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeerCredentialUpdateParamsWithTimeout creates a new PeerCredentialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeerCredentialUpdateParamsWithTimeout(timeout time.Duration) *PeerCredentialUpdateParams {
	var ()
	return &PeerCredentialUpdateParams{

		timeout: timeout,
	}
}

// NewPeerCredentialUpdateParamsWithContext creates a new PeerCredentialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeerCredentialUpdateParamsWithContext(ctx context.Context) *PeerCredentialUpdateParams {
	var ()
	return &PeerCredentialUpdateParams{

		Context: ctx,
	}
}

// NewPeerCredentialUpdateParamsWithHTTPClient creates a new PeerCredentialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeerCredentialUpdateParamsWithHTTPClient(client *http.Client) *PeerCredentialUpdateParams {
	var ()
	return &PeerCredentialUpdateParams{
		HTTPClient: client,
	}
}

/*PeerCredentialUpdateParams contains all the parameters to send to the API endpoint
for the peer credential update operation typically these are written to a http.Request
*/
type PeerCredentialUpdateParams struct {

	/*Credential
	  Peerを認証するためのクレデンシャルです

	*/
	Credential *models.PeerCredential
	/*PeerID
	  他のピアがこのピアへ接続するときに使われるIDです。Peerオブジェクトの特定にも利用します

	*/
	PeerID string
	/*Token
	  Peerオブジェクトを利用するために必要なトークンです。他のユーザのリソースに対する誤操作防止のために指定します

	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the peer credential update params
func (o *PeerCredentialUpdateParams) WithTimeout(timeout time.Duration) *PeerCredentialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peer credential update params
func (o *PeerCredentialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peer credential update params
func (o *PeerCredentialUpdateParams) WithContext(ctx context.Context) *PeerCredentialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peer credential update params
func (o *PeerCredentialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peer credential update params
func (o *PeerCredentialUpdateParams) WithHTTPClient(client *http.Client) *PeerCredentialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peer credential update params
func (o *PeerCredentialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the peer credential update params
func (o *PeerCredentialUpdateParams) WithCredential(credential *models.PeerCredential) *PeerCredentialUpdateParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the peer credential update params
func (o *PeerCredentialUpdateParams) SetCredential(credential *models.PeerCredential) {
	o.Credential = credential
}

// WithPeerID adds the peerID to the peer credential update params
func (o *PeerCredentialUpdateParams) WithPeerID(peerID string) *PeerCredentialUpdateParams {
	o.SetPeerID(peerID)
	return o
}

// SetPeerID adds the peerId to the peer credential update params
func (o *PeerCredentialUpdateParams) SetPeerID(peerID string) {
	o.PeerID = peerID
}

// WithToken adds the token to the peer credential update params
func (o *PeerCredentialUpdateParams) WithToken(token string) *PeerCredentialUpdateParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the peer credential update params
func (o *PeerCredentialUpdateParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PeerCredentialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {
		if err := r.SetBodyParam(o.Credential); err != nil {
			return err
		}
	}

	// path param peer_id
	if err := r.SetPathParam("peer_id", o.PeerID); err != nil {
		return err
	}

	// query param token
	qrToken := o.Token
	qToken := qrToken
	if qToken != "" {
		if err := r.SetQueryParam("token", qToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
