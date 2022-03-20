package grid_test

import (
	"github.com/blazejsewera/notipie/core/internal/model"
)

func newTestAppNotification() model.AppNotification {
	return model.AppNotification{
		HashableNetNotification: model.HashableNetNotification{
			AppName:    "TestApp",
			Title:      "Test AppNotification",
			Subtitle:   "Test Subtitle",
			Body:       "Test Body",
			ExtURI:     "testExtUri",
			ReadURI:    "testReadUri",
			ArchiveURI: "testArchiveUri",
		},
		Timestamp: "2022-02-19T11:00:00.000Z",
		ID:        "1",
		Read:      false,
	}
}

func newTestClientNotification() model.ClientNotification {
	an := newTestAppNotification()
	return model.ClientNotification{
		HashableNetNotification: an.HashableNetNotification,
		ID:                      an.ID,
		Timestamp:               an.Timestamp,
		Read:                    an.Read,
	}
}