package cmd

import (
	"os"
	"path"
)

func getSQLPath() string {
	wd, err := os.Getwd()
	if err != nil {
		wd = ""
	}

	sqlLoc := path.Join(wd, "migrations", "sql")
	_, err = os.Stat(sqlLoc)
	if err == nil {
		return sqlLoc
	}

	if !os.IsNotExist(err) {
		return ""
	}

	sqlLoc = path.Join(wd, "sql")
	_, err = os.Stat(sqlLoc)
	if err == nil {
		return sqlLoc
	}

	if !os.IsNotExist(err) {
		return ""
	}

	if err := os.MkdirAll(sqlLoc, os.ModePerm); err != nil {
		return ""
	}

	return sqlLoc
}
