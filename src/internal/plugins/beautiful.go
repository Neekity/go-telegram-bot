package plugins

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"neekity.com/go-telegram-bot/src/internal/config"
	"net/http"
	"strconv"
	"time"
)

type RandomPic struct {
	Preview string `json:"preview"`
	Id      int    `json:"id"`
}

type CountPic struct {
	Count int `json:"count"`
}

func GetRandomPic(commamd string, count string) []string {
	resp, err := http.Get(config.Conf.PhotoConfigs[commamd].PreviewUrl + count)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	var randomPics []RandomPic
	if err := json.Unmarshal(body, &randomPics); err != nil {
		log.Panic(err)
	}
	var result []string
	for _, randomPic := range randomPics {
		result = append(result, config.Conf.PhotoConfigs[commamd].Url+randomPic.Preview)
	}
	return result
}

func GetRandomPic2(commamd string, count string) []string {
	resp, err := http.Get(config.Conf.PhotoConfigs[commamd].PreviewUrl + "count")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	var picAllCount []CountPic
	if err := json.Unmarshal(body, &picAllCount); err != nil {
		log.Panic(err)
	}
	reCount, _ := strconv.Atoi(count)
	randomPicIds := generateRandomNumber(1, picAllCount[0].Count, reCount)
	log.Println(randomPicIds)
	var result []string
	for _, picId := range randomPicIds {
		resp, err = http.Get(config.Conf.PhotoConfigs[commamd].PreviewUrl + strconv.Itoa(picId))
		log.Println(config.Conf.PhotoConfigs[commamd].PreviewUrl + string(rune(picId)))
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Panic(err)
		}
		var randomPic []RandomPic
		if err := json.Unmarshal(body, &randomPic); err != nil {
			log.Panic(err)
		}
		log.Println(randomPic)
		result = append(result, config.Conf.PhotoConfigs[commamd].Url+randomPic[0].Preview)
	}
	return result
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
