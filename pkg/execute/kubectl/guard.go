package kubectl

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
)

type Resource struct {
	Name       string
	Namespaced bool
}

type CommandGuard struct {
	k8sCli       kubernetes.Interface
	discoveryCli discovery.DiscoveryInterface
}

func (g *CommandGuard) GetAllowedResourcesForVerb(selectedVerb string, allConfiguredResources []string) ([]Resource, error) {
	resList, err := g.discoveryCli.ServerPreferredResources()

	res
}
