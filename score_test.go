package osugoi

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGetRecentScores(t *testing.T) {
	const data string = "{\"scores\":[{\"classic_total_score\":3616240,\"preserve\":true,\"processed\":true,\"ranked\":true,\"maximum_statistics\":{\"great\":513,\"legacy_combo_increase\":172},\"mods\":[{\"acronym\":\"DT\"},{\"acronym\":\"CL\"}],\"statistics\":{\"ok\":89,\"meh\":7,\"miss\":14,\"great\":403},\"total_score_without_mods\":394915,\"beatmap_id\":4195108,\"best_id\":null,\"id\":4907621995,\"rank\":\"B\",\"type\":\"solo_score\",\"user_id\":17848512,\"accuracy\":0.845679,\"build_id\":null,\"ended_at\":\"2025-05-27T10:20:00Z\",\"has_replay\":true,\"is_perfect_combo\":false,\"legacy_perfect\":false,\"legacy_score_id\":4843524826,\"legacy_total_score\":2260501,\"max_combo\":322,\"passed\":true,\"pp\":115.635,\"ruleset_id\":0,\"started_at\":null,\"total_score\":417030,\"replay\":true,\"current_user_attributes\":{\"pin\":null}},{\"classic_total_score\":975762,\"preserve\":true,\"processed\":true,\"ranked\":true,\"maximum_statistics\":{\"great\":364,\"legacy_combo_increase\":272},\"mods\":[{\"acronym\":\"CL\"}],\"statistics\":{\"ok\":80,\"meh\":21,\"miss\":13,\"great\":250},\"total_score_without_mods\":230199,\"beatmap_id\":5095584,\"best_id\":null,\"id\":4907621996,\"rank\":\"C\",\"type\":\"solo_score\",\"user_id\":37941977,\"accuracy\":0.769689,\"build_id\":null,\"ended_at\":\"2025-05-27T10:20:00Z\",\"has_replay\":true,\"is_perfect_combo\":false,\"legacy_perfect\":false,\"legacy_score_id\":4843524827,\"legacy_total_score\":402000,\"max_combo\":78,\"passed\":true,\"pp\":3.44624,\"ruleset_id\":0,\"started_at\":null,\"total_score\":220991,\"replay\":true,\"current_user_attributes\":{\"pin\":null}},{\"classic_total_score\":294139,\"preserve\":true,\"processed\":true,\"ranked\":true,\"maximum_statistics\":{\"great\":128,\"legacy_combo_increase\":65},\"mods\":[{\"acronym\":\"CL\"}],\"statistics\":{\"ok\":13,\"miss\":5,\"great\":110},\"total_score_without_mods\":483557,\"beatmap_id\":871027,\"best_id\":null,\"id\":4907621997,\"rank\":\"B\",\"type\":\"solo_score\",\"user_id\":37074992,\"accuracy\":0.893229,\"build_id\":null,\"ended_at\":\"2025-05-27T10:20:00Z\",\"has_replay\":false,\"is_perfect_combo\":false,\"legacy_perfect\":false,\"legacy_score_id\":4843524828,\"legacy_total_score\":157322,\"max_combo\":63,\"passed\":true,\"pp\":32.9272,\"ruleset_id\":0,\"started_at\":null,\"total_score\":464215,\"replay\":false,\"current_user_attributes\":{\"pin\":null}},{\"classic_total_score\":405269,\"preserve\":true,\"processed\":true,\"ranked\":true,\"maximum_statistics\":{\"great\":125,\"legacy_combo_increase\":21},\"mods\":[{\"acronym\":\"CL\"}],\"statistics\":{\"ok\":21,\"meh\":2,\"great\":102},\"total_score_without_mods\":693300,\"beatmap_id\":1639736,\"best_id\":null,\"id\":4907621998,\"rank\":\"B\",\"type\":\"solo_score\",\"user_id\":20625118,\"accuracy\":0.874667,\"build_id\":null,\"ended_at\":\"2025-05-27T10:20:00Z\",\"has_replay\":false,\"is_perfect_combo\":true,\"legacy_perfect\":true,\"legacy_score_id\":4843524829,\"legacy_total_score\":396706,\"max_combo\":146,\"passed\":true,\"pp\":44.6868,\"ruleset_id\":0,\"started_at\":null,\"total_score\":665568,\"replay\":false,\"current_user_attributes\":{\"pin\":null}}], \"cursor_string\":\"random_values!\"}" 
	//" < for Vim tokenizer, because data variable has a 3k long string

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/scores" {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	}))
	defer mockServer.Close()

	client := NewClient(mockServer.URL)

	scores, cursor, err := client.GetRecentScores(Standard)
	if err != nil {
		t.Errorf("GetRecentScores failed: %v", err)	
	}

	t.Logf("Cursor String: %s\n", cursor)
	t.Logf("%+v\n", (*scores)[len(*scores) - 1])
}
