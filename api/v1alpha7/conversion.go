/*
Copyright 2023 The Kubernetes Authors.

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

package v1alpha7

import (
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha8"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/utils/conversion"
)

var _ ctrlconversion.Convertible = &OpenStackCluster{}

var v1alpha7OpenStackClusterRestorer = conversion.RestorerFor[*OpenStackCluster]{
	"spec": conversion.HashedFieldRestorer[*OpenStackCluster, OpenStackClusterSpec]{
		GetField: func(c *OpenStackCluster) *OpenStackClusterSpec {
			return &c.Spec
		},
	},
	"status": conversion.HashedFieldRestorer[*OpenStackCluster, OpenStackClusterStatus]{
		GetField: func(c *OpenStackCluster) *OpenStackClusterStatus {
			return &c.Status
		},
	},
}

var v1alpha8OpenStackClusterRestorer = conversion.RestorerFor[*infrav1.OpenStackCluster]{
	"status": conversion.HashedFieldRestorer[*infrav1.OpenStackCluster, infrav1.OpenStackClusterStatus]{
		GetField: func(c *infrav1.OpenStackCluster) *infrav1.OpenStackClusterStatus {
			return &c.Status
		},
	},
}

func (r *OpenStackCluster) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackCluster)

	compare := &OpenStackCluster{}
	return conversion.ConvertAndRestore(
		r, dst, compare,
		Convert_v1alpha7_OpenStackCluster_To_v1alpha8_OpenStackCluster, Convert_v1alpha8_OpenStackCluster_To_v1alpha7_OpenStackCluster,
		v1alpha7OpenStackClusterRestorer, v1alpha8OpenStackClusterRestorer,
	)
}

func (r *OpenStackCluster) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackCluster)

	compare := &infrav1.OpenStackCluster{}
	return conversion.ConvertAndRestore(
		src, r, compare,
		Convert_v1alpha8_OpenStackCluster_To_v1alpha7_OpenStackCluster, Convert_v1alpha7_OpenStackCluster_To_v1alpha8_OpenStackCluster,
		v1alpha8OpenStackClusterRestorer, v1alpha7OpenStackClusterRestorer,
	)
}

var _ ctrlconversion.Convertible = &OpenStackClusterList{}

func (r *OpenStackClusterList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackClusterList)

	return Convert_v1alpha7_OpenStackClusterList_To_v1alpha8_OpenStackClusterList(r, dst, nil)
}

func (r *OpenStackClusterList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackClusterList)

	return Convert_v1alpha8_OpenStackClusterList_To_v1alpha7_OpenStackClusterList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackClusterTemplate{}

var v1alpha7OpenStackClusterTemplateRestorer = conversion.RestorerFor[*OpenStackClusterTemplate]{
	"spec": conversion.HashedFieldRestorer[*OpenStackClusterTemplate, OpenStackClusterSpec]{
		GetField: func(c *OpenStackClusterTemplate) *OpenStackClusterSpec {
			return &c.Spec.Template.Spec
		},
	},
}

var v1alpha8OpenStackClusterTemplateRestorer = conversion.RestorerFor[*infrav1.OpenStackClusterTemplate]{}

func (r *OpenStackClusterTemplate) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackClusterTemplate)

	compare := &OpenStackClusterTemplate{}
	return conversion.ConvertAndRestore(
		r, dst, compare,
		Convert_v1alpha7_OpenStackClusterTemplate_To_v1alpha8_OpenStackClusterTemplate, Convert_v1alpha8_OpenStackClusterTemplate_To_v1alpha7_OpenStackClusterTemplate,
		v1alpha7OpenStackClusterTemplateRestorer, v1alpha8OpenStackClusterTemplateRestorer,
	)
}

func (r *OpenStackClusterTemplate) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackClusterTemplate)

	compare := &infrav1.OpenStackClusterTemplate{}
	return conversion.ConvertAndRestore(
		src, r, compare,
		Convert_v1alpha8_OpenStackClusterTemplate_To_v1alpha7_OpenStackClusterTemplate, Convert_v1alpha7_OpenStackClusterTemplate_To_v1alpha8_OpenStackClusterTemplate,
		v1alpha8OpenStackClusterTemplateRestorer, v1alpha7OpenStackClusterTemplateRestorer,
	)
}

var _ ctrlconversion.Convertible = &OpenStackMachine{}

var v1alpha7OpenStackMachineRestorer = conversion.RestorerFor[*OpenStackMachine]{
	"spec": conversion.HashedFieldRestorer[*OpenStackMachine, OpenStackMachineSpec]{
		GetField: func(c *OpenStackMachine) *OpenStackMachineSpec {
			return &c.Spec
		},
		FilterField: func(s *OpenStackMachineSpec) *OpenStackMachineSpec {
			// Despite being spec fields, ProviderID and InstanceID
			// are both set by the machine controller. If these are
			// the only changes to the spec, we still want to
			// restore the rest of the spec to its original state.
			if s.ProviderID != nil || s.InstanceID != nil {
				f := *s
				f.ProviderID = nil
				f.InstanceID = nil
				return &f
			}
			return s
		},
	},
}

var v1alpha8OpenStackMachineRestorer = conversion.RestorerFor[*infrav1.OpenStackMachine]{
	"spec": conversion.HashedFieldRestorer[*infrav1.OpenStackMachine, infrav1.OpenStackMachineSpec]{
		GetField: func(c *infrav1.OpenStackMachine) *infrav1.OpenStackMachineSpec {
			return &c.Spec
		},
	},
}

func (r *OpenStackMachine) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachine)

	compare := &OpenStackMachine{}
	return conversion.ConvertAndRestore(
		r, dst, compare,
		Convert_v1alpha7_OpenStackMachine_To_v1alpha8_OpenStackMachine, Convert_v1alpha8_OpenStackMachine_To_v1alpha7_OpenStackMachine,
		v1alpha7OpenStackMachineRestorer, v1alpha8OpenStackMachineRestorer,
	)
}

func (r *OpenStackMachine) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachine)

	compare := &infrav1.OpenStackMachine{}
	return conversion.ConvertAndRestore(
		src, r, compare,
		Convert_v1alpha8_OpenStackMachine_To_v1alpha7_OpenStackMachine, Convert_v1alpha7_OpenStackMachine_To_v1alpha8_OpenStackMachine,
		v1alpha8OpenStackMachineRestorer, v1alpha7OpenStackMachineRestorer,
	)
}

var _ ctrlconversion.Convertible = &OpenStackMachineList{}

func (r *OpenStackMachineList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineList)
	return Convert_v1alpha7_OpenStackMachineList_To_v1alpha8_OpenStackMachineList(r, dst, nil)
}

func (r *OpenStackMachineList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineList)
	return Convert_v1alpha8_OpenStackMachineList_To_v1alpha7_OpenStackMachineList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplate{}

var v1alpha7OpenStackMachineTemplateRestorer = conversion.RestorerFor[*OpenStackMachineTemplate]{
	"spec": conversion.HashedFieldRestorer[*OpenStackMachineTemplate, OpenStackMachineSpec]{
		GetField: func(c *OpenStackMachineTemplate) *OpenStackMachineSpec {
			return &c.Spec.Template.Spec
		},
	},
}

var v1alpha8OpenStackMachineTemplateRestorer = conversion.RestorerFor[*infrav1.OpenStackMachineTemplate]{
	"spec": conversion.HashedFieldRestorer[*infrav1.OpenStackMachineTemplate, infrav1.OpenStackMachineSpec]{
		GetField: func(c *infrav1.OpenStackMachineTemplate) *infrav1.OpenStackMachineSpec {
			return &c.Spec.Template.Spec
		},
	},
}

func (r *OpenStackMachineTemplate) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineTemplate)

	compare := &OpenStackMachineTemplate{}
	return conversion.ConvertAndRestore(
		r, dst, compare,
		Convert_v1alpha7_OpenStackMachineTemplate_To_v1alpha8_OpenStackMachineTemplate, Convert_v1alpha8_OpenStackMachineTemplate_To_v1alpha7_OpenStackMachineTemplate,
		v1alpha7OpenStackMachineTemplateRestorer, v1alpha8OpenStackMachineTemplateRestorer,
	)
}

func (r *OpenStackMachineTemplate) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineTemplate)

	compare := &infrav1.OpenStackMachineTemplate{}
	return conversion.ConvertAndRestore(
		src, r, compare,
		Convert_v1alpha8_OpenStackMachineTemplate_To_v1alpha7_OpenStackMachineTemplate, Convert_v1alpha7_OpenStackMachineTemplate_To_v1alpha8_OpenStackMachineTemplate,
		v1alpha8OpenStackMachineTemplateRestorer, v1alpha7OpenStackMachineTemplateRestorer,
	)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplateList{}

func (r *OpenStackMachineTemplateList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineTemplateList)
	return Convert_v1alpha7_OpenStackMachineTemplateList_To_v1alpha8_OpenStackMachineTemplateList(r, dst, nil)
}

func (r *OpenStackMachineTemplateList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineTemplateList)
	return Convert_v1alpha8_OpenStackMachineTemplateList_To_v1alpha7_OpenStackMachineTemplateList(src, r, nil)
}
