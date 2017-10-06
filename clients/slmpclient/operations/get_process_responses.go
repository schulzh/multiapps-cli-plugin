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

// GetProcessReader is a Reader for the GetProcess structure.
type GetProcessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProcessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProcessOK creates a GetProcessOK with default headers values
func NewGetProcessOK() *GetProcessOK {
	return &GetProcessOK{}
}

/*GetProcessOK handles this case with default header values.

OK
*/
type GetProcessOK struct {
	Payload *models.Process
}

func (o *GetProcessOK) Error() string {
	return fmt.Sprintf("[GET /processes/{processId}][%d] getProcessOK  %+v", 200, o.Payload)
}

func (o *GetProcessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Process)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
