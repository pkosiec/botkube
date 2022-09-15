package recommendation

import (
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/events"
	"github.com/kubeshop/botkube/pkg/ptr"
)

const (
	podsResourceName    = "v1/pods"
	ingressResourceName = "networking.k8s.io/v1/ingresses"
)

// ResourceEventsForConfig returns the resource event map for a given source recommendations config.
func ResourceEventsForConfig(recCfg config.Recommendations) map[string]config.EventType {
	resNames := make(map[string]config.EventType)

	if ptr.IsTrue(recCfg.Ingress.TLSSecretValid) || ptr.IsTrue(recCfg.Ingress.BackendServiceValid) {
		resNames[ingressResourceName] = config.CreateEvent
	}

	if ptr.IsTrue(recCfg.Pod.NoLatestImageTag) || ptr.IsTrue(recCfg.Pod.LabelsSet) {
		resNames[podsResourceName] = config.CreateEvent
	}

	return resNames
}

// FilterSourcesForEvent checks if user doesn't listen to events for a given resource, apart from enabled recommendations and filters
// the source bindings.
func FilterSourcesForEvent(recCfg config.Recommendations, sources map[string]config.Sources, sourceBindings []string, event events.Event) []string {
	if event.HasRecommendationsOrWarnings() {
		// shouldn't be skipped
		return sourceBindings
	}

	res := ResourceEventsForConfig(recCfg)
	recommEventType, ok := res[event.Resource]
	if !ok {
		// this event doesn't relate to recommendations, finish early
		return sourceBindings
	}

	if event.Type != recommEventType {
		// this event doesn't relate to recommendations, finish early
		return sourceBindings
	}

	// Resource + event type matches the ones configured from recommendation.
	// Check if user listens to this event.
	var sourceBindingsToNotify []string
	for _, key := range sourceBindings {
		source, exists := sources[key]
		if !exists {
			continue
		}

		// sources are appended, so we need to check the first source that has a given resource with event
		if source.Kubernetes.Resources.IsAllowed(event.Resource, event.Namespace, event.Type) {
			sourceBindingsToNotify = append(sourceBindingsToNotify, key)
		}
	}

	// The event is related to recommendations informers. No recommendations there, so it should be skipped.
	return []string{}
}
