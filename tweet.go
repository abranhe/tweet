package tweet

import (
	"net/http"
	"net/url"
	"strings"
)

type TwitterCredentials struct {
	Token string
	Ct0   string
	Auth  string
}

func Tweet(message string, credentials TwitterCredentials) error {
	endpoint, err := url.Parse("https://api.twitter.com/1.1/statuses/update.json")

	if err != nil {
		return err
	}

	body := url.Values{}

	body.Add("skip_status", "1")
	body.Add("include_cards", "1")
	body.Add("include_can_dm", "1")
	body.Add("include_blocking", "1")
	body.Add("include_mute_edge", "1")
	body.Add("include_blocked_by", "1")
	body.Add("include_reply_count", "1")
	body.Add("include_followed_by", "1")
	body.Add("include_want_retweets", "1")
	body.Add("include_can_media_tag", "1")
	body.Add("include_profile_interstitial_type", "1")
	body.Add("batch_mode", "off")
	body.Add("trim_user", "false")
	body.Add("tweet_mode", "extended")
	body.Add("cards_platform", "Web-12")
	body.Add("include_quote_count", "true")
	body.Add("simple_quoted_tweet", "true")
	body.Add("include_ext_alt_text", "true")
	body.Add("include_ext_media_color", "true")
	body.Add("auto_populate_reply_metadata", "false")
	body.Add("include_ext_media_availability", "true")

	body.Add("status", message)

	endpoint.RawQuery = body.Encode()

	res, err := http.NewRequest("POST", endpoint.String(), strings.NewReader(body.Encode()))

	if err != nil {
		return err
	}

	res.Header.Set("accept", "*/*")
	res.Header.Set("accept-encoding", "gzip, deflate, br")
	res.Header.Set("accept-language", "en-US,en;q=0.9")
	res.Header.Set("Authorization", "Bearer "+credentials.Token)
	res.Header.Set("authority", "twitter.com")
	res.Header.Set("content-length", "504")
	res.Header.Set("content-type", "application/x-www-form-urlencoded")
	res.Header.Set("cookie", "auth_token="+credentials.Auth+"; ct0="+credentials.Ct0)
	res.Header.Set("origin", "https://twitter.com")
	res.Header.Set("referer", "https://twitter.com/compose/tweet")
	res.Header.Set("sec-fetch-dest", "empty")
	res.Header.Set("sec-fetch-mode", "cors")
	res.Header.Set("sec-fetch-site", "same-origin")
	res.Header.Set("sec-gpc", "1")
	res.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	res.Header.Set("x-csrf-token", credentials.Ct0)
	res.Header.Set("x-twitter-active-user", "yes")
	res.Header.Set("x-twitter-auth-type", "OAuth2Session")
	res.Header.Set("x-twitter-client-language", "en")

	defer res.Body.Close()

	return nil
}
