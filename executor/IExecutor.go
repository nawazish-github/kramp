package executor

import "github.com/nawazish-github/kramp/models"

type IExecutor interface {
	Fetch(searchString string) models.Response
}
