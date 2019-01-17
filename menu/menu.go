package menu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/iancoleman/orderedmap"
)

type Menu struct {
	orderedmap.OrderedMap
}

var Cfn Menu

func init() {
	f, err := os.Open("CfnMenu.json")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &Cfn)
	if err != nil {
		panic(err)
	}
}

func (m Menu) Build(name string) (*orderedmap.OrderedMap, error) {
	thing, ok := m.Get(name)

	if !ok {
		return orderedmap.New(), fmt.Errorf("No such type: %s", name)
	}

	return buildThing(name, thing.(orderedmap.OrderedMap))
}

func buildThing(name string, thing orderedmap.OrderedMap) (*orderedmap.OrderedMap, error) {
	var props orderedmap.OrderedMap

	if o, ok := thing.Get("Output"); ok {
		props = o.(orderedmap.OrderedMap)
	} else {
		props = *orderedmap.New()
	}

	for _, key := range thing.Keys() {
		question, _ := thing.Get(key)

		q, err := buildQuestion(key, question.(orderedmap.OrderedMap))
		if err != nil {
			return orderedmap.New(), err
		}

		mungeOrderedMap(&props, q)
	}

	output := orderedmap.New()
	output.Set("Type", name)
	output.Set("Properties", props)

	return output, nil
}

func buildQuestion(name string, question orderedmap.OrderedMap) (*orderedmap.OrderedMap, error) {
	var output orderedmap.OrderedMap

	if o, ok := question.Get("Output"); ok {
		output = o.(orderedmap.OrderedMap)
	} else {
		output = *orderedmap.New()
	}

	if o, ok := question.Get("Options"); ok {
		options := o.(orderedmap.OrderedMap)
		f, _ := options.Get(options.Keys()[0])
		first := f.(orderedmap.OrderedMap)

		mungeOrderedMap(&output, &first)
	}

	return &output, nil
}
