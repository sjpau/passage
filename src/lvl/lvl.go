package lvl

import (
	"embed"
	"encoding/json"
	"io/ioutil"

	"github.com/sjpau/passage/src/help"
)

type Level struct {
	Name   string  `json:"name"`
	Layout [][]int `json:"layout"`
}

func LoadFromJSON(fs *embed.FS, filePath string) (*Level, error) {
	file, e := fs.Open(filePath)
	help.Check(e)
	defer file.Close()

	dataBytes, e := ioutil.ReadAll(file)
	help.Check(e)

	var level Level
	e = json.Unmarshal(dataBytes, &level)
	help.Check(e)

	return &level, nil
}
