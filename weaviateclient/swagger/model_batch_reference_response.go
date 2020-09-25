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

type BatchReferenceResponse struct {
	// Long-form beacon-style URI to identify the source of the cross-ref including the property name. Should be in the form of weaviate://localhost/<kinds>/<uuid>/<className>/<propertyName>, where <kinds> must be one of 'actions', 'things' and <className> and <propertyName> must represent the cross-ref property of source class to be used.
	From string `json:"from,omitempty"`
	// Short-form URI to point to the cross-ref. Should be in the form of weaviate://localhost/things/<uuid> for the example of a local cross-ref to a thing
	To     string                        `json:"to,omitempty"`
	Result *BatchReferenceResponseResult `json:"result,omitempty"`
}
