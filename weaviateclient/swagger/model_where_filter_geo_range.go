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

// filter within a distance of a georange
type WhereFilterGeoRange struct {
	GeoCoordinates *GeoCoordinates              `json:"geoCoordinates,omitempty"`
	Distance       *WhereFilterGeoRangeDistance `json:"distance,omitempty"`
}
