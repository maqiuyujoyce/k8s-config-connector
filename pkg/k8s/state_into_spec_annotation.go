// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// StateIntoSpecValue contains the required 'defaultValue' field and the
// optional 'userOverride' field.
type StateIntoSpecValue struct {
	defaultValue string
	userOverride *string
}

func NewStateIntoSpecValue(defaultValue string, userOverride *string) (*StateIntoSpecValue, error) {
	if !isAcceptedValue(defaultValue, StateIntoSpecAnnotationValues) {
		return nil, fmt.Errorf("invalid default value '%v' for '%v' annotation, need to be one of {%v}", defaultValue, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	if userOverride != nil && !isAcceptedValue(*userOverride, StateIntoSpecAnnotationValues) {
		return nil, fmt.Errorf("invalid user override value '%v' for '%v' annotation, need to be one of {%v}", userOverride, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	return &StateIntoSpecValue{
		defaultValue: defaultValue,
		userOverride: userOverride,
	}, nil
}

func (v *StateIntoSpecValue) GetValue() string {
	if v.userOverride == nil {
		return v.defaultValue
	}
	return *v.userOverride
}

// ValidateOrDefaultStateIntoSpecAnnotation defaults the annotation to the
// passed in defaultValue if it is unset and validates the value of the
// 'state-into-spec' annotation.
func ValidateOrDefaultStateIntoSpecAnnotation(obj metav1.Object, gvk schema.GroupVersionKind, defaultValue string) error {
	_, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		defaultStateIntoSpecAnnotation(obj, gvk, defaultValue)
	}
	return validateStateIntoSpecAnnotation(obj, gvk)
}

func defaultStateIntoSpecAnnotation(obj metav1.Object, gvk schema.GroupVersionKind, defaultValue string) {
	if defaultValue == StateAbsentInSpec && !ResourceSupportsStateAbsentInSpec(gvk.Kind) {
		SetAnnotation(StateIntoSpecAnnotation, StateMergeIntoSpec, obj)
		return
	}
	SetAnnotation(StateIntoSpecAnnotation, defaultValue, obj)
	return
}

// ResourceSupportsStateAbsentInSpec returns true for resource kinds which
// allow the 'state-into-spec' annotation to be set to 'absent'.
func ResourceSupportsStateAbsentInSpec(kind string) bool {
	switch kind {
	// Setting 'state-into-spec' to 'absent' for ComputeAddress may hide 'spec.address' field from users and cause breaking change.
	case "ComputeAddress":
		return false
	}
	return true
}

func validateStateIntoSpecAnnotation(obj metav1.Object, gvk schema.GroupVersionKind) error {
	val, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		return fmt.Errorf("couldn't find the value for '%v' annotation", StateIntoSpecAnnotation)
	}

	if !isAcceptedValue(val, StateIntoSpecAnnotationValues) {
		return fmt.Errorf("invalid value '%v' for '%v' annotation, can be one of {%v}", val, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}

	if val == StateAbsentInSpec && !ResourceSupportsStateAbsentInSpec(gvk.Kind) {
		return fmt.Errorf("kind '%v' does not support having annotation '%v' set to value '%v'", gvk.Kind, StateIntoSpecAnnotation, val)
	}
	return nil
}

func isAcceptedValue(val string, acceptedValues []string) bool {
	for _, v := range acceptedValues {
		if val == v {
			return true
		}
	}
	return false
}
