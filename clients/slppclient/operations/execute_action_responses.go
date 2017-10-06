// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExecuteActionReader is a Reader for the ExecuteAction structure.
type ExecuteActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewExecuteActionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExecuteActionNoContent creates a ExecuteActionNoContent with default headers values
func NewExecuteActionNoContent() *ExecuteActionNoContent {
	return &ExecuteActionNoContent{}
}

/*ExecuteActionNoContent handles this case with default header values.

Executed
*/
type ExecuteActionNoContent struct {
}

func (o *ExecuteActionNoContent) Error() string {
	return fmt.Sprintf("[POST /actions/{actionId}][%d] executeActionNoContent ", 204)
}

func (o *ExecuteActionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
