package main
 
import (
    "github.com/alphaloan/vehicle/datastore"
)
 
func main() {
    datastore.InitializeDatabase("db/migration", "sqlite3://alphaloan.db")
}