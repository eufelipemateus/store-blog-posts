package interfaces

import "github.com/eufelipemateus/go-wordpress"

type PostParams struct {
	Slug   string `json:"slug"`
	Status string `json:"status"`
}

type RenderedType struct {
	Rendered string `json:"rendered"`
}

type Post struct {
	ID      uint   `json:"id"  gorm:"primary_key"`
	Created string `json:"date"`
	//	CreatedGmt    string       `json:"date_gmt"`
	//	Guid          string       `json:"guid"`
	Update string `json:"modified"`
	//	UpdateGmt     string       `json:"modified_gmt"`
	Slug          string `json:"slug"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Link          string `json:"link"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	Excerpt       string `json:"excerpt"`
	Author        User   `gorm:"serializer:json"  json:"author"`
	FeaturedMedia Media  `gorm:"serializer:json" json:"faturedMedia"`
	//	CommentStatus string  `json:"comment_status"`
	//	PingStatus    string  `json:"ping_status"`
	//Sticky     bool   `json:"sticky"`
	//Template   string `json:"template"`
	//Format     string `json:"format"`
	//Meta       string `json:"meta"`
	Categories []Term `gorm:"serializer:json" json:"categories"`
	Tags       []Term `gorm:"serializer:json" json:"tags"`
	//Acf        string `json:"acf"`
	Metadata wordpress.YoastHead `gorm:"serializer:json" json:"metadata"`
}

type User struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Slug        string `json:"slug"`
	AvatarUrl   string `json:"avatar_url"`
	Nickname    string `json:"nickname"`
}

type Media struct {
	Title     string `json:"title"`
	SourceUrl string `json:"source_url"`
}

type Term struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Parent *Term  `json:"parent"`
}
