# osugoi!
osugoi! is a small sdk for public osu! API written in Go language.

## Installation
Currently the only way to install this module is by cloning this repository
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
[o] Search Beatmapset  
[ ] Cursor String optional parameter  
[x] Lookup Beatmapset  
[x] Get Beatmapset  
[!] Download beatmapset (This is osu!lazer specific and you cannot access this endpoint normally)

### Events
[x] List Events
[x] Optional query parameters

### OAuth tokens
[x] Revoke current token

### Scores
[x] Get Scores  
[!] Reorder pinned score (This is osu!lazer specific and you cannot access this endpoint notmally)
[!] Unpin score (This is osu!lazer specific and you cannot access this endpoint notmally)
[!] Pin score (This is osu!lazer specific and you cannot access this endpoint notmally)
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
