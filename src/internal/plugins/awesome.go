package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neekity.com/go-telegram-bot/src/internal/config"
	"net/http"
	"sync"
)

type RandomResource struct {
	FileUrl string `json:"file_url"`
	FileExt string `json:"file_ext"`
}

func GetRankResource(command string, query string) ([]RandomResource, error) {
	var results []RandomResource
	url := fmt.Sprintf(config.Conf.ResourceConfigs[command].PreviewUrl, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func GetRandomResource(command string, query string) ([]RandomResource, error) {
	var results []RandomResource
	var wg sync.WaitGroup
	var erRrr error
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			url := fmt.Sprintf(config.Conf.ResourceConfigs[command].PreviewUrl, query)
			resp, err := http.Get(url)
			if err != nil {
				erRrr = err
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				erRrr = err
				return
			}
			var randomResource RandomResource
			if err := json.Unmarshal(body, &randomResource); err != nil {
				erRrr = err
				return
			}
			results = append(results, randomResource)
		}()
	}

	wg.Wait()

	return results, erRrr
}
