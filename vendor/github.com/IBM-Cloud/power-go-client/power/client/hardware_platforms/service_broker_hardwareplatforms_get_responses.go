// Code generated by go-swagger; DO NOT EDIT.

package hardware_platforms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerHardwareplatformsGetReader is a Reader for the ServiceBrokerHardwareplatformsGet structure.
type ServiceBrokerHardwareplatformsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerHardwareplatformsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerHardwareplatformsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerHardwareplatformsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerHardwareplatformsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerHardwareplatformsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerHardwareplatformsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerHardwareplatformsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/hardware-platforms] serviceBroker.hardwareplatforms.get", response, response.Code())
	}
}

// NewServiceBrokerHardwareplatformsGetOK creates a ServiceBrokerHardwareplatformsGetOK with default headers values
func NewServiceBrokerHardwareplatformsGetOK() *ServiceBrokerHardwareplatformsGetOK {
	return &ServiceBrokerHardwareplatformsGetOK{}
}

/*
ServiceBrokerHardwareplatformsGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerHardwareplatformsGetOK struct {
	Payload models.HardwarePlatforms
}

// IsSuccess returns true when this service broker hardwareplatforms get o k response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker hardwareplatforms get o k response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get o k response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker hardwareplatforms get o k response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker hardwareplatforms get o k response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker hardwareplatforms get o k response
func (o *ServiceBrokerHardwareplatformsGetOK) Code() int {
	return 200
}

func (o *ServiceBrokerHardwareplatformsGetOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetOK) GetPayload() models.HardwarePlatforms {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHardwareplatformsGetBadRequest creates a ServiceBrokerHardwareplatformsGetBadRequest with default headers values
func NewServiceBrokerHardwareplatformsGetBadRequest() *ServiceBrokerHardwareplatformsGetBadRequest {
	return &ServiceBrokerHardwareplatformsGetBadRequest{}
}

/*
ServiceBrokerHardwareplatformsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerHardwareplatformsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker hardwareplatforms get bad request response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker hardwareplatforms get bad request response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get bad request response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker hardwareplatforms get bad request response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker hardwareplatforms get bad request response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker hardwareplatforms get bad request response
func (o *ServiceBrokerHardwareplatformsGetBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerHardwareplatformsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHardwareplatformsGetUnauthorized creates a ServiceBrokerHardwareplatformsGetUnauthorized with default headers values
func NewServiceBrokerHardwareplatformsGetUnauthorized() *ServiceBrokerHardwareplatformsGetUnauthorized {
	return &ServiceBrokerHardwareplatformsGetUnauthorized{}
}

/*
ServiceBrokerHardwareplatformsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerHardwareplatformsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker hardwareplatforms get unauthorized response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker hardwareplatforms get unauthorized response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get unauthorized response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker hardwareplatforms get unauthorized response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker hardwareplatforms get unauthorized response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker hardwareplatforms get unauthorized response
func (o *ServiceBrokerHardwareplatformsGetUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerHardwareplatformsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHardwareplatformsGetForbidden creates a ServiceBrokerHardwareplatformsGetForbidden with default headers values
func NewServiceBrokerHardwareplatformsGetForbidden() *ServiceBrokerHardwareplatformsGetForbidden {
	return &ServiceBrokerHardwareplatformsGetForbidden{}
}

/*
ServiceBrokerHardwareplatformsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerHardwareplatformsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker hardwareplatforms get forbidden response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker hardwareplatforms get forbidden response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get forbidden response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker hardwareplatforms get forbidden response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker hardwareplatforms get forbidden response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker hardwareplatforms get forbidden response
func (o *ServiceBrokerHardwareplatformsGetForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerHardwareplatformsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetForbidden) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHardwareplatformsGetNotFound creates a ServiceBrokerHardwareplatformsGetNotFound with default headers values
func NewServiceBrokerHardwareplatformsGetNotFound() *ServiceBrokerHardwareplatformsGetNotFound {
	return &ServiceBrokerHardwareplatformsGetNotFound{}
}

/*
ServiceBrokerHardwareplatformsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerHardwareplatformsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker hardwareplatforms get not found response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker hardwareplatforms get not found response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get not found response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker hardwareplatforms get not found response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker hardwareplatforms get not found response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker hardwareplatforms get not found response
func (o *ServiceBrokerHardwareplatformsGetNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerHardwareplatformsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetNotFound) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHardwareplatformsGetInternalServerError creates a ServiceBrokerHardwareplatformsGetInternalServerError with default headers values
func NewServiceBrokerHardwareplatformsGetInternalServerError() *ServiceBrokerHardwareplatformsGetInternalServerError {
	return &ServiceBrokerHardwareplatformsGetInternalServerError{}
}

/*
ServiceBrokerHardwareplatformsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerHardwareplatformsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker hardwareplatforms get internal server error response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker hardwareplatforms get internal server error response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get internal server error response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker hardwareplatforms get internal server error response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker hardwareplatforms get internal server error response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker hardwareplatforms get internal server error response
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}