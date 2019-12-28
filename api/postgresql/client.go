package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq" // Postgres driver dependency.
)

const (
	// databaseDriver defines the driver
	// used for establishing connection
	// to remote database endpoint.
	databaseDriver = "postgres"
)

// ClientConfig defines the Client configuration.
type ClientConfig struct {
	Endpoint string
	Port     int
	Database string
	Username string
	Password string
	SslMode  string
}

// Client defines a database client.
type Client struct {
	config   *ClientConfig
	database *sqlx.DB
}

// NewClient returns new storage client
// from the provided parameters.
func NewClient(ctx context.Context, config *ClientConfig) (*Client, error) {
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		config.Endpoint,
		config.Port,
		config.Database,
		config.Username,
		config.Password,
		config.SslMode,
	)

	db, err := sqlx.ConnectContext(ctx, databaseDriver, connStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Postgres: %s", err)
	}
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	client := Client{
		config:   config,
		database: db,
	}

	return &client, nil
}

// DropDatabase drops database mathing the provided parameters.
func (c *Client) DropDatabase(ctx context.Context, name string) (*sql.Result, error) {

	// Using fmt.Sprintf to create the drop database statement
	// as currently, the placeholders are not supported.
	//
	// See https://github.com/lib/pq/issues/694 for more details.
	query := fmt.Sprintf("DROP DATABASE %s ;", name)

	result, err := c.database.ExecContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute drop query: %s", err)
	}

	return &result, nil
}

// CreateDatabase creates new database with the provided parameters.
func (c *Client) CreateDatabase(ctx context.Context, name string) (*sql.Result, error) {

	// Using fmt.Sprintf to create the create database statement
	// as currently, the placeholders are not supported.
	//
	// See https://github.com/lib/pq/issues/694 for more details.
	query := fmt.Sprintf("CREATE DATABASE %s ;", name)

	result, err := c.database.ExecContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute create query: %s", err)
	}

	return &result, nil
}

// Execute executes query on a remote database endpoint.
func (c *Client) Execute(ctx context.Context, query string) (sql.Result, error) {
	return c.database.ExecContext(ctx, query)
}

// Close closes the remote database endpoint connection.
func (c *Client) Close() error {
	return c.database.Close()
}

// String implements fmt.Stringer interface.
func (c *Client) String() string {
	databaseInfo := fmt.Sprintf("host=%s port=%d dbname=%s",
		c.config.Endpoint,
		c.config.Port,
		c.config.Database,
	)

	return databaseInfo
}
