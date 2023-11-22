package trafficapp

import (
	"github.com/ergo-services/ergo/gen"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/buntdb"
	"github.com/vmihailenco/msgpack/v5"
	"testing"
	"time"
	"traffic/cmd"
)

// This test is mainly a experiment on how to properly test an actor in isolation. Note that this test does NOT
// touch the process parameter, need to figure out how to test that properly in some other test, for example to check
// for sending messages.
func TestStoreAndQuery(t *testing.T) {
	// Set up DB in-memory
	db, err := buntdb.Open(":memory:")
	if err != nil {
		t.FailNow()
	}
	defer db.Close()

	testee := createStorage(db)
	serverStatus := testee.HandleInfo(nil, main.PositionUpdate{
		ID:  "abc-123",
		Lon: 12.12,
		Lat: 34.34,
		At:  time.Now().UnixMilli(),
	})
	assert.Equal(t, gen.ServerStatusOK, serverStatus)

	geoPos := struct {
		lon, lat float64
	}{}
	posUpdate := main.PositionUpdate{}
	err = db.View(func(tx *buntdb.Tx) error {
		vehicle, err := tx.Get("vehicle:abc-123")
		assert.NoError(t, err)
		assert.NoError(t, msgpack.Unmarshal([]byte(vehicle), &posUpdate))

		pos, err := tx.Get("fleet:abc-123:pos")
		assert.NoError(t, err)

		geoPos.lon, geoPos.lat, _ = main.coordToLonLat(pos)
		return nil
	})
	assert.NoError(t, err)

	assert.Equal(t, "abc-123", posUpdate.ID)
	assert.Equal(t, geoPos.lon, 12.12)
	assert.Equal(t, geoPos.lat, 34.34)
}
