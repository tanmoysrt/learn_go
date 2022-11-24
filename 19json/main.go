package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	EncodeJson()
	decodeJson()
}

type course struct {
	Name     string `json:"course-name"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lcoCourses := []course{
		{Name: "Go", Price: 200, Platform: "LCO", Password: "1234", Tags: []string{"beginner", "intermediate", "advance"}},
		{Name: "Python", Price: 200, Platform: "LCO", Password: "1234", Tags: nil},
		{Name: "Django", Price: 200, Platform: "LCO", Password: "1234", Tags: []string{"beginner", "intermediate", "advance"}},
	}

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, " ", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}


func decodeJson() {
	jsonData := []byte(`[{"course-name":"Go","Price":200,"website":"LCO","tags":["beginner","intermediate","advance"]},{"course-name":"Python","Price":200,"website":"LCO","tags":null},{"course-name":"Django","Price":200,"website":"LCO","tags":["beginner","intermediate","advance"]}]`)

	var courses []course
	err := json.Unmarshal(jsonData, &courses)
	if err != nil {
		panic(err)
	}

	fmt.Println(courses)
	for _, course := range courses {
		fmt.Println(course.Name)
	}
}