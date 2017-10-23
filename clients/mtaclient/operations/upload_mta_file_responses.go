// Code generated by go-swagger; DO NOT EDIT.

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

// UploadMtaFileReader is a Reader for the UploadMtaFile structure.
type UploadMtaFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadMtaFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUploadMtaFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadMtaFileCreated creates a UploadMtaFileCreated with default headers values
func NewUploadMtaFileCreated() *UploadMtaFileCreated {
	return &UploadMtaFileCreated{}
}

/*UploadMtaFileCreated handles this case with default header values.

Created
*/
type UploadMtaFileCreated struct {
	Payload *models.FileMetadata
}

func (o *UploadMtaFileCreated) Error() string {
	return fmt.Sprintf("[POST /files][%d] uploadMtaFileCreated  %+v", 201, o.Payload)
}

func (o *UploadMtaFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
