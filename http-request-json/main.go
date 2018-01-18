package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	m "tauki.com/practice/http-request-json/model"
	"strconv"
)

const(
	token = `imle3KGX1ilyFUpeVHgfqqVzGSsBDXgudE7u6OH4BkLdMwTUmqcWTBlX5TJdt7dA`
	tokenType = `X-CSRFToken`
	url = "https://backenddev.shophobe.com/api/references/"
	categories = url + "categories/"
	temp = 100
)

func main() {
	var category []m.Category

	req, err := http.NewRequest("GET", categories, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add(tokenType, token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// read categories
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("_____CAT_____")
	json.NewDecoder(resp.Body).Decode(&category)
	fmt.Println(category)

	fmt.Println("_____SUB_____")
	// read subcategories
	for index, i := range category {
		var result []m.Category
		for j := i.ID +1; j < temp ; j++ {
			sURL := categories + strconv.Itoa(i.ID) + "/" + strconv.Itoa(j)
			req, err = http.NewRequest("GET", sURL, nil)
			if err != nil {
				panic(err)
			}
			fmt.Println(sURL)

			resp, err = client.Do(req)
			if err != nil {
				panic(err)
			}

			var data m.Category
			json.NewDecoder(resp.Body).Decode(&data)

			if data.Detail == "Not found." {
				break
			}

			result = append(result, data)
		}
		category[index].Sub = result

	}

	d, _ := json.MarshalIndent(&category, " ", " ")
	fmt.Println(string(d))
}
