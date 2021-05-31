package config

import (
	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Config struct {
	Data DataBases `hlc:"data"`
}

type DataBases struct {
	UserName string `hcl:"username"`
	Password string `hcl:"password"`
	DataPass string `hcl:"datapass"`
}

func LoadConfigHCL(filePath string, target interface{}) (err error) {
	defer func() { err = errors.Wrap(err, "config.LoadConfigHCL") }()
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = errors.Wrap(err, `config: read`)
		return
	}
	err = errors.Wrap(hcl.Unmarshal(b, target), `config: parse`)
	return
}