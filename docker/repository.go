package docker

import "fmt"

// Repository is a single Docker repository & its current tag.
type Repository struct {
	// The repository name.
	Name string
	// The current tag.
	Tag string
}

// CheckDockerRepoForNewerTag checks the repository tags for something newer
// and returns an error if something newer exists.
func CheckDockerRepoForNewerTag(repo Repository) error {
	tags, err := LatestTags(repo.Name)
	if err != nil {
		return internalError{repo: repo, errMessage: err.Error()}
	}
	if len(tags) == 0 {
		return internalError{repo: repo, errMessage: "no tags found"}
	}
	if repo.Tag != tags[0].Name {
		return newerVersionError{repo: repo, newerTag: tags[0]}
	}
	return nil
}

type internalError struct {
	repo       Repository
	errMessage string
}

func (e internalError) Error() string {
	return fmt.Sprintf(
		"Repository %s could not be checked: %s",
		e.repo.Name,
		e.errMessage,
	)
}

type newerVersionError struct {
	repo     Repository
	newerTag Tag
}

func (e newerVersionError) Error() string {
	return fmt.Sprintf(
		"Repository %s has a new tag %s created on %s",
		e.repo.Name,
		e.newerTag.Name,
		e.newerTag.LastUpdated,
	)
}
