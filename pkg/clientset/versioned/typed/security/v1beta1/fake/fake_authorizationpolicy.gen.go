// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthorizationPolicies implements AuthorizationPolicyInterface
type FakeAuthorizationPolicies struct {
	Fake *FakeSecurityV1beta1
	ns   string
}

var authorizationpoliciesResource = schema.GroupVersionResource{Group: "security.istio.io", Version: "v1beta1", Resource: "authorizationpolicies"}

var authorizationpoliciesKind = schema.GroupVersionKind{Group: "security.istio.io", Version: "v1beta1", Kind: "AuthorizationPolicy"}

// Get takes name of the authorizationPolicy, and returns the corresponding authorizationPolicy object, and an error if there is any.
func (c *FakeAuthorizationPolicies) Get(name string, options v1.GetOptions) (result *v1beta1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(authorizationpoliciesResource, c.ns, name), &v1beta1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AuthorizationPolicy), err
}

// List takes label and field selectors, and returns the list of AuthorizationPolicies that match those selectors.
func (c *FakeAuthorizationPolicies) List(opts v1.ListOptions) (result *v1beta1.AuthorizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(authorizationpoliciesResource, authorizationpoliciesKind, c.ns, opts), &v1beta1.AuthorizationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AuthorizationPolicyList{ListMeta: obj.(*v1beta1.AuthorizationPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.AuthorizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authorizationPolicies.
func (c *FakeAuthorizationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(authorizationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a authorizationPolicy and creates it.  Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *FakeAuthorizationPolicies) Create(authorizationPolicy *v1beta1.AuthorizationPolicy) (result *v1beta1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(authorizationpoliciesResource, c.ns, authorizationPolicy), &v1beta1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AuthorizationPolicy), err
}

// Update takes the representation of a authorizationPolicy and updates it. Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *FakeAuthorizationPolicies) Update(authorizationPolicy *v1beta1.AuthorizationPolicy) (result *v1beta1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(authorizationpoliciesResource, c.ns, authorizationPolicy), &v1beta1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AuthorizationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAuthorizationPolicies) UpdateStatus(authorizationPolicy *v1beta1.AuthorizationPolicy) (*v1beta1.AuthorizationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(authorizationpoliciesResource, "status", c.ns, authorizationPolicy), &v1beta1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AuthorizationPolicy), err
}

// Delete takes name of the authorizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAuthorizationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(authorizationpoliciesResource, c.ns, name), &v1beta1.AuthorizationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthorizationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(authorizationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.AuthorizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched authorizationPolicy.
func (c *FakeAuthorizationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(authorizationpoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AuthorizationPolicy), err
}
