package handler

import (
	"os"
	"sync"

	"github.com/golang-graphql/connection"
	"github.com/golang-graphql/repository"
	"github.com/golang-graphql/repository/postgres/query"
	"github.com/golang-graphql/usecase"
)

var uc usecase.Usecase
var ucOnce sync.Once
var saltText = os.Getenv("SALT_PASSWORD")

func GetUsecase() usecase.Usecase {
	ucOnce.Do(func() {
		uc = usecase.NewUsecase(
			getRepository(),
			saltText,
		)
	})

	return uc
}

var repo repository.Repository
var repoOnce sync.Once

func getRepository() repository.Repository {
	repoOnce.Do(func() {
		conn := connection.NewPostgresConnection()
		repo = repository.NewRepository(getQuery(conn))
	})

	return repo
}

var qry *query.Queries
var oneQuery sync.Once

func getQuery(conn connection.Connection) *query.Queries {
	oneQuery.Do(func() {
		qry = query.New(conn.DB())
	})

	return qry
}
