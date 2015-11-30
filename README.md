## What Is It?
The `transfer` file will take a csv file (in our case collection.csv) which should contain 2 columns (track, artist) and search Spotify then add it to your collection.

## How do I use it?
1. Clone this repo
2. Add in a new CSV which contains 2 columns (track, artist). Altenatively you can modify the code to work with your format.
3. Obtain simplejson by running `go get github.com/bitly/go-simplejson`
4. Obtain a Spotify OAuth token for saving tracks. You can obtain one simply by going to (https://developer.spotify.com/web-api/console/put-current-user-saved-tracks/).
5. Add the new token into the `transfer.go` file.
6. Run `go build transfer.go`
7. Run the `transfer` file and wait. This could take some time depending on how large your library is.
