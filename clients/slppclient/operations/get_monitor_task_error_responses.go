package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/SAP/cf-mta-plugin/clients/models"
)

// GetMonitorTaskErrorReader is a Reader for the GetMonitorTaskError structure.
type GetMonitorTaskErrorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMonitorTaskErrorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMonitorTaskErrorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMonitorTaskErrorOK creates a GetMonitorTaskErrorOK with default headers values
func NewGetMonitorTaskErrorOK() *GetMonitorTaskErrorOK {
	return &GetMonitorTaskErrorOK{}
}

/*GetMonitorTaskErrorOK handles this case with default header values.

OK
*/
type GetMonitorTaskErrorOK struct {
	Payload *models.Error
}

func (o *GetMonitorTaskErrorOK) Error() string {
	return fmt.Sprintf("[GET /monitor/{taskId}/error][%d] getMonitorTaskErrorOK  %+v", 200, o.Payload)
}

func (o *GetMonitorTaskErrorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
