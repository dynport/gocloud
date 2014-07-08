package digitalocean

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {
	Convey("Loading", t, func() {
		cl, e := NewFromEnv()
		So(e, ShouldBeNil)
		So(cl, ShouldNotBeNil)

		images, e := cl.Images()
		So(e, ShouldBeNil)
		So(images, ShouldNotBeNil)

		for _, d := range images.Images {
			b, e := json.Marshal(d)
			So(e, ShouldBeNil)
			t.Logf("%s", string(b))
		}
	})
}
