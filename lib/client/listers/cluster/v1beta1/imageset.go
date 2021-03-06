/*
Copyright 2019 Gravitational, Inc.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/gravitational/gravity/lib/apis/cluster/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImageSetLister helps list ImageSets.
type ImageSetLister interface {
	// List lists all ImageSets in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ImageSet, err error)
	// ImageSets returns an object that can list and get ImageSets.
	ImageSets(namespace string) ImageSetNamespaceLister
	ImageSetListerExpansion
}

// imageSetLister implements the ImageSetLister interface.
type imageSetLister struct {
	indexer cache.Indexer
}

// NewImageSetLister returns a new ImageSetLister.
func NewImageSetLister(indexer cache.Indexer) ImageSetLister {
	return &imageSetLister{indexer: indexer}
}

// List lists all ImageSets in the indexer.
func (s *imageSetLister) List(selector labels.Selector) (ret []*v1beta1.ImageSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ImageSet))
	})
	return ret, err
}

// ImageSets returns an object that can list and get ImageSets.
func (s *imageSetLister) ImageSets(namespace string) ImageSetNamespaceLister {
	return imageSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImageSetNamespaceLister helps list and get ImageSets.
type ImageSetNamespaceLister interface {
	// List lists all ImageSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.ImageSet, err error)
	// Get retrieves the ImageSet from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.ImageSet, error)
	ImageSetNamespaceListerExpansion
}

// imageSetNamespaceLister implements the ImageSetNamespaceLister
// interface.
type imageSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ImageSets in the indexer for a given namespace.
func (s imageSetNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ImageSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ImageSet))
	})
	return ret, err
}

// Get retrieves the ImageSet from the indexer for a given namespace and name.
func (s imageSetNamespaceLister) Get(name string) (*v1beta1.ImageSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("imageset"), name)
	}
	return obj.(*v1beta1.ImageSet), nil
}
