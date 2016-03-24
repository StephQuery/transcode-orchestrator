package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nytm/video-transcoding-api/provider"
)

// swagger:route GET /providers providers listProviders
//
// Describe available providers in the API, including their name, capabilities
// and health state.
//
//     Responses:
//       200: listProviders
//       500: genericError
func (s *TranscodingService) listProviders(r *http.Request) gizmoResponse {
	return newListProvidersResponse(provider.ListProviders())
}

// swagger:route GET /providers/{name} providers getProvider
//
// Describe available providers in the API, including their name, capabilities
// and health state.
//
//     Responses:
//       200: provider
//       404: providerNotFound
//       500: genericError
func (s *TranscodingService) getProvider(r *http.Request) gizmoResponse {
	var params getProviderInput
	params.loadParams(mux.Vars(r))
	description, err := provider.DescribeProvider(params.Name, s.config)
	switch err {
	case nil:
		return newGetProviderResponse(description)
	case provider.ErrProviderNotFound:
		return newProviderNotFoundResponse(err)
	default:
		return newErrorResponse(err)
	}
}