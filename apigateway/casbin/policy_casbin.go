package casbin

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

const (
	host     = "postgres3"
	port     = "5432"
	dbname   = "casbin"
	username = "postgres"
	password = "1111" //3333 husan
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	fmt.Println("fsdfadsfiadsfjasd")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS casbin;")
	if err != nil {
		return nil, err
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password)
	fmt.Println(conn)
	adapter, err := xormadapter.NewAdapter("postgres", conn)
	if err != nil {
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("./casbin/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	policies := [][]string{
		// user user service
		{"customer", "/auth/register", "POST"},
		{"customer", "/auth/profile", "GET"},
		{"customer", "/auth/profile", "PUT"},
		{"customer", "/auth/users/:id", "DELETE"},
		{"customer", "/auth/users/:id", "GET"},

		// admin user service
		{"admin", "/auth/register", "POST"},
		{"admin", "/auth/profile", "GET"},
		{"admin", "/auth/profile", "PUT"},
		{"admin", "/auth/users/:id", "DELETE"},
		{"admin", "/auth/users/:id", "GET"},

		// user learning service
		{"customer", "/learning/languages", "GET"},
		{"customer", "/learning/languages", "POST"},
		{"customer", "/learning/lessons", "GET"},
		{"customer", "/learning/lessons/:id", "GET"},
		{"customer", "/learning/lessons/complate", "POST"},
		{"customer", "/learning/exercises:code", "GET"},
		{"customer", "/learning/exercises/:id", "POST"},
		{"customer", "/learning/vocabulary", "GET"},
		{"customer", "/learning/vocabulary", "POST"},

		// admin learning service
		{"admin", "learning/languages", "GET"},
		{"admin", "learning/languages", "POST"},
		{"admin", "learning/lessons", "GET"},
		{"admin", "learning/lessons/:id", "GET"},
		{"admin", "learning/lessons/complate", "POST"},
		{"admin", "learning/exercises", "GET"},
		{"admin", "learning/exercises/:id", "POST"},
		{"admin", "learning/vocabulary", "GET"},
		{"admin", "learning/vocabulary", "POST"},
	}
	fmt.Println("sadfadsjfadsjfiadsifads")
	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	err = enforcer.SavePolicy()
	if err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}
	fmt.Println("sdsdfjdsijfiasd")
	return enforcer, nil
}
