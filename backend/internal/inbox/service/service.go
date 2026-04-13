package service

import (
	"context"
	"errors"
	"mindeflow-app/backend/internal/inbox"
	"mindeflow-app/backend/internal/inbox/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, input inbox.CreateInput) (inbox.InboxItem, error) {
	if input.Title == "" {
		return inbox.InboxItem{}, errors.New("title is required")
	}
	return s.repo.Create(ctx, input)
}

func (s *Service) List(ctx context.Context, filter inbox.ListFilter) (inbox.ListResult, error) {
	if filter.Limit <= 0 {
		filter.Limit = 20
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}

	return s.repo.List(ctx, filter)
}
