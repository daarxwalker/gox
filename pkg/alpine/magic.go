package alpine

import "encoding/json"

type magic struct {
	Event string
	El    string
	Refs  string
	Store string
	Root  string
	Data  string
}

var (
	Magic = magic{
		Event: "$event",
		El:    "$el",
		Refs:  "$refs",
		Store: "$store",
		Root:  "$root",
		Data:  "$data",
	}
)

func (m magic) Id(name string) string {
	return "$id('" + name + "')"
}

func (m magic) Watch(name, script string) string {
	return "$watch('" + name + "', " + script + ")"
}

func (m magic) Dispatch(name string, data map[string]any) string {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "$dispatch('" + name + "')"
	}
	return "$dispatch('" + name + "', " + string(dataBytes) + ")"
}

func (m magic) NextTick(script string, await ...bool) string {
	r := "$nextTick(" + script + ")"
	if len(await) > 0 && await[0] {
		r = "await " + r
	}
	return r
}
