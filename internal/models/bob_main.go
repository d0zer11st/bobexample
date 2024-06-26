// Code generated by BobGen psql v0.26.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"

	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

var TableNames = struct {
	Examples string
}{
	Examples: "example",
}

var ColumnNames = struct {
	Examples exampleColumnNames
}{
	Examples: exampleColumnNames{
		ID:  "id",
		Mac: "mac",
	},
}

var (
	SelectWhere = Where[*dialect.SelectQuery]()
	InsertWhere = Where[*dialect.InsertQuery]()
	UpdateWhere = Where[*dialect.UpdateQuery]()
	DeleteWhere = Where[*dialect.DeleteQuery]()
)

func Where[Q psql.Filterable]() struct {
	Examples exampleWhere[Q]
} {
	return struct {
		Examples exampleWhere[Q]
	}{
		Examples: ExampleWhere[Q](),
	}
}

var (
	SelectJoins = getJoins[*dialect.SelectQuery]
	UpdateJoins = getJoins[*dialect.UpdateQuery]
	DeleteJoins = getJoins[*dialect.DeleteQuery]
)

type joinSet[Q any] struct {
	InnerJoin Q
	LeftJoin  Q
	RightJoin Q
}

type joins[Q dialect.Joinable] struct{}

func getJoins[Q dialect.Joinable](ctx context.Context) joins[Q] {
	return joins[Q]{}
}
