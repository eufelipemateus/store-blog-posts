package functions

import (
	"github.com/eufelipemateus/store-blog-posts/interfaces"
	"github.com/eufelipemateus/store-blog-posts/utils"
)

var Posts []*interfaces.Post

func getAuthor(userId int) interfaces.User {
	user, _, _, err := client.Users().Get(userId, nil)
	utils.Check(err)
	return interfaces.User{
		Name:        user.Name,
		Url:         user.URL,
		Description: user.Description,
		Link:        user.Link,
		Slug:        user.Slug,
		AvatarUrl:   user.AvatarURL,
		Nickname:    user.Username,
	}
}
func getMedia(idMedia int) interfaces.Media {

	if idMedia == 0 {
		return interfaces.Media{}
	}

	media, _, _, err := client.Media().Get(idMedia, nil)
	utils.Check(err)
	return interfaces.Media{
		Title:     media.Title.Rendered,
		SourceUrl: media.SourceURL,
	}
}

func getTags(tagsId []int) []interfaces.Term {
	var tags []interfaces.Term
	for i := range tagsId {
		resp, _, _, err := client.Terms().Tag().Get(tagsId[i], nil)
		utils.Check(err)
		tag := interfaces.Term{
			Name: resp.Name,
			Slug: resp.Slug,
		}
		tags = append(tags, tag)
	}
	return tags
}

func getParents(parent int) *interfaces.Term {
	var first *interfaces.Term
	var parentAtual *interfaces.Term = nil

	for parent > 0 {
		respParent, _, _, err := client.Terms().Category().Get(parent, nil)
		utils.Check(err)
		parentValue := interfaces.Term{
			Name:   respParent.Name,
			Slug:   respParent.Slug,
			Parent: parentAtual,
		}

		if first == nil {
			first = &parentValue
		}

		parentAtual = &parentValue
		parent = respParent.Parent
	}
	return first
}

func getCategories(categoriesIDs []int) []interfaces.Term {
	var categories []interfaces.Term
	for i := range categoriesIDs {
		resp, _, _, err := client.Terms().Category().Get(categoriesIDs[i], nil)
		utils.Check(err)
		category := interfaces.Term{
			Name:   resp.Name,
			Slug:   resp.Slug,
			Parent: getParents(resp.Parent),
		}
		categories = append(categories, category)
	}
	return categories
}

func GetPosts() {
	posts, _, _, err := client.Posts().List(interfaces.PostParams{Status: "publish"})
	utils.Check(err)
	for i := range posts {
		p := interfaces.Post{
			ID:            uint(posts[i].ID),
			Created:       posts[i].Date,
			Update:        posts[i].Modified,
			Slug:          posts[i].Slug,
			Type:          posts[i].Type,
			Link:          posts[i].Link,
			Title:         posts[i].Title.Rendered,
			Content:       posts[i].Content.Rendered,
			Excerpt:       posts[i].Excerpt.Rendered,
			Author:        getAuthor(posts[i].Author),
			FeaturedMedia: getMedia(posts[i].FeaturedMedia),
			Tags:          getTags(posts[i].Tags),
			Categories:    getCategories(posts[i].Categories),
			Metadata:      posts[i].YoastHead,
		}
		Posts = append(Posts, &p)
	}
}
