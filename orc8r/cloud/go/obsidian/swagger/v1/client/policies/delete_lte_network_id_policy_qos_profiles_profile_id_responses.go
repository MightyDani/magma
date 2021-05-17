// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// DeleteLTENetworkIDPolicyQosProfilesProfileIDReader is a Reader for the DeleteLTENetworkIDPolicyQosProfilesProfileID structure.
type DeleteLTENetworkIDPolicyQosProfilesProfileIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDPolicyQosProfilesProfileIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDPolicyQosProfilesProfileIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent creates a DeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent with default headers values
func NewDeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent() *DeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent {
	return &DeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent{}
}

/*DeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent handles this case with default header values.

Success
*/
type DeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent struct {
}

func (o *DeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/policy_qos_profiles/{profile_id}][%d] deleteLteNetworkIdPolicyQosProfilesProfileIdNoContent ", 204)
}

func (o *DeleteLTENetworkIDPolicyQosProfilesProfileIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDPolicyQosProfilesProfileIDDefault creates a DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault with default headers values
func NewDeleteLTENetworkIDPolicyQosProfilesProfileIDDefault(code int) *DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault {
	return &DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault{
		_statusCode: code,
	}
}

/*DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID policy qos profiles profile ID default response
func (o *DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/policy_qos_profiles/{profile_id}][%d] DeleteLTENetworkIDPolicyQosProfilesProfileID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDPolicyQosProfilesProfileIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
