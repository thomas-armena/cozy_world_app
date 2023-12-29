package client

type FakeClient struct {
	Client

	entityId      int32
	writeCallback func([]byte)
}

func NewFakeClient(writeCallback func([]byte)) *FakeClient {
	return &FakeClient{
		writeCallback: writeCallback,
	}
}

func (f *FakeClient) Write(data []byte) {
	f.writeCallback(data)
}

func (f *FakeClient) AssignEntityId(id int32) {
	f.entityId = id
}
