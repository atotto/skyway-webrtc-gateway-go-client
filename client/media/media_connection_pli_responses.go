// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// MediaConnectionPliReader is a Reader for the MediaConnectionPli structure.
type MediaConnectionPliReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaConnectionPliReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewMediaConnectionPliCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewMediaConnectionPliBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewMediaConnectionPliForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewMediaConnectionPliNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewMediaConnectionPliMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewMediaConnectionPliNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 408:
		result := NewMediaConnectionPliRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMediaConnectionPliCreated creates a MediaConnectionPliCreated with default headers values
func NewMediaConnectionPliCreated() *MediaConnectionPliCreated {
	return &MediaConnectionPliCreated{}
}

/*MediaConnectionPliCreated handles this case with default header values.

Created
*/
type MediaConnectionPliCreated struct {
	Payload interface{}
}

func (o *MediaConnectionPliCreated) Error() string {
	return fmt.Sprintf("[POST /media/connections/{media_connection_id}/pli][%d] mediaConnectionPliCreated  %+v", 201, o.Payload)
}

func (o *MediaConnectionPliCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaConnectionPliBadRequest creates a MediaConnectionPliBadRequest with default headers values
func NewMediaConnectionPliBadRequest() *MediaConnectionPliBadRequest {
	return &MediaConnectionPliBadRequest{}
}

/*MediaConnectionPliBadRequest handles this case with default header values.

Invalid input
*/
type MediaConnectionPliBadRequest struct {
	Payload *MediaConnectionPliBadRequestBody
}

func (o *MediaConnectionPliBadRequest) Error() string {
	return fmt.Sprintf("[POST /media/connections/{media_connection_id}/pli][%d] mediaConnectionPliBadRequest  %+v", 400, o.Payload)
}

func (o *MediaConnectionPliBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MediaConnectionPliBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaConnectionPliForbidden creates a MediaConnectionPliForbidden with default headers values
func NewMediaConnectionPliForbidden() *MediaConnectionPliForbidden {
	return &MediaConnectionPliForbidden{}
}

/*MediaConnectionPliForbidden handles this case with default header values.

Forbidden(不正な操作を行った場合)
*/
type MediaConnectionPliForbidden struct {
}

func (o *MediaConnectionPliForbidden) Error() string {
	return fmt.Sprintf("[POST /media/connections/{media_connection_id}/pli][%d] mediaConnectionPliForbidden ", 403)
}

func (o *MediaConnectionPliForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaConnectionPliNotFound creates a MediaConnectionPliNotFound with default headers values
func NewMediaConnectionPliNotFound() *MediaConnectionPliNotFound {
	return &MediaConnectionPliNotFound{}
}

/*MediaConnectionPliNotFound handles this case with default header values.

Not Found(media_connection_idが間違っている場合)
*/
type MediaConnectionPliNotFound struct {
}

func (o *MediaConnectionPliNotFound) Error() string {
	return fmt.Sprintf("[POST /media/connections/{media_connection_id}/pli][%d] mediaConnectionPliNotFound ", 404)
}

func (o *MediaConnectionPliNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaConnectionPliMethodNotAllowed creates a MediaConnectionPliMethodNotAllowed with default headers values
func NewMediaConnectionPliMethodNotAllowed() *MediaConnectionPliMethodNotAllowed {
	return &MediaConnectionPliMethodNotAllowed{}
}

/*MediaConnectionPliMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type MediaConnectionPliMethodNotAllowed struct {
}

func (o *MediaConnectionPliMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /media/connections/{media_connection_id}/pli][%d] mediaConnectionPliMethodNotAllowed ", 405)
}

func (o *MediaConnectionPliMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaConnectionPliNotAcceptable creates a MediaConnectionPliNotAcceptable with default headers values
func NewMediaConnectionPliNotAcceptable() *MediaConnectionPliNotAcceptable {
	return &MediaConnectionPliNotAcceptable{}
}

/*MediaConnectionPliNotAcceptable handles this case with default header values.

Not Acceptable
*/
type MediaConnectionPliNotAcceptable struct {
}

func (o *MediaConnectionPliNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /media/connections/{media_connection_id}/pli][%d] mediaConnectionPliNotAcceptable ", 406)
}

func (o *MediaConnectionPliNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaConnectionPliRequestTimeout creates a MediaConnectionPliRequestTimeout with default headers values
func NewMediaConnectionPliRequestTimeout() *MediaConnectionPliRequestTimeout {
	return &MediaConnectionPliRequestTimeout{}
}

/*MediaConnectionPliRequestTimeout handles this case with default header values.

Request Timeout
*/
type MediaConnectionPliRequestTimeout struct {
}

func (o *MediaConnectionPliRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /media/connections/{media_connection_id}/pli][%d] mediaConnectionPliRequestTimeout ", 408)
}

func (o *MediaConnectionPliRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*MediaConnectionPliBadRequestBody media connection pli bad request body
swagger:model MediaConnectionPliBadRequestBody
*/
type MediaConnectionPliBadRequestBody struct {

	// command type
	// Required: true
	CommandType *string `json:"command_type"`

	// params
	// Required: true
	Params *MediaConnectionPliBadRequestBodyParams `json:"params"`
}

// Validate validates this media connection pli bad request body
func (o *MediaConnectionPliBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommandType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MediaConnectionPliBadRequestBody) validateCommandType(formats strfmt.Registry) error {

	if err := validate.Required("mediaConnectionPliBadRequest"+"."+"command_type", "body", o.CommandType); err != nil {
		return err
	}

	return nil
}

func (o *MediaConnectionPliBadRequestBody) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("mediaConnectionPliBadRequest"+"."+"params", "body", o.Params); err != nil {
		return err
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaConnectionPliBadRequest" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MediaConnectionPliBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MediaConnectionPliBadRequestBody) UnmarshalBinary(b []byte) error {
	var res MediaConnectionPliBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MediaConnectionPliBadRequestBodyParams media connection pli bad request body params
swagger:model MediaConnectionPliBadRequestBodyParams
*/
type MediaConnectionPliBadRequestBodyParams struct {

	// errors
	// Required: true
	Errors []*MediaConnectionPliBadRequestBodyParamsErrorsItems0 `json:"errors"`
}

// Validate validates this media connection pli bad request body params
func (o *MediaConnectionPliBadRequestBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MediaConnectionPliBadRequestBodyParams) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("mediaConnectionPliBadRequest"+"."+"params"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mediaConnectionPliBadRequest" + "." + "params" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *MediaConnectionPliBadRequestBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MediaConnectionPliBadRequestBodyParams) UnmarshalBinary(b []byte) error {
	var res MediaConnectionPliBadRequestBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MediaConnectionPliBadRequestBodyParamsErrorsItems0 media connection pli bad request body params errors items0
swagger:model MediaConnectionPliBadRequestBodyParamsErrorsItems0
*/
type MediaConnectionPliBadRequestBodyParamsErrorsItems0 struct {

	// field
	Field string `json:"field,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this media connection pli bad request body params errors items0
func (o *MediaConnectionPliBadRequestBodyParamsErrorsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MediaConnectionPliBadRequestBodyParamsErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MediaConnectionPliBadRequestBodyParamsErrorsItems0) UnmarshalBinary(b []byte) error {
	var res MediaConnectionPliBadRequestBodyParamsErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
