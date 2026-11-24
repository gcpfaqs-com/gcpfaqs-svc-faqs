package services_test

import (
	"gcpfaqs-svc-faqs/internal/domain/services"
	"gcpfaqs-svc-faqs/internal/domain/test"
	"gcpfaqs-svc-faqs/internal/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldDeleteAResource(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
	repository := test.NewStubRepository()

	sut := services.NewDeleteFaqUseCase(repository)
	err := sut.Do(id)

	assert.Nil(t, err)
	assert.True(t, repository.DeleteHaveBeenCalled(), sut)
}

func TestDeleteItShouldAnErrorOnInvalidId(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba"
	repository := test.NewStubRepository()

	sut := services.NewDeleteFaqUseCase(repository)
	err := sut.Do(id)

	assert.IsType(t, err, types.ErrInvalidID)
}
