package test

import (
	"os"
	"os/exec"
)

// SetupWeavaite run docker compose up
func SetupWeavaite() error {
	app := "docker-compose"
	arguments := []string{
		"up",
		"-d",
	}
	return command(app, arguments)
}

// TearDownWeavaite run docker-compose down
func TearDownWeavaite() error {
	app := "docker-compose"
	arguments := []string{
		"down",
	}
	return command(app, arguments)
}

func command(app string, arguments []string) error {
	mydir, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd := exec.Command(app, arguments...)
	cmd.Dir = mydir + "/../"
	err = cmd.Start()
	return err
}