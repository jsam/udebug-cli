// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Input input
// swagger:model Input
type Input struct {

	// date
	Date string `json:"Date,omitempty"`

	// votes
	Votes interface{} `json:"Votes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this input
func (m *Input) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Input) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Input) UnmarshalBinary(b []byte) error {
	var res Input
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}