package neo4j

import (
	"context"
	"hackathon-backend/utils/logger"
	"hackathon-backend/utils/variables"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driver neo4j.DriverWithContext

func Init() neo4j.DriverWithContext {
	ctx := context.Background()
	dbUri := variables.MustGetenv("NEO4J_URI")
	dbUser := variables.MustGetenv("NEO4J_USERNAME")
	dbPassword := variables.MustGetenv("NEO4J_PASSWORD")

	// Create a new driver instance
	var err error
	driver, err = neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		logger.Error(err)
	}

	// Check connection
	if err = driver.VerifyConnectivity(ctx); err != nil {
		logger.Error(err)
	}

	logger.Info("Connected to Neo4j")

	return driver
}

func Exec(query string, params map[string]interface{}) ([]*neo4j.Record, error) {
	ctx := context.Background()

	// Run cypher query
	result, err := neo4j.ExecuteQuery(ctx, driver, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		logger.Error("neo4j query error: ", err)
		return nil, err
	}

	return result.Records, nil
}
