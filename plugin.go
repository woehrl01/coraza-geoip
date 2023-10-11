package plugin

import (
	"github.com/corazawaf/coraza/v3/experimental/plugins"
)

func RegisterGeoDatabase(database []byte) {
	plugins.RegisterOperator("geoLookupEmbedded", newGeolookupCreator(database))
}

