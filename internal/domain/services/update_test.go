package services_test

import (
	"gcpfaqs-svc-faqs/internal/domain/services"
	"gcpfaqs-svc-faqs/internal/domain/test"
	"gcpfaqs-svc-faqs/internal/domain/types"

	"testing"

	"github.com/stretchr/testify/assert"
)

var updateId string = "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
var updateTitle string = "any valid title"
var updateContent string = "any valid content"
var updatePoster string = "any valid poster"
var updateTags []string = []string{"atag"}
var updateVideoUrl string = "any valid url"
var updateVideoId string = "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
var updatePlaylistId string = "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"

func TestItShouldUpdateAResource(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.Nil(t, err)
	assert.True(t, repository.UpdateHaveBeenCalled(), sut)
}

func TestUpdateItShouldResultAnErrorOnInvalidTitle(t *testing.T) {
	updateTitle = ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTitle, err)
}

func TestUpdateItShouldResultAnErrorOnInvalidContent(t *testing.T) {
	updateContent = ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidContent, err)
}

func TestUpdateItShouldResultAnErrorOnInvalidPoster(t *testing.T) {
	updatePoster = ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidPoster, err)
}
func TestUpdateItShouldResultAnErrorOnInvalidTags(t *testing.T) {
	updateTags = []string{}
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidTags, err)
}

func TestUpdateItShouldResultAnErrorOnInvalidVideoUrl(t *testing.T) {
	updateVideoUrl = ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidVideoURL, err)
}

func TestUpdateItShouldResultAnErrorOnInvalidVideoId(t *testing.T) {
	updateVideoId = ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidID, err)
}

func TestUpdateItShouldResultAnErrorOnInvalidVideoPlaylistId(t *testing.T) {
	updatePlaylistId = ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateFaqUseCase(repository)
	err := sut.Do(
		updateId,
		updateTitle,
		updateContent,
		updatePoster,
		updateTags,
		updateVideoUrl,
		updatePlaylistId,
		updateVideoId)

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidVideoPlaylistID, err)
}
