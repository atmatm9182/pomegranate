package blueprint

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const DefaultBlueprintPath = "pomegranate.yml"

func Parse(configPath string) (blueprint Blueprint, err error) {
	var configFile *os.File
	configFile, err = os.Open(configPath)
	if err != nil {
		return
	}
	defer configFile.Close()

	var stat os.FileInfo
	stat, err = configFile.Stat()
	if err != nil {
		return
	}
	
    var configData []byte
	if stat.IsDir() {
		configPath = path.Join(configPath, DefaultBlueprintPath)
		configData, err = os.ReadFile(configPath)
	} else {
		configData = make([]byte, stat.Size())
		_, err = configFile.Read(configData)
	}

    if err != nil {
        return
    }

    err = yaml.Unmarshal(configData, &blueprint)
	blueprint.absolutePath = path.Dir(configPath)
    return
}
