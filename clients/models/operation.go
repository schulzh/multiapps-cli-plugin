package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Operation operation
// swagger:model operation
type Operation struct {
	XMLName xml.Name `xml:"ongoing-operation"`

	// acquired lock
	// Required: true
	AcquiredLock *bool `xml:"acquired-lock,omitempty"`

	// mta id
	MtaID string `xml:"mta-id,omitempty"`

	// process id
	// Required: true
	ProcessID *string `xml:"process-id,omitempty"`

	// process type
	// Required: true
	ProcessType ProcessType `xml:"process-type,omitempty"`

	// space id
	// Required: true
	SpaceID *string `xml:"space-id,omitempty"`

	// started at
	// Required: true
	StartedAt *string `xml:"started-at,omitempty"`

	// state
	// Required: true
	State SlpTaskStateEnum `xml:"state,omitempty"`

	// user
	// Required: true
	User *string `xml:"user"`
}

// Validate validates this operation
func (m *Operation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcquiredLock(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcessID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcessType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpaceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Operation) validateAcquiredLock(formats strfmt.Registry) error {

	if err := validate.Required("acquired-lock", "body", m.AcquiredLock); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validateProcessID(formats strfmt.Registry) error {

	if err := validate.Required("process-id", "body", m.ProcessID); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validateProcessType(formats strfmt.Registry) error {

	if err := m.ProcessType.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validateSpaceID(formats strfmt.Registry) error {

	if err := validate.Required("space-id", "body", m.SpaceID); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("started-at", "body", m.StartedAt); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validateState(formats strfmt.Registry) error {

	if err := m.State.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *Operation) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}
