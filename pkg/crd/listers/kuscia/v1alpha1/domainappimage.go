// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainAppImageLister helps list DomainAppImages.
// All objects returned here must be treated as read-only.
type DomainAppImageLister interface {
	// List lists all DomainAppImages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainAppImage, err error)
	// DomainAppImages returns an object that can list and get DomainAppImages.
	DomainAppImages(namespace string) DomainAppImageNamespaceLister
	DomainAppImageListerExpansion
}

// domainAppImageLister implements the DomainAppImageLister interface.
type domainAppImageLister struct {
	indexer cache.Indexer
}

// NewDomainAppImageLister returns a new DomainAppImageLister.
func NewDomainAppImageLister(indexer cache.Indexer) DomainAppImageLister {
	return &domainAppImageLister{indexer: indexer}
}

// List lists all DomainAppImages in the indexer.
func (s *domainAppImageLister) List(selector labels.Selector) (ret []*v1alpha1.DomainAppImage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainAppImage))
	})
	return ret, err
}

// DomainAppImages returns an object that can list and get DomainAppImages.
func (s *domainAppImageLister) DomainAppImages(namespace string) DomainAppImageNamespaceLister {
	return domainAppImageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainAppImageNamespaceLister helps list and get DomainAppImages.
// All objects returned here must be treated as read-only.
type DomainAppImageNamespaceLister interface {
	// List lists all DomainAppImages in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainAppImage, err error)
	// Get retrieves the DomainAppImage from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainAppImage, error)
	DomainAppImageNamespaceListerExpansion
}

// domainAppImageNamespaceLister implements the DomainAppImageNamespaceLister
// interface.
type domainAppImageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainAppImages in the indexer for a given namespace.
func (s domainAppImageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainAppImage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainAppImage))
	})
	return ret, err
}

// Get retrieves the DomainAppImage from the indexer for a given namespace and name.
func (s domainAppImageNamespaceLister) Get(name string) (*v1alpha1.DomainAppImage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainappimage"), name)
	}
	return obj.(*v1alpha1.DomainAppImage), nil
}