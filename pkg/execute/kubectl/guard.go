package kubectl

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/discovery"
)

type Resource struct {
	// Plural
	Name       string
	Namespaced bool

	SlashSeparatedInCommand bool
}

type CommandGuard struct {
	log          logrus.FieldLogger
	discoveryCli discovery.DiscoveryInterface
}

func NewCommandGuard(log logrus.FieldLogger, discoveryCli discovery.DiscoveryInterface) *CommandGuard {
	return &CommandGuard{log: log, discoveryCli: discoveryCli}
}

func (g *CommandGuard) GetAllowedResourcesForVerb(selectedVerb string, allConfiguredResources []string) ([]Resource, error) {
	//resList, err := g.discoveryCli.ServerPreferredResources()

	return nil, nil
}

func (g *CommandGuard) GetResourceDetails(selectedVerb, resourceType string) (Resource, error) {
	return Resource{}, nil
}

//func (g *CommandGuard) GetCommandsForResource(resourceType string, allConfiguredVerbs []string) ([], error) {
//	return nil, nil
//}
