// Copyright 2020 The Tekton Authors
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

package annotation

import "encoding/json"

const (
	Result = "results.tekton.dev/result"
	Record = "results.tekton.dev/record"
	Log    = "results.tekton.dev/log"
)

type Annotation struct {
	Name  string
	Value string
}

// Add creates a jsonpatch path used for adding result / record identifiers
// an object's annotations field.
func Add(annotation ...Annotation) ([]byte, error) {
	annotations := make(map[string]string)
	for _, annotation := range annotation {
		if len(annotation.Value) > 0 {
			annotations[annotation.Name] = annotation.Value
		}
	}
	data := map[string]interface{}{
		"metadata": map[string]interface{}{
			"annotations": annotations,
		},
	}
	return json.Marshal(data)
}
