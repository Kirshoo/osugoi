package osugoi

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
)

// https://osu.ppy.sh/beatmapsets/2349118#osu/5054112
const testBeatmapId = 5054112

func TestGetBeatmap(t *testing.T) {
	const data string = "{\"beatmapset_id\":2349118,\"difficulty_rating\":4.9,\"id\":5054112,\"mode\":\"osu\",\"status\":\"ranked\",\"total_length\":204,\"user_id\":7326908,\"version\":\"Amateurre's Insane\",\"accuracy\":8,\"ar\":8.8,\"bpm\":174,\"convert\":false,\"count_circles\":442,\"count_sliders\":330,\"count_spinners\":2,\"cs\":4,\"deleted_at\":null,\"drain\":5,\"hit_length\":199,\"is_scoreable\":true,\"last_updated\":\"2025-05-17T22:09:01Z\",\"mode_int\":0,\"passcount\":279,\"playcount\":990,\"ranked\":1,\"url\":\"https:\\/\\/osu.ppy.sh\\/beatmaps\\/5054112\",\"checksum\":\"2d6634f712f305a907f462c438ffcea7\",\"beatmapset\":{\"artist\":\"Ryokuoushoku Shakai\",\"artist_unicode\":\"\\u7dd1\\u9ec4\\u8272\\u793e\\u4f1a\",\"covers\":{\"cover\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/cover.jpg?1747519762\",\"cover@2x\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/cover@2x.jpg?1747519762\",\"card\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/card.jpg?1747519762\",\"card@2x\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/card@2x.jpg?1747519762\",\"list\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/list.jpg?1747519762\",\"list@2x\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/list@2x.jpg?1747519762\",\"slimcover\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/slimcover.jpg?1747519762\",\"slimcover@2x\":\"https:\\/\\/assets.ppy.sh\\/beatmaps\\/2349118\\/covers\\/slimcover@2x.jpg?1747519762\"},\"creator\":\"FuJu\",\"favourite_count\":34,\"genre_id\":4,\"hype\":null,\"id\":2349118,\"language_id\":3,\"nsfw\":false,\"offset\":0,\"play_count\":4666,\"preview_url\":\"\\/\\/b.ppy.sh\\/preview\\/2349118.mp3\",\"source\":\"\",\"spotlight\":false,\"status\":\"ranked\",\"title\":\"Scarlet\",\"title_unicode\":\"\\u30b9\\u30ab\\u30fc\\u30ec\\u30c3\\u30c8\",\"track_id\":null,\"user_id\":10773882,\"video\":false,\"bpm\":174,\"can_be_hyped\":false,\"deleted_at\":null,\"discussion_enabled\":true,\"discussion_locked\":false,\"is_scoreable\":true,\"last_updated\":\"2025-05-17T22:09:00Z\",\"legacy_thread_url\":\"https:\\/\\/osu.ppy.sh\\/community\\/forums\\/topics\\/2061954\",\"nominations_summary\":{\"current\":2,\"eligible_main_rulesets\":[\"osu\"],\"required_meta\":{\"main_ruleset\":2,\"non_main_ruleset\":1}},\"ranked\":1,\"ranked_date\":\"2025-05-25T19:03:02Z\",\"rating\":8.75,\"storyboard\":false,\"submitted_date\":\"2025-04-03T14:56:53Z\",\"tags\":\"nakazawa mnyui gelidium amateurre japanese rock singalong \\u30ea\\u30e7\\u30af\\u30b7\\u30e3\\u30ab ryokushaka \\u7a74\\u898b\\u771f\\u543e shingo anami \\u5c0f\\u6797\\u58f1\\u8a93 kobayashi issei \\u5ddd\\u53e3\\u572d\\u592a keita kawaguchi \\u9577\\u5c4b\\u6674\\u5b50 nagaya haruko peppe\",\"availability\":{\"download_disabled\":false,\"more_information\":null},\"ratings\":[0,1,1,0,0,0,0,0,1,1,12]},\"current_user_playcount\":0,\"failtimes\":{\"fail\":[0,0,0,0,9,9,9,0,0,0,9,0,9,0,9,0,0,0,9,0,0,0,0,0,0,0,18,0,9,0,18,9,0,0,0,0,0,0,0,0,0,0,0,0,0,9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],\"exit\":[0,0,0,0,54,9,18,18,45,0,18,0,18,0,9,18,0,9,9,9,9,0,0,0,9,18,9,18,18,18,9,18,0,0,0,0,27,9,0,0,0,9,0,9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,9,0,9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,9,9,9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]},\"max_combo\":1139,\"owners\":[{\"id\":7326908,\"username\":\"Amateurre\"}]}"
	// " < Vim tokenizer breaks because data string is 3k long

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		beatmapPath := fmt.Sprintf("/api/v2/beatmaps/%d", testBeatmapId)
		if r.URL.Path != beatmapPath {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	}))
	defer mockServer.Close()

	client := NewClient(mockServer.URL)

	_, err := client.GetBeatmap(testBeatmapId)
	if err != nil {
		t.Errorf("error getting beatmap %d: %v",
			testBeatmapId, err)
	}
}
