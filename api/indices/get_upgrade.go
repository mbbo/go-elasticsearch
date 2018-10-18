// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// GetUpgradeOption is a non-required GetUpgrade option that gets applied to an HTTP request.
type GetUpgradeOption func(r *transport.Request)

// WithGetUpgradeIndex - a comma-separated list of index names; use "_all" or empty string to perform the operation on all indices.
func WithGetUpgradeIndex(index []string) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradeAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithGetUpgradeAllowNoIndices(allowNoIndices bool) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// GetUpgradeExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type GetUpgradeExpandWildcards int

const (
	// GetUpgradeExpandWildcardsOpen can be used to set GetUpgradeExpandWildcards to "open"
	GetUpgradeExpandWildcardsOpen = iota
	// GetUpgradeExpandWildcardsClosed can be used to set GetUpgradeExpandWildcards to "closed"
	GetUpgradeExpandWildcardsClosed = iota
	// GetUpgradeExpandWildcardsNone can be used to set GetUpgradeExpandWildcards to "none"
	GetUpgradeExpandWildcardsNone = iota
	// GetUpgradeExpandWildcardsAll can be used to set GetUpgradeExpandWildcards to "all"
	GetUpgradeExpandWildcardsAll = iota
)

// WithGetUpgradeExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithGetUpgradeExpandWildcards(expandWildcards GetUpgradeExpandWildcards) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradeIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithGetUpgradeIgnoreUnavailable(ignoreUnavailable bool) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradeErrorTrace - include the stack trace of returned errors.
func WithGetUpgradeErrorTrace(errorTrace bool) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradeFilterPath - a comma-separated list of filters used to reduce the respone.
func WithGetUpgradeFilterPath(filterPath []string) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradeHuman - return human readable values for statistics.
func WithGetUpgradeHuman(human bool) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradeIgnore - ignores the specified HTTP status codes.
func WithGetUpgradeIgnore(ignore []int) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradePretty - pretty format the returned JSON response.
func WithGetUpgradePretty(pretty bool) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// WithGetUpgradeSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithGetUpgradeSourceParam(sourceParam string) GetUpgradeOption {
	return func(r *transport.Request) {
	}
}

// GetUpgrade - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-upgrade.html for more info.
//
// options: optional parameters.
func (i *Indices) GetUpgrade(options ...GetUpgradeOption) (*GetUpgradeResponse, error) {
	req := i.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := i.transport.Do(req)
	return &GetUpgradeResponse{resp}, err
}

// GetUpgradeResponse is the response for GetUpgrade.
type GetUpgradeResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetUpgradeResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}