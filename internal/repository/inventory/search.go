package inventory

import (
	"context"
	"fmt"
	"log"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/repository/inventory/model"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) GetInventoryByID(ctx context.Context, tx pgx.Tx, ownerID int64) (model.Inventory, error) {
	builder := sq.Select(
			repository.MerchTitleColumn,
			repository.InventoryQuantityColumn,
		).PlaceholderFormat(sq.Dollar).
		From(repository.InventoryTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
				repository.MerchTableName,
				repository.InventoryTableName,
				repository.InventoryIdMerchColumn,
				repository.MerchTableName,
				repository.MerchIdColumn,
			),
		).Where(sq.Eq{
			fmt.Sprintf("%s.%s", 
			repository.InventoryTableName,
			repository.InventoryIdUserColumn,
		): ownerID})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errs.NewDBError(err.Error(), err)
	}
	log.Println(query, args)
	var rows pgx.Rows
	if tx != nil {
		rows, err = tx.Query(ctx, query, args...)
	} else {
		rows, err = r.db.Query(ctx, query, args...)
	}
	if err != nil {
		return nil, errs.NewDBError(err.Error(), err)
	}
	inventory := model.Inventory{}
	var title string
	var quantity int
	for rows.Next() {
		err = rows.Scan(
			&title,
			&quantity,
		)
		inventory[title] = quantity
	}
	return inventory, nil
}