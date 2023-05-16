package utils

import (
	"strings"

	config "github.com/eufelipemateus/store-blog-posts/configs"
)

func PublicUploads(urlUploads string) string {
	return strings.Replace(
		urlUploads,
		config.GetApp().URL_WORDPRESS+config.GetApp().WP_UPLOADS,
		config.GetApp().URL_PUBLIC_UPLOADS,
		1,
	)

}
