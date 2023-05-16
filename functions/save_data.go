package functions

import (
	"context"
	"fmt"

	"github.com/eufelipemateus/store-blog-posts/database"
	"github.com/eufelipemateus/store-blog-posts/database/query"
	"github.com/eufelipemateus/store-blog-posts/utils"
	"gorm.io/gorm/clause"
)

func SaveData() {
	ctx := context.Background()
	q := query.Use(database.DB)
	tx := q.Begin()
	txCtx := tx.WithContext(ctx)

	err := txCtx.Post.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(Posts, len(Posts))
	utils.Check(err)

	tx.Commit()
}
