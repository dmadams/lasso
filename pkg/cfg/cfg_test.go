package cfg

import (
	"testing"
	// "github.com/dmadams/lasso/pkg/structs"
	// log "github.com/Sirupsen/logrus"
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	cfg CfgT
)

func init() {
	// log.SetLevel(log.DebugLevel)
}

func TestConfigParsing(t *testing.T) {

	UnmarshalKey("lasso", &cfg)
	log.Debugf("cfgPort %d", cfg.Port)
	log.Debugf("cfgDomains %s", cfg.Domains[0])

	assert.Equal(t, cfg.Port, 9090)
	assert.Equal(t, cfg.Cookie.Name, "bnfSSO")

	assert.NotEmpty(t, cfg.JWT.MaxAge)

}
