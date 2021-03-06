package fake

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeQuotas implements QuotaInterface
type FakeQuotas struct {
	Fake *FakeGardenV1beta1
	ns   string
}

var quotasResource = schema.GroupVersionResource{Group: "garden.sapcloud.io", Version: "v1beta1", Resource: "quotas"}

var quotasKind = schema.GroupVersionKind{Group: "garden.sapcloud.io", Version: "v1beta1", Kind: "Quota"}

// Get takes name of the quota, and returns the corresponding quota object, and an error if there is any.
func (c *FakeQuotas) Get(name string, options v1.GetOptions) (result *v1beta1.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(quotasResource, c.ns, name), &v1beta1.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Quota), err
}

// List takes label and field selectors, and returns the list of Quotas that match those selectors.
func (c *FakeQuotas) List(opts v1.ListOptions) (result *v1beta1.QuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(quotasResource, quotasKind, c.ns, opts), &v1beta1.QuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.QuotaList{}
	for _, item := range obj.(*v1beta1.QuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested quotas.
func (c *FakeQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(quotasResource, c.ns, opts))

}

// Create takes the representation of a quota and creates it.  Returns the server's representation of the quota, and an error, if there is any.
func (c *FakeQuotas) Create(quota *v1beta1.Quota) (result *v1beta1.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(quotasResource, c.ns, quota), &v1beta1.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Quota), err
}

// Update takes the representation of a quota and updates it. Returns the server's representation of the quota, and an error, if there is any.
func (c *FakeQuotas) Update(quota *v1beta1.Quota) (result *v1beta1.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(quotasResource, c.ns, quota), &v1beta1.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Quota), err
}

// Delete takes name of the quota and deletes it. Returns an error if one occurs.
func (c *FakeQuotas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(quotasResource, c.ns, name), &v1beta1.Quota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(quotasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.QuotaList{})
	return err
}

// Patch applies the patch and returns the patched quota.
func (c *FakeQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Quota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(quotasResource, c.ns, name, data, subresources...), &v1beta1.Quota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Quota), err
}
