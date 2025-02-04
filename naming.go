// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
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

package eiffelevents

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
)

// EventStructName returns the name of the Go struct used to represent
// a particular version of an event type.
func EventStructName(eventType string, eventVersion *semver.Version) string {
	return fmt.Sprintf("%sV%d", strings.TrimPrefix(strings.TrimSuffix(eventType, "Event"), "Eiffel"), eventVersion.Major())
}
