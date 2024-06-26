package generic

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/formancehq/payments/cmd/connectors/internal/connectors"
	"github.com/formancehq/payments/cmd/connectors/internal/connectors/configtemplate"
)

const (
	pageSize             = 100
	defaultPollingPeriod = 2 * time.Minute
)

type Config struct {
	Name          string              `json:"name" yaml:"name" bson:"name"`
	APIKey        string              `json:"apiKey" yaml:"apiKey" bson:"apiKey"`
	Endpoint      string              `json:"endpoint" yaml:"endpoint" bson:"endpoint"`
	PollingPeriod connectors.Duration `json:"pollingPeriod" yaml:"pollingPeriod" bson:"pollingPeriod"`
}

func (c Config) String() string {
	return fmt.Sprintf("endpoint=%s", c.Endpoint)
}

func (c Config) Validate() error {
	if c.Name == "" {
		return ErrMissingName
	}

	return nil
}

func (c Config) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c Config) ConnectorName() string {
	return c.Name
}

func (c Config) BuildTemplate() (string, configtemplate.Config) {
	cfg := configtemplate.NewConfig()

	cfg.AddParameter("name", configtemplate.TypeString, name.String(), false)
	cfg.AddParameter("endpoint", configtemplate.TypeString, "", false)
	cfg.AddParameter("apiKey", configtemplate.TypeString, "", false)
	cfg.AddParameter("pollingPeriod", configtemplate.TypeDurationNs, defaultPollingPeriod.String(), false)

	return name.String(), cfg
}
