package upstream

type Upstream struct {
	installer Installer
	pkg       Package
}

type Package struct {
	Name    string
	Version string
}

func NewUpstream(installer Installer, pkg Package) *Upstream {
	return &Upstream{
		installer: installer,
		pkg:       pkg,
	}
}

func (u *Upstream) Installer() Installer {
	return u.installer
}

func (u *Upstream) Package() Package {
	return u.pkg
}
