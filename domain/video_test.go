package domain_test

import (
		"testing"
		"github.com/stretchr/testify/require"
		"encoder/domain"
	)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}