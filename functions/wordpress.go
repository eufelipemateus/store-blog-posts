package functions

import (
	config "github.com/eufelipemateus/go-get-blog-posts/configs"
	"github.com/eufelipemateus/go-wordpress"
)

const WPAPI = "/wp-json"

const WPV2 = "/wp/v2"

var client *wordpress.Client

func GetClient() {
	client = wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: config.GetApp().URL + WPAPI + WPV2,
	})
}