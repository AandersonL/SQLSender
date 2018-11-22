package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/moraes/config"
	"os"
	"time"
)

func Query(instruction *string, db *sql.DB) (string, error) {

	query, err := db.Query(*instruction)

	if err != nil {
		return "", err
	}
	var dbRes string
	for query.Next() {
		query.Scan(&dbRes)
	}
	return dbRes, err
}

func LoadConfig() string {
	cfg, err := os.Open("config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on open config file!\n")
		os.Exit(1)
	}

	cfg_str := ""
	buff := bufio.NewScanner(cfg)
	for buff.Scan() {
		cfg_str = cfg_str + buff.Text() + "\n"
	}
	return cfg_str
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [sql_script1, sql_script2, sql_script3...]\n", os.Args[0])
		os.Exit(1)
	}
	cfg, err := config.ParseYaml(LoadConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on load config: %s\n", err)
		os.Exit(1)
	}
	user, _ := cfg.String("database.user")
	pass, _ := cfg.String("database.pass")
	host, _ := cfg.String("database.host")
	port, _ := cfg.String("database.port")
	db_name, _ := cfg.String("database.db_name")
	db_type, _ := cfg.String("database.db_type")

	db, _ := sql.Open(db_type, fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, db_name))

	err = db.Ping()
	if err != nil {
		fmt.Println("Database not responding, check your config file or either server is up and running")
		os.Exit(1)
	}

	// Opening script(s)
	start := time.Now()
	for _, file := range os.Args[1:] {
		fileFd, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cant open file %s\n", file)
			os.Exit(1)
		}
		scan := bufio.NewScanner(fileFd)
		query := ""
		for scan.Scan() {
			line := scan.Text()
			query = query + line
			if len(query) > 0 && query[len(query)-1] == ';' {
				_, err := Query(&query, db)
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
					os.Exit(1)
				}
				query = ""
			}
		}
	}

	fmt.Printf("Finished in %s\n", time.Since(start))
	defer db.Close()
}
