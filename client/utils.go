package client

import (
	"context"
	"fmt"
)

const (
	queryLastMigration = `
		SELECT
			*
		FROM
			versioner.migration
		ORDER BY created DESC
		LIMIT 1;
	`

	queryRegisterMigration = `
		INSERT INTO versioner.migration (
			name,
			path,
			applied_by
		) VALUES (
			:name,
			:path,
			:applied_by
		);
	`
)

// LastMigration retrieves the last successfully applied migration name.
func (c *Client) LastMigration(ctx context.Context) (*Migration, error) {
	tx, err := c.database.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to being transaction: %s", err)
	}

	migrations := []Migration{}
	err = tx.SelectContext(ctx, &migrations, queryLastMigration)
	if err != nil {
		return nil, fmt.Errorf("failed select result: %s", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %s", err)
	}

	if len(migrations) == 0 {
		return nil, nil
	}

	return &migrations[0], nil
}

// RegisterMigration registers migration as applied on a remote PostgreSQL endpoint.
func (c *Client) RegisterMigration(ctx context.Context, migration *Migration) error {
	tx, err := c.database.Beginx()
	if err != nil {
		return fmt.Errorf("failed to being transaction: %s", err)
	}

	if _, err = tx.NamedExecContext(ctx, queryRegisterMigration, migration); err != nil {
		if errRowBack := tx.Rollback(); errRowBack != nil {
			return fmt.Errorf("failed to rollback transaction: %s", errRowBack)
		}
		return fmt.Errorf("failed to execute insert query: %s", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %s", err)
	}

	return nil
}
