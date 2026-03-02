package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskMap struct {
	Map             map[int]Task
	NextAvailableID int
}

/*
Load() returns address to newly created TaskMap
Process
1. check to see if tasks.json exists
2. if doesnt exist -> create it and return empty taskmap and error
2. if does exist -> decode from tasks.json and store in tm
3. return tm
*/
func Load(filename string) (*TaskMap, error) {
	tm := &TaskMap{
		Map:             make(map[int]Task),
		NextAvailableID: 1,
	}
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return tm, err
		}
		file.Close()
	}
	if err != nil {
		fmt.Println("Error opening existing file!")
		return tm, nil
	} else {
		//have to decode here!!!
		file, err = os.Open(filename)
		if err != nil {
			return tm, nil
		}
		if err := json.NewDecoder(file).Decode(tm); err != nil {
			fmt.Println("Error decoding json to map")
			return tm, nil
		}
		file.Close()
	}
	return tm, nil //address of newly created Taskmap
}

/*
1. simply write to taskmap pointer encoded data
*/

func Save(filename string, tm *TaskMap) (bool, error) {
	//no need to check if file exists because Load() ran before and check is done there
	//decode map
	file, _ := os.Create(filename)
	if err := json.NewEncoder(file).Encode(tm); err != nil {
		return false, err
	}
	return true, nil
}
