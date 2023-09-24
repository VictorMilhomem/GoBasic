package parser

type Environment struct {
	Values map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{
		Values: make(map[string]interface{}),
	}
}

func (e *Environment) Get(name string) (interface{}, bool) {
	val, ok := e.Values[name]
	return val, ok
}

func (e *Environment) Set(name string, val interface{}) interface{} {
	e.Values[name] = val
	return val
}

func CopyEnv(env *Environment) *Environment {
	newEnv := NewEnvironment()
	for k, v := range env.Values {
		newEnv.Values[k] = v
	}
	return newEnv
}
