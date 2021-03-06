// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SiteDefaultHooksData site default hooks data
//
// swagger:model siteDefaultHooksData
type SiteDefaultHooksData struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`
}

// Validate validates this site default hooks data
func (m *SiteDefaultHooksData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SiteDefaultHooksData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteDefaultHooksData) UnmarshalBinary(b []byte) error {
	var res SiteDefaultHooksData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
