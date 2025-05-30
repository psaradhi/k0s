/*
Copyright 2020 k0s authors

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

package v1beta1

import "k8s.io/utils/ptr"

var _ Validateable = (*ClusterTelemetry)(nil)

// ClusterTelemetry holds telemetry related settings
type ClusterTelemetry struct {
	// +kubebuilder:default=false
	Enabled *bool `json:"enabled,omitempty"`
}

func (t *ClusterTelemetry) IsEnabled() bool {
	return t != nil && t.Enabled != nil && *t.Enabled
}

// DefaultClusterTelemetry default settings
func DefaultClusterTelemetry() *ClusterTelemetry {
	return &ClusterTelemetry{
		Enabled: ptr.To(false),
	}
}

// Validate stub for Validateable interface
func (c *ClusterTelemetry) Validate() []error {
	return nil
}
