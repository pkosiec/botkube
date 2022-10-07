package kubectl

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/discovery"
	"k8s.io/utils/strings/slices"
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

var ErrVerbNotFound = errors.New("verb not found")

func (g *CommandGuard) GetResourceDetails(selectedVerb, resourceType string) (Resource, error) {
	resList, err := g.discoveryCli.ServerPreferredResources()
	if err != nil {
		return Resource{}, fmt.Errorf("while getting server resources: %w", err)
	}

	for _, item := range resList {
		for _, res := range item.APIResources {
			if res.Name != resourceType {
				continue
			}

			// TODO:
			// 	- remove watch
			//  - add describe when get is available
			//  - handle logs
			if !slices.Contains(res.Verbs, selectedVerb) {
				return Resource{}, ErrVerbNotFound
			}

			return Resource{
				Name:                    res.Name,
				Namespaced:              res.Namespaced,
				SlashSeparatedInCommand: false, //FIXME
			}, nil
		}
		fmt.Printf("%+v\n\n", item)
	}

	return Resource{}, nil
}
