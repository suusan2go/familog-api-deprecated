package model

func (db *DB) Migration() {
	db.AutoMigrate(&Diary{})
}
