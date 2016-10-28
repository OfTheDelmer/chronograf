package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/chronograf/models"
)

/*DeleteSourcesIDKapacitorsKapaIDNoContent kapacitor has been removed.

swagger:response deleteSourcesIdKapacitorsKapaIdNoContent
*/
type DeleteSourcesIDKapacitorsKapaIDNoContent struct {
}

// NewDeleteSourcesIDKapacitorsKapaIDNoContent creates DeleteSourcesIDKapacitorsKapaIDNoContent with default headers values
func NewDeleteSourcesIDKapacitorsKapaIDNoContent() *DeleteSourcesIDKapacitorsKapaIDNoContent {
	return &DeleteSourcesIDKapacitorsKapaIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteSourcesIDKapacitorsKapaIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteSourcesIDKapacitorsKapaIDNotFound Unknown Data source or Kapacitor id

swagger:response deleteSourcesIdKapacitorsKapaIdNotFound
*/
type DeleteSourcesIDKapacitorsKapaIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSourcesIDKapacitorsKapaIDNotFound creates DeleteSourcesIDKapacitorsKapaIDNotFound with default headers values
func NewDeleteSourcesIDKapacitorsKapaIDNotFound() *DeleteSourcesIDKapacitorsKapaIDNotFound {
	return &DeleteSourcesIDKapacitorsKapaIDNotFound{}
}

// WithPayload adds the payload to the delete sources Id kapacitors kapa Id not found response
func (o *DeleteSourcesIDKapacitorsKapaIDNotFound) WithPayload(payload *models.Error) *DeleteSourcesIDKapacitorsKapaIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete sources Id kapacitors kapa Id not found response
func (o *DeleteSourcesIDKapacitorsKapaIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSourcesIDKapacitorsKapaIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSourcesIDKapacitorsKapaIDDefault Unexpected internal service error

swagger:response deleteSourcesIdKapacitorsKapaIdDefault
*/
type DeleteSourcesIDKapacitorsKapaIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSourcesIDKapacitorsKapaIDDefault creates DeleteSourcesIDKapacitorsKapaIDDefault with default headers values
func NewDeleteSourcesIDKapacitorsKapaIDDefault(code int) *DeleteSourcesIDKapacitorsKapaIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSourcesIDKapacitorsKapaIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete sources ID kapacitors kapa ID default response
func (o *DeleteSourcesIDKapacitorsKapaIDDefault) WithStatusCode(code int) *DeleteSourcesIDKapacitorsKapaIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete sources ID kapacitors kapa ID default response
func (o *DeleteSourcesIDKapacitorsKapaIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete sources ID kapacitors kapa ID default response
func (o *DeleteSourcesIDKapacitorsKapaIDDefault) WithPayload(payload *models.Error) *DeleteSourcesIDKapacitorsKapaIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete sources ID kapacitors kapa ID default response
func (o *DeleteSourcesIDKapacitorsKapaIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSourcesIDKapacitorsKapaIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
