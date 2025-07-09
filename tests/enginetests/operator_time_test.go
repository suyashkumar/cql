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

package enginetests

import (
	"context"
	"testing"

	"github.com/google/cql/interpreter"
	"github.com/google/cql/parser"
	"github.com/google/cql/result"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestEquivalentTime(t *testing.T) {
	tests := []struct {
		name       string
		cql        string
		wantResult result.Value
	}{
		{
			name:       "T10:20:30.123 ~ T10:20:30.123",
			cql:        "T10:20:30.123 ~ T10:20:30.123",
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "T10:20:30.123 ~ T10:20:30.124",
			cql:        "T10:20:30.123 ~ T10:20:30.124",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "T10:20:30.123 ~ T10:20",
			cql:        "T10:20:30.123 ~ T10:20",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "T10:20 ~ T10:20",
			cql:        "T10:20 ~ T10:20",
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "T10:20 ~ T10:21",
			cql:        "T10:20 ~ T10:21",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "null as Time ~ null as Time",
			cql:        "null as Time ~ null as Time",
			wantResult: newOrFatal(t, true),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantResult, getTESTRESULT(t, results), protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}
		})
	}
}

func TestCompareTime(t *testing.T) {
	tests := []struct {
		name       string
		cql        string
		wantResult result.Value
	}{
		// Greater
		{
			name:       "T11 > T10",
			cql:        "T11 > T10",
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "T10 > T11",
			cql:        "T10 > T11",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "T10 > T10",
			cql:        "T10 > T10",
			wantResult: newOrFatal(t, false),
		},
		// Greater or Equal
		{
			name:       "T11 >= T10",
			cql:        "T11 >= T10",
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "T10 >= T11",
			cql:        "T10 >= T11",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "T10 >= T10",
			cql:        "T10 >= T10",
			wantResult: newOrFatal(t, true),
		},
		// Less
		{
			name:       "T10 < T11",
			cql:        "T10 < T11",
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "T11 < T10",
			cql:        "T11 < T10",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "T10 < T10",
			cql:        "T10 < T10",
			wantResult: newOrFatal(t, false),
		},
		// Less or Equal
		{
			name:       "T10 <= T11",
			cql:        "T10 <= T11",
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "T11 <= T10",
			cql:        "T11 <= T10",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "T10 <= T10",
			cql:        "T10 <= T10",
			wantResult: newOrFatal(t, true),
		},
		// Precision
		{
			name:       "T10:20 > T10",
			cql:        "T10:20 > T10",
			wantResult: newOrFatal(t, nil),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantResult, getTESTRESULT(t, results), protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}
		})
	}
}
