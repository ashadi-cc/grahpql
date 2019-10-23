package pg

import "testing"

func TestGetInstace(t *testing.T) {
	if testing.Short() {
		t.Skip("skip db test")
	}
	_, err := connectDB()

	if err != nil {
		t.Fatalf("Cannot connect to database %s", err.Error())
	}
}
