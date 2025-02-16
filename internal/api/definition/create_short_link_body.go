// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateShortLinkBody create short link body
//
// swagger:model CreateShortLinkBody
type CreateShortLinkBody struct {

	// original link
	// Required: true
	OriginalLink *string `json:"original_link"`
}

// Validate validates this create short link body
func (m *CreateShortLinkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOriginalLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShortLinkBody) validateOriginalLink(formats strfmt.Registry) error {

	if err := validate.Required("original_link", "body", m.OriginalLink); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create short link body based on context it is used
func (m *CreateShortLinkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateShortLinkBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShortLinkBody) UnmarshalBinary(b []byte) error {
	var res CreateShortLinkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
