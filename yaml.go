package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func getYamlInfo(yamlContent string) (*KubernetesAPI, error) {

	// Start by removing all lines with templating in to create sane yaml
	//cleanYaml := ""
	//for _, line := range strings.Split(yamlContent, "\n") {
	//	if !strings.Contains(line, "{{") {
	//		cleanYaml += line + "\n"
	//	}
	//}

	var m KubernetesAPI
	err := yaml.Unmarshal([]byte(yamlContent), &m)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal: %v \n---\n%v", err, yamlContent)
	}

	if m.Kind == "" {
		return nil, fmt.Errorf("yaml file with kind missing")
	} else if m.Metadata.Name == "" {
		return nil, fmt.Errorf("yaml file with name missing")
	}

	return &m, nil
}
