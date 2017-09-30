/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	wardle "k8s.io/sample-apiserver/pkg/apis/wardle"
	internalclientset "k8s.io/sample-apiserver/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/sample-apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/sample-apiserver/pkg/client/listers_generated/wardle/internalversion"
	time "time"
)

// FischerInformer provides access to a shared informer and lister for
// Fischers.
type FischerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.FischerLister
}

type fischerInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewFischerInformer constructs a new informer for Fischer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFischerInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Wardle().Fischers().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Wardle().Fischers().Watch(options)
			},
		},
		&wardle.Fischer{},
		resyncPeriod,
		indexers,
	)
}

func defaultFischerInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFischerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *fischerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&wardle.Fischer{}, defaultFischerInformer)
}

func (f *fischerInformer) Lister() internalversion.FischerLister {
	return internalversion.NewFischerLister(f.Informer().GetIndexer())
}
