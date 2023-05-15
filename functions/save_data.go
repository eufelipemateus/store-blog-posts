package functions

import (
	"context"

	"github.com/eufelipemateus/store-blog-posts/database"
	"github.com/eufelipemateus/store-blog-posts/database/query"
	"gorm.io/gorm/clause"
)

func SaveData() {
	ctx := context.Background()
	q := query.Use(database.DB)
	tx := q.Begin()
	txCtx := tx.WithContext(ctx)

	txCtx.Post.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(Posts, len(Posts))
	tx.Commit()
}
