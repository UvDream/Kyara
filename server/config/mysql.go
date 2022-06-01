package config

type Mysql struct {
	DBConfig `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
