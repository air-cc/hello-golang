package service

import (
	"mblog/blog"
	"testing"

	mblog "mblog/test/mocks/blog"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlog := mblog.NewMockBlog(ctrl)
	mockBlog.EXPECT().ListPosts().Return([]blog.Post{})

	service := NewService(mockBlog)

	data, err := service.ListPosts()

	assert.Equal(t, nil, err)
	assert.Equal(t, []blog.Post{}, data)
}
