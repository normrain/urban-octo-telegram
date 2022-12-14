package apis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Auth() {
	url := "https://api.cloudmore.com/connect/token"
	method := "POST"

	payload := strings.NewReader("grant_type=password&username=<<username>>&password=<<password>>")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic <<Token>>")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
	fmt.Println(string(body))
}
