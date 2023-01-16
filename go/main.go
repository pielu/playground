package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Generic struct {
	Age      int                      `json:"age"`
	EyeColor string                   `json:"eyeColor"`
	Friends  []map[string]interface{} `json:"friends"`
	Gender   string                   `json:"gender"`
	Id       string                   `json:"_id"`
	Index    int                      `json:"index"`
	IsActive bool                     `json:"isActive"`
	Name     string                   `json:"name"`
}

type Person struct {
	Age       int    `json:"age"`
	EyeColor  string `json:"eyeColor"`
	FirstName string `json:"firstName"`
	Gender    string `json:"gender"`
	Id        string `json:"_id"`
	Index     int    `json:"index"`
	IsActive  bool   `json:"isActive"`
	LastName  string `json:"lastName"`
}

// func (p *Person) update(uk string, uv string) *Person {

// }

type SourcesResult struct {
	age      int
	eyeColor string
	name     string
	gender   string
	isActive bool
}

type Sources struct {
	Generics []Generic
	People   []Person
}

func (s Sources) find(fkv map[string]interface{}) []SourcesResult {
	var srs []SourcesResult

	for _, v := range s.Generics {
		rv := reflect.ValueOf(v)
		for fk, fv := range fkv {
			rvfv := rv.FieldByName(fk).Interface()
			if rvfv == fv {
				srs = append(srs, SourcesResult{
					age:      v.Age,
					eyeColor: v.EyeColor,
					name:     v.Name,
					gender:   v.Gender,
					isActive: v.IsActive,
				})
			}
		}
	}

	for _, v := range s.People {
		rv := reflect.ValueOf(v)
		for fk, fv := range fkv {
			rvfv := rv.FieldByName(fk).Interface()
			if rvfv == fv {
				srs = append(srs, SourcesResult{
					age:      v.Age,
					eyeColor: v.EyeColor,
					name:     fmt.Sprintf("%s %s", v.FirstName, v.LastName),
					gender:   v.Gender,
					isActive: v.IsActive,
				})
			}
		}
	}

	return srs
}

func makeSources(ns map[string]string) *Sources {
	s := new(Sources)

	for k, v := range ns {
		bs, err := os.ReadFile(v)
		if err != nil {
			fmt.Println(err)
		}

		switch k {
		case "Generic":
			err := json.Unmarshal(bs, &s.Generics)
			if err != nil {
				fmt.Println(err)
			}
		case "People":
			err := json.Unmarshal(bs, &s.People)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return s
}

func main() {
	ns := makeSources(map[string]string{
		"Generic": "generic.json",
		"People":  "people.json",
	})
	srs := ns.find(map[string]interface{}{"EyeColor": "brown"})
	for _, v := range srs {
		fmt.Println(v)
	}
}
