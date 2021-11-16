package key_value

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type data struct {
	KeyValueList []*Entity
	fileName     string
}

var Data = &data{
	KeyValueList: []*Entity{},
	fileName:     "",
}

func (d *data) addValue(value *Entity) {
	d.KeyValueList = append(d.KeyValueList, value)
}

func (d *data) removeValue(value *Entity) {
	for index, keyValue := range d.KeyValueList {
		if keyValue.Key == value.Key {
			d.KeyValueList = append(d.KeyValueList[:index], d.KeyValueList[index+1:]...)
			break
		}
	}
}

func (d *data) getValue(key string) *Entity {
	for _, keyValue := range d.KeyValueList {
		if keyValue.Key == key {
			return keyValue
		}
	}
	return nil
}

func (d *data) flushAll() {
	d.KeyValueList = []*Entity{}
}

func (d *data) intervalSave(seconds int) {
	t := time.NewTicker(time.Second * time.Duration(seconds))
	for range t.C {
		jsonData, err := json.Marshal(Data.KeyValueList)
		if err != nil {
			fmt.Println("Error while intervalSave json.Marshal:", err)
			return
		}
		d.saveDataToFile(string(jsonData))
	}
}

func (d *data) saveDataToFile(jsonData string) {
	var fileName = "tmp/" + fmt.Sprint(time.Now().UnixNano()) + "-data.json"
	Data.fileName = fileName

	err := ioutil.WriteFile(fileName, []byte(jsonData), 0777)
	if err != nil {
		fmt.Println("Error while writing data:", err)
	}

	files, err := ioutil.ReadDir("tmp/")
	if err != nil {
		fmt.Println("Error while listing files:", err)
	}

	for _, file := range files {
		if "tmp/"+file.Name() != Data.fileName {
			err = os.Remove("tmp/" + file.Name())
			if err != nil {
				fmt.Println("Error while removing files:", err)
			}
		}
	}
}

func (d *data) extractDataFromFile() {
	files, err := ioutil.ReadDir("tmp/")
	if err != nil {
		fmt.Println("Error while listing files:", err)
	}
	for _, file := range files {
		data, err := ioutil.ReadFile("tmp/" + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		var list []*Entity
		err = json.Unmarshal(data, &list)
		if err != nil {
			fmt.Println(err)
		}
		Data.KeyValueList = list
	}

}

func decodeJson(d string) *Entity {
	var keyValue Entity
	json.Unmarshal([]byte(d), &keyValue)
	return &keyValue
}

func encodeJson(e Entity) string {
	data, err := json.Marshal(e)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(data[:])
}
