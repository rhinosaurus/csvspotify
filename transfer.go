package main

import (
	"encoding/csv"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var key = "" // <-- Your key from Spotify goes here

func addSongToSpotify(id string) {
	url := "https://api.spotify.com/v1/me/tracks?ids=" + id
	req, err := http.NewRequest("PUT", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erm, something went wrong during save...")
		fmt.Println(resp)
		return
	}
	fmt.Println("Song Added!")
}

func searchSpotify(trackQuery string, artist string) {
	// First we retrieve the song
	resp, err := http.Get("https://api.spotify.com/v1/search?q=" + trackQuery + "&type=track")
	if err != nil {
		fmt.Println("Could not hit it...")
		return
	} else {
		bdy, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		js, err := simplejson.NewJson(bdy)
		if err != nil {
			return
		}

		// Iterate over the tracks
		tracks := js.Get("tracks").Get("items")
		ids := ""
		for i, song := range tracks.MustArray() {
			track, _ := song.(map[string]interface{})
			trackArtist := js.Get("tracks").Get("items").GetIndex(i).Get("artists").GetIndex(0).Get("name").MustString()
			if trackArtist == artist {
				fmt.Println("Found " + artist + " - " + trackQuery)
				ids += track["id"].(string) + ","
			}
		}

		if len(ids) > 0 {
			time.Sleep(time.Second * time.Duration(2))
			addSongToSpotify(ids)
		}
	}
}

func main() {
	file, err := os.Open("collection.csv")
	if err != nil {
		fmt.Println("Could not open CSV")
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	lineCount := 0

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Having issues reading the CSV...but will continue...")
			continue
		}

		searchSpotify(record[0], record[1])
		lineCount += 1
	}
}
