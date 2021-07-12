/*
Copyright 2020 The KubeSphere Authors.

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

// xCode generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "devops.kubesphere.io/plugin/pkg/api/devops/v1alpha1"
	scheme "devops.kubesphere.io/plugin/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// S2iRunsGetter has a method to return a S2iRunInterface.
// A group's client should implement this interface.
type S2iRunsGetter interface {
	S2iRuns(namespace string) S2iRunInterface
}

// S2iRunInterface has methods to work with S2iRun resources.
type S2iRunInterface interface {
	Create(ctx context.Context, s2iRun *v1alpha1.S2iRun, opts v1.CreateOptions) (*v1alpha1.S2iRun, error)
	Update(ctx context.Context, s2iRun *v1alpha1.S2iRun, opts v1.UpdateOptions) (*v1alpha1.S2iRun, error)
	UpdateStatus(ctx context.Context, s2iRun *v1alpha1.S2iRun, opts v1.UpdateOptions) (*v1alpha1.S2iRun, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.S2iRun, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.S2iRunList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S2iRun, err error)
	S2iRunExpansion
}

// s2iRuns implements S2iRunInterface
type s2iRuns struct {
	client rest.Interface
	ns     string
}

// newS2iRuns returns a S2iRuns
func newS2iRuns(c *DevopsV1alpha1Client, namespace string) *s2iRuns {
	return &s2iRuns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s2iRun, and returns the corresponding s2iRun object, and an error if there is any.
func (c *s2iRuns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S2iRuns that match those selectors.
func (c *s2iRuns) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.S2iRunList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S2iRunList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2iruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s2iRuns.
func (c *s2iRuns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s2iruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a s2iRun and creates it.  Returns the server's representation of the s2iRun, and an error, if there is any.
func (c *s2iRuns) Create(ctx context.Context, s2iRun *v1alpha1.S2iRun, opts v1.CreateOptions) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s2iruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s2iRun).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a s2iRun and updates it. Returns the server's representation of the s2iRun, and an error, if there is any.
func (c *s2iRuns) Update(ctx context.Context, s2iRun *v1alpha1.S2iRun, opts v1.UpdateOptions) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(s2iRun.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s2iRun).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *s2iRuns) UpdateStatus(ctx context.Context, s2iRun *v1alpha1.S2iRun, opts v1.UpdateOptions) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(s2iRun.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(s2iRun).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the s2iRun and deletes it. Returns an error if one occurs.
func (c *s2iRuns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2iruns").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s2iRuns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2iruns").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched s2iRun.
func (c *s2iRuns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S2iRun, err error) {
	result = &v1alpha1.S2iRun{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s2iruns").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
