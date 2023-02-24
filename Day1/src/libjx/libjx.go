package libjx

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"errors"
)

type Ingredient struct {
	Name      string `json:"ingredient_name"           xml:"itemname"`
	Count     string `json:"ingredient_count"          xml:"itemcount"`
	Unit      string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type Cake struct {
	Name        string       `json:"name"        xml:"name"`
	Time        string       `json:"time"        xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Cakes struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

func (c Cake) FindIngredient(name string) (*Ingredient, bool) {
	for _, ingredient := range c.Ingredients {
		if ingredient.Name == name {
			return &ingredient, true
		}
	}
	return nil, false
}

func Compare(oldCakes, newCakes Cakes) []string {
	var differences []string
	cakeMap := make(map[string]Cake)
	for _, c := range oldCakes.Cakes {
		cakeMap[c.Name] = c
	}
	for _, c := range newCakes.Cakes {
		if _, ok := cakeMap[c.Name]; !ok {
			differences = append(differences, "ADDED cake \""+c.Name+"\"")
			continue
		}
		oldCake := cakeMap[c.Name]
		if c.Time != oldCake.Time {
			differences = append(differences, "CHANGED cooking time for cake \""+c.Name+"\" - \""+c.Time+"\" instead of \""+oldCake.Time+"\"")
		}
		for _, ing := range c.Ingredients {
			if oldIng, ok := oldCake.FindIngredient(ing.Name); !ok {
				differences = append(differences, "ADDED ingredient \""+ing.Name+"\" for cake \""+c.Name+"\"")
			} else {
				if ing.Unit != oldIng.Unit {
					differences = append(differences, "CHANGED unit for ingredient \""+ing.Name+"\" for cake \""+c.Name+"\" - \""+ing.Unit+"\" instead of \""+oldIng.Unit+"\"")
				}
				if ing.Count != oldIng.Count {
					differences = append(differences, "CHANGED unit count for ingredient \""+ing.Name+"\" for cake \""+c.Name+"\" - \""+ing.Count+"\" instead of \""+oldIng.Count+"\"")
				}
			}
		}
		for _, ing := range oldCake.Ingredients {
			if _, ok := c.FindIngredient(ing.Name); !ok {
				differences = append(differences, "REMOVED ingredient \""+ing.Name+"\" for cake \""+c.Name+"\"")
			}
		}
		delete(cakeMap, c.Name)
	}
	for c := range cakeMap {
		differences = append(differences, "REMOVED cake \""+c+"\"")
	}
	return differences
}

func (cakes Cakes) Print() {
	for _, cake := range cakes.Cakes {
		fmt.Println("Name:", cake.Name)
		fmt.Println("Time:", cake.Time)
		fmt.Println("Ingredients:")
		for _, ingredient := range cake.Ingredients {
			fmt.Println("-", ingredient.Name, ingredient.Count, ingredient.Unit)
		}
		fmt.Println("---")
	}
}

func extension(fileName string) string {
	return strings.TrimLeft(filepath.Ext(fileName), ".")
}

func jsonDeCoder(cakes *Cakes, fileName string) error {
	newData, err := stripJsonComments(fileName)
	if err != nil { return err }
	return json.Unmarshal(newData, &cakes)
}

func xmlDeCoder(cakes *Cakes, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil { return err }
	defer file.Close()
	return xml.NewDecoder(file).Decode(&cakes)
}


func DeCoder(fileName string) (cakes Cakes, err error) {
	switch extension(fileName) {
	case "json":
		return cakes, jsonDeCoder(&cakes, fileName)
	case "xml":
		return cakes, xmlDeCoder(&cakes, fileName)
	default:
		return cakes, errors.New("Unsupported file format")
	}
}

func EnCoder(cakes Cakes, fileName string) (err error) {
	var b []byte
	var file *os.File
	switch extension(fileName) {
	case "xml":
		file, err = os.Create(fmt.Sprintf("%s.%s", fileName, "json"))
		defer file.Close()
		if err != nil {
			return
		}
		b, err = json.MarshalIndent(cakes, "", "\t")
	case "json":
		file, err = os.Create(fmt.Sprintf("%s.%s", fileName, "xml"))
		defer file.Close()
		if err != nil {
			return
		}
		b, err = xml.MarshalIndent(cakes, "", "\t")
	default:
		return errors.New("Unsupported file format")
	}
	if err == nil { _,err = file.Write(b) }
	return err
}
