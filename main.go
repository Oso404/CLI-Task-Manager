package main

import (
	"cli-project/models"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
project CLI Task Manager
*build a command line task manager that
1. Adds tasks -> add "task to be added"
2. Lists tasks ->  list (a task consists of an ID and description and completion status)
3. Marks tasks complete -> complete id
4. Deletes tasks -> delete id
5. Saves tasks to json file
*/

func main() {
	file_name := "tasks.json"

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			panic("Insufficient number of arguments")
		}

		taskMap, err := models.Load(file_name)
		if err != nil {
			panic("trouble loading")
		}
		tm := taskMap.Map
		tm[taskMap.NextAvailableID] = models.NewTask(os.Args[2], (taskMap.NextAvailableID))
		fmt.Println("After adding  map is ", tm)
		taskMap.NextAvailableID = taskMap.NextAvailableID + 1
		models.Save(file_name, taskMap)
	case "list":
		taskMap, err := models.Load(file_name)
		if err != nil {
			panic("trouble loading")
		}
		tm := taskMap.Map
		keys := make([]int, 0)
		for id, _ := range tm {
			keys = append(keys, id)
		}
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Println("Task:", tm[k].Description, "|| Modify with ID:", tm[k].ID, "|| Completion Status:", tm[k].Complete)
		}

	case "delete":
		if len(os.Args) < 3 {
			panic("Missing one argument!")
		}
		id_to_delete, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic("Third argument non integer")
		}
		tm, _ := models.Load(file_name)
		m := tm.Map
		delete(m, id_to_delete)
		models.Save(file_name, tm)
	case "complete":
		if len(os.Args) < 3 {
			panic("Missing one argument!")
		}
		id_to_complete, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic("Third argument non integer")
		}
		tm, _ := models.Load(file_name)
		t := tm.Map[id_to_complete]
		t.Complete = true
		tm.Map[id_to_complete] = t
		/*
			earlier i had:
			tm.Map[id_to_complete].Complete = true
			tm.Map[id_to_complete] returns COPY of struct..we need to modify actual
			workaround -> create new Task Struct and assign to
		*/
		models.Save(file_name, tm)
	default:
		fmt.Println("Unknown input")
	}

	// switch command {
	// case "add":
	// 	if arguments_len < 3 {
	// 		panic("Expected one more argument!")
	// 	}
	// 	_, err := os.Stat(file_name) //we are checking to see if tasks.json exists and focusing on error
	// 	if os.IsNotExist(err) {      //if error is of type isNotExist
	// 		//tasks.json doesnt exist
	// 		tasks["1"] = models.NewTask(args[2], tasks_len) // updating tasks map with new task
	// 		file, err := os.Create(file_name)               //creating tasks.json file
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		defer file.Close()                  //will run when main() finishes due to defer ensuring file will close
	// 		json.NewEncoder(file).Encode(tasks) //NewEncoder is place where encoded material will go and we will encode tasks into json
	// 	} else {
	// 		//tasks.json exists
	// 		data, err := os.ReadFile(file_name) //read tasks.json in order to retrieve json data (Readfile returns []byte)
	// 		if err != nil {
	// 			panic("Unable to open tasks.json")
	// 		}
	// 		err = json.Unmarshal(data, &tasks) //Unmarshal converts bytes to Go and returns data to tasks
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		for _, T := range tasks { //iterate tasks map to verify if tasks already exists
	// 			if models.ValueExists(&T, args[2]) {
	// 				panic("Task already exists in map!")
	// 			}
	// 		}
	// 		///retrieve nextAvailableID for more secure update!

	// 		tasks[strconv.Itoa(len(tasks)+1)] = models.NewTask(args[2], len(tasks)+1) //update tasks with new ID and newTask
	// 		f, e := os.OpenFile(file_name, os.O_WRONLY|os.O_TRUNC, 0644)
	// 		if e != nil {
	// 			panic("Error opening tasks.json!")
	// 		}
	// 		json.NewEncoder(f).Encode(tasks)
	// 	}

	// case "list":
	// 	_, err := os.Stat(file_name)
	// 	if os.IsNotExist(err) {
	// 		panic("tasks.json doesnt exist!")
	// 	}
	// 	tasks := make(map[string]models.Task)
	// 	data, err := os.ReadFile(file_name)
	// 	if err != nil {
	// 		panic("Unable to read tasks.json")
	// 	}
	// 	err = json.Unmarshal(data, &tasks)
	// 	if err != nil {
	// 		panic("Unable to read data from tasks.json")
	// 	}
	// 	fmt.Println(tasks)

	// case "delete":
	// 	if arguments_len < 3 {
	// 		fmt.Println("Expected 1 more argument")
	// 		return
	// 	}
	// 	_, err := os.Stat(file_name)
	// 	if os.IsNotExist(err) {
	// 		panic("tasks.json doesn't exist!")
	// 	}
	// 	tasks := make(map[string]models.Task)
	// 	data, err := os.ReadFile(file_name)
	// 	if err != nil {
	// 		panic("Unable to read tasks.json")
	// 	}
	// 	err = json.Unmarshal(data, &tasks)
	// 	if err != nil {
	// 		panic("Unable to read data from tasks.json")
	// 	}
	// 	if _, ok := tasks[args[2]]; ok {
	// 		//remove key
	// 		fmt.Println("Removing from tasks id", args[2])
	// 		delete(tasks, args[2])
	// 	} else {
	// 		panic("Unable to delete....invalid ID")
	// 	}

	// 	/*
	// 		BUG!!!
	// 		as of now can have multiple elements with same ID
	// 		orig: have 2 elements (id 1 & 2)...delete 1...add task....both IDs will contain 2
	// 	*/
	// 	f, e := os.OpenFile(file_name, os.O_WRONLY|os.O_TRUNC, 0644)
	// 	if e != nil {
	// 		panic("Error opening tasks.json!")
	// 	}
	// 	json.NewEncoder(f).Encode(tasks)

	// case "complete":
	// 	fmt.Println("Mark complete command detected")
	// 	if arguments_len < 3 {
	// 		fmt.Println("Expected 1 more argument")
	// 		return
	// 	}
	// 	fmt.Println("Marking as complete->", args[2])

	// default:
	// 	fmt.Println("Command unknown")
	// }

}
