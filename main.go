package main

import "db/collections"

func main() {
	st := collections.Store{Collection: ""}
	str := Str{
		Test:  "eqew",
		Test2: 44,
	}

	st.Insert(str)
}

type Str struct {
	Test  string `json:"test" schema:"test"`
	Test2 int    `json:"test_2" schema:"test_2"`
}
