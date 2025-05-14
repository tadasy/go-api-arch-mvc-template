package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitEnv(t *testing.T) {
	err := LoadEnv()
	assert.Nil(t, err)

	assert.Equal(t, "development", Config.Env)
	// DBHost
	assert.Equal(t, "0.0.0.0", Config.DBHost)
	// DBPort
	assert.Equal(t, 3306, Config.DBPort)
	// DBDriver
	assert.Equal(t, "mysql", Config.DBDriver)
	// DBName
	assert.Equal(t, "api_database", Config.DBName)
	// DBUser
	assert.Equal(t, "app", Config.DBUser)
	// DBPassword
	assert.Equal(t, "password", Config.DBPassword)

	assert.Equal(t, true, Config.IsDevelopment())
}
