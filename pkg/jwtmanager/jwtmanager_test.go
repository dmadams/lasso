package jwtmanager

import (
	"testing"

	"github.com/dmadams/lasso/pkg/cfg"
	"github.com/dmadams/lasso/pkg/structs"
	// log "github.com/Sirupsen/logrus"
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	u1 = structs.User{
		Email:         "test@testing.com",
		EmailVerified: true,
		Name:          "Test Name",
	}

	lc LassoClaims
)

func init() {
	// log.SetLevel(log.DebugLevel)

	lc = LassoClaims{
		u1.Email,
		Sites,
		StandardClaims,
	}
}

func TestCreateUserTokenStringAndParseToEmail(t *testing.T) {

	uts := CreateUserTokenString(u1)
	assert.NotEmpty(t, uts)

	utsParsed, err := ParseTokenString(uts)
	if utsParsed == nil || err != nil {
		t.Error(err)
	} else {
		log.Debugf("test parsed token string %v", utsParsed)
		ptemail, _ := PTokenToEmail(utsParsed)
		assert.Equal(t, u1.Email, ptemail)
	}

}

func TestClaims(t *testing.T) {
	cfg.ParseConfig()

	log.Debugf("jwt config %s %d", string(cfg.Cfg.JWT.Secret), cfg.Cfg.JWT.MaxAge)
	assert.NotEmpty(t, cfg.Cfg.JWT.Secret)
	assert.NotEmpty(t, cfg.Cfg.JWT.MaxAge)

	// now := time.Now()
	// d := time.Duration(ExpiresAtMinutes) * time.Minute
	// log.Infof("lc d %s", d.String())
	// lc.StandardClaims.ExpiresAt = now.Add(time.Duration(ExpiresAtMinutes) * time.Minute).Unix()
	// log.Infof("lc expiresAt %d", now.Unix()-lc.StandardClaims.ExpiresAt)
	uts := CreateUserTokenString(u1)
	utsParsed, _ := ParseTokenString(uts)
	assert.True(t, SiteInToken("naga.bnf.net", utsParsed))

}
