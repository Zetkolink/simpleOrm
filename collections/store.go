package collections

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type Store struct {
	Collection string
	primaryKey string
	*sql.DB
}

func NewStore(collection string, db *sql.DB) *Store {
	return &Store{Collection: collection, DB: db}
}

func (s *Store) Insert(ctx context.Context, str interface{}) (int, error) {
	keys, values := s.getMap(str)
	query := "INSERT INTO " + s.Collection + "(" + strings.Join(keys, ",") + ") VALUES (?" + strings.Repeat("?", len(keys)-1) + ") RETURNING id"
	fmt.Println(query)
	_, err := s.DB.QueryContext(ctx, query, values...)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (s *Store) getMap(str interface{}) ([]string, []interface{}) {
	sVal := reflect.ValueOf(str)
	sType := reflect.TypeOf(str)
	values := make([]interface{}, sVal.NumField())
	keys := make([]string, sVal.NumField())

	for i := 0; i < sVal.NumField(); i++ {
		schemaName := sType.Field(i).Tag.Get(SCHEMA)
		field := sVal.Field(i)
		keys[i] = schemaName
		switch field.Kind() {
		case reflect.String:
			values[i] = field.String()
		case reflect.Int:
			values[i] = field.Int()
		}
	}

	return keys, values
}
