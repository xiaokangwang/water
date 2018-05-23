package water

import "os"

func DropTAPInterface(fd os.File) *Interface {
	return &Interface{isTAP: true, name: "/dev/net/tun", fd: int(fd.Fd()), ReadWriteCloser: &fd}
}
