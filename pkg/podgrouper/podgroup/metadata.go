// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package podgroup

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Metadata struct {
	Annotations       map[string]string
	Labels            map[string]string
	PriorityClassName string
	Queue             string
	Namespace         string
	Name              string
	MinAvailable      int32
	Owner             metav1.OwnerReference

	PreferredTopologyLevel string
	RequiredTopologyLevel  string
	Topology               string
}
