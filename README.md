# osugoi!
osugoi! is a small sdk for public osu! API written in Go language.

## Usage
Currently the only way to use this module is by cloning this repository
and including the following in your go.mod
```
require github.com/Kirshoo/osugoi v0.0.0
replace github.com/Kirshoo/osugoi => path/to/cloned/repo
```

## API Coverage
### Yet to be covered
- Changelog
- Chat
- Comments
- Events
- Home
- Matches
- Multiplayer
- News
- Notification
- Ranking
- Wiki

### Authentication
[x] Authorization Code Grant  
[x] Authorization Token Refresh  
[x] Client Credentials Grant  

### Beatmap Packs
[x] Get Beatmap Packs  
[x] Get Beatmap Pack (specified using pack id)  
[x] Optional query parameters  

### Beatmaps
[x] Lookup Beatmap  
[x] Optional Lookup parameters  
[x] Get a User Beatmap score  
[x] Get a User Beatmap scores  
[x] Optional query parameters for user beatmap scores  
[x] Get top Beatmap scores  
[x] Get Beatmaps  
[x] Optional query parameters for Beatmaps  
[x] Get Beatmap (specified by beatmap id)  
[x] Get Beatmap Attributes  

### Beatmapset Discussions
[ ] Get Discussion Posts  
[ ] Get Discussion Votes  
[ ] Get Discussions  
[ ] Optional query parameters  

### Beatmapsets
[ ] Search Beatmapset  
[ ] Cursor String optional parameter  
[ ] Lookup Beatmapset  
[ ] Get Beatmapset  
[ ] Download beatmapset  

### OAuth tokens
[x] Revoke current token

### Scores
[x] Get Scores  
[ ] Reorder pinned score  
[ ] Unpin score  
[ ] Pin score  
[x] Optional query parameters  

### Users
[ ] Get Own Data  
[ ] Get User Kudosu  
[ ] Get User Scores  
[ ] Get User Beatmaps  
[ ] Get User recent activity  
[ ] Get User  
[ ] Get Users  
[ ] Optional query parameters  
