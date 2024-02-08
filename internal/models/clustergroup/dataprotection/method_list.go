/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse Response from listing DataProtections.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.ListDataProtectionsResponse
type VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse struct {

	// List of dataprotections.
	DataProtections []*VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection `json:"dataProtections"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupDataprotectionListDataProtectionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}