package models

var MessagesStorage *MessageStorage

type MessageStorage struct {
	SenderMessages    map[string]*Message
	RecipientMessages map[string]*Message
}

func NewMessageStorage() {
	MessagesStorage = &MessageStorage{
		SenderMessages:    make(map[string]*Message),
		RecipientMessages: make(map[string]*Message),
	}
}

func (ms *MessageStorage) AddMessage(message Message) {
	ms.SenderMessages[message.SenderID] = &message
	ms.RecipientMessages[message.RecipientID] = &message

}
func (ms *MessageStorage) GetMessageBySender(id string) *Message {
	return ms.SenderMessages[id]
}

func (ms *MessageStorage) GetMessageByRecipient(id string) *Message {
	return ms.RecipientMessages[id]
}
