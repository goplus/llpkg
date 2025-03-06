package upstream

type Installer interface {
	Name() string
	Config() map[string]string
	Install(pkg Package, dir string) error
	Search(pkg Package) (string, error)
}
