package resource_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/concourse/github-release-resource"
	"github.com/google/go-github/github"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGithubReleaseResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Github Release Resource Suite")
}

func newRepositoryRelease(id int, version string) *github.RepositoryRelease {
	return &github.RepositoryRelease{
		TagName:    github.String(version),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(false),
		ID:         github.Int(id),
	}
}

func newPreReleaseRepositoryRelease(id int, version string) *github.RepositoryRelease {
	return &github.RepositoryRelease{
		TagName:    github.String(version),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(true),
		ID:         github.Int(id),
	}
}
func newDraftRepositoryRelease(id int, version string) *github.RepositoryRelease {
	return &github.RepositoryRelease{
		TagName:    github.String(version),
		Draft:      github.Bool(true),
		Prerelease: github.Bool(false),
		ID:         github.Int(id),
	}
}

func newDraftWithNilTagRepositoryRelease(id int) *github.RepositoryRelease {
	return &github.RepositoryRelease{
		Draft:      github.Bool(true),
		Prerelease: github.Bool(false),
		ID:         github.Int(id),
	}
}

func exampleTimeStamp(day int) time.Time {
	return time.Date(2018, time.January, day, 0, 0, 0, 0, time.UTC)
}

func newRepositoryReleaseWithCreatedTime(id int, version string, day int) *github.RepositoryRelease {
	return &github.RepositoryRelease{
		TagName:    github.String(version),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(false),
		ID:         github.Int(id),
		CreatedAt:  &github.Timestamp{exampleTimeStamp(day)},
	}
}

func newRepositoryReleaseWithPublishedTime(id int, version string, day int) *github.RepositoryRelease {
	return &github.RepositoryRelease{
		TagName:     github.String(version),
		Draft:       github.Bool(false),
		Prerelease:  github.Bool(false),
		ID:          github.Int(id),
		PublishedAt: &github.Timestamp{exampleTimeStamp(day)},
	}
}

func newRepositoryReleaseWithCreatedAndPublishedTime(id int, version string, createdDay int, publishedDay int) *github.RepositoryRelease {
	return &github.RepositoryRelease{
		TagName:     github.String(version),
		Draft:       github.Bool(false),
		Prerelease:  github.Bool(false),
		ID:          github.Int(id),
		CreatedAt:   &github.Timestamp{exampleTimeStamp(createdDay)},
		PublishedAt: &github.Timestamp{exampleTimeStamp(publishedDay)},
	}
}

func newVersionWithTimestamp(id int, tag string, day int) resource.Version {
	return resource.Version{
		ID:        strconv.Itoa(id),
		Tag:       tag,
		Timestamp: exampleTimeStamp(day),
	}
}
