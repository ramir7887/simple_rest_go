package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=89.223.65.66 port=1234 dbname=onboarding sslmode=disable user=onboarding password=secret_pass"
	}
	os.Exit(m.Run())
}
