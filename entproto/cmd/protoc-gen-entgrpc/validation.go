// Copyright 2019-present Facebook
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

package main

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema/field"
)

func typeNeedsValidator(d entproto.FieldMap) bool {
	for _, fld := range d {
		if fieldNeedsValidator(fld) {
			return true
		}
	}
	return false
}

func fieldNeedsValidator(d *entproto.FieldMappingDescriptor) bool {
	// TODO: check if edge
	if d.EntField == nil {
		return false
	}
	switch d.EntField.Type.Type {
	// TODO: rm TypeString, leave UUID once https://github.com/ent/contrib/pull/54 is merged
	case field.TypeUUID, field.TypeString:
		return true
	default:
		return false
	}
}
