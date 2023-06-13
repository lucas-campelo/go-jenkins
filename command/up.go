package command

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

const (
	DOCKER_DIR         string = "/tmp/go-jenkins/docker"
	DOCKERFILE_URL     string = "https://raw.githubusercontent.com/lucas-campelo/go-jenkins/main/docker/Dockerfile"
	DOCKER_COMPOSE_URL string = "https://raw.githubusercontent.com/lucas-campelo/go-jenkins/main/docker/docker-compose.yaml"
)

func Up() error {
	// Create docker directory
	if err := os.MkdirAll(DOCKER_DIR, 0755); err != nil {
		return err
	}

	// Create Dockerfile
	if err := createDockerfile(); err != nil {
		return err
	}

	// Create docker-compose.yaml
	if err := createDockerCompose(); err != nil {
		return err
	}

	os.Chdir(DOCKER_DIR)

	dockerCommand := exec.Command("docker", "compose", "up", "--build", "-d")

	dockerCommand.Stdout = os.Stdout
	dockerCommand.Stderr = os.Stderr

	if err := dockerCommand.Run(); err != nil {
		return err
	}

	return nil
}

func createDockerfile() error {
	response, err := http.Get(DOCKERFILE_URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	dockerfileBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	dockerfile, err := os.Create(DOCKER_DIR + "/Dockerfile")
	if err != nil {
		return err
	}

	if _, err = dockerfile.Write(dockerfileBytes); err != nil {
		return err
	}

	return nil
}

func createDockerCompose() error {
	response, err := http.Get(DOCKER_COMPOSE_URL)
	if err != nil {
		return err
	}

	dockerComposeBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	dockerCompose, err := os.Create(DOCKER_DIR + "/docker-compose.yaml")
	if err != nil {
		return err
	}

	if _, err = dockerCompose.Write(dockerComposeBytes); err != nil {
		return err
	}

	return nil
}
