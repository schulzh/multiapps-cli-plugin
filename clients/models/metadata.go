package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Metadata Container element for specifying metadata
// swagger:model Metadata
type Metadata struct {
	XMLName xml.Name `xml:"http://www.sap.com/lmsl/slp Metadata"`

	// slmpversion
	Slmpversion string `xml:"slmpversion,omitempty"`

	// slppversion
	Slppversion string `xml:"slppversion,omitempty"`
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
