package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// DeleteCrossConnectReader is a Reader for the DeleteCrossConnect structure.
type DeleteCrossConnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCrossConnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCrossConnectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteCrossConnectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCrossConnectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteCrossConnectPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteCrossConnectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteCrossConnectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCrossConnectNoContent creates a DeleteCrossConnectNoContent with default headers values
func NewDeleteCrossConnectNoContent() *DeleteCrossConnectNoContent {
	return &DeleteCrossConnectNoContent{}
}

/*DeleteCrossConnectNoContent handles this case with default header values.

The cross-connect is being deleted.
*/
type DeleteCrossConnectNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteCrossConnectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /crossConnects/{crossConnectId}][%d] deleteCrossConnectNoContent ", 204)
}

func (o *DeleteCrossConnectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteCrossConnectUnauthorized creates a DeleteCrossConnectUnauthorized with default headers values
func NewDeleteCrossConnectUnauthorized() *DeleteCrossConnectUnauthorized {
	return &DeleteCrossConnectUnauthorized{}
}

/*DeleteCrossConnectUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCrossConnectUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCrossConnectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /crossConnects/{crossConnectId}][%d] deleteCrossConnectUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCrossConnectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCrossConnectNotFound creates a DeleteCrossConnectNotFound with default headers values
func NewDeleteCrossConnectNotFound() *DeleteCrossConnectNotFound {
	return &DeleteCrossConnectNotFound{}
}

/*DeleteCrossConnectNotFound handles this case with default header values.

Not Found
*/
type DeleteCrossConnectNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCrossConnectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /crossConnects/{crossConnectId}][%d] deleteCrossConnectNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCrossConnectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCrossConnectPreconditionFailed creates a DeleteCrossConnectPreconditionFailed with default headers values
func NewDeleteCrossConnectPreconditionFailed() *DeleteCrossConnectPreconditionFailed {
	return &DeleteCrossConnectPreconditionFailed{}
}

/*DeleteCrossConnectPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DeleteCrossConnectPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCrossConnectPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /crossConnects/{crossConnectId}][%d] deleteCrossConnectPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteCrossConnectPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCrossConnectInternalServerError creates a DeleteCrossConnectInternalServerError with default headers values
func NewDeleteCrossConnectInternalServerError() *DeleteCrossConnectInternalServerError {
	return &DeleteCrossConnectInternalServerError{}
}

/*DeleteCrossConnectInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteCrossConnectInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteCrossConnectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /crossConnects/{crossConnectId}][%d] deleteCrossConnectInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCrossConnectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCrossConnectDefault creates a DeleteCrossConnectDefault with default headers values
func NewDeleteCrossConnectDefault(code int) *DeleteCrossConnectDefault {
	return &DeleteCrossConnectDefault{
		_statusCode: code,
	}
}

/*DeleteCrossConnectDefault handles this case with default header values.

An error has occurred.
*/
type DeleteCrossConnectDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete cross connect default response
func (o *DeleteCrossConnectDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCrossConnectDefault) Error() string {
	return fmt.Sprintf("[DELETE /crossConnects/{crossConnectId}][%d] DeleteCrossConnect default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCrossConnectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}