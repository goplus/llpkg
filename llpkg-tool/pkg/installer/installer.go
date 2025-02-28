package installer

type Installer interface {
	Name() string
	Config() map[string]string
	Install(dir string) error
}
