package key_value

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func Initialize() {
	http.HandleFunc("/set", setKeyHandler)
	http.HandleFunc("/get/", getKeyHandler)
	http.HandleFunc("/flush", flushHandler)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		Data.extractDataFromFile()

		Data.intervalSave(10)

		wg.Done()
	}()

}

func setKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := "Wrong method has arrived\n"
		log.Print(err, r.Method)
		fmt.Fprintf(w, err)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		fmt.Fprintf(w, "Invalid json has arrived\n")
		return
	}
	keyValue := decodeJson(string(b))

	existed := Data.getValue(keyValue.Key)
	if existed != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Key exists")
		return
	}

	Data.addValue(keyValue.BeforeCreate())
	json.NewEncoder(w).Encode(Data.KeyValueList)
}

func getKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := "Wrong method has arrived"
		log.Print(err, r.Method)
		fmt.Fprintf(w, err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	match := keyExp.FindStringSubmatch(r.URL.Path)
	if len(match) == 0 {
		fmt.Fprintf(w, "Wrong url\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := make(map[string]string)
	for i, name := range keyExp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	keyValue := Data.getValue(result["key"])
	if keyValue == nil {
		fmt.Fprintf(w, "Wrong key")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(keyValue)
}

func flushHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		err := "Wrong method has arrived"
		log.Print(err, r.Method)
		fmt.Fprintf(w, err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	Data.flushAll()
}
