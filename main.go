package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"

	glack "github.com/get-go/glack"
)

type quote struct {
	Quote string
}

var ronurl = "http://ron-swanson-quotes.herokuapp.com/v2/quotes"
var profileurl = flag.String("profile-url", "https://twitter.com/ronswansonsays/profile_image?size=normal", "Ron's profile picture URL. Can be slack emoji example ':ron-swanson:'.")
var token = flag.String("token", os.Getenv("SLACK_TOKEN"), "Slack API token to use. Defaults to 'SLACK_TOKEN' environment variable.")
var channel = flag.String("channel", "#random", "Slack channel to post to")

func main() {
	flag.Parse()

	g := glack.New(*token)

	r, _ := http.Get(ronurl)
	defer r.Body.Close()

	var q quote
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &q)

	g.Send(&glack.Message{Channel: *channel, Icon: *profileurl, Username: "Ron Swanson", Message: q.Quote})
}
