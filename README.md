osugoi! is a small wrapper around public osu! API written in Go language.

## Usage
In order to use osugoi!, you will have to download:  
```
go get github.com/Kirshoo/osugoi
```  
  
After that, you will be able to import and use it inside your project.

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
- Revoke OAuth token
- Ranking
- Wiki

### Authentication
[ ] Authorization Code Grant  
[ ] Authorization Token Refresh  
[x] Client Credentials Grant  

### Beatmap Packs
[x] Get Beatmap Packs  
[x] Get Beatmap Pack (specified using pack id)  
[ ] Optional query parameters  

### Beatmaps
[x] Lookup Beatmap  
[x] Optional Lookup parameters  
[x] Get a User Beatmap score  
[x] Get a User Beatmap scores  
[ ] Optional query parameters for user beatmap scores  
[ ] Get top Beatmap scores  
[ ] Get Beatmaps  
[ ] Optional query parameters for Beatmaps  
[x] Get Beatmap (specified by beatmap id)  
[ ] Get Beatmap Attributes  

### Beatmapset Discussions
[ ] Get Discussion Posts  
[ ] Get Discussion Votes  
[ ] Get Discussions  
[ ] Optional query parameters  

### Beatmapsets
[x] Search Beatmapset  
[x] Cursor String optional parameter  
[ ] Lookup Beatmapset  
[ ] Get Beatmapset  
[ ] Download beatmapset  

### Scores
[o] Get Scores  
[ ] Reorder pinned score  
[ ] Unpin score  
[ ] Pin score  
[ ] Optional query parameters  

### Users
[ ] Get Own Data  
[o] Get User Kudosu  
[o] Get User Scores  
[o] Get User Beatmaps  
[ ] Get User recent activity  
[o] Get User  
[o] Get Users  
[ ] Optional query parameters  
