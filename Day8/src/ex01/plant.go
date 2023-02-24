package main

import (
    "fmt"
    // "unsafe"
    "reflect"
)

type Plant interface{}

type UnknownPlant struct {
    FlowerType  string
    LeafType    string
    Color       int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
    FlowerColor int
    LeafType    string
    Height      int `unit:"inches"`
}

func DescribePlant(plant interface{}) {
	val := reflect.ValueOf(plant)
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		key := typeField.Name

		var tag string
		if colorTag, ok := typeField.Tag.Lookup("color_scheme"); ok {
			tag = fmt.Sprintf("(color_scheme=%s)", colorTag)
		}
		if unitTag, ok := typeField.Tag.Lookup("unit"); ok {
			tag = fmt.Sprintf("(unit=%s)", unitTag)
		}
		fmt.Printf("%s: %v %s\n", key, valueField.Interface(), tag)
	}
}

func main() {
    t := UnknownPlant{
        FlowerType: "ee",
        LeafType: "rr",
        Color: 1,
    }
    DescribePlant(t)
}
