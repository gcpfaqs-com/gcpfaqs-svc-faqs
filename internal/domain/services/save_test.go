package services_test

import (
	"gcpfaqs-svc-faqs/internal/domain/services"
	"gcpfaqs-svc-faqs/internal/domain/test"
	"gcpfaqs-svc-faqs/internal/domain/types"

	"testing"

	"github.com/stretchr/testify/assert"
)

var newTitle string = "any valid title"
var newContent string = "any valid content"
var newPoster string = "any valid poster"
var newTags []string = []string{"atag"}
var newVideoUrl string = "any valid url"
var newVideoId string = "0fda1566-a5ff-4d61-a96b-9e030ba75020"
var newVideoPlaylistId string = "0fda1566-a5ff-4d61-a96b-9e030ba75020"

func TestItShouldSaveAResource(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewSaveFaqUseCase(repository)
	_, err := sut.Do(
		newTitle,
		newContent,
		newPoster,
		newTags,
		newVideoUrl,
		newVideoPlaylistId,
		newVideoId)

	assert.Nil(t, err)
	assert.True(t, repository.SaveHaveBeenCalled(), sut)
}

func TestItShouldResultAnErrorOnInvalidTitle(t *testing.T) {
	newTitle = ""
	repository := test.NewStubRepository()

	sut := services.NewSaveFaqUseCase(repository)
	_, err := sut.Do(
		newTitle,
		newContent,
		newPoster,
		newTags,
		newVideoUrl,
		newVideoPlaylistId,
		newVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTitle, err)
}

func TestItShouldResultAnErrorOnInvalidContent(t *testing.T) {
	newContent = ""
	repository := test.NewStubRepository()

	sut := services.NewSaveFaqUseCase(repository)
	_, err := sut.Do(
		newTitle,
		newContent,
		newPoster,
		newTags,
		newVideoUrl,
		newVideoPlaylistId,
		newVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidContent, err)
}

func TestItShouldResultAnErrorOnInvalidPoster(t *testing.T) {
	newPoster = ""
	repository := test.NewStubRepository()

	sut := services.NewSaveFaqUseCase(repository)
	_, err := sut.Do(
		newTitle,
		newContent,
		newPoster,
		newTags,
		newVideoUrl,
		newVideoPlaylistId,
		newVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidPoster, err)
}
func TestItShouldResultAnErrorOnInvalidTags(t *testing.T) {
	newTags = []string{}
	repository := test.NewStubRepository()

	sut := services.NewSaveFaqUseCase(repository)
	_, err := sut.Do(
		newTitle,
		newContent,
		newPoster,
		newTags,
		newVideoUrl,
		newVideoPlaylistId,
		newVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTags, err)
}
