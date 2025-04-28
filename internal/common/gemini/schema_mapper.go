package gemini

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

// ToGenaiSchema converts a struct into a *genai.Schema.
func StructToSchema(i interface{}) *genai.Schema {
	// 타입 Ptr인지 조사 후 Ptr일 경우 Value로
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// Schema 생성
	return schemaFromType(t)
}

// schemaFromType generates genai.Schema based on the Go type.
func schemaFromType(t reflect.Type) *genai.Schema {
	// 타입 Ptr인지 조사 후 Ptr일 경우 Value로
	if t.Kind() == reflect.Ptr {
		return schemaFromType(t.Elem())
	}

	switch t.Kind() {
	// Struct kind일 경우 Field 순차 조회하여 Schema 생성
	case reflect.Struct:
		s := &genai.Schema{
			Type:       genai.TypeObject,
			Properties: make(map[string]*genai.Schema),
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			// 공개 Field만 필터링 (Private field의 경우 빈 값 들어감)
			if f.PkgPath != "" {
				continue
			}

			// 태그 내부의 JSON 필드 명 조회
			// 명시되어있지 않은 경우 Field name 사용
			name := f.Tag.Get("json")
			if name == "" {
				name = f.Name
			} else {
				name = strings.Split(name, ",")[0]
			}

			// 해당 필드에 대해 Schema 생성 후 태그 적용
			child := schemaFromType(f.Type)
			applyTags(child, f.Tag.Get("genai"))
			s.Properties[name] = child

			// Required tag 있을 경우 Required 목록에 포함
			if hasTag(f.Tag.Get("genai"), "required") {
				s.Required = append(s.Required, name)
			}
		}
		return s

	// Slice or Array kind일 경우
	case reflect.Slice, reflect.Array:
		item := schemaFromType(t.Elem())
		return &genai.Schema{
			Type:  genai.TypeArray,
			Items: item,
		}

	// 그 외 원시 타입들
	default:
		return primitiveGenaiSchema(t)
	}
}

// primitiveGenaiSchema maps Go basic kinds to genai.Schema types.
func primitiveGenaiSchema(t reflect.Type) *genai.Schema {
	s := &genai.Schema{}
	switch t.Kind() {
	case reflect.String:
		s.Type = genai.TypeString
	case reflect.Bool:
		s.Type = genai.TypeBoolean
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s.Type = genai.TypeInteger
	case reflect.Float32, reflect.Float64:
		s.Type = genai.TypeNumber
	default:
		s.Type = genai.TypeString
	}
	return s
}

// applyTags parses `genai` struct tag and applies to the Schema.
func applyTags(s *genai.Schema, tag string) {
	// 빈 태그일 경우 생략
	if tag == "" {
		return
	}

	for _, part := range strings.Split(tag, ";") {
		if part == "" {
			continue
		}
		kv := strings.SplitN(part, "=", 2)
		key := kv[0]
		val := ""
		if len(kv) == 2 {
			val = kv[1]
		}

		switch key {
		case "format":
			s.Format = val
		case "description":
			s.Description = val
		case "nullable":
			s.Nullable = true
		case "enum":
			s.Enum = strings.Split(val, ",")
		}
	}
}

// hasTag checks if a semicolon-delimited tag contains the given key.
func hasTag(tag, key string) bool {
	if tag == "" {
		return false
	}
	for _, part := range strings.Split(tag, ";") {
		if part == key || strings.HasPrefix(part, key+"=") {
			return true
		}
	}
	return false
}

func JsonToStruct(jsonStr string, out any) error {
	rv := reflect.ValueOf(out)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("out must be a non‐nil pointer, got %T", out)
	}
	return json.Unmarshal([]byte(jsonStr), out)
}
