package provider

import (
	"sort"
	"sync"

	"github.com/fatih/color"
)

// ModelsMap Maps provider name to model
var ModelsMap = map[string]Model{
	Types.AWS: NewAwsModel(),
}

// Types provider types
var Types = &providerRegistry{
	AWS: "AWS",
}

type providerRegistry struct {
	AWS string
}

// Provider Hosted application provider
type Provider struct {
	Name      string      `json:"name"`
	Resources []*Resource `json:"resources"`
	*ResourceCommands
}

// Resource Provider resource (i.e. cloudformation)
type Resource struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Body     string `json:"body,omitempty"`
	Priority int    `json:"priority,omitempty"`
	*ResourceCommands
}

// ResourceCommands Custom resource commands
type ResourceCommands struct {
	PreProvision  []string `json:"preProvision"`
	PostProvision []string `json:"postProvision"`
}

// Model provider model interface
type Model interface {
	Provision(*Resource)
	Destroy(*Resource)
	CheckStatus()
}

// Provision initiates resource provisioning of provider map
func Provision(provisionMap map[string]map[int][]*Resource) {
	providersWg := &sync.WaitGroup{}
	providersWg.Add(len(provisionMap))
	defer providersWg.Wait()

	for provider, resourcesMap := range provisionMap {
		color.Green("Provision: %s", provider)
		go provisionProvider(provider, resourcesMap, providersWg)
	}
}

func provisionProvider(p string, rm map[int][]*Resource, wg *sync.WaitGroup) {
	defer wg.Done()
	model := ModelsMap[p]
	keys := make([]int, len(rm))
	for k := range rm {
		// if k != 0 {
		keys[k] = k
		// }

	}
	sort.Ints(keys)
	// keys = append(keys, 0)
	for key := range keys {
		color.Green("Batch: %b", key)
		resourceBatch := rm[key]
		resourceBatchWaitGroup := &sync.WaitGroup{}
		resourceBatchWaitGroup.Add(len(resourceBatch))
		defer resourceBatchWaitGroup.Wait()
		for _, resources := range resourceBatch {
			go provisionResource(resources, model, resourceBatchWaitGroup)
		}
	}
}

func provisionResource(r *Resource, m Model, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Provision(r)
}
