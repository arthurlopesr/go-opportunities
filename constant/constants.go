package constant

const (
	APIBasePath = "/api/v1/opportunities"
	APIPort     = ":8080"

	// APIOpeningPath Opening related paths
	APIOpeningPath  = "/opening"
	APIOpeningsPath = "/openings"

	// DatabasePath Database
	DatabasePath = "./db/main.db"

	// ErrToConnectDatabase Error messages
	ErrToConnectDatabase   = "Failed to connect to SQLite database: %v"
	ErrAutoMigrateDatabase = "Failed to migrate SQLite database: %v"
	ErrCreateDatabase      = "Failed to create database directory: %v"

	// ErrCreateOpenings CreateOpenings error
	ErrCreateOpenings   = "Failed to create opening: "
	ErrValidateOpenings = "Failed to validate opening: "
	ErrInvalidJson      = "Invalid JSON format"
)
