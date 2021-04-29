package collector

type Source struct {
	Path string
}

type YAMLSource interface {
	Get(source Source) (string, error)
}
