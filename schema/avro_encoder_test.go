package schema_test

import (
	"reflect"
	"testing"

	"github.com/bborbe/kafka-version-collector/schema"
)

func TestAvroEncoder(t *testing.T) {
	testcases := []struct {
		name            string
		schemaId        uint32
		content         []byte
		expectedLength  int
		expectedError   error
		expectedContent []byte
	}{
		{
			name:            "simple",
			schemaId:        123,
			content:         []byte("hello"),
			expectedLength:  5 + 5,
			expectedError:   nil,
			expectedContent: append([]byte{0, 0, 0, 0, 123}, []byte("hello")...),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := schema.AvroEncoder{
				SchemaId: 123,
				Content:  []byte("hello"),
			}
			length := encoder.Length()
			if tc.expectedLength != length {
				t.Errorf("expect length %d but got %d", tc.expectedLength, length)
			}
			content, err := encoder.Encode()
			if tc.expectedError != err {
				t.Errorf("expect length %d but got %d", tc.expectedLength, length)
			}
			if !reflect.DeepEqual(tc.expectedContent, content) {
				t.Errorf("expected content %v  but got %v", tc.expectedContent, content)
			}
		})
	}
}
