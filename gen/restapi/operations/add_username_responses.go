// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddUsernameCreatedCode is the HTTP code returned for type AddUsernameCreated
const AddUsernameCreatedCode int = 201

/*AddUsernameCreated username added

swagger:response addUsernameCreated
*/
type AddUsernameCreated struct {

	/*
	  In: Body
	*/
	Payload *AddUsernameCreatedBody `json:"body,omitempty"`
}

// NewAddUsernameCreated creates AddUsernameCreated with default headers values
func NewAddUsernameCreated() *AddUsernameCreated {

	return &AddUsernameCreated{}
}

// WithPayload adds the payload to the add username created response
func (o *AddUsernameCreated) WithPayload(payload *AddUsernameCreatedBody) *AddUsernameCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add username created response
func (o *AddUsernameCreated) SetPayload(payload *AddUsernameCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUsernameCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUsernameConflictCode is the HTTP code returned for type AddUsernameConflict
const AddUsernameConflictCode int = 409

/*AddUsernameConflict username already exists

swagger:response addUsernameConflict
*/
type AddUsernameConflict struct {
}

// NewAddUsernameConflict creates AddUsernameConflict with default headers values
func NewAddUsernameConflict() *AddUsernameConflict {

	return &AddUsernameConflict{}
}

// WriteResponse to the client
func (o *AddUsernameConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// AddUsernameInternalServerErrorCode is the HTTP code returned for type AddUsernameInternalServerError
const AddUsernameInternalServerErrorCode int = 500

/*AddUsernameInternalServerError internal server error

swagger:response addUsernameInternalServerError
*/
type AddUsernameInternalServerError struct {
}

// NewAddUsernameInternalServerError creates AddUsernameInternalServerError with default headers values
func NewAddUsernameInternalServerError() *AddUsernameInternalServerError {

	return &AddUsernameInternalServerError{}
}

// WriteResponse to the client
func (o *AddUsernameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
