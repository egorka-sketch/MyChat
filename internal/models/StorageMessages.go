package models

var MessagesStorage *MessageStorage

type MessageStorage struct {
	SenderMessages    map[string][]*Message
	RecipientMessages map[string][]*Message
}

func NewMessageStorage() {
	MessagesStorage = &MessageStorage{
		SenderMessages:    make(map[string][]*Message),
		RecipientMessages: make(map[string][]*Message),
	}
}

func (ms *MessageStorage) AddMessage(message Message) {
	ms.SenderMessages[message.SenderID] = append(ms.SenderMessages[message.SenderID], &message)
	ms.RecipientMessages[message.RecipientID] = append(ms.RecipientMessages[message.RecipientID], &message)

}

func (ms *MessageStorage) GetMessageBySender(senderID string) []*Message {
	return ms.SenderMessages[senderID]
}
func (ms *MessageStorage) GetMessageByRecipient(recipientID string) []*Message {
	return ms.RecipientMessages[recipientID]
}
func (ms *MessageStorage) GetAllMessages(UserID string) []*Message {
	var messages []*Message
	for _, message := range ms.SenderMessages[UserID] {
		messages = append(messages, message)
	}
	for _, message := range ms.RecipientMessages[UserID] {
		messages = append(messages, message)
	}
	return messages
}
