package helper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAutoBindConfig(t *testing.T) {
	err := AutoBindConfig("config.yml")
	assert.Nil(t, err)

	keys := viper.AllKeys()
	assert.Greater(t, len(keys), 0)
	assert.Equal(t, "localhost", viper.GetString("postgres.host"))
	assert.Equal(t, 5432, viper.GetInt("postgres.port"))
	assert.Equal(t, "postgres", viper.GetString("postgres.username"))
	assert.Equal(t, 1234, viper.GetInt("postgres.password"))
	assert.Equal(t, "Hieu", viper.GetString("postgres.database"))
	assert.Equal(t, "disabled", viper.GetString("postgres.ssl_mode"))
	assert.Equal(t, "Asia/Ho_Chi_Minh", viper.GetString("postgres.time_zone"))
}
