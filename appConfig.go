package main

import "path"

type AppConfig struct {
	dataPath string
	db       []*Toko
}

func newConfig() *AppConfig {
	const DATA_PATH = "data"
	const DATA_FILE_NAME = "db.json"

	var tokoCollection = make([]*Toko, 0)
	return &AppConfig{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		db:       tokoCollection,
	}
}

func (c *AppConfig) getDB() []*Toko {
	return c.db
}

func (c *AppConfig) getDbPath() string {
	return c.dataPath
}