package main

import (
	"fmt"
	"sort"
)

func mapIter() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	for k, v := range scene {
		fmt.Println(k, v)
	}

	var sceneList []string

	for k := range scene {
		sceneList = append(sceneList, k)
	}

	sort.Strings(sceneList)
	fmt.Println(sceneList)

}

func main() {
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	mapIter()
}
