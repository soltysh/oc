// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/openshift/api/machineconfiguration/v1alpha1"
	machineconfigurationv1alpha1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1alpha1"
	scheme "github.com/openshift/client-go/machineconfiguration/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PinnedImageSetsGetter has a method to return a PinnedImageSetInterface.
// A group's client should implement this interface.
type PinnedImageSetsGetter interface {
	PinnedImageSets() PinnedImageSetInterface
}

// PinnedImageSetInterface has methods to work with PinnedImageSet resources.
type PinnedImageSetInterface interface {
	Create(ctx context.Context, pinnedImageSet *v1alpha1.PinnedImageSet, opts v1.CreateOptions) (*v1alpha1.PinnedImageSet, error)
	Update(ctx context.Context, pinnedImageSet *v1alpha1.PinnedImageSet, opts v1.UpdateOptions) (*v1alpha1.PinnedImageSet, error)
	UpdateStatus(ctx context.Context, pinnedImageSet *v1alpha1.PinnedImageSet, opts v1.UpdateOptions) (*v1alpha1.PinnedImageSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PinnedImageSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PinnedImageSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PinnedImageSet, err error)
	Apply(ctx context.Context, pinnedImageSet *machineconfigurationv1alpha1.PinnedImageSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PinnedImageSet, err error)
	ApplyStatus(ctx context.Context, pinnedImageSet *machineconfigurationv1alpha1.PinnedImageSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PinnedImageSet, err error)
	PinnedImageSetExpansion
}

// pinnedImageSets implements PinnedImageSetInterface
type pinnedImageSets struct {
	client rest.Interface
}

// newPinnedImageSets returns a PinnedImageSets
func newPinnedImageSets(c *MachineconfigurationV1alpha1Client) *pinnedImageSets {
	return &pinnedImageSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the pinnedImageSet, and returns the corresponding pinnedImageSet object, and an error if there is any.
func (c *pinnedImageSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PinnedImageSet, err error) {
	result = &v1alpha1.PinnedImageSet{}
	err = c.client.Get().
		Resource("pinnedimagesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PinnedImageSets that match those selectors.
func (c *pinnedImageSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PinnedImageSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PinnedImageSetList{}
	err = c.client.Get().
		Resource("pinnedimagesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pinnedImageSets.
func (c *pinnedImageSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("pinnedimagesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a pinnedImageSet and creates it.  Returns the server's representation of the pinnedImageSet, and an error, if there is any.
func (c *pinnedImageSets) Create(ctx context.Context, pinnedImageSet *v1alpha1.PinnedImageSet, opts v1.CreateOptions) (result *v1alpha1.PinnedImageSet, err error) {
	result = &v1alpha1.PinnedImageSet{}
	err = c.client.Post().
		Resource("pinnedimagesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pinnedImageSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a pinnedImageSet and updates it. Returns the server's representation of the pinnedImageSet, and an error, if there is any.
func (c *pinnedImageSets) Update(ctx context.Context, pinnedImageSet *v1alpha1.PinnedImageSet, opts v1.UpdateOptions) (result *v1alpha1.PinnedImageSet, err error) {
	result = &v1alpha1.PinnedImageSet{}
	err = c.client.Put().
		Resource("pinnedimagesets").
		Name(pinnedImageSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pinnedImageSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *pinnedImageSets) UpdateStatus(ctx context.Context, pinnedImageSet *v1alpha1.PinnedImageSet, opts v1.UpdateOptions) (result *v1alpha1.PinnedImageSet, err error) {
	result = &v1alpha1.PinnedImageSet{}
	err = c.client.Put().
		Resource("pinnedimagesets").
		Name(pinnedImageSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pinnedImageSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the pinnedImageSet and deletes it. Returns an error if one occurs.
func (c *pinnedImageSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("pinnedimagesets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pinnedImageSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("pinnedimagesets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched pinnedImageSet.
func (c *pinnedImageSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PinnedImageSet, err error) {
	result = &v1alpha1.PinnedImageSet{}
	err = c.client.Patch(pt).
		Resource("pinnedimagesets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied pinnedImageSet.
func (c *pinnedImageSets) Apply(ctx context.Context, pinnedImageSet *machineconfigurationv1alpha1.PinnedImageSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PinnedImageSet, err error) {
	if pinnedImageSet == nil {
		return nil, fmt.Errorf("pinnedImageSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(pinnedImageSet)
	if err != nil {
		return nil, err
	}
	name := pinnedImageSet.Name
	if name == nil {
		return nil, fmt.Errorf("pinnedImageSet.Name must be provided to Apply")
	}
	result = &v1alpha1.PinnedImageSet{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("pinnedimagesets").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *pinnedImageSets) ApplyStatus(ctx context.Context, pinnedImageSet *machineconfigurationv1alpha1.PinnedImageSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PinnedImageSet, err error) {
	if pinnedImageSet == nil {
		return nil, fmt.Errorf("pinnedImageSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(pinnedImageSet)
	if err != nil {
		return nil, err
	}

	name := pinnedImageSet.Name
	if name == nil {
		return nil, fmt.Errorf("pinnedImageSet.Name must be provided to Apply")
	}

	result = &v1alpha1.PinnedImageSet{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("pinnedimagesets").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
