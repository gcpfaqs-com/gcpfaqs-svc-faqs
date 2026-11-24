package services_test

import (
	"gcpfaqs-svc-faqs/internal/domain/models"
	"gcpfaqs-svc-faqs/internal/domain/services"
	"gcpfaqs-svc-faqs/internal/domain/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldResultAList(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewListFaqUseCase(repository)
	list, err := sut.Do()

	assert.Nil(t, err)
	assert.IsType(t, []*models.Faq{}, list)
	assert.True(t, repository.ListHaveBeenCalled(), sut)
}
