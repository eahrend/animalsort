package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const csv = `id,animal,region
1001,elephant,India
1002,elephant,Africa
1003,elephant,India
1004,tiger,Africa
1005,tiger,Africa
1006,tiger,Africa
1007,tiger,Siberia
1008,tiger,Siberia
1009,zebra,Africa
1010,zebra,Africa
1011,zebra,Africa
1012,zebra,Africa
1013,zebra,Africa
1014,zebra,Africa
1015,zebra,Africa
1016,lion,Africa
1017,lion,Africa
1018,lion,Africa
1019,lion,Africa
1020,kangaroo,Australia
1021,kangaroo,Australia
1022,kangaroo,Australia`

func main() {
	err := getFieldCount("animal", csv)
	if err != nil {
		panic(err)
	}
}

type OutputData struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func getFieldCount(fieldName, table string) error {
	strSplit := strings.Split(table, "\n")
	fieldID := -1
	for k, v := range strings.Split(strSplit[0], ",") {
		if strings.ToLower(v) == strings.ToLower(fieldName) {
			fieldID = k
		}
	}
	if fieldID == -1 {
		return fmt.Errorf("failed to find key")
	}
	outputMap := map[string]int{}
	for k, v := range strSplit {
		if k == 0 {
			continue
		}
		animalData := strings.Split(v, ",")
		keyName := animalData[fieldID]
		if _, ok := outputMap[keyName]; ok {
			outputMap[keyName]++
		} else {
			outputMap[keyName] = 1
		}
	}
	outputArray := []OutputData{}
	for k, v := range outputMap {
		newOPData := OutputData{
			Key:   k,
			Value: v,
		}
		outputArray = append(outputArray, newOPData)
	}
	sort.Slice(outputArray, func(i, j int) bool {
		return outputArray[i].Value > outputArray[j].Value
	})
	b, err := json.MarshalIndent(outputArray, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
