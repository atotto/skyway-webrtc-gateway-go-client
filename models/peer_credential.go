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

// PeerCredential peer credential
// swagger:model PeerCredential
type PeerCredential struct {

	// HMACを利用して生成する認証用トークンです
	// Required: true
	// Format: byte
	AuthToken *strfmt.Base64 `json:"authToken"`

	// 現在のUNIXタイムスタンプです
	// Required: true
	Timestamp *int64 `json:"timestamp"`

	// クレデンシャルの生存期間です。timestampからttl秒経過したとき、クレデンシャルが失効します。
	// Required: true
	// Maximum: 90000
	// Minimum: 600
	TTL *int64 `json:"ttl"`
}

// Validate validates this peer credential
func (m *PeerCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerCredential) validateAuthToken(formats strfmt.Registry) error {

	if err := validate.Required("authToken", "body", m.AuthToken); err != nil {
		return err
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *PeerCredential) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *PeerCredential) validateTTL(formats strfmt.Registry) error {

	if err := validate.Required("ttl", "body", m.TTL); err != nil {
		return err
	}

	if err := validate.MinimumInt("ttl", "body", int64(*m.TTL), 600, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("ttl", "body", int64(*m.TTL), 90000, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerCredential) UnmarshalBinary(b []byte) error {
	var res PeerCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}