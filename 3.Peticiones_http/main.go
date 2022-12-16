package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	USERID int    `json:"userId"`
	ID     int    `json:"id"`
	TITLE  string `json:"title"`
	BODY   string `json:"body"`
}

func main() {
	posts, err := getPosts()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range posts {
		// if p.ID == 6 {
		// 	post, err := getPost(p.ID)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	log.Println(post)
		// }
		if p.USERID == 8 {
			post, err := savePost(p.USERID, "Testing Title", "hola holaaa")
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Saved new post %v", post)
		}
	}

}

func getPost(id int) (*Post, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// content, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	post := &Post{}

	err = json.NewDecoder(resp.Body).Decode(post)
	if err != nil {
		return nil, err
	}

	// log.Println(string(content))
	// log.Println(post)
	return post, nil
}

func getPosts() ([]*Post, error) {
	posts := []*Post{}
	url := "https://jsonplaceholder.typicode.com/posts/"

	res, err := http.Get(url)
	if err != nil {
		return posts, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&posts)
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func savePost(userId int, title string, body string) (*Post, error) {
	url := "https://jsonplaceholder.typicode.com/posts"
	post := &Post{
		USERID: userId,
		TITLE:  title,
		BODY:   body,
	}

	content, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(content)

	response, err := http.Post(url, "application/json", buffer)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(post)
	if err != nil {
		return nil, err
	}
	return post, err
}
