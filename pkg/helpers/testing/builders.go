package testing

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	clusterapiv1alpha1 "github.com/open-cluster-management/api/cluster/v1alpha1"
)

const (
	clusterSetLabel = "cluster.open-cluster-management.io/clusterset"
	placementLabel  = "cluster.open-cluster-management.io/placement"
)

type placementBuilder struct {
	placement *clusterapiv1alpha1.Placement
}

func NewPlacement(namespace, name string) *placementBuilder {
	return &placementBuilder{
		placement: &clusterapiv1alpha1.Placement{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      name,
			},
		},
	}
}

func (b *placementBuilder) WithUID(uid string) *placementBuilder {
	b.placement.UID = types.UID(uid)
	return b
}

func (b *placementBuilder) WithNOC(noc int32) *placementBuilder {
	b.placement.Spec.NumberOfClusters = &noc
	return b
}

func (b *placementBuilder) WithClusterSets(clusterSets []string) *placementBuilder {
	b.placement.Spec.ClusterSets = clusterSets
	return b
}

func (b *placementBuilder) WithDeletionTimestamp() *placementBuilder {
	now := metav1.Now()
	b.placement.DeletionTimestamp = &now
	return b
}

func (b *placementBuilder) AddPredicate(labelSelector *metav1.LabelSelector, claimSelector *clusterapiv1alpha1.ClusterClaimSelector) *placementBuilder {
	predicate := clusterapiv1alpha1.ClusterPredicate{
		RequiredClusterSelector: clusterapiv1alpha1.ClusterSelector{},
	}

	if labelSelector != nil {
		predicate.RequiredClusterSelector.LabelSelector = *labelSelector
	}

	if claimSelector != nil {
		predicate.RequiredClusterSelector.ClaimSelector = *claimSelector
	}

	if b.placement.Spec.Predicates == nil {
		b.placement.Spec.Predicates = []clusterapiv1alpha1.ClusterPredicate{}
	}
	b.placement.Spec.Predicates = append(b.placement.Spec.Predicates, predicate)

	return b
}

func (b *placementBuilder) Build() *clusterapiv1alpha1.Placement {
	return b.placement
}

type placementDecisionBuilder struct {
	placementDecision *clusterapiv1alpha1.PlacementDecision
}

func NewPlacementDecision(namespace, name string) *placementDecisionBuilder {
	return &placementDecisionBuilder{
		placementDecision: &clusterapiv1alpha1.PlacementDecision{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      name,
			},
		},
	}
}

func (b *placementDecisionBuilder) WithController(uid string) *placementDecisionBuilder {
	controller := true
	b.placementDecision.OwnerReferences = append(b.placementDecision.OwnerReferences, metav1.OwnerReference{
		Controller: &controller,
		UID:        types.UID(uid),
	})
	return b
}

func (b *placementDecisionBuilder) WithPlacementLabel(placementName string) *placementDecisionBuilder {
	if b.placementDecision.Labels == nil {
		b.placementDecision.Labels = map[string]string{}
	}
	b.placementDecision.Labels[placementLabel] = placementName
	return b
}

func (b *placementDecisionBuilder) WithDecisions(decisions []clusterapiv1alpha1.ClusterDecision) *placementDecisionBuilder {
	b.placementDecision.Status.Decisions = decisions
	return b
}

func (b *placementDecisionBuilder) Build() *clusterapiv1alpha1.PlacementDecision {
	return b.placementDecision
}
