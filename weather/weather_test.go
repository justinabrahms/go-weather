package weather

import (
	"testing"
)

func TestApiKey(t *testing.T) {
	c := CreateClient("test-api-key")
	if (c.getApiUrlForFeature("hourly") != 	"http://api.wunderground.com/api/test-api-key/hourly/q") {
		t.Errorf("Hourly feature isn't the correct URL; result: %s", c.getApiUrlForFeature("hourly"))
	}
}