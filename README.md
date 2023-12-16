# Play On Request
HTTP server that plays an mp3 file via `paplay` when it gets a request to `/play` and stops with `/stop`.
## Usage
```console
$ ./play-on-request /path/to/filename.mp3 # Starts the server at port 1234
```
```console
$ curl localhost:1234/play # Starts playback
$ curl localhost:1234/stop # Stops playback
```
