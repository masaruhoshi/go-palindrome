package main

import "testing"

func TestGetSettingsToReturnDefaultValues(t *testing.T) {
	settings := GetSettings()

	Expect(t, ":8080", settings.ServicePort)
	Expect(t, "localhost", settings.HostName)
	Expect(t, "gopal", settings.DbName)
}