package main

import "flag"

var configFile = flag.String("f", "./config/config.yaml", "set config file which viper will loading.")

func main() {
	flag.Parse()

	app, err := InitializeApp(*configFile)
	if err != nil {
:		panic(err)
	}

	err = app.Start()
	if err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
