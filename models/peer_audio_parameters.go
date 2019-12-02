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

// PeerAudioParameters peer audio parameters
// swagger:model PeerAudioParameters
type PeerAudioParameters struct {

	// band width
	// Required: true
	BandWidth *int64 `json:"band_width"`

	// codec
	// Required: true
	Codec *string `json:"codec"`

	// Mediaを特定するためのIDです
	// Required: true
	MediaID *string `json:"media_id"`

	// SDP内部でdynamic assignされるメディアのペイロード番号を指定
	PayloadType int64 `json:"payload_type,omitempty"`

	// RTCP転送ポートを特定するためのIDです
	RtcpID string `json:"rtcp_id,omitempty"`

	// sampling rate
	SamplingRate int64 `json:"sampling_rate,omitempty"`
}

// Validate validates this peer audio parameters
func (m *PeerAudioParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCodec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerAudioParameters) validateBandWidth(formats strfmt.Registry) error {

	if err := validate.Required("band_width", "body", m.BandWidth); err != nil {
		return err
	}

	return nil
}

func (m *PeerAudioParameters) validateCodec(formats strfmt.Registry) error {

	if err := validate.Required("codec", "body", m.Codec); err != nil {
		return err
	}

	return nil
}

func (m *PeerAudioParameters) validateMediaID(formats strfmt.Registry) error {

	if err := validate.Required("media_id", "body", m.MediaID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerAudioParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerAudioParameters) UnmarshalBinary(b []byte) error {
	var res PeerAudioParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
