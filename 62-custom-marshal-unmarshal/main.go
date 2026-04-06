package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	marshal()
	unmarshal()
}

type CarType int

const (
	_ CarType = iota
	Hatchback
	Sedan
	SUV
)

var (
	stringToCarType = map[string]CarType{
		"hatchback": Hatchback,
		"sedan":     Sedan,
		"suv":       SUV,
	}

	carTypeToString = map[CarType]string{
		Hatchback: "hatchback",
		Sedan:     "sedan",
		SUV:       "suv",
	}
)

func (c CarType) String() string {
	return carTypeToString[c]
}

func (c CarType) MarshalJSON() ([]byte, error) {
	t, ok := carTypeToString[c]
	if !ok {
		return nil, errors.New("invalid car type provided")
	}
	return []byte(`"` + t + `"`), nil
}

func marshal() {
	carType := Hatchback
	bytes, err := json.Marshal(carType)
	if err != nil {
		fmt.Println("error marshalling car type:", err)
		panic(err)
	}

	fmt.Println(string(bytes))
}

func (c *CarType) UnmarshalJSON(bytes []byte) error {
	var str string
	if err := json.Unmarshal(bytes, &str); err != nil {
		return err
	}
	t, ok := stringToCarType[str]
	if !ok {
		return fmt.Errorf("invalid car type: %s", str)
	}
	*c = t
	return nil
}

func unmarshal() {
	var unmarshalledCarType CarType
	err := json.Unmarshal([]byte(`"suv"`), &unmarshalledCarType)
	if err != nil {
		fmt.Println("error unmarshalling car type:", err)
		panic(err)
	}
	fmt.Println(unmarshalledCarType)
}
