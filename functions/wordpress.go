package functions

import (
	"github.com/eufelipemateus/go-wordpress"
	config "github.com/eufelipemateus/store-blog-posts/configs"
)

const WPAPI = "/wp-json"

const WPV2 = "/wp/v2"

var client *wordpress.Client

func GetClient() {
	client = wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: config.GetApp().URL_WORDPRESS + WPAPI + WPV2,
	})
}
