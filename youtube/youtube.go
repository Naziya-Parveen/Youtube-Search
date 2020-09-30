package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"google.golang.org/genproto/googleapis/type/datetime"
)

//Response structure
type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

//Item details
type Item struct {
	Kind    string  `json:"kind"`
	ID      ID      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

//ID of video
type ID struct {
	Kind    string `json:"kind"`
	Videoid string `json:"videoId"`
}

//Snippet content
type Snippet struct {
	// Publishedat time `json:"publishedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//GetVideos to get videos
func GetVideos() ([]Item, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/search", nil)
	if err != nil {
		fmt.Println(err)
		return []Item{}, err
	}

	q := req.URL.Query()
	q.Add("key", "AIzaSyCKE5FPsQQn8xSzVpEh_-2k0AQ2YLh5wYM")
	q.Add("part", "snippet")
	q.Add("order", "date")
	q.Add("type", "video")
	q.Add("q", "cricket")
	q.Add("publishedAfter", "2020-01-01T00:00:00Z")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.RawQuery)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []Item{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response status : ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []Item{}, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Item{}, err
	}

	return response.Items, nil
}
