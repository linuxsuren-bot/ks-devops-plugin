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
	internalinterfaces "devops.kubesphere.io/plugin/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// S2iBinaries returns a S2iBinaryInformer.
	S2iBinaries() S2iBinaryInformer
	// S2iBuilders returns a S2iBuilderInformer.
	S2iBuilders() S2iBuilderInformer
	// S2iBuilderTemplates returns a S2iBuilderTemplateInformer.
	S2iBuilderTemplates() S2iBuilderTemplateInformer
	// S2iRuns returns a S2iRunInformer.
	S2iRuns() S2iRunInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// S2iBinaries returns a S2iBinaryInformer.
func (v *version) S2iBinaries() S2iBinaryInformer {
	return &s2iBinaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// S2iBuilders returns a S2iBuilderInformer.
func (v *version) S2iBuilders() S2iBuilderInformer {
	return &s2iBuilderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// S2iBuilderTemplates returns a S2iBuilderTemplateInformer.
func (v *version) S2iBuilderTemplates() S2iBuilderTemplateInformer {
	return &s2iBuilderTemplateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// S2iRuns returns a S2iRunInformer.
func (v *version) S2iRuns() S2iRunInformer {
	return &s2iRunInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
