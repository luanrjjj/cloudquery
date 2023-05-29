package resources

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/luanrjjj/cq-source-xqcd/client"
	"github.com/luanrjjj/cq-source-xqcd/internal/xkcd"
)

func TestComicsTable(t *testing.T) {
	var comic xkcd.Comic
	if err := faker.FakeObject(&comic); err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		d, _ := json.Marshal(comic)
		_, _ = w.Write(d)
	}))
	defer ts.Close()

	client.TestHelper(t, ComicsTable(), ts)
}
