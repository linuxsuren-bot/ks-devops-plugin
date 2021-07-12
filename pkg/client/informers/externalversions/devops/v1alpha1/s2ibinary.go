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

// xCode generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	devopsv1alpha1 "devops.kubesphere.io/plugin/pkg/api/devops/v1alpha1"
	versioned "devops.kubesphere.io/plugin/pkg/client/clientset/versioned"
	internalinterfaces "devops.kubesphere.io/plugin/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "devops.kubesphere.io/plugin/pkg/client/listers/devops/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// S2iBinaryInformer provides access to a shared informer and lister for
// S2iBinaries.
type S2iBinaryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.S2iBinaryLister
}

type s2iBinaryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewS2iBinaryInformer constructs a new informer for S2iBinary type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewS2iBinaryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredS2iBinaryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredS2iBinaryInformer constructs a new informer for S2iBinary type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredS2iBinaryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevopsV1alpha1().S2iBinaries(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevopsV1alpha1().S2iBinaries(namespace).Watch(context.TODO(), options)
			},
		},
		&devopsv1alpha1.S2iBinary{},
		resyncPeriod,
		indexers,
	)
}

func (f *s2iBinaryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredS2iBinaryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *s2iBinaryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&devopsv1alpha1.S2iBinary{}, f.defaultInformer)
}

func (f *s2iBinaryInformer) Lister() v1alpha1.S2iBinaryLister {
	return v1alpha1.NewS2iBinaryLister(f.Informer().GetIndexer())
}
