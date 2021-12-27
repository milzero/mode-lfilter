package filter

type ModelContainer struct {
	FilterModel Model `yaml:"filterModel"`
}
type Item struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
type Items struct {
	Item Item `yaml:"item"`
}
type Filter struct {
	Desc     string  `yaml:"desc"`
	Function string  `yaml:"function"`
	Type     string  `yaml:"type"`
	Method   string  `yaml:"methon"`
	Items    []Items `yaml:"items"`
}
type Filters struct {
	Filter Filter `yaml:"filte"`
}
type Model struct {
	Version   string    `yaml:"version"`
	Namespace string    `yaml:"namespace"`
	Desc      string    `yaml:"desc"`
	Filters   []Filters `yaml:"filtes"`
}
