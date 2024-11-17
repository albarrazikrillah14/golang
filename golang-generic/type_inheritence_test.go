package golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (v *MyVicePresident) GetName() string {
	return v.Name
}

func (v *MyVicePresident) GetVicePresidentName() string {
	return v.Name
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Albarra", GetName[Manager](&MyManager{Name: "Albarra"}))
	assert.Equal(t, "Albarra", GetName[*MyVicePresident](&MyVicePresident{Name: "Albarra"}))
}
