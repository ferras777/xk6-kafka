package kafka

import (
	"errors"
)

type Serializer func(configuration Configuration, topic string, data interface{}, keyOrValue string, schema string) ([]byte, error)

func GetSerializer(serializer string) Serializer {
	switch serializer {
	case "org.apache.kafka.common.serialization.ByteArraySerializer":
		return SerializeByteArray
	case "org.apache.kafka.common.serialization.StringSerializer":
		return SerializeString
	case "io.confluent.kafka.serializers.KafkaAvroSerializer":
		return SerializeAvro
	default:
		return SerializeAvro
	}
}

func SerializeByteArray(configuration Configuration, topic string, data interface{}, keyOrValue string, schema string) ([]byte, error) {
	switch data.(type) {
	case []interface{}:
		bArray := data.([]interface{})
		arr := make([]byte, len(bArray))
		for i, u := range bArray {
			arr[i] = byte(u.(int64))
		}
		return arr, nil
	default:
		return nil, errors.New("Invalid data type provided for byte array serializer (requires []byte)")
	}
}

func SerializeString(configuration Configuration, topic string, data interface{}, keyOrValue string, schema string) ([]byte, error) {
	switch data.(type) {
	case string:
		return []byte(data.(string)), nil
	default:
		return nil, errors.New("Invalid data type provided for string serializer (requires string)")
	}
}
