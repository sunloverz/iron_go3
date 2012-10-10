package config_test

import (
	"github.com/iron-io/go.iron/config"
	. "github.com/sdegutis/go.bdd"
	"testing"
)

func init() {
	defer PrintSpecReport()
	Describe("gets config", func() {
		It("gets default configs", func() {
			s := config.Config("iron_undefined")
			Expect(s.Host, ToEqual, "undefined-aws-us-east-1.iron.io")
		})
	})
}

func TestEverything(t *testing.T) {}
