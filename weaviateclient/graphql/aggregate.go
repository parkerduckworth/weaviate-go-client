package graphql

import (
	"context"
	"fmt"
	"github.com/semi-technologies/weaviate-go-client/weaviateclient/connection"
	"github.com/semi-technologies/weaviate-go-client/weaviateclient/paragons"
	"github.com/semi-technologies/weaviate/entities/models"
	"strings"
)

// Aggregate allows the building of an aggregation query
type Aggregate struct {
	connection *connection.Connection
}

// Things aggregate things
func (a *Aggregate) Things() *AggregateBuilder {
	return &AggregateBuilder{
		connection:   a.connection,
		semanticKind: paragons.SemanticKindThings,
	}
}

// Actions aggregate actions
func (a *Aggregate) Actions() *AggregateBuilder {
	return &AggregateBuilder{
		connection:   a.connection,
		semanticKind: paragons.SemanticKindActions,
	}
}

// AggregateBuilder for the aggregate GraphQL query string
type AggregateBuilder struct {
	connection rest
	semanticKind paragons.SemanticKind
	fields string
	className string
}

// WithFields that should be included in the aggregation query e.g. `meta{count}`
func (ab *AggregateBuilder) WithFields(fields string) *AggregateBuilder {
	ab.fields = fields
	return ab
}

// WithClassName that should be aggregated
func (ab *AggregateBuilder) WithClassName(name string) *AggregateBuilder {
	ab.className = name
	return ab
}

// Do execute the aggregation query
func (ab *AggregateBuilder) Do(ctx context.Context) (*models.GraphQLResponse, error) {
	return runGraphQLQuery(ctx, ab.connection, ab.build())
}

// build the query string
func (ab *AggregateBuilder) build() string {
	semanticKind := strings.Title(string(ab.semanticKind))
	return 	fmt.Sprintf("{Aggregate{%v{%v{%v}}}}", semanticKind, ab.className, ab.fields)
}