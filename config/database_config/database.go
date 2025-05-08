package databaseconfig

import "os"

type config struct {
	USERNAME string
	PASSWORD string
	DATABASE string
	HOST     string
	PORT     string
}

func DatabaseConfig() *config {

	username := os.Getenv("MYSQL_USER")
	if username == "" {
		username = "your_username"
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "your_password"
	}

	database := os.Getenv("MYSQL_DATABASE")
	if database == "" {
		database = "your_database_name"
	}

	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}

	return &config{
		USERNAME: username,
		PASSWORD: password,
		DATABASE: database,
		HOST:     host,
		PORT:     port,
	}
}
