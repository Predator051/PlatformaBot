// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AdminsOfGroupListRequest struct {
	ID          int64
	ChatID      pgtype.Int8
	GroupListID pgtype.Int4
	Username    pgtype.Text
	FirstName   pgtype.Text
	SecondName  pgtype.Text
}

type GroupList struct {
	ID   int64
	Name string
}

type GroupListsAdmin struct {
	ID          int64
	ChatID      pgtype.Int8
	GroupListID pgtype.Int4
}

type SubscriptionToGroupList struct {
	ID          int64
	ChatID      pgtype.Int8
	GroupListID pgtype.Int4
	Username    pgtype.Text
	Title       pgtype.Text
	ChatType    pgtype.Text
}