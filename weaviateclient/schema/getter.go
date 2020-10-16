package schema

import (
	"context"
	"github.com/semi-technologies/weaviate-go-client/weaviateclient/clienterrors"
	"github.com/semi-technologies/weaviate-go-client/weaviateclient/connection"
	clientModels "github.com/semi-technologies/weaviate-go-client/weaviateclient/models"
	"net/http"
)

// Getter builder to get the current schema loaded in weaviate
type Getter struct {
	connection *connection.Connection
}

// Do get and return the weaviate schema
func (sg *Getter) Do (ctx context.Context) (*clientModels.SchemaDump, error){
	responseData, err := sg.connection.RunREST(ctx, "/schema", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	if responseData.StatusCode == 200 {
		var fullSchema clientModels.SchemaDump
		decodeErr := responseData.DecodeBodyIntoTarget(&fullSchema)
		if decodeErr != nil {
			return nil, decodeErr
		}
		return &fullSchema, nil
	}
	return nil, clienterrors.NewUnexpectedStatusCodeError(responseData.StatusCode, string(responseData.Body))
}
