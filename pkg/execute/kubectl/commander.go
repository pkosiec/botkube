package kubectl

import (
	"fmt"
	"github.com/kubeshop/botkube/pkg/events"
)

type Commander struct {
	merger Merger
	guard  CommandGuard
}

type Command struct {
	Name string
	Cmd  string
}

func (c *Commander) GetCommandsForEvent(event events.Event, executorBindings []string) ([]Command, error) {
	enabledKubectls := c.merger.MergeForNamespace(executorBindings, event.Namespace)

	if _, exists := enabledKubectls.AllowedKubectlResource[event.Resource]; !exists {
		// resource not allowed
		return nil, nil
	}

	allowedVerbs := enabledKubectls.AllowedKubectlVerb

	var commands []Command
	for verb := range allowedVerbs {
		res, err := c.guard.GetResourceDetails(verb, event.Resource)
		if err != nil {
			// TODO:
			continue
		}

		var resource string
		if res.SlashSeparatedInCommand {
			resource = fmt.Sprintf("%s/%s", event.Resource, event.Name)
		} else {
			resource = fmt.Sprintf("%s %s", event.Resource, event.Name)
		}

		var namespace string
		if res.Namespaced {
			namespace = event.Namespace
		}

		commands = append(commands, Command{
			Name: verb,
			Cmd:  fmt.Sprintf("%s", verb, resource, namespace),
		})
	}

	// TODO: Filter verbs using CommandGuard

	return nil, nil
}
