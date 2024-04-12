package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	err = errors.New("Bad request")
	logrus.WithError(err).Warn("on reading response body")
	fmt.Println("some")
	log := logrus.New()

}
