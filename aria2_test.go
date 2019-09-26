package aria2_test

import "testing"

func TestClient_GetVersion(t *testing.T) {
	ver, features, err := client.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("version: ", ver)
	t.Log("features: ", features)
}
