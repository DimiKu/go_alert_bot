package handlers

import (
	"go_alert_bot/pkg/db_operations"
)

type AddOption func(*common.SerializedTaskMetaData) error

func NewandleFunc(storage *db_operations.Storage)
