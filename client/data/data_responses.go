// Code generated by go-swagger; DO NOT EDIT.

package data

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

	models "github.com/atotto/skyway-webrtc-gateway-go-client/models"
)

// DataReader is a Reader for the Data structure.
type DataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDataCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDataMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewDataNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 408:
		result := NewDataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDataCreated creates a DataCreated with default headers values
func NewDataCreated() *DataCreated {
	return &DataCreated{}
}

/*DataCreated handles this case with default header values.

successful operation
*/
type DataCreated struct {
	Payload *models.DataSockParameters
}

func (o *DataCreated) Error() string {
	return fmt.Sprintf("[POST /data][%d] dataCreated  %+v", 201, o.Payload)
}

func (o *DataCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSockParameters)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataBadRequest creates a DataBadRequest with default headers values
func NewDataBadRequest() *DataBadRequest {
	return &DataBadRequest{}
}

/*DataBadRequest handles this case with default header values.

Invalid input
*/
type DataBadRequest struct {
	Payload *DataBadRequestBody
}

func (o *DataBadRequest) Error() string {
	return fmt.Sprintf("[POST /data][%d] dataBadRequest  %+v", 400, o.Payload)
}

func (o *DataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DataBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataMethodNotAllowed creates a DataMethodNotAllowed with default headers values
func NewDataMethodNotAllowed() *DataMethodNotAllowed {
	return &DataMethodNotAllowed{}
}

/*DataMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type DataMethodNotAllowed struct {
}

func (o *DataMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /data][%d] dataMethodNotAllowed ", 405)
}

func (o *DataMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataNotAcceptable creates a DataNotAcceptable with default headers values
func NewDataNotAcceptable() *DataNotAcceptable {
	return &DataNotAcceptable{}
}

/*DataNotAcceptable handles this case with default header values.

Not Acceptable
*/
type DataNotAcceptable struct {
}

func (o *DataNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /data][%d] dataNotAcceptable ", 406)
}

func (o *DataNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataRequestTimeout creates a DataRequestTimeout with default headers values
func NewDataRequestTimeout() *DataRequestTimeout {
	return &DataRequestTimeout{}
}

/*DataRequestTimeout handles this case with default header values.

Request Timeout
*/
type DataRequestTimeout struct {
}

func (o *DataRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /data][%d] dataRequestTimeout ", 408)
}

func (o *DataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DataBadRequestBody data bad request body
swagger:model DataBadRequestBody
*/
type DataBadRequestBody struct {

	// command type
	// Required: true
	CommandType *string `json:"command_type"`

	// params
	// Required: true
	Params *DataBadRequestBodyParams `json:"params"`
}

// Validate validates this data bad request body
func (o *DataBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *DataBadRequestBody) validateCommandType(formats strfmt.Registry) error {

	if err := validate.Required("dataBadRequest"+"."+"command_type", "body", o.CommandType); err != nil {
		return err
	}

	return nil
}

func (o *DataBadRequestBody) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("dataBadRequest"+"."+"params", "body", o.Params); err != nil {
		return err
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataBadRequest" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DataBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DataBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DataBadRequestBodyParams data bad request body params
swagger:model DataBadRequestBodyParams
*/
type DataBadRequestBodyParams struct {

	// errors
	// Required: true
	Errors []*DataBadRequestBodyParamsErrorsItems0 `json:"errors"`
}

// Validate validates this data bad request body params
func (o *DataBadRequestBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DataBadRequestBodyParams) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("dataBadRequest"+"."+"params"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataBadRequest" + "." + "params" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DataBadRequestBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataBadRequestBodyParams) UnmarshalBinary(b []byte) error {
	var res DataBadRequestBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DataBadRequestBodyParamsErrorsItems0 data bad request body params errors items0
swagger:model DataBadRequestBodyParamsErrorsItems0
*/
type DataBadRequestBodyParamsErrorsItems0 struct {

	// field
	Field string `json:"field,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this data bad request body params errors items0
func (o *DataBadRequestBodyParamsErrorsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DataBadRequestBodyParamsErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataBadRequestBodyParamsErrorsItems0) UnmarshalBinary(b []byte) error {
	var res DataBadRequestBodyParamsErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
