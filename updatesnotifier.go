package updatesnotifier

import "github.com/parkr/updates-notifier/docker"

// Config is the updates configuration. It declares the current state
// you're running.
type Config struct {
	DockerRepositories []docker.Repository
}

// Run runs the updates notifier for the given configuration.
func Run(cfg Config) []error {
	errs := []error{}

	for _, dockerRepo := range cfg.DockerRepositories {
		if err := docker.CheckDockerRepoForNewerTag(dockerRepo); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
