package blueprint

import (
	"os"

	"gopkg.in/yaml.v3"
)

const DefaultBlueprintPath = "pomegranate.yml"

func Parse(configPath string) (blueprint Blueprint, err error) {
    var configData []byte
    configData, err = os.ReadFile(configPath)
    if err != nil {
        return
    }

    err = yaml.Unmarshal(configData, &blueprint)
    return
}
