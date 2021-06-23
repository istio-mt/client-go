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

package v1beta1

import (
	"time"

	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuthorizationPoliciesGetter has a method to return a AuthorizationPolicyInterface.
// A group's client should implement this interface.
type AuthorizationPoliciesGetter interface {
	AuthorizationPolicies(namespace string) AuthorizationPolicyInterface
}

// AuthorizationPolicyInterface has methods to work with AuthorizationPolicy resources.
type AuthorizationPolicyInterface interface {
	Create(*v1beta1.AuthorizationPolicy) (*v1beta1.AuthorizationPolicy, error)
	Update(*v1beta1.AuthorizationPolicy) (*v1beta1.AuthorizationPolicy, error)
	UpdateStatus(*v1beta1.AuthorizationPolicy) (*v1beta1.AuthorizationPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.AuthorizationPolicy, error)
	List(opts v1.ListOptions) (*v1beta1.AuthorizationPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.AuthorizationPolicy, err error)
	AuthorizationPolicyExpansion
}

// authorizationPolicies implements AuthorizationPolicyInterface
type authorizationPolicies struct {
	client rest.Interface
	ns     string
}

// newAuthorizationPolicies returns a AuthorizationPolicies
func newAuthorizationPolicies(c *SecurityV1beta1Client, namespace string) *authorizationPolicies {
	return &authorizationPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the authorizationPolicy, and returns the corresponding authorizationPolicy object, and an error if there is any.
func (c *authorizationPolicies) Get(name string, options v1.GetOptions) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AuthorizationPolicies that match those selectors.
func (c *authorizationPolicies) List(opts v1.ListOptions) (result *v1beta1.AuthorizationPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.AuthorizationPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authorizationPolicies.
func (c *authorizationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a authorizationPolicy and creates it.  Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *authorizationPolicies) Create(authorizationPolicy *v1beta1.AuthorizationPolicy) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Body(authorizationPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a authorizationPolicy and updates it. Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *authorizationPolicies) Update(authorizationPolicy *v1beta1.AuthorizationPolicy) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(authorizationPolicy.Name).
		Body(authorizationPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *authorizationPolicies) UpdateStatus(authorizationPolicy *v1beta1.AuthorizationPolicy) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(authorizationPolicy.Name).
		SubResource("status").
		Body(authorizationPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the authorizationPolicy and deletes it. Returns an error if one occurs.
func (c *authorizationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authorizationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched authorizationPolicy.
func (c *authorizationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("authorizationpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
