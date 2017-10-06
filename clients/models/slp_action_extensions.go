package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SlpActionExtensions slp action extensions
// swagger:model slp_action_extensions
type SlpActionExtensions string

// Validate validates this slp action extensions
func (m SlpActionExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Pattern("", "body", string(m), `slp.action.[A-Z]+`); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
