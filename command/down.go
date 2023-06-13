package command

import (
	"os"
	"os/exec"
)

const GO_JENKINS_DIR string = "/tmp/go-jenkins"

func Down() error {
	err := os.Chdir(DOCKER_DIR)
	if err != nil {
		return err
	}

	dockerCommand := exec.Command("docker", "compose", "down")

	dockerCommand.Stdout = os.Stdout
	dockerCommand.Stderr = os.Stderr

	if err := dockerCommand.Run(); err != nil {
		return err
	}

	if err := cleanUp(); err != nil {
		return err
	}

	return nil
}

func cleanUp() error {
	if err := os.RemoveAll(GO_JENKINS_DIR); err != nil {
		return err
	}

	return nil
}
