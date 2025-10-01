// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package interpreter

import (
	"context"
	"testing"

	"github.com/google/cql/model"
	"github.com/google/cql/result"
	"github.com/google/cql/types"
	"github.com/google/go-cmp/cmp"
)

func TestToBoolean(t *testing.T) {
	tests := []struct {
		name       string
		expression model.IExpression
		want       result.Value
	}{
		{
			name: "String true to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'true'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "String t to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'t'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "String YES to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'YES'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "String y to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'y'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "String 1 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'1'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "String false to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'false'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "String f to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'f'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "String NO to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'NO'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "String n to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'n'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "String 0 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'0'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "Invalid String to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("'invalid'", types.String),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(nil),
		},
		{
			name: "Integer 1 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("1", types.Integer),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "Integer 0 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("0", types.Integer),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "Integer 2 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("2", types.Integer),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(nil),
		},
		{
			name: "Long 1 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("1L", types.Long),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "Long 0 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("0L", types.Long),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "Long 2 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("2L", types.Long),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(nil),
		},
		{
			name: "Decimal 1.0 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("1.0", types.Decimal),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(true),
		},
		{
			name: "Decimal 0.0 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("0.0", types.Decimal),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(false),
		},
		{
			name: "Decimal 2.0 to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("2.0", types.Decimal),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(nil),
		},
		{
			name: "Null to Boolean",
			expression: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("null", types.Any),
					Expression: model.ResultType(types.Boolean),
				},
			},
			want: result.New(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			i := &interpreter{}
			got, err := i.evalExpression(test.expression)
			if err != nil {
				t.Fatalf("evalExpression(%v) returned error: %v", test.expression, err)
			}
			if diff := cmp.Diff(test.want, got, cmp.AllowUnexported(result.Value{})); diff != "" {
				t.Errorf("evalExpression(%v) returned diff (-want +got):\n%s", test.expression, diff)
			}
		})
	}
}
