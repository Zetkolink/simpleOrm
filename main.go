package main

import (
	"context"
	"db/collections"
	"time"
)

func main() {
	st := collections.Store{Collection: "test"}
	str := Change{
		Id:        1,
		ProjectId: 527,
		Filename:  "osm.change",
		UpdateAt:  time.Time{},
		UpdateBy:  44,
		Comment:   "Test",
		Active:    false,
	}

	st.GetAll(context.Background(), str)
}

// Change osm change.
type Change struct {
	Id        int       `json:"id" schema:"id" primary:"id"`
	ProjectId int       `json:"project_id" schema:"project_id"`
	Filename  string    `json:"filename" schema:"filename"`
	Active    bool      `json:"active" schema:"active"`
	UpdateAt  time.Time `json:"update_at" schema:"update_at"`
	UpdateBy  int       `json:"update_by" schema:"update_by"`
	Comment   string    `json:"comment" schema:"comment"`
}
