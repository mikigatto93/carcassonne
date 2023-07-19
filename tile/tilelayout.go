package tile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Terrain uint8

const (
	Grass Terrain = 1 // 0 is used as absence of a value
	City
	Street
	Water
	Monastery
)

type TileLayout [][]Terrain
type tileFileStructureElem = struct {
	Quantity uint8      `json:"quantity"`
	Layout   TileLayout `json:"layout"`
}
type tileFileStructure map[string]tileFileStructureElem

var TileData tileFileStructure = make(tileFileStructure)

func parseTileLayoutFile(path string, jsonBody *tileFileStructure) {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	jsonErr := json.Unmarshal(byteValue, jsonBody)
	if jsonErr != nil {
		fmt.Printf("%v", jsonErr)
	}
}

func LoadAllTilesets() {
	parseTileLayoutFile("River1.json", &TileData)
	//fmt.Println(TileLayouts)
}
