package generate

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// These files get rendered as part of the build process.

// If you add a file to `sliver/` it won't be automatically included
// as part of the build by the server, you have to add it here.

var (
	srcFiles = []string{
		"constants/constants.go",

		"handlers/generic-rpc-handlers.go",
		"handlers/generic-tun-handlers.go",
		"handlers/special-handlers.go",
		"handlers/handlers_darwin.go",
		"handlers/handlers_linux.go",
		"handlers/handlers_windows.go",
		"handlers/handlers.go",

		"limits/limits.go",
		"limits/limits_windows.go",
		"limits/limits_darwin.go",
		"limits/limits_linux.go",

		"platform/win/win_windows.go",
		"platform/win/types_amd64.go",
		"platform/win/winmm_windows.go",
		"platform/win/win_amd64.go",
		"platform/win/gdi32_windows.go",
		"platform/win/oleaut32_windows.go",
		"platform/win/iphlpapi_windows.go",
		"platform/win/ws2_32_windows.go",
		"platform/win/pdh_windows.go",
		"platform/win/types_windows.go",
		"platform/win/version_windows.go",
		"platform/win/win_nonwindows_windows.go",
		"platform/win/advapi32_windows.go",
		"platform/win/psapi_windows.go",
		"platform/win/types_386.go",
		"platform/win/win_windows_windows.go",
		"platform/win/shell32_windows.go",
		"platform/win/comctl32_windows.go",
		"platform/win/kernel32_windows.go",
		"platform/win/comdlg32_windows.go",
		"platform/win/user32_windows.go",
		"platform/win/imm32_windows.go",
		"platform/win/gdiplus_windows.go",
		"platform/win/rpcrt4_windows.go",
		"platform/win/uxtheme_windows.go",
		"platform/win/win_386.go",
		"platform/win/ole32_windows.go",
		"platform/win/opengl32_windows.go",
		"platform/win/const_windows.go",
		"platform/win/shlwapi_windows.go",
		"platform/win/crypt32_windows.go",

		"procdump/dump.go",
		"procdump/dump_windows.go",
		"procdump/dump_linux.go",
		"procdump/dump_darwin.go",

		"proxy/provider_darwin.go",
		"proxy/provider_linux.go",
		"proxy/provider_windows.go",
		"proxy/provider.go",
		"proxy/proxy.go",
		"proxy/url.go",

		"ps/ps.go",
		"ps/ps_windows.go",
		"ps/ps_linux.go",
		"ps/ps_darwin.go",

		"shell/shell.go",
		"shell/shell_windows.go",
		"shell/shell_darwin.go",
		"shell/pty/pty_darwin.go",
		"shell/shell_linux.go",
		"shell/pty/pty_linux.go",

		"shell/pty/run.go",
		"shell/pty/util.go",
		"shell/pty/doc.go",
		"shell/pty/types.go",
		"shell/pty/ztypes_386.go",
		"shell/pty/ztypes_amd64.go",
		"shell/pty/ioctl.go",
		"shell/pty/ioctl_bsd.go",
		"shell/pty/ioctl_darwin.go",
		"shell/pty/pty_unsupported.go",

		"taskrunner/task.go",
		"taskrunner/task_windows.go",
		"taskrunner/task_darwin.go",
		"taskrunner/task_linux.go",

		"priv/priv.go",
		"priv/priv_windows.go",

		"transports/crypto.go",
		"transports/tcp-mtls.go",
		"transports/tcp-http.go",
		"transports/udp-dns.go",
		"transports/udp-dns_linux.go",
		"transports/udp-dns_darwin.go",
		"transports/udp-dns_windows.go",
		"transports/transports.go",

		"version/version.go",
		"version/version_windows.go",
		"version/version_linux.go",
		"version/version_darwin.go",

		"winhttp/winhttp.go",

		"sliver.go",
	}
)
