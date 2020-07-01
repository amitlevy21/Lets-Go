package track

import (
	"fmt"
	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	proto "github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

func Feed(url string) gtfs.FeedMessage {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
	if err != nil {
		log.Fatal(err)
	}
	return feed
}

func getKey() string {
	viper.SetConfigFile("../.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	key, ok := viper.Get("API_KEY").(string)
	if !ok {
		log.Fatalf("Error while getting api key")
	}
	return key
}
func Entities() []*gtfs.FeedEntity {
	key := getKey()
	feed := Feed(fmt.Sprintf("http://datamine.mta.info/mta_esi.php?key=%s&feed_id=1", key))
	return feed.Entity
}

func Delays() []*gtfs.TripDescriptor {
	delays := []*gtfs.TripDescriptor{}
	for _, entity := range Entities() {
		tripUpdate := entity.TripUpdate
		if tripUpdate != nil {
			delay := tripUpdate.Delay
			trip := tripUpdate.Trip
			if delay != nil && *delay != 0 {
				delays = append(delays, trip)
			}
		}
	}
	return delays
}
