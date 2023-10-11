package plugin

import (
	"github.com/corazawaf/coraza/v3/experimental/plugins"
)

func init() {
	plugins.RegisterOperator("geoLookupEmbedded", newGeolookup)
}
