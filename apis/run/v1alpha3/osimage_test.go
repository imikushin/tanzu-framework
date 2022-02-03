// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachineImageInfo_DeepCopy(t *testing.T) {
	const amiID = "ami-0f2e5eec7ae0a1986"
	const isFoo = true
	var ref interface{} = map[string]interface{}{
		"id":    amiID,
		"isFoo": isFoo,
	}
	imageInfo := &MachineImageInfo{
		Type: "aws",
		Ref:  ref,
	}
	imageInfoCopy := imageInfo.DeepCopy()
	assert.Equal(t, imageInfo, imageInfoCopy)
	assert.Equal(t, amiID, imageInfoCopy.Ref.(map[string]interface{})["id"])
	assert.Equal(t, isFoo, imageInfoCopy.Ref.(map[string]interface{})["isFoo"])
}
