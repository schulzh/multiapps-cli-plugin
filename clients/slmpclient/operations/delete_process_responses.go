package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProcessReader is a Reader for the DeleteProcess structure.
type DeleteProcessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProcessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProcessNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProcessNoContent creates a DeleteProcessNoContent with default headers values
func NewDeleteProcessNoContent() *DeleteProcessNoContent {
	return &DeleteProcessNoContent{}
}

/*DeleteProcessNoContent handles this case with default header values.

Deleted
*/
type DeleteProcessNoContent struct {
}

func (o *DeleteProcessNoContent) Error() string {
	return fmt.Sprintf("[DELETE /processes/{processId}][%d] deleteProcessNoContent ", 204)
}

func (o *DeleteProcessNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
