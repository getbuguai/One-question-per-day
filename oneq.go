package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello BuGuai !!! ")

	url := "https://leetcode-cn.com/graphql"

	method := "POST"

	payload := strings.NewReader(`{

    "operationName": "questionOfToday",

    "variables": {},

    "query": "query questionOfToday {\n  todayRecord {\n    question {\n      questionFrontendId\n      questionTitleSlug\n      __typename\n    }\n    lastSubmission {\n      id\n      __typename\n    }\n    date\n    userStatus\n    __typename\n  }\n}\n"

}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "csrftoken=sToEjoc2KtDkRWTQpZTqUwhU78uNv1haCwJBWVDBb5HY45GGfEYwiJBubwUFFscD")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	oneq := struct {
		Data struct {
			TodayRecord []struct {
				Question struct {
					QuestionFrontendId string `json:"questionFrontendId"`
					QuestionTitleSlug  string `json:"questionTitleSlug"`
					Typename           string `json:"__typename"`
				} `json:"question"`
			} `json:"todayRecord"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(body, &oneq)
	if err != nil {
		fmt.Println(err)
		return
	}

	oneQUrl := fmt.Sprintf("https://leetcode-cn.com/problems/%s/", oneq.Data.TodayRecord[0].Question.QuestionTitleSlug)
	fmt.Printf("----------------------\n> Number: %s \n> URL: %s \n", oneq.Data.TodayRecord[0].Question.QuestionFrontendId, oneQUrl)
}
