// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type fieldSetter struct {
	fieldMap map[string]string
}

func (fs fieldSetter) SetFields(o *unstructured.Unstructured, annotationValues map[string]interface{}) error {
	for field, value := range annotationValues {
		path := fs.fieldMap[field]
		fieldPath := strings.Split(path, ".")
		err := unstructured.SetNestedField(o.UnstructuredContent(), value, fieldPath...)
		if err != nil {
			return err
		}
	}
	return nil
}
