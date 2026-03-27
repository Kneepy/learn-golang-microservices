package queries

import (
	_ "embed"
	"strings"
)

//go:embed user.sql
var userSQL string

//go:embed init.sql
var initSQL string

type UserQueries struct {
	CreateUser     string
	GetUserByID    string
	GetUserByEmail string
	UpdateUser     string
	DeleteUser     string

	CheckEmailExist string

	UpdateUserStatus string
	GetUserStatus    string

	SearchUser string
}

type InitQueries struct {
	InitDB        string
	InitUserTable string
}

// ParseQueries извлекает именованные запросы из SQL файла
// Формат: -- name: QueryName :one|many|exec
func ParseQueries(sqlContent string) map[string]string {
	queries := make(map[string]string)

	lines := strings.Split(sqlContent, "\n")
	var currentQuery strings.Builder
	var currentName string
	var isQuery bool

	for _, line := range lines {
		if strings.HasPrefix(line, "-- name:") {
			if isQuery && currentName != "" {
				queries[currentName] = strings.TrimSpace(currentQuery.String())
				currentQuery.Reset()
			}

			parts := strings.Fields(line)

			if len(parts) >= 3 {
				currentName = strings.TrimSpace(parts[2])
				isQuery = true
				currentQuery.Reset()
			}

			continue
		}

		if isQuery {
			currentQuery.WriteString(line)
			currentQuery.WriteRune('\n')
		}
	}

	if isQuery && currentName != "" {
		queries[currentName] = strings.TrimSpace(currentQuery.String())
		currentQuery.Reset()
	}

	return queries
}

var (
	userQueries *UserQueries
	initQueries *InitQueries
)

func GetUserQueries() *UserQueries {
	if userQueries != nil {
		return userQueries
	}

	userQueriesMap := ParseQueries(userSQL)

	userQueries = &UserQueries{
		CreateUser:       userQueriesMap["CreateUser"],
		GetUserByID:      userQueriesMap["GetUserById"],
		GetUserByEmail:   userQueriesMap["GetUserByEmail"],
		UpdateUser:       userQueriesMap["UpdateUser"],
		DeleteUser:       userQueriesMap["DeleteUser"],
		CheckEmailExist:  userQueriesMap["CheckEmailExist"],
		UpdateUserStatus: userQueriesMap["UpdateUserStatus"],
		GetUserStatus:    userQueriesMap["GetUserStatus"],
		SearchUser:       userQueriesMap["SearchUser"],
	}

	return userQueries
}

func GetInitQueries() *InitQueries {
	if initQueries != nil {
		return initQueries
	}

	initQueriesMap := ParseQueries(initSQL)

	initQueries = &InitQueries{
		InitDB:        initQueriesMap["InitDB"],
		InitUserTable: initQueriesMap["InitUserTable"],
	}

	return initQueries
}
