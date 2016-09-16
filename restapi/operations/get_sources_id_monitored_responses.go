package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*GetSourcesIDMonitoredOK An array of permissions

swagger:response getSourcesIdMonitoredOK
*/
type GetSourcesIDMonitoredOK struct {

	// In: body
	Payload *models.Services `json:"body,omitempty"`
}

// NewGetSourcesIDMonitoredOK creates GetSourcesIDMonitoredOK with default headers values
func NewGetSourcesIDMonitoredOK() *GetSourcesIDMonitoredOK {
	return &GetSourcesIDMonitoredOK{}
}

// WithPayload adds the payload to the get sources Id monitored o k response
func (o *GetSourcesIDMonitoredOK) WithPayload(payload *models.Services) *GetSourcesIDMonitoredOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id monitored o k response
func (o *GetSourcesIDMonitoredOK) SetPayload(payload *models.Services) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDMonitoredOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDMonitoredNotFound Data source id does not exist.

swagger:response getSourcesIdMonitoredNotFound
*/
type GetSourcesIDMonitoredNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDMonitoredNotFound creates GetSourcesIDMonitoredNotFound with default headers values
func NewGetSourcesIDMonitoredNotFound() *GetSourcesIDMonitoredNotFound {
	return &GetSourcesIDMonitoredNotFound{}
}

// WithPayload adds the payload to the get sources Id monitored not found response
func (o *GetSourcesIDMonitoredNotFound) WithPayload(payload *models.Error) *GetSourcesIDMonitoredNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id monitored not found response
func (o *GetSourcesIDMonitoredNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDMonitoredNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDMonitoredDefault Unexpected internal service error

swagger:response getSourcesIdMonitoredDefault
*/
type GetSourcesIDMonitoredDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDMonitoredDefault creates GetSourcesIDMonitoredDefault with default headers values
func NewGetSourcesIDMonitoredDefault(code int) *GetSourcesIDMonitoredDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSourcesIDMonitoredDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sources ID monitored default response
func (o *GetSourcesIDMonitoredDefault) WithStatusCode(code int) *GetSourcesIDMonitoredDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sources ID monitored default response
func (o *GetSourcesIDMonitoredDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sources ID monitored default response
func (o *GetSourcesIDMonitoredDefault) WithPayload(payload *models.Error) *GetSourcesIDMonitoredDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources ID monitored default response
func (o *GetSourcesIDMonitoredDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDMonitoredDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}