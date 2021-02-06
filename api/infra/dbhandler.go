package infra

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/tidwall/buntdb"
)

type DBHandler struct {
	db *buntdb.DB
}

func NewDBHandler() *DBHandler {
	inMemoryEnv := os.Getenv("DATABASE_IN_MEMORY")
	inMemory, err := strconv.ParseBool(inMemoryEnv)
	if err != nil {
		log.Fatal("Could not parse DATABASE_IN_MEMORY environment variable.")
	}
	if inMemory {
		db, err := buntdb.Open(":memory:")
		if err != nil {
			log.Fatal(err)
		}
		return &DBHandler{
			db,
		}
	}
	dbDirEnv := os.Getenv("DATABASE_DIR")
	dbDir := filepath.Dir(filepath.Clean(dbDirEnv))

	db, err := buntdb.Open(filepath.Join(dbDir, "notenote.db"))
	if err != nil {
		log.Fatal(err)
	}
	return &DBHandler{
		db,
	}
}
