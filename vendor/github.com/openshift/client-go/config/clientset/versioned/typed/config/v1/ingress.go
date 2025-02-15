// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	configv1 "github.com/openshift/api/config/v1"
	applyconfigurationsconfigv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// IngressesGetter has a method to return a IngressInterface.
// A group's client should implement this interface.
type IngressesGetter interface {
	Ingresses() IngressInterface
}

// IngressInterface has methods to work with Ingress resources.
type IngressInterface interface {
	Create(ctx context.Context, ingress *configv1.Ingress, opts metav1.CreateOptions) (*configv1.Ingress, error)
	Update(ctx context.Context, ingress *configv1.Ingress, opts metav1.UpdateOptions) (*configv1.Ingress, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, ingress *configv1.Ingress, opts metav1.UpdateOptions) (*configv1.Ingress, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*configv1.Ingress, error)
	List(ctx context.Context, opts metav1.ListOptions) (*configv1.IngressList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *configv1.Ingress, err error)
	Apply(ctx context.Context, ingress *applyconfigurationsconfigv1.IngressApplyConfiguration, opts metav1.ApplyOptions) (result *configv1.Ingress, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, ingress *applyconfigurationsconfigv1.IngressApplyConfiguration, opts metav1.ApplyOptions) (result *configv1.Ingress, err error)
	IngressExpansion
}

// ingresses implements IngressInterface
type ingresses struct {
	*gentype.ClientWithListAndApply[*configv1.Ingress, *configv1.IngressList, *applyconfigurationsconfigv1.IngressApplyConfiguration]
}

// newIngresses returns a Ingresses
func newIngresses(c *ConfigV1Client) *ingresses {
	return &ingresses{
		gentype.NewClientWithListAndApply[*configv1.Ingress, *configv1.IngressList, *applyconfigurationsconfigv1.IngressApplyConfiguration](
			"ingresses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *configv1.Ingress { return &configv1.Ingress{} },
			func() *configv1.IngressList { return &configv1.IngressList{} },
		),
	}
}
