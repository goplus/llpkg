package upstream

type Upstream interface {
	Installer() Installer
	Package() Package
	Validate() error
}

type defaultUpstream struct {
	installer Installer
	pkg       Package
}

func NewDefaultUpstream(installer Installer, pkg Package) Upstream {
	return &defaultUpstream{
		installer: installer,
		pkg:       pkg,
	}
}

func (u *defaultUpstream) Installer() Installer {
	return u.installer
}

func (u *defaultUpstream) Package() Package {
	return u.pkg
}

func (u *defaultUpstream) Validate() error {
	return u.installer.Validate(u.pkg)
}

type Installer interface {
	Name() string
	Config() map[string]string
	Install(pkg Package, dir string) error
	Search(pkg Package) (string, error)
	Validate(pkg Package) error // validate remote exists
}

type Package struct {
	Name    string
	Version string
}
