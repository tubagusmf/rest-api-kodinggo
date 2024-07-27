package console

import (
	"database/sql"
	"golang-rest-api-articles/internal/config"
	"golang-rest-api-articles/internal/helper"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var (
	direction string
	step      int = 1
)

func init() {
	rootCmd.AddCommand(migrationCmd)

	migrationCmd.Flags().StringVarP(&direction, "direction", "d", "up", "Migration direction")

	migrationCmd.Flags().IntVarP(&step, "step", "s", 1, "Migration step")
}

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Migration management database",
	Run:   migrateDB,
}

func migrateDB(cmd *cobra.Command, args []string) {
	config.LoadWithViper()

	connDB, err := sql.Open("mysql", helper.GetConnectionString())
	if err != nil {
		log.Panicf("Gagal terkoneksi ke database: %s", err.Error())
	}

	defer connDB.Close()

	migrations := &migrate.FileMigrationSource{Dir: "./db/migrations"}

	var n int

	if direction == "down" {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Up, step)
	}
	if err != nil {
		log.Panicf("Gagal melakukan migration: %s", err.Error())
	}

	log.Printf("Sukses melakukan migration(s): %d", n)
}
