package types

import (
	"errors"
	"strings"

	"d3tech.com/platform/utils"
)

type CustomError struct {
	Name       string
	Code       string `yaml:"code"`
	HttpStatus int    `yaml:"httpStatus"`
	Message    string `yaml:"message"`
}

func (c *CustomError) Validate() error {
	if c.Code == "" {
		return errors.New("missing code")
	}
	c.Code = utils.UKC(c.Code)
	c.Name = utils.UCC(strings.ToLower(c.Code))
	if c.HttpStatus == 0 {
		c.HttpStatus = 500
	}
	if c.Message == "" {
		return errors.New("missing message")
	}
	return nil
}

func (c *CustomError) PrettyName() string {
	return utils.UCC(c.Code)
}
