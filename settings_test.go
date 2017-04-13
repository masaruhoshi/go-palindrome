package main

import "testing"

func TestGetSettingsToReturnDefaultValues(t *testing.T) {
	settings := GetSettings()

	Expect(t, "localhost", settings.HostName)
	Expect(t, "gopal", settings.DbName)
}