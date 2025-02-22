// Copyright The OpenTelemetry Authors
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

package diff

import (
	"reflect"
	"testing"
)

func TestDiffMaps(t *testing.T) {
	type args struct {
		current map[string]string
		new     map[string]string
	}
	tests := []struct {
		name string
		args args
		want Changes[string]
	}{
		{
			name: "basic replacement",
			args: args{
				current: map[string]string{
					"current": "one",
				},
				new: map[string]string{
					"new": "another",
				},
			},
			want: Changes[string]{
				additions: map[string]string{
					"new": "another",
				},
				removals: map[string]string{
					"current": "one",
				},
			},
		},
		{
			name: "single addition",
			args: args{
				current: map[string]string{
					"current": "one",
				},
				new: map[string]string{
					"current": "one",
					"new":     "another",
				},
			},
			want: Changes[string]{
				additions: map[string]string{
					"new": "another",
				},
				removals: map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maps(tt.args.current, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
