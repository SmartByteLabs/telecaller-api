package test

import (
	"testing"

	"github.com/princeparmar/telecaller-app/database"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	dbo := database.DatabaseManager.GetORM()
	assert.Nilf(t, dbo.DB().Ping(), "database ping fail")
}
