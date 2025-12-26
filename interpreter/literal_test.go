// Copyright 2024 Google LLC
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

package interpreter

import (
	"fmt"
	"testing"

	"github.com/google/cql/model"
	"github.com/google/cql/types"
)

func BenchmarkEvalInstance(b *testing.B) {
	const numElements = 1000
	elements := make([]*model.InstanceElement, numElements)
	for i := 0; i < numElements; i++ {
		elements[i] = &model.InstanceElement{
			Name:  fmt.Sprintf("elem%d", i),
			Value: &model.Literal{Value: "1", Expression: model.ResultType(types.Integer)},
		}
	}
	instance := &model.Instance{
		Elements:  elements,
		ClassType: &types.Named{TypeName: "Test.TestClass"},
	}
	i := &interpreter{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := i.evalInstance(instance)
		if err != nil {
			b.Fatalf("evalInstance failed: %v", err)
		}
	}
}
