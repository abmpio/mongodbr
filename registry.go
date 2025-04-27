package mongodbr

import (
	"fmt"
	"reflect"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

var (
	_tUUID       = reflect.TypeOf(uuid.UUID{})
	_uuidSubtype = byte(0x04)
)

// 自定义uuid的序列化
func uuidEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != _tUUID {
		return bsoncodec.ValueEncoderError{Name: "uuidEncodeValue", Types: []reflect.Type{_tUUID}, Received: val}
	}
	b := val.Interface().(uuid.UUID)
	return vw.WriteBinaryWithSubtype(b[:], _uuidSubtype)
}

// 自定义uuid的反序列化
func uuidDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != _tUUID {
		return bsoncodec.ValueDecoderError{Name: "uuidDecodeValue", Types: []reflect.Type{_tUUID}, Received: val}
	}

	var data []byte
	var subtype byte
	var err error
	switch vrType := vr.Type(); vrType {
	case bson.TypeBinary:
		data, subtype, err = vr.ReadBinary()
		if subtype != _uuidSubtype {
			return fmt.Errorf("unsupported binary subtype %v for UUID", subtype)
		}
	case bson.TypeNull:
		err = vr.ReadNull()
	case bson.TypeUndefined:
		err = vr.ReadUndefined()
	default:
		return fmt.Errorf("cannot decode %v into a UUID", vrType)
	}

	if err != nil {
		return err
	}
	uuid2, err := uuid.FromBytes(data)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(uuid2))
	return nil
}

// // 自定义 time.Time 反序列化（用于反序列化 primitive.DateTime -> time.Time）
// func timeDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
// 	// 验证 BSON 值的类型为 DateTime
// 	if vr.Type() != bson.TypeDateTime {
// 		return fmt.Errorf("cannot decode non-date-time BSON value into time.Time")
// 	}

// 	// 读取 BSON DateTime 并转换为 time.Time
// 	// 将 BSON DateTime 转换为 time.Time
// 	ms, err := vr.ReadDateTime()
// 	if err != nil {
// 		return err
// 	}
// 	t := time.Unix(0, ms*int64(time.Millisecond))

// 	// 根据不同的目标类型进行处理
// 	switch val.Kind() {
// 	case reflect.Interface:
// 		// 如果是 interface{} 类型，直接将 time.Time 赋给该接口
// 		val.Set(reflect.ValueOf(t))

// 	case reflect.Struct:
// 		// 如果是 time.Time 类型，直接设置值
// 		if val.Type() == reflect.TypeOf(time.Time{}) {
// 			val.Set(reflect.ValueOf(t))
// 		} else {
// 			return bsoncodec.ValueDecoderError{Name: "timeDecoder", Types: []reflect.Type{reflect.TypeOf(time.Time{})}, Received: val}
// 		}

// 	case reflect.Ptr:
// 		// 如果是 *time.Time 类型
// 		if val.Type().Elem() == reflect.TypeOf(time.Time{}) {
// 			if val.IsNil() {
// 				// 创建新的 time.Time
// 				val.Set(reflect.New(val.Type().Elem()))
// 			}
// 			// 设置解码后的值
// 			val.Elem().Set(reflect.ValueOf(t))
// 		} else {
// 			return bsoncodec.ValueDecoderError{Name: "timeDecoder", Types: []reflect.Type{reflect.TypeOf(time.Time{}), reflect.TypeOf(&time.Time{})}, Received: val}
// 		}

// 	default:
// 		return bsoncodec.ValueDecoderError{Name: "timeDecoder", Types: []reflect.Type{reflect.TypeOf(time.Time{}), reflect.TypeOf(&time.Time{})}, Received: val}
// 	}
// 	return nil
// }
