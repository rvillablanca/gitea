// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTopic(t *testing.T) {
	assert.NoError(t, PrepareTestDatabase())

	topics, err := FindTopics(&FindTopicOptions{})
	assert.NoError(t, err)
	assert.EqualValues(t, 3, len(topics))

	topics, err = FindTopics(&FindTopicOptions{
		Limit: 2,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(topics))

	topics, err = FindTopics(&FindTopicOptions{
		RepoID: 1,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 3, len(topics))

	assert.NoError(t, SaveTopics(2, "golang"))
	topics, err = FindTopics(&FindTopicOptions{})
	assert.NoError(t, err)
	assert.EqualValues(t, 3, len(topics))

	topics, err = FindTopics(&FindTopicOptions{
		RepoID: 2,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(topics))

	assert.NoError(t, SaveTopics(2, "golang", "gitea"))
	topic, err := GetTopicByName("gitea")
	assert.NoError(t, err)
	assert.EqualValues(t, 1, topic.RepoCount)

	topics, err = FindTopics(&FindTopicOptions{})
	assert.NoError(t, err)
	assert.EqualValues(t, 4, len(topics))

	topics, err = FindTopics(&FindTopicOptions{
		RepoID: 2,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(topics))
}
