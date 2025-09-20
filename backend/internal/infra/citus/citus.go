package infra_citus

import (
	"context"
	"fmt"

	"github.com/dijer/otus-highload/backend/internal/config"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
)

func InitCitus(ctx context.Context, dbRouter *infra_database.DBRouter, nodes []config.CitusNode, dbCfg config.DatabaseConf) error {
	if _, err := dbRouter.Master.ExecContext(ctx,
		`CREATE EXTENSION IF NOT EXISTS citus`,
	); err != nil {
		return err
	}

	if _, err := dbRouter.Master.ExecContext(ctx,
		`SELECT citus_set_coordinator_host($1)`, "db"); err != nil {
		return fmt.Errorf("cannot set coordinator host: %w", err)
	}

	rows, err := dbRouter.Master.QueryContext(ctx, "SELECT nodename, nodeport FROM pg_dist_node")
	if err != nil {
		return err
	}
	defer rows.Close()

	registered := map[string]struct{}{}
	for rows.Next() {
		var name string
		var port int
		if err := rows.Scan(&name, &port); err != nil {
			return err
		}
		registered[fmt.Sprintf("%s:%d", name, port)] = struct{}{}
	}

	for _, node := range nodes {
		key := fmt.Sprintf("%s:%d", node.Host, node.Port)
		if _, ok := registered[key]; !ok {
			_, err := dbRouter.Master.ExecContext(ctx,
				`SELECT * FROM citus_add_node($1, $2)`,
				node.Host, node.Port)
			if err != nil {
				return fmt.Errorf("cannot add worker node: %w", err)
			}
		}
	}

	return nil
}
