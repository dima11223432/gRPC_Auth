package migrator

import "flag"

func main() {
	var storagePath, migrationPath, migrationTable string
	flag.StringVar(&storagePath, "storage", "", "path to storage")
	flag.StringVar(&migrationPath, "migration", "", "path to migration")
	flag.StringVar(&migrationTable, "table", "", "migration table")
	flag.Parse()
}
