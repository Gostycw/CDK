package capability

import (
	"testing"
)

func TestCapHexParser(t *testing.T) {
	var testcase = map[string]string{
		"00000000a80425fb": "CAP_CHOWN,CAP_DAC_OVERRIDE,CAP_FOWNER,CAP_FSETID,CAP_KILL,CAP_SETGID,CAP_SETUID,CAP_SETPCAP,CAP_NET_BIND_SERVICE,CAP_NET_RAW,CAP_SYS_CHROOT,CAP_MKNOD,CAP_AUDIT_WRITE,CAP_SETFCAP",
		"00000000a80c25fb" : "CAP_CHOWN,CAP_DAC_OVERRIDE,CAP_FOWNER,CAP_FSETID,CAP_KILL,CAP_SETGID,CAP_SETUID,CAP_SETPCAP,CAP_NET_BIND_SERVICE,CAP_NET_RAW,CAP_SYS_CHROOT,CAP_SYS_PTRACE,CAP_MKNOD,CAP_AUDIT_WRITE,CAP_SETFCAP",
		"00000000a82425fb": "CAP_CHOWN,CAP_DAC_OVERRIDE,CAP_FOWNER,CAP_FSETID,CAP_KILL,CAP_SETGID,CAP_SETUID,CAP_SETPCAP,CAP_NET_BIND_SERVICE,CAP_NET_RAW,CAP_SYS_CHROOT,CAP_SYS_ADMIN,CAP_MKNOD,CAP_AUDIT_WRITE,CAP_SETFCAP",
	}
	for k, v := range testcase {
		if CapHexToText(k) != v {
			t.Errorf("CapHexParser error parse %s: %s", k, v)
		}
	}
}
