// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/scylladb/scylla-operator/pkg/externalapi/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeThanosRulers implements ThanosRulerInterface
type FakeThanosRulers struct {
	Fake *FakeMonitoringV1
	ns   string
}

var thanosrulersResource = v1.SchemeGroupVersion.WithResource("thanosrulers")

var thanosrulersKind = v1.SchemeGroupVersion.WithKind("ThanosRuler")

// Get takes name of the thanosRuler, and returns the corresponding thanosRuler object, and an error if there is any.
func (c *FakeThanosRulers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ThanosRuler, err error) {
	emptyResult := &v1.ThanosRuler{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(thanosrulersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ThanosRuler), err
}

// List takes label and field selectors, and returns the list of ThanosRulers that match those selectors.
func (c *FakeThanosRulers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ThanosRulerList, err error) {
	emptyResult := &v1.ThanosRulerList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(thanosrulersResource, thanosrulersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ThanosRulerList{ListMeta: obj.(*v1.ThanosRulerList).ListMeta}
	for _, item := range obj.(*v1.ThanosRulerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested thanosRulers.
func (c *FakeThanosRulers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(thanosrulersResource, c.ns, opts))

}

// Create takes the representation of a thanosRuler and creates it.  Returns the server's representation of the thanosRuler, and an error, if there is any.
func (c *FakeThanosRulers) Create(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.CreateOptions) (result *v1.ThanosRuler, err error) {
	emptyResult := &v1.ThanosRuler{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(thanosrulersResource, c.ns, thanosRuler, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ThanosRuler), err
}

// Update takes the representation of a thanosRuler and updates it. Returns the server's representation of the thanosRuler, and an error, if there is any.
func (c *FakeThanosRulers) Update(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.UpdateOptions) (result *v1.ThanosRuler, err error) {
	emptyResult := &v1.ThanosRuler{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(thanosrulersResource, c.ns, thanosRuler, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ThanosRuler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeThanosRulers) UpdateStatus(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.UpdateOptions) (result *v1.ThanosRuler, err error) {
	emptyResult := &v1.ThanosRuler{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(thanosrulersResource, "status", c.ns, thanosRuler, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ThanosRuler), err
}

// Delete takes name of the thanosRuler and deletes it. Returns an error if one occurs.
func (c *FakeThanosRulers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(thanosrulersResource, c.ns, name, opts), &v1.ThanosRuler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeThanosRulers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(thanosrulersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ThanosRulerList{})
	return err
}

// Patch applies the patch and returns the patched thanosRuler.
func (c *FakeThanosRulers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ThanosRuler, err error) {
	emptyResult := &v1.ThanosRuler{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(thanosrulersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ThanosRuler), err
}
