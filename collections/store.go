package collections

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
)

// Store структура для работы коллекциями.
type Store struct {

	// Collection название таблицы "schema.name" || "name".
	Collection string

	// primaryKey первичный ключ.
	primaryKey string

	// primaryStructKey название свойства первичного ключа структуры.
	primaryStructKey string

	*sql.DB
}

// NewStore создание экземпляра структуры.
func NewStore(collection string, db *sql.DB) *Store {
	return &Store{Collection: collection, DB: db}
}

// Insert вставка структуры в таблицу.
func (s *Store) Insert(ctx context.Context, str interface{}) (id int, err error) {
	keys, values := s.getMap(str, false)
	query := "INSERT INTO " + s.Collection + "(" + strings.Join(keys, ",") + ") VALUES (?" + strings.Repeat(",?", len(keys)-1) + ") RETURNING id"
	err = s.DB.QueryRowContext(ctx, query, values...).Scan(&id)
	if err != nil {
		return
	}
	return
}

// Delete удаление записи по первичному ключу.
func (s *Store) Delete(ctx context.Context, str interface{}) (id int, err error) {
	query := "DELETE FROM " + s.Collection + " WHERE " + s.getPrimaryKey(str) + " = ?"
	err = s.DB.QueryRowContext(ctx, query, s.getColumnVal(str, s.getPrimaryStructKey(str))).Scan(&id)
	if err != nil {
		return
	}
	return
}

// Update обновление всех полей записи из структуры с фильтром по первичном ключу.
func (s *Store) Update(ctx context.Context, str interface{}) (id int, err error) {
	keys, values := s.getMap(str, false)
	query := "UPDATE " + s.Collection + " SET " + s.getSetLine(keys) + " WHERE " + s.getPrimaryKey(str) + " = ?"
	err = s.DB.QueryRowContext(ctx, query, append(values, s.getColumnVal(str, s.getPrimaryStructKey(str)))).Scan(&id)
	if err != nil {
		return
	}
	return
}

// Get получение записи по первичному ключу.
func (s *Store) Get(ctx context.Context, str interface{}) (interface{}, error) {
	keys, _ := s.getMap(str, true)
	result := ""
	query := "SELECT " + s.getJsonBuild(keys) + " FROM " + s.Collection + " WHERE " + s.getPrimaryKey(str) + " = ?"
	err := s.DB.QueryRowContext(ctx, query, s.getColumnVal(str, s.getPrimaryStructKey(str))).Scan(&result)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), str)
	if err != nil {
		return nil, err
	}
	return str, nil
}
