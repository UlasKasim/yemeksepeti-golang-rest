package key_value

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

//data is struct field for holding in-memory data
type data struct {
	KeyValueList []*Entity
	fileName     string
}

//Data is object of Data
var Data = &data{
	KeyValueList: []*Entity{},
	fileName:     "",
}

//addValue add entity to KeyValueList
func (d *data) addValue(value *Entity) {
	d.KeyValueList = append(d.KeyValueList, value)
}

//removeValue remove entity from KeyValueList
func (d *data) removeValue(value *Entity) {
	for index, keyValue := range d.KeyValueList {
		//key is unique in list, so we can search by key equals
		if keyValue.Key == value.Key {
			//removing item from slice is just find index of item and merge before and after
			d.KeyValueList = append(d.KeyValueList[:index], d.KeyValueList[index+1:]...)
			break
		}
	}
}

//getValue gets Entity from KeyValueList by key parameter
func (d *data) getValue(key string) *Entity {
	for _, keyValue := range d.KeyValueList {
		//key is unique in list, so we can search by key equals
		if keyValue.Key == key {
			return keyValue
		}
	}
	return nil
}

//flushAll removes all Entities from KeyValueList
func (d *data) flushAll() {
	d.KeyValueList = []*Entity{}
}

//intervalSave saves KeyValueList to file as Json with interval of seconds parameter
func (d *data) intervalSave(seconds int) {
	//Create Ticker for tick of every interval
	t := time.NewTicker(time.Second * time.Duration(seconds))

	//Check every tick
	for range t.C {
		//KeyValueList -> json
		jsonData, err := json.Marshal(Data.KeyValueList)
		if err != nil {
			//Returns if error on marshal
			fmt.Println("Error while intervalSave json.Marshal:", err)
			return
		}
		//saving data to file as string
		d.saveDataToFile(string(jsonData))
	}
}

//saveDataToFile saves data to file in "tmp" folder
//
//While saving data, also deletes other files to prevent swelling.
//Only holds created file
//
//File name contains the number of nanoseconds elapsed since January 1, 1970 UTC
func (d *data) saveDataToFile(jsonData string) {
	//Path for file, includes name of the file
	var fileName = "tmp/" + fmt.Sprint(time.Now().UnixNano()) + "-data.json"
	Data.fileName = fileName

	//Writing file to path, 0777 for read and write access
	err := ioutil.WriteFile(fileName, []byte(jsonData), 0777)
	if err != nil {
		//Returns if error on writing
		fmt.Println("Error while writing data:", err)
		return
	}

	//After writing file, look all files on "tmp" folder
	files, err := ioutil.ReadDir("tmp/")
	if err != nil {
		//Returns if error on Reading directory
		fmt.Println("Error while listing files:", err)
		return
	}

	for _, file := range files {
		//Delete if not equals to written file
		if "tmp/"+file.Name() != Data.fileName {
			err = os.Remove("tmp/" + file.Name())
			if err != nil {
				//Returns if error on Removing file
				fmt.Println("Error while removing files:", err)
			}
		}
	}
}

//extractDataFromFile extracts KeyValueList from file
func (d *data) extractDataFromFile() {
	//look all files on "tmp" folder.
	//Should contain one file, just to be sure we get all files
	files, err := ioutil.ReadDir("tmp/")
	if err != nil {
		//Returns if error on Reading directory
		fmt.Println("Error while listing files:", err)
		return
	}
	if len(files) == 0 {
		//Returns if files is empty
		fmt.Println("Files list empty:")
		return
	}
	//If only one file on pack, extract data
	if len(files) == 1 {
		setKeyValueListFromFile(files[0])
		return
	}

}

//setKeyValueListFromFile sets KeyValueList from parameter file info
func setKeyValueListFromFile(file fs.FileInfo) {
	//Read data from file path
	data, err := ioutil.ReadFile("tmp/" + file.Name())
	if err != nil {
		fmt.Println(err)
	}
	//Creating empty Entity list and load with unmarshal file
	var list []*Entity
	err = json.Unmarshal(data, &list)
	if err != nil {
		fmt.Println(err)
	}
	//Simply set KeyValueList with created Entity list
	Data.KeyValueList = list
}
