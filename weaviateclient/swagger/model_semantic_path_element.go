/*
 * Weaviate
 *
 * Open Source Search Graph (GraphQL/RESTful/P2P)
 *
 * API version: 0.22.17
 * Contact: hello@semi.technology
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// On link on the semantic path chain
type SemanticPathElement struct {
	Concept            string  `json:"concept,omitempty"`
	DistanceToQuery    float32 `json:"distanceToQuery,omitempty"`
	DistanceToResult   float32 `json:"distanceToResult,omitempty"`
	DistanceToPrevious float32 `json:"distanceToPrevious,omitempty"`
	DistanceToNext     float32 `json:"distanceToNext,omitempty"`
}
