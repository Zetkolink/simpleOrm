package collections

import (
	"reflect"
	"strconv"
	"strings"
)

// getMap конвертации структуры в два массива ключей по тегу SCHEMA и значений.
func (s *Store) getMap(str interface{}, withPrimary bool) ([]string, []interface{}) {
	sVal := reflect.ValueOf(str)
	sType := reflect.TypeOf(str)
	length := sVal.NumField()
	if !withPrimary {
		length--
	}
	values := make([]interface{}, length)
	keys := make([]string, length)
	j := 0
	for i := 0; i < sVal.NumField(); i++ {
		schemaName := sType.Field(i).Tag.Get(SCHEMA)
		primary := sType.Field(i).Tag.Get(PRIMARY)
		if schemaName != "" {
			if primary == "" || withPrimary {
				field := sVal.Field(i).Interface()
				keys[i-j] = schemaName
				values[i-j] = field
			} else if primary != "" {
				j++
			}
		}
	}
	return keys, values
}

// getMap конвертации структуры в два массива ключей по тегу SCHEMA и значений.
func (s *Store) getKeys(str interface{}) []string {
	sVal := reflect.ValueOf(str)
	sType := reflect.TypeOf(str)
	length := sVal.NumField()
	keys := make([]string, length)
	j := 0
	for i := 0; i < sVal.NumField(); i++ {
		schemaName := sType.Field(i).Tag.Get(SCHEMA)
		if schemaName != "" {
			keys[i-j] = schemaName
		}
	}
	return keys
}

// getColumnVal получение значения свойства структуры.
func (s *Store) getColumnVal(str interface{}, field string) interface{} {
	sVal := reflect.ValueOf(str)

	value := sVal.FieldByName(field)
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Int:
		return value.Int()
	}

	return nil
}

// getPrimaryKey получение первичного ключа.
func (s *Store) getPrimaryKey(str interface{}) string {
	if s.primaryKey != "" {
		return s.primaryKey
	}

	s.setPrimary(str)

	return s.primaryKey
}

// getPrimaryStructKey получение названия свойства первичного ключа структуры.
func (s *Store) getPrimaryStructKey(str interface{}) string {
	if s.primaryStructKey != "" {
		return s.primaryStructKey
	}

	s.setPrimary(str)

	return s.primaryStructKey
}

// setPrimary устанавливает названия первичного ключа таблицы и структуры.
func (s *Store) setPrimary(str interface{}) {
	sVal := reflect.ValueOf(str)
	sType := reflect.TypeOf(str)

	for i := 0; i < sVal.NumField(); i++ {
		primary := sType.Field(i).Tag.Get(PRIMARY)
		primaryStructKey := sType.Field(i).Name
		if primary != "" {
			s.primaryKey = primary
			s.primaryStructKey = primaryStructKey
		}
	}
}

// getSetLine получение строки для обновления записи.
func (s *Store) getSetLine(keys []string) (sets string) {
	for i := 0; i < len(keys); i++ {
		if sets != "" {
			sets += ", " + keys[i] + " = ?"
		} else {
			sets += keys[i] + " = ?"
		}
	}
	return
}

// getPlaceholder получение плейсхолдера.
func (s *Store) getPlaceholder(keys []string) (sets string) {
	for i := 2; i <= len(keys); i++ {
		sets += ",$" + strconv.Itoa(i)
	}
	return
}

// getJsonBuild получение строки для построения JSON ответа при запросе.
func (s *Store) getJsonBuild(keys []string) string {
	items := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		items[i] = "'" + keys[i] + "'," + keys[i]
	}
	return "json_build_object(" + strings.Join(items, ",") + ")"
}
