package kubectl

import (
	"fmt"
	"github.com/kubeshop/botkube/pkg/events"
	"strings"
)

type Commander struct {
	merger *Merger
	guard  *CommandGuard
}

func NewCommander(merger *Merger, guard *CommandGuard) *Commander {
	return &Commander{merger: merger, guard: guard}
}

type Command struct {
	Name string
	Cmd  string
}

func (c *Commander) GetCommandsForEvent(event events.Event, executorBindings []string) ([]Command, error) {
	enabledKubectls := c.merger.MergeForNamespace(executorBindings, event.Namespace)

	resourceTypeParts := strings.Split(event.Resource, "/")
	resourceName := resourceTypeParts[len(resourceTypeParts)-1]

	if _, exists := enabledKubectls.AllowedKubectlResource[resourceName]; !exists {
		// resource not allowed
		return nil, nil
	}

	allowedVerbs := enabledKubectls.AllowedKubectlVerb

	var commands []Command
	for verb := range allowedVerbs {
		res, err := c.guard.GetResourceDetails(verb, resourceName)
		if err != nil {
			// TODO:
			continue
		}

		var resourceSubstr string
		if res.SlashSeparatedInCommand {
			resourceSubstr = fmt.Sprintf("%s/%s", resourceName, event.Name)
		} else {
			resourceSubstr = fmt.Sprintf("%s %s", resourceName, event.Name)
		}

		var namespaceSubstr string
		if res.Namespaced {
			namespaceSubstr = fmt.Sprintf(" --namespace %s", event.Namespace)
		}

		commands = append(commands, Command{
			Name: verb,
			Cmd:  fmt.Sprintf("%s %s%s", verb, resourceSubstr, namespaceSubstr),
		})
	}

	return commands, nil
}
