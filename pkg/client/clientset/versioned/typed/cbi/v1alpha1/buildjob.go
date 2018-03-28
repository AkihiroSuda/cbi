/*
Copyright The CBI Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/containerbuilding/cbi/pkg/apis/cbi/v1alpha1"
	scheme "github.com/containerbuilding/cbi/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BuildJobsGetter has a method to return a BuildJobInterface.
// A group's client should implement this interface.
type BuildJobsGetter interface {
	BuildJobs(namespace string) BuildJobInterface
}

// BuildJobInterface has methods to work with BuildJob resources.
type BuildJobInterface interface {
	Create(*v1alpha1.BuildJob) (*v1alpha1.BuildJob, error)
	Update(*v1alpha1.BuildJob) (*v1alpha1.BuildJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BuildJob, error)
	List(opts v1.ListOptions) (*v1alpha1.BuildJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BuildJob, err error)
	BuildJobExpansion
}

// buildJobs implements BuildJobInterface
type buildJobs struct {
	client rest.Interface
	ns     string
}

// newBuildJobs returns a BuildJobs
func newBuildJobs(c *CbiV1alpha1Client, namespace string) *buildJobs {
	return &buildJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the buildJob, and returns the corresponding buildJob object, and an error if there is any.
func (c *buildJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.BuildJob, err error) {
	result = &v1alpha1.BuildJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BuildJobs that match those selectors.
func (c *buildJobs) List(opts v1.ListOptions) (result *v1alpha1.BuildJobList, err error) {
	result = &v1alpha1.BuildJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested buildJobs.
func (c *buildJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("buildjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a buildJob and creates it.  Returns the server's representation of the buildJob, and an error, if there is any.
func (c *buildJobs) Create(buildJob *v1alpha1.BuildJob) (result *v1alpha1.BuildJob, err error) {
	result = &v1alpha1.BuildJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("buildjobs").
		Body(buildJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a buildJob and updates it. Returns the server's representation of the buildJob, and an error, if there is any.
func (c *buildJobs) Update(buildJob *v1alpha1.BuildJob) (result *v1alpha1.BuildJob, err error) {
	result = &v1alpha1.BuildJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("buildjobs").
		Name(buildJob.Name).
		Body(buildJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the buildJob and deletes it. Returns an error if one occurs.
func (c *buildJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *buildJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched buildJob.
func (c *buildJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BuildJob, err error) {
	result = &v1alpha1.BuildJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("buildjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}