// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OryAccessControlPolicy oryAccessControlPolicy specifies an ORY Access Policy document.
//
// swagger:model oryAccessControlPolicy
type OryAccessControlPolicy struct {

	// Actions is an array representing all the actions this ORY Access Policy applies to.
	Actions []string `json:"actions"`

	// Conditions represents a keyed object of conditions under which this ORY Access Policy is active.
	Conditions interface{} `json:"conditions,omitempty"`

	// Description is an optional, human-readable description.
	Description string `json:"description,omitempty"`

	// Effect is the effect of this ORY Access Policy. It can be "allow" or "deny".
	Effect string `json:"effect,omitempty"`

	// ID is the unique identifier of the ORY Access Policy. It is used to query, update, and remove the ORY Access Policy.
	ID string `json:"id,omitempty"`

	// Resources is an array representing all the resources this ORY Access Policy applies to.
	Resources []string `json:"resources"`

	// Subjects is an array representing all the subjects this ORY Access Policy applies to.
	Subjects []string `json:"subjects"`
}

// Validate validates this ory access control policy
func (m *OryAccessControlPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OryAccessControlPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OryAccessControlPolicy) UnmarshalBinary(b []byte) error {
	var res OryAccessControlPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
