// Copyright 2020 Google LLC
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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeNetworkFirewallPolicyAssociations implements ComputeNetworkFirewallPolicyAssociationInterface
type FakeComputeNetworkFirewallPolicyAssociations struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computenetworkfirewallpolicyassociationsResource = v1beta1.SchemeGroupVersion.WithResource("computenetworkfirewallpolicyassociations")

var computenetworkfirewallpolicyassociationsKind = v1beta1.SchemeGroupVersion.WithKind("ComputeNetworkFirewallPolicyAssociation")

// Get takes name of the computeNetworkFirewallPolicyAssociation, and returns the corresponding computeNetworkFirewallPolicyAssociation object, and an error if there is any.
func (c *FakeComputeNetworkFirewallPolicyAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeNetworkFirewallPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computenetworkfirewallpolicyassociationsResource, c.ns, name), &v1beta1.ComputeNetworkFirewallPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeNetworkFirewallPolicyAssociation), err
}

// List takes label and field selectors, and returns the list of ComputeNetworkFirewallPolicyAssociations that match those selectors.
func (c *FakeComputeNetworkFirewallPolicyAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeNetworkFirewallPolicyAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computenetworkfirewallpolicyassociationsResource, computenetworkfirewallpolicyassociationsKind, c.ns, opts), &v1beta1.ComputeNetworkFirewallPolicyAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeNetworkFirewallPolicyAssociationList{ListMeta: obj.(*v1beta1.ComputeNetworkFirewallPolicyAssociationList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeNetworkFirewallPolicyAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeNetworkFirewallPolicyAssociations.
func (c *FakeComputeNetworkFirewallPolicyAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computenetworkfirewallpolicyassociationsResource, c.ns, opts))

}

// Create takes the representation of a computeNetworkFirewallPolicyAssociation and creates it.  Returns the server's representation of the computeNetworkFirewallPolicyAssociation, and an error, if there is any.
func (c *FakeComputeNetworkFirewallPolicyAssociations) Create(ctx context.Context, computeNetworkFirewallPolicyAssociation *v1beta1.ComputeNetworkFirewallPolicyAssociation, opts v1.CreateOptions) (result *v1beta1.ComputeNetworkFirewallPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computenetworkfirewallpolicyassociationsResource, c.ns, computeNetworkFirewallPolicyAssociation), &v1beta1.ComputeNetworkFirewallPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeNetworkFirewallPolicyAssociation), err
}

// Update takes the representation of a computeNetworkFirewallPolicyAssociation and updates it. Returns the server's representation of the computeNetworkFirewallPolicyAssociation, and an error, if there is any.
func (c *FakeComputeNetworkFirewallPolicyAssociations) Update(ctx context.Context, computeNetworkFirewallPolicyAssociation *v1beta1.ComputeNetworkFirewallPolicyAssociation, opts v1.UpdateOptions) (result *v1beta1.ComputeNetworkFirewallPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computenetworkfirewallpolicyassociationsResource, c.ns, computeNetworkFirewallPolicyAssociation), &v1beta1.ComputeNetworkFirewallPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeNetworkFirewallPolicyAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeNetworkFirewallPolicyAssociations) UpdateStatus(ctx context.Context, computeNetworkFirewallPolicyAssociation *v1beta1.ComputeNetworkFirewallPolicyAssociation, opts v1.UpdateOptions) (*v1beta1.ComputeNetworkFirewallPolicyAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computenetworkfirewallpolicyassociationsResource, "status", c.ns, computeNetworkFirewallPolicyAssociation), &v1beta1.ComputeNetworkFirewallPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeNetworkFirewallPolicyAssociation), err
}

// Delete takes name of the computeNetworkFirewallPolicyAssociation and deletes it. Returns an error if one occurs.
func (c *FakeComputeNetworkFirewallPolicyAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computenetworkfirewallpolicyassociationsResource, c.ns, name, opts), &v1beta1.ComputeNetworkFirewallPolicyAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeNetworkFirewallPolicyAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computenetworkfirewallpolicyassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeNetworkFirewallPolicyAssociationList{})
	return err
}

// Patch applies the patch and returns the patched computeNetworkFirewallPolicyAssociation.
func (c *FakeComputeNetworkFirewallPolicyAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeNetworkFirewallPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computenetworkfirewallpolicyassociationsResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeNetworkFirewallPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeNetworkFirewallPolicyAssociation), err
}
