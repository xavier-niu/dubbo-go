package hessian2

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestGetJavaName(t *testing.T) {
	_, err := GetJavaName(nil)
	assert.Equal(t, NilError, err)

	typ, err := GetJavaName(true)
	assert.Equal(t, "boolean", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(1)
	assert.Equal(t, "long", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName(int64(1))
	assert.Equal(t, "long", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(int32(1))
	assert.Equal(t, "int", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(int16(1))
	assert.Equal(t, "short", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName(int8(1))
	assert.Equal(t, "short", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(uint(1))
	assert.Equal(t, "unsigned long", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName(uint64(1))
	assert.Equal(t, "unsigned long", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(uint32(1))
	assert.Equal(t, "unsigned int", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(uint16(1))
	assert.Equal(t, "unsigned short", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName(byte('a'))
	assert.Equal(t, "char", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(float32(1.0))
	assert.Equal(t, "float", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(1.0)
	assert.Equal(t, "double", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName("hello")
	assert.Equal(t, "java.lang.String", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName([]string{"hello"})
	assert.Equal(t, "java.lang.String[]", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName([]*mockPOJOPtr{{}})
	assert.Equal(t, "org.apache.dubbo.mockPOJOPtr[]", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName([]mockPOJO{{}})
	assert.Equal(t, "org.apache.dubbo.mockPOJO[]", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName([][]string{{"hello"}})
	assert.Equal(t, "java.lang.String[][]", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName([][]*mockPOJOPtr{{&mockPOJOPtr{}}})
	assert.Equal(t, "org.apache.dubbo.mockPOJOPtr[][]", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName([1]string{"hello"})
	assert.Equal(t, "java.lang.String[]", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(map[string]string{"key1": "value1"})
	assert.Equal(t, "java.util.Map", typ)
	assert.Nil(t, err)

	typ, err = GetJavaName(mockPOJO{})
	assert.Equal(t, "org.apache.dubbo.mockPOJO", typ)
	assert.Nil(t, err)
	typ, err = GetJavaName(&mockPOJOPtr{})
	assert.Equal(t, "org.apache.dubbo.mockPOJOPtr", typ)
	assert.Nil(t, err)

	_, err = GetJavaName(&mockNonPOJO{})
	assert.Equal(t, UnexpectedTypeError, err)
}

type mockPOJOPtr struct {
	TestField string
}

func (m *mockPOJOPtr) JavaClassName() string {
	return "org.apache.dubbo.mockPOJOPtr"
}

type mockPOJO struct{}

func (m mockPOJO) JavaClassName() string {
	return "org.apache.dubbo.mockPOJO"
}

type mockNonPOJO struct{}
