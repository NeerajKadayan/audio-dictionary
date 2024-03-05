# audio-dictionary

Clone the repo & plug n play, hit go run main.go in root directory of your terminal.
You can use postman or web browser to view your results

Contains following APIs:

localhost:8080/ping -> health check

localhost:8080/track_search -> search tracks
query param:
key   - q_track
value - can be name of the song you are trying to search for

localhost:8080/get_lyrics
query param:
key   - track_id
value -  an int equivalent for track_id

localhost:8080/top_track
query param:
key   - country
value - name of the country in 2 chars
This api provides region based top track

localhost:8080/artist_info
query param:
key - artist_id
value - an int equivalent of artist_id