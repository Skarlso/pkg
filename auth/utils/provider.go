/*
Copyright 2025 The Flux authors

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

package authutils

import (
	"github.com/fluxcd/pkg/auth"
	"github.com/fluxcd/pkg/auth/aws"
	"github.com/fluxcd/pkg/auth/azure"
	"github.com/fluxcd/pkg/auth/gcp"
)

// ProviderByName looks up the implemented providers by name.
func ProviderByName(name string) auth.Provider {
	switch name {
	case aws.ProviderName:
		return aws.Provider{}
	case azure.ProviderName:
		return azure.Provider{}
	case gcp.ProviderName:
		return gcp.Provider{}
	default:
		return nil
	}
}
