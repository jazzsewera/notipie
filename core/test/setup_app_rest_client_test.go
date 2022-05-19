package test

import (
	"encoding/json"
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/netutil"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"
)

type appRestClient struct {
	appID   string
	t       testing.TB
	pushURL url.URL
	cl      *http.Client
}

type appIDRes struct {
	AppID string `json:"appId"`
}

func newAppRestClient(t testing.TB, port int) *appRestClient {
	return &appRestClient{
		t:       t,
		pushURL: pushURL(port),
		cl:      &http.Client{Timeout: 200 * time.Millisecond},
	}
}

func (c *appRestClient) pushNotification(notification model.AppNotification) {
	c.t.Helper()

	status, res, err := netutil.PostReq(c.cl, c.pushURL, "application/json", notification.ToJSON())
	if err != nil {
		c.t.Fatal(err)
	}

	assertStatusCreated(c.t, status)

	c.appID, err = appIdFromRes(res)
	if err != nil {
		c.t.Fatal(err)
	}

	c.t.Log("appRestClient: pushNotification: success\tappID:", c.appID)
}

func assertStatusCreated(t testing.TB, statusCode int) {
	assert.Equal(t, http.StatusCreated, statusCode)
}

func appIdFromRes(r io.Reader) (string, error) {
	a := appIDRes{}
	d := json.NewDecoder(r)
	err := d.Decode(&a)
	if err != nil {
		return "", fmt.Errorf("unmarshal app response: %s", err)
	}
	return a.AppID, nil
}
