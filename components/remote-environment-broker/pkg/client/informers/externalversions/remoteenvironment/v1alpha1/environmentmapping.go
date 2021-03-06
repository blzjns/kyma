// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	remoteenvironment_v1alpha1 "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/apis/remoteenvironment/v1alpha1"
	versioned "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/client/listers/remoteenvironment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EnvironmentMappingInformer provides access to a shared informer and lister for
// EnvironmentMappings.
type EnvironmentMappingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EnvironmentMappingLister
}

type environmentMappingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEnvironmentMappingInformer constructs a new informer for EnvironmentMapping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEnvironmentMappingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEnvironmentMappingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEnvironmentMappingInformer constructs a new informer for EnvironmentMapping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEnvironmentMappingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RemoteenvironmentV1alpha1().EnvironmentMappings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RemoteenvironmentV1alpha1().EnvironmentMappings(namespace).Watch(options)
			},
		},
		&remoteenvironment_v1alpha1.EnvironmentMapping{},
		resyncPeriod,
		indexers,
	)
}

func (f *environmentMappingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEnvironmentMappingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *environmentMappingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&remoteenvironment_v1alpha1.EnvironmentMapping{}, f.defaultInformer)
}

func (f *environmentMappingInformer) Lister() v1alpha1.EnvironmentMappingLister {
	return v1alpha1.NewEnvironmentMappingLister(f.Informer().GetIndexer())
}
