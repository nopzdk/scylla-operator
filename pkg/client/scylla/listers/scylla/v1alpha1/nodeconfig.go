// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// NodeConfigLister helps list NodeConfigs.
// All objects returned here must be treated as read-only.
type NodeConfigLister interface {
	// List lists all NodeConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeConfig, err error)
	// Get retrieves the NodeConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeConfig, error)
	NodeConfigListerExpansion
}

// nodeConfigLister implements the NodeConfigLister interface.
type nodeConfigLister struct {
	listers.ResourceIndexer[*v1alpha1.NodeConfig]
}

// NewNodeConfigLister returns a new NodeConfigLister.
func NewNodeConfigLister(indexer cache.Indexer) NodeConfigLister {
	return &nodeConfigLister{listers.New[*v1alpha1.NodeConfig](indexer, v1alpha1.Resource("nodeconfig"))}
}
