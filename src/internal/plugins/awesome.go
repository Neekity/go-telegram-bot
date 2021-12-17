package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"neekity.com/go-telegram-bot/src/internal/config"
	"net/http"
	"sync"
)

type RandomResource struct {
	FileUrl string `json:"file_url"`
	FileExt string `json:"file_ext"`
}

func GetRandomResource(command string, query string) []RandomResource {
	var results []RandomResource
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			url := fmt.Sprintf(config.Conf.ResourceConfigs[command].PreviewUrl, query)
			resp, err := http.Get(url)
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}
			var randomResource RandomResource
			if err := json.Unmarshal(body, &randomResource); err != nil {
				log.Panic(err)
			}
			results = append(results, randomResource)
		}()
	}

	wg.Wait()

	return results
}
