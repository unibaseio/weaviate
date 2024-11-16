//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TenantResponse attributes representing a single tenant response within weaviate
//
// swagger:model TenantResponse
type TenantResponse struct {
	Tenant

	// The list of nodes that owns that tenant data.
	BelongsToNodes []string `json:"belongsToNodes"`

	// Experimental. The data version of the tenant is a monotonically increasing number starting from 0 which is incremented each time a tenant's data is offloaded to cloud storage.
	// Example: 3
	// Minimum: 0
	DataVersion *int64 `json:"dataVersion,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TenantResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Tenant
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Tenant = aO0

	// AO1
	var dataAO1 struct {
		BelongsToNodes []string `json:"belongsToNodes"`

		DataVersion *int64 `json:"dataVersion,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BelongsToNodes = dataAO1.BelongsToNodes

	m.DataVersion = dataAO1.DataVersion

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TenantResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Tenant)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		BelongsToNodes []string `json:"belongsToNodes"`

		DataVersion *int64 `json:"dataVersion,omitempty"`
	}

	dataAO1.BelongsToNodes = m.BelongsToNodes

	dataAO1.DataVersion = m.DataVersion

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tenant response
func (m *TenantResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Tenant
	if err := m.Tenant.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantResponse) validateDataVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.DataVersion) { // not required
		return nil
	}

	if err := validate.MinimumInt("dataVersion", "body", *m.DataVersion, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this tenant response based on the context it is used
func (m *TenantResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Tenant
	if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TenantResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantResponse) UnmarshalBinary(b []byte) error {
	var res TenantResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
