package key_value

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

//Initialize initialize module with endpoints
//
//Contains Waitgroup for interval save
func Initialize() {
	http.HandleFunc("/set", setKeyHandler)
	http.HandleFunc("/get/", getKeyHandler)
	http.HandleFunc("/flush", flushHandler)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		Data.extractDataFromFile()

		Data.intervalSave(60 * time.Second)

		wg.Done()
	}()

}

//Set setting keyValue Entity to InMemory KeyValueList
//	POST method listener
//	Returns StatusMethodNotAllowed if wrong method
//	Returns StatusBadRequest if invalid json
//	Returns StatusBadRequest if key exists
//	Returns StatusOk with written entity
func setKeyHandler(w http.ResponseWriter, r *http.Request) {
	//Check if method is `POST`
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := "Wrong method has arrived\n"
		fmt.Fprintf(w, err)
		return
	}
	//Read body of request and check json is valid
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid json has arrived\n")
		return
	}
	//Decode Json to Entity
	keyValue := decodeJson(string(b)).BeforeCreate()

	//Getting Entity from KeyValueList with key parameter
	existed := Data.getValue(keyValue.Key)
	if existed != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Key exists")
		return
	}

	//Add entity to KeyValueList
	Data.addValue(keyValue)

	//Send back added Entity
	json.NewEncoder(w).Encode(keyValue)
}

//Get getting keyValue Entity from key parameter
//	GET method listener
//	Returns StatusMethodNotAllowed if wrong method
//	Returns StatusBadRequest if wrong url or key parameter
//	Returns StatusNoContent if key doesnt exits
//	Returns StatusOk with found entity
func getKeyHandler(w http.ResponseWriter, r *http.Request) {
	//Check if method is `GET`
	if r.Method != http.MethodGet {
		err := "Wrong method has arrived"
		fmt.Fprintf(w, err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//looking for key from request url path
	match := keyExp.FindStringSubmatch(r.URL.Path)
	if len(match) == 0 {
		fmt.Fprintf(w, "Wrong url\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Creating a map and storing key with match
	result := make(map[string]string)
	for i, name := range keyExp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	//Looking for key from map and getting entity
	keyValue := Data.getValue(result["key"])
	if keyValue == nil {
		fmt.Fprintf(w, "Wrong key")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	//Send entity
	json.NewEncoder(w).Encode(keyValue)
}

func flushHandler(w http.ResponseWriter, r *http.Request) {
	//Check if method is `DELETE`
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := "Wrong method has arrived"
		fmt.Fprintf(w, err)
		return
	}
	//Flush all data
	Data.flushAll()

	//Send successful message
	w.Write([]byte("Succesfully flushed all data"))
}
