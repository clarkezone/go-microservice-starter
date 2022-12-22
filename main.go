/*
Copyright © 2022 clarkezone
*/
package main

import (
	"github.com/clarkezone/gomicroservicestarter/cmd"
	clarkezoneLog "github.com/clarkezone/gomicroservicestarter/pkg/log"
	"github.com/sirupsen/logrus"
)

func main() {
	clarkezoneLog.Init(logrus.WarnLevel)
	cmd.Execute()
}
