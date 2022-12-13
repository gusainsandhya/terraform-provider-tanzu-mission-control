/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
Code generated by go-swagger; DO NOT EDIT.
*/

package policyrecipeimagecommonmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag Tag
//
// Allowed image tag, wildcards are supported (for example: v1.*). No validation is performed on tag if the field is empty.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag
type VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag struct {

	// The negate flag used to exclude certain tag patterns.
	Negate *bool `json:"negate,omitempty"`

	// The value (support wildcard) is used to validate against the tag of the image.
	Value string `json:"value,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
