// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// DetachBootVolumeReader is a Reader for the DetachBootVolume structure.
type DetachBootVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachBootVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDetachBootVolumeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDetachBootVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDetachBootVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDetachBootVolumeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDetachBootVolumePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDetachBootVolumeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDetachBootVolumeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDetachBootVolumeNoContent creates a DetachBootVolumeNoContent with default headers values
func NewDetachBootVolumeNoContent() *DetachBootVolumeNoContent {
	return &DetachBootVolumeNoContent{}
}

/*DetachBootVolumeNoContent handles this case with default header values.

The request was accepted by the server
*/
type DetachBootVolumeNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DetachBootVolumeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /bootVolumeAttachments/{bootVolumeAttachmentId}][%d] detachBootVolumeNoContent ", 204)
}

func (o *DetachBootVolumeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDetachBootVolumeUnauthorized creates a DetachBootVolumeUnauthorized with default headers values
func NewDetachBootVolumeUnauthorized() *DetachBootVolumeUnauthorized {
	return &DetachBootVolumeUnauthorized{}
}

/*DetachBootVolumeUnauthorized handles this case with default header values.

Unauthorized
*/
type DetachBootVolumeUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DetachBootVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /bootVolumeAttachments/{bootVolumeAttachmentId}][%d] detachBootVolumeUnauthorized  %+v", 401, o.Payload)
}

func (o *DetachBootVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachBootVolumeNotFound creates a DetachBootVolumeNotFound with default headers values
func NewDetachBootVolumeNotFound() *DetachBootVolumeNotFound {
	return &DetachBootVolumeNotFound{}
}

/*DetachBootVolumeNotFound handles this case with default header values.

Not Found
*/
type DetachBootVolumeNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DetachBootVolumeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /bootVolumeAttachments/{bootVolumeAttachmentId}][%d] detachBootVolumeNotFound  %+v", 404, o.Payload)
}

func (o *DetachBootVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachBootVolumeConflict creates a DetachBootVolumeConflict with default headers values
func NewDetachBootVolumeConflict() *DetachBootVolumeConflict {
	return &DetachBootVolumeConflict{}
}

/*DetachBootVolumeConflict handles this case with default header values.

Conflict
*/
type DetachBootVolumeConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DetachBootVolumeConflict) Error() string {
	return fmt.Sprintf("[DELETE /bootVolumeAttachments/{bootVolumeAttachmentId}][%d] detachBootVolumeConflict  %+v", 409, o.Payload)
}

func (o *DetachBootVolumeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachBootVolumePreconditionFailed creates a DetachBootVolumePreconditionFailed with default headers values
func NewDetachBootVolumePreconditionFailed() *DetachBootVolumePreconditionFailed {
	return &DetachBootVolumePreconditionFailed{}
}

/*DetachBootVolumePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DetachBootVolumePreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DetachBootVolumePreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /bootVolumeAttachments/{bootVolumeAttachmentId}][%d] detachBootVolumePreconditionFailed  %+v", 412, o.Payload)
}

func (o *DetachBootVolumePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachBootVolumeInternalServerError creates a DetachBootVolumeInternalServerError with default headers values
func NewDetachBootVolumeInternalServerError() *DetachBootVolumeInternalServerError {
	return &DetachBootVolumeInternalServerError{}
}

/*DetachBootVolumeInternalServerError handles this case with default header values.

Internal Server Error
*/
type DetachBootVolumeInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DetachBootVolumeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /bootVolumeAttachments/{bootVolumeAttachmentId}][%d] detachBootVolumeInternalServerError  %+v", 500, o.Payload)
}

func (o *DetachBootVolumeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachBootVolumeDefault creates a DetachBootVolumeDefault with default headers values
func NewDetachBootVolumeDefault(code int) *DetachBootVolumeDefault {
	return &DetachBootVolumeDefault{
		_statusCode: code,
	}
}

/*DetachBootVolumeDefault handles this case with default header values.

An error has occurred.
*/
type DetachBootVolumeDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the detach boot volume default response
func (o *DetachBootVolumeDefault) Code() int {
	return o._statusCode
}

func (o *DetachBootVolumeDefault) Error() string {
	return fmt.Sprintf("[DELETE /bootVolumeAttachments/{bootVolumeAttachmentId}][%d] DetachBootVolume default  %+v", o._statusCode, o.Payload)
}

func (o *DetachBootVolumeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}