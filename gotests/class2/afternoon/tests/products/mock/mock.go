package mock

import "encoding/json"

type StoreMock struct {
	ReadWasCalled  bool
	WriteWasCalled bool
}

func (s *StoreMock) Read(data interface{}) error {
	s.ReadWasCalled = true

	encoded, _ := json.Marshal([]Product{{
		Id:    1,
		Name:  "Before Update",
		Price: 100,
	}})

	return json.Unmarshal(encoded, &data)
}

func (s *StoreMock) Write(data interface{}) error {
	s.WriteWasCalled = true

	return nil
}
