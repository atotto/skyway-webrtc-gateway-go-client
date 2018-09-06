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

// PeerConnectOptions peer connect options
// swagger:model PeerConnectOptions
type PeerConnectOptions struct {

	// options
	Options *PeerConnectOptionsOptions `json:"options,omitempty"`

	// params
	// Required: true
	Params *PeerConnectOptionsParams `json:"params"`

	// 接続元のPeerのidを指定します
	// Required: true
	PeerID *string `json:"peer_id"`

	// redirect params
	RedirectParams *DataConnectionRedirectOptions `json:"redirect_params,omitempty"`

	// 接続対象のpeer_idを指定します
	// Required: true
	TargetID *string `json:"target_id"`

	// Peerオブジェクトを利用するために必要なトークンです。他のユーザのリソースに対する誤操作防止のために指定します
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this peer connect options
func (m *PeerConnectOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerConnectOptions) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *PeerConnectOptions) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("params", "body", m.Params); err != nil {
		return err
	}

	if m.Params != nil {
		if err := m.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

func (m *PeerConnectOptions) validatePeerID(formats strfmt.Registry) error {

	if err := validate.Required("peer_id", "body", m.PeerID); err != nil {
		return err
	}

	return nil
}

func (m *PeerConnectOptions) validateRedirectParams(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirectParams) { // not required
		return nil
	}

	if m.RedirectParams != nil {
		if err := m.RedirectParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redirect_params")
			}
			return err
		}
	}

	return nil
}

func (m *PeerConnectOptions) validateTargetID(formats strfmt.Registry) error {

	if err := validate.Required("target_id", "body", m.TargetID); err != nil {
		return err
	}

	return nil
}

func (m *PeerConnectOptions) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerConnectOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerConnectOptions) UnmarshalBinary(b []byte) error {
	var res PeerConnectOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PeerConnectOptionsOptions peer connect options options
// swagger:model PeerConnectOptionsOptions
type PeerConnectOptionsOptions struct {

	// dc init
	DcInit *DcInit `json:"dcInit,omitempty"`

	// 接続相手に渡される任意の文字列です
	Metadata string `json:"metadata,omitempty"`

	// 送信時のシリアライズ方法を指定します。'BINARY', 'JSON', 'NONE'のいずれかが選択可能です
	Serialization string `json:"serialization,omitempty"`
}

// Validate validates this peer connect options options
func (m *PeerConnectOptionsOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDcInit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerConnectOptionsOptions) validateDcInit(formats strfmt.Registry) error {

	if swag.IsZero(m.DcInit) { // not required
		return nil
	}

	if m.DcInit != nil {
		if err := m.DcInit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options" + "." + "dcInit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerConnectOptionsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerConnectOptionsOptions) UnmarshalBinary(b []byte) error {
	var res PeerConnectOptionsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PeerConnectOptionsParams peer connect options params
// swagger:model PeerConnectOptionsParams
type PeerConnectOptionsParams struct {

	// data id
	DataID string `json:"data_id,omitempty"`
}

// Validate validates this peer connect options params
func (m *PeerConnectOptionsParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PeerConnectOptionsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerConnectOptionsParams) UnmarshalBinary(b []byte) error {
	var res PeerConnectOptionsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}