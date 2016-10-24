package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/chronograf/models"
)

/*DeleteUsersUserIDExplorationsExplorationIDNoContent Exploration session has been removed

swagger:response deleteUsersUserIdExplorationsExplorationIdNoContent
*/
type DeleteUsersUserIDExplorationsExplorationIDNoContent struct {
}

// NewDeleteUsersUserIDExplorationsExplorationIDNoContent creates DeleteUsersUserIDExplorationsExplorationIDNoContent with default headers values
func NewDeleteUsersUserIDExplorationsExplorationIDNoContent() *DeleteUsersUserIDExplorationsExplorationIDNoContent {
	return &DeleteUsersUserIDExplorationsExplorationIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteUsersUserIDExplorationsExplorationIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteUsersUserIDExplorationsExplorationIDNotFound Data source id, user, or exploration does not exist.

swagger:response deleteUsersUserIdExplorationsExplorationIdNotFound
*/
type DeleteUsersUserIDExplorationsExplorationIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUsersUserIDExplorationsExplorationIDNotFound creates DeleteUsersUserIDExplorationsExplorationIDNotFound with default headers values
func NewDeleteUsersUserIDExplorationsExplorationIDNotFound() *DeleteUsersUserIDExplorationsExplorationIDNotFound {
	return &DeleteUsersUserIDExplorationsExplorationIDNotFound{}
}

// WithPayload adds the payload to the delete users user Id explorations exploration Id not found response
func (o *DeleteUsersUserIDExplorationsExplorationIDNotFound) WithPayload(payload *models.Error) *DeleteUsersUserIDExplorationsExplorationIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete users user Id explorations exploration Id not found response
func (o *DeleteUsersUserIDExplorationsExplorationIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUsersUserIDExplorationsExplorationIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteUsersUserIDExplorationsExplorationIDDefault Unexpected internal service error

swagger:response deleteUsersUserIdExplorationsExplorationIdDefault
*/
type DeleteUsersUserIDExplorationsExplorationIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUsersUserIDExplorationsExplorationIDDefault creates DeleteUsersUserIDExplorationsExplorationIDDefault with default headers values
func NewDeleteUsersUserIDExplorationsExplorationIDDefault(code int) *DeleteUsersUserIDExplorationsExplorationIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUsersUserIDExplorationsExplorationIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete users user ID explorations exploration ID default response
func (o *DeleteUsersUserIDExplorationsExplorationIDDefault) WithStatusCode(code int) *DeleteUsersUserIDExplorationsExplorationIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete users user ID explorations exploration ID default response
func (o *DeleteUsersUserIDExplorationsExplorationIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete users user ID explorations exploration ID default response
func (o *DeleteUsersUserIDExplorationsExplorationIDDefault) WithPayload(payload *models.Error) *DeleteUsersUserIDExplorationsExplorationIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete users user ID explorations exploration ID default response
func (o *DeleteUsersUserIDExplorationsExplorationIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUsersUserIDExplorationsExplorationIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}