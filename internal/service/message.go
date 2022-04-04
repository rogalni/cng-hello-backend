package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/repository"
	"github.com/rogalni/cng-hello-backend/pkg/log"
)

type MessageService struct {
	messageRepository repository.IMessageRepository
}

func NewMessageService() *MessageService {
	return &MessageService{
		messageRepository: repository.NewMessageRepository(),
	}
}

func (ms MessageService) GetMessages(ctx context.Context) ([]*model.Message, error) {
	log.Ctx(ctx).Info().Msg("Get message from service with trace infos")
	return ms.messageRepository.GetMessages(ctx)
}

func (ms MessageService) GetMessage(ctx context.Context, id uuid.UUID) (*model.Message, error) {
	log.Ctx(ctx).Info().Msg("Get message from service with trace infos")
	return ms.messageRepository.GetMessage(ctx, id)
}

func (ms MessageService) CreateMessage(ctx context.Context, m *model.Message) error {
	log.Ctx(ctx).Info().Msg("Create message from service with trace infos")
	return ms.messageRepository.CreateMessage(ctx, m)
}

func (ms MessageService) DeleteMessage(ctx context.Context, id uuid.UUID) error {
	log.Ctx(ctx).Info().Msg("Delete message from service with trace infos")
	return ms.messageRepository.DeleteMessage(ctx, id)
}
