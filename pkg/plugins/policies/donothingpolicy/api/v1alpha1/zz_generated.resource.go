// Generated by tools/resource-gen.
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	"fmt"

	"github.com/kumahq/kuma/pkg/core/resources/model"
)

const (
	DoNothingPolicyType model.ResourceType = "DoNothingPolicy"
)

var _ model.Resource = &DoNothingPolicyResource{}

type DoNothingPolicyResource struct {
	Meta model.ResourceMeta
	Spec *DoNothingPolicy
}

func NewDoNothingPolicyResource() *DoNothingPolicyResource {
	return &DoNothingPolicyResource{
		Spec: &DoNothingPolicy{},
	}
}

func (t *DoNothingPolicyResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *DoNothingPolicyResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *DoNothingPolicyResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *DoNothingPolicyResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*DoNothingPolicy)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &DoNothingPolicy{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *DoNothingPolicyResource) Descriptor() model.ResourceTypeDescriptor {
	return DoNothingPolicyResourceTypeDescriptor
}

func (t *DoNothingPolicyResource) Validate() error {
	if v, ok := interface{}(t).(interface{ validate() error }); !ok {
		return nil
	} else {
		return v.validate()
	}
}

var _ model.ResourceList = &DoNothingPolicyResourceList{}

type DoNothingPolicyResourceList struct {
	Items      []*DoNothingPolicyResource
	Pagination model.Pagination
}

func (l *DoNothingPolicyResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *DoNothingPolicyResourceList) GetItemType() model.ResourceType {
	return DoNothingPolicyType
}

func (l *DoNothingPolicyResourceList) NewItem() model.Resource {
	return NewDoNothingPolicyResource()
}

func (l *DoNothingPolicyResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*DoNothingPolicyResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*DoNothingPolicyResource)(nil), r)
	}
}

func (l *DoNothingPolicyResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

var DoNothingPolicyResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:           DoNothingPolicyType,
	Resource:       NewDoNothingPolicyResource(),
	ResourceList:   &DoNothingPolicyResourceList{},
	Scope:          model.ScopeMesh,
	KDSFlags:       model.FromGlobalToZone,
	WsPath:         "donothingpolicies",
	KumactlArg:     "donothingpolicy",
	KumactlListArg: "donothingpolicies",
	AllowToInspect: true,
	IsPolicy:       true,
	DisplayName:    "Do Nothing Policies",
}
