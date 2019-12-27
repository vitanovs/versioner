package postgresql

import (
	"fmt"
)

// DatabaseVersion retrieves the current database version.
func (c *Client) DatabaseVersion() (string, error) {
	exists, err := c.functionExists("core", "schema_version")
	if err != nil {
		return "", fmt.Errorf("failed determining if version function exists: %s", err)
	}

	if !exists {
		// No migrations have been applied, yet.
		return "", nil
	}

	query := "SELECT core.schema_version() AS version;"
	var version []string
	err = c.database.Select(&version, query)
	if err != nil {
		return "", fmt.Errorf("failed selecting version function: %s", err)
	}

	return version[0], nil
}

func (c *Client) functionExists(schemaName string, funcName string) (bool, error) {
	query := `
	SELECT
		COUNT(*) = 1
	FROM
		information_schema.routines
	WHERE
		routine_type = 'FUNCTION' AND
		specific_schema = $1 AND
		routine_name = $2 ;
	`

	var result []bool
	err := c.database.Select(&result, query, schemaName, funcName)
	if err != nil {
		return false, err
	}

	return result[0], nil
}
