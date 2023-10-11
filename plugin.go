package plugin

import (
	"github.com/corazawaf/coraza/v3/experimental/plugins"
	"github.com/oschwald/geoip2-golang"
)

func RegisterGeoDatabase(database []byte) {
	db, err := geoip2.FromBytes(database)
	if err != nil {
		panic(err)
	}

	plugins.RegisterOperator("geoLookupEmbedded", newGeolookupCreator(db))
}

