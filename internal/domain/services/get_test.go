package services_test

import (
	"gcpfaqs-svc-faqs/internal/domain/models"
	"gcpfaqs-svc-faqs/internal/domain/services"
	"gcpfaqs-svc-faqs/internal/domain/test"
	"gcpfaqs-svc-faqs/internal/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldResultAResource(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
	repository := test.NewStubRepository()

	sut := services.NewGetFaqUseCase(repository)
	resource, err := sut.Do(id)

	assert.Nil(t, err)
	assert.IsType(t, &models.Faq{}, resource)
	assert.True(t, repository.GetHaveBeenCalled(), sut)
}

func TestGetItShouldAnErrorOnInvalidId(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba"
	repository := test.NewStubRepository()

	sut := services.NewGetFaqUseCase(repository)
	_, err := sut.Do(id)

	assert.IsType(t, err, types.ErrInvalidID)
}
