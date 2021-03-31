// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MediaConnectionStatusMessage media connection status message
// swagger:model MediaConnectionStatusMessage
type MediaConnectionStatusMessage struct {

	// コネクションが開始されたときに、コネクションと関連付けされるメタデータです
	// Required: true
	Metadata interface{} `json:"metadata"`

	// MediaConnectionがアクティブなときtrueになります
	// Required: true
	Open *bool `json:"open"`

	// 接続相手のPeerのpeer_idです
	// Required: true
	RemoteID *string `json:"remote_id"`
}

// Validate validates this media connection status message
func (m *MediaConnectionStatusMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaConnectionStatusMessage) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	return nil
}

func (m *MediaConnectionStatusMessage) validateOpen(formats strfmt.Registry) error {

	if err := validate.Required("open", "body", m.Open); err != nil {
		return err
	}

	return nil
}

func (m *MediaConnectionStatusMessage) validateRemoteID(formats strfmt.Registry) error {

	if err := validate.Required("remote_id", "body", m.RemoteID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediaConnectionStatusMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaConnectionStatusMessage) UnmarshalBinary(b []byte) error {
	var res MediaConnectionStatusMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
