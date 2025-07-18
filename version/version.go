/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
)

// ParseVersion behaves as semver.StrictNewVersion, with as sole exception
// that it allows versions with a preceding "v" (i.e. v1.2.3).
func ParseVersion(v string) (*semver.Version, error) {
	vLessV := strings.TrimPrefix(v, "v")
	if _, err := semver.StrictNewVersion(vLessV); err != nil {
		return nil, err
	}
	return semver.NewVersion(v)
}

// Sort filters the given strings based on the provided semver range
// and sorts them in descending order.
func Sort(c *semver.Constraints, vs []string) []string {
	var versions []*semver.Version
	for _, v := range vs {
		if pv, err := ParseVersion(v); err == nil && (c == nil || c.Check(pv)) {
			versions = append(versions, pv)
		}
	}
	sort.Sort(sort.Reverse(semver.Collection(versions)))
	sorted := make([]string, 0, len(versions))
	for _, v := range versions {
		sorted = append(sorted, v.Original())
	}
	return sorted
}
