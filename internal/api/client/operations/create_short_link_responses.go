// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/Maksim646/ozon_shortner/internal/api/definition"
)

// CreateShortLinkReader is a Reader for the CreateShortLink structure.
type CreateShortLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateShortLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateShortLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateShortLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateShortLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /shortner_link] CreateShortLink", response, response.Code())
	}
}

// NewCreateShortLinkOK creates a CreateShortLinkOK with default headers values
func NewCreateShortLinkOK() *CreateShortLinkOK {
	return &CreateShortLinkOK{}
}

/*
CreateShortLinkOK describes a response with status code 200, with default header values.

Create Short Link Response
*/
type CreateShortLinkOK struct {
	Payload *models.ShortLink
}

// IsSuccess returns true when this create short link o k response has a 2xx status code
func (o *CreateShortLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create short link o k response has a 3xx status code
func (o *CreateShortLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create short link o k response has a 4xx status code
func (o *CreateShortLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create short link o k response has a 5xx status code
func (o *CreateShortLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create short link o k response a status code equal to that given
func (o *CreateShortLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create short link o k response
func (o *CreateShortLinkOK) Code() int {
	return 200
}

func (o *CreateShortLinkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /shortner_link][%d] createShortLinkOK %s", 200, payload)
}

func (o *CreateShortLinkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /shortner_link][%d] createShortLinkOK %s", 200, payload)
}

func (o *CreateShortLinkOK) GetPayload() *models.ShortLink {
	return o.Payload
}

func (o *CreateShortLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShortLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateShortLinkBadRequest creates a CreateShortLinkBadRequest with default headers values
func NewCreateShortLinkBadRequest() *CreateShortLinkBadRequest {
	return &CreateShortLinkBadRequest{}
}

/*
CreateShortLinkBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateShortLinkBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create short link bad request response has a 2xx status code
func (o *CreateShortLinkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create short link bad request response has a 3xx status code
func (o *CreateShortLinkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create short link bad request response has a 4xx status code
func (o *CreateShortLinkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create short link bad request response has a 5xx status code
func (o *CreateShortLinkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create short link bad request response a status code equal to that given
func (o *CreateShortLinkBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create short link bad request response
func (o *CreateShortLinkBadRequest) Code() int {
	return 400
}

func (o *CreateShortLinkBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /shortner_link][%d] createShortLinkBadRequest %s", 400, payload)
}

func (o *CreateShortLinkBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /shortner_link][%d] createShortLinkBadRequest %s", 400, payload)
}

func (o *CreateShortLinkBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateShortLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateShortLinkInternalServerError creates a CreateShortLinkInternalServerError with default headers values
func NewCreateShortLinkInternalServerError() *CreateShortLinkInternalServerError {
	return &CreateShortLinkInternalServerError{}
}

/*
CreateShortLinkInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateShortLinkInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this create short link internal server error response has a 2xx status code
func (o *CreateShortLinkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create short link internal server error response has a 3xx status code
func (o *CreateShortLinkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create short link internal server error response has a 4xx status code
func (o *CreateShortLinkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create short link internal server error response has a 5xx status code
func (o *CreateShortLinkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create short link internal server error response a status code equal to that given
func (o *CreateShortLinkInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create short link internal server error response
func (o *CreateShortLinkInternalServerError) Code() int {
	return 500
}

func (o *CreateShortLinkInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /shortner_link][%d] createShortLinkInternalServerError %s", 500, payload)
}

func (o *CreateShortLinkInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /shortner_link][%d] createShortLinkInternalServerError %s", 500, payload)
}

func (o *CreateShortLinkInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateShortLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
