package osugoi

import "testing"

func TestGetRecentScores(t *testing.T) {
	scores, cursor, err := testClient.GetRecentScores(Standard)
	if err != nil {
		t.Errorf("GetRecentScores failed: %v", err)	
	}

	t.Logf("Cursor String: %s\n", cursor)
	t.Logf("%+v\n", (*scores)[4])
}
