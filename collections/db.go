package collections

import "database/sql"

// Connect initialize connect to DB
func Connect(connLine string, driverName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, connLine)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
