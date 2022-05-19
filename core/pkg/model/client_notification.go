package model

import (
	"encoding/json"
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/timeformat"
)

type ClientNotification struct {
	HashableNetNotification
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Read      bool   `json:"read,omitempty"`
}

func (c ClientNotification) ToJSON() ([]byte, error) {
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("marshal ClientNotification: %s", err)
	}
	return jsonBytes, nil
}

func ClientNotificationFromJSON(jsonBytes []byte) (ClientNotification, error) {
	clientNotification := ClientNotification{}
	err := json.Unmarshal(jsonBytes, &clientNotification)
	if err != nil {
		return ClientNotification{}, err
	}
	if !clientNotification.validate() {
		return ClientNotification{}, fmt.Errorf(NotEnoughInfoInNotificationErrorMessage)
	}
	return clientNotification, nil
}

func ClientNotificationFromDomain(n domain.Notification) ClientNotification {
	timestamp := n.Timestamp.Format(timeformat.RFC3339Milli)
	return ClientNotification{
		HashableNetNotification: HashableNetNotification{
			AppName:    n.App.Name,
			AppID:      n.App.ID,
			AppImgURI:  n.App.IconURI,
			Title:      n.Title,
			Subtitle:   n.Subtitle,
			Body:       n.Body,
			ExtURI:     n.ExtURI,
			ReadURI:    n.ReadURI,
			ArchiveURI: n.ArchiveURI,
		},
		ID:        n.ID,
		Timestamp: timestamp,
		Read:      false,
	} // TODO: implement urgency
}

func (c ClientNotification) validate() bool {
	if c.AppName == "" || c.Title == "" || c.Timestamp == "" {
		return false
	}
	return true
}
