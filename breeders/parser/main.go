package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	t, _ := useUnmarshal()
	fmt.Println("Unmarshal: ", t.Id, t.Title)
	t2, _ := useDecode()
	fmt.Println("Decode: ", t2.Id, t2.Title)
	todo := Todo{
		Id:        44,
		UserId:    55,
		Title:     "just test",
		Completed: false,
	}
	s, _ := useMarshal(todo)
	fmt.Println("Marshal", s)
	s1, _ := useEncode(todo)
	fmt.Println("Encode", s1)
}
func useMarshal(t Todo) (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(b), nil
}
func useEncode(t Todo) (string, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(t)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return buf.String(), nil
}
func useUnmarshal() (*Todo, error) {
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &todo, nil
}
func useDecode() (*Todo, error) {
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer r.Body.Close()

	var todo Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &todo, nil
}
