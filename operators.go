package plugin

import (
	"fmt"
	"net"
	"strconv"

	"github.com/corazawaf/coraza/v3/collection"
	"github.com/corazawaf/coraza/v3/experimental/plugins/plugintypes"
	"github.com/corazawaf/coraza/v3/types/variables"
	"github.com/oschwald/geoip2-golang"
)


type geo struct {
	db 		 *geoip2.Reader
}

func newGeolookup(options plugintypes.OperatorOptions) (plugintypes.Operator, error) {
	data := options.Arguments

	databaseFile := data
	if databaseFile == "" {
		return nil, fmt.Errorf("missing database file")
	}

	db, err := geoip2.Open(databaseFile)
	if err != nil {
		return nil, err
	}

	return &geo{db: db}, nil
}

func (o *geo) Evaluate(tx plugintypes.TransactionState, value string) bool {
	ip := net.ParseIP(value)
	r, err := o.db.City(ip)
	if err != nil {
		return false
	}

	var col collection.Map
	if c, ok := tx.Collection(variables.Geo).(collection.Map); !ok {
		tx.DebugLogger().Error().Msg("collection in getset is not a map")
		return false
	} else {
		col = c
	}
	if col == nil {
		tx.DebugLogger().Error().Msg("collection in getset is nil")
		return false
	}

	col.Set("country_code", []string{r.Country.IsoCode})
	col.Set("country_name", []string{r.Country.Names["en"]})
	col.Set("country_continent", []string{r.Continent.Names["en"]})
	col.Set("region", []string{""})
	col.Set("city", []string{r.City.Names["en"]})
	col.Set("postal_code", []string{r.Postal.Code})
	col.Set("latitude", []string{strconv.FormatFloat(r.Location.Latitude, 'f', 10, 64)})
	col.Set("longitude", []string{strconv.FormatFloat(r.Location.Longitude, 'f', 10, 64)})

	return true
}
