package plugins

type LinuxPlatform struct {}

func (self *LinuxPlatform) OsType() string {
	return "linux"
}

func (self *LinuxPlatform) Run() {

}
