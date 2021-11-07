package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregateBuilder(t *testing.T) {

	t.Run("Simple Explore", func(t *testing.T) {
		conMock := &MockRunREST{}

		builder := AggregateBuilder{
			connection: conMock,
		}

		fields := []Field{
			{
				Name: "meta",
				Fields: []Field{
					{
						Name: "count",
					},
				},
			},
		}
		query := builder.WithClassName("Pizza").WithFields(fields).build()

		expected := `{Aggregate{Pizza{meta{count}}}}`
		assert.Equal(t, expected, query)
	})

	t.Run("Group by", func(t *testing.T) {
		conMock := &MockRunREST{}
		builder := AggregateBuilder{
			connection: conMock,
		}

		fields := []Field{
			{
				Name: "groupedBy",
				Fields: []Field{
					{
						Name: "value",
					},
				},
			},
			{
				Name: "name",
				Fields: []Field{
					{
						Name: "count",
					},
				},
			},
		}

		query := builder.WithClassName("Pizza").WithFields(fields).WithGroupBy("name").build()

		expected := `{Aggregate{Pizza(groupBy: "name"){groupedBy{value} name{count}}}}`
		assert.Equal(t, expected, query)
	})

	t.Run("Missuse", func(t *testing.T) {
		conMock := &MockRunREST{}
		builder := AggregateBuilder{
			connection: conMock,
		}
		query := builder.build()
		assert.NotEmpty(t, query, "Check that there is no panic if query is not validly build")
	})

}
