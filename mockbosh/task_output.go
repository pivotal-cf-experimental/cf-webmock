package mockbosh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	. "github.com/onsi/gomega"
	"github.com/pivotal-cf-experimental/cf-webmock/http"
)

type taskOutputMock struct {
	*mockhttp.MockHttp
}

func TaskOutput(taskId int) *taskOutputMock {
	mock := &taskOutputMock{MockHttp: mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/tasks/%d/output?type=result", taskId))}
	return mock
}

func (t *taskOutputMock) RespondsWithVMsOutput(vms interface{}) *mockhttp.MockHttp {
	output := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(output)

	for _, line := range interfaceSlice(vms) {
		Expect(encoder.Encode(line)).ToNot(HaveOccurred())
	}

	return t.RespondsWith(string(output.Bytes()))
}

func (t *taskOutputMock) RespondsWithTaskOutput(taskOutput interface{}) *mockhttp.MockHttp {
	output := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(output)

	for _, line := range interfaceSlice(taskOutput) {
		Expect(encoder.Encode(line)).ToNot(HaveOccurred())
	}

	return t.RespondsWith(string(output.Bytes()))
}

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("needs to be called with a slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}