// Code generated by 'go generate'; DO NOT EDIT.

package Unhook

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modKernel32 = windows.NewLazySystemDLL("Kernel32.dll")
	modPsapi    = windows.NewLazySystemDLL("Psapi.dll")

	procCloseHandle          = modKernel32.NewProc("CloseHandle")
	procCreateFileA          = modKernel32.NewProc("CreateFileA")
	procCreateFileMappingW   = modKernel32.NewProc("CreateFileMappingW")
	procGetCurrentProcess    = modKernel32.NewProc("GetCurrentProcess")
	procGetModuleHandleA     = modKernel32.NewProc("GetModuleHandleA")
	procMapViewOfFile        = modKernel32.NewProc("MapViewOfFile")
	procVirtualProtect       = modKernel32.NewProc("VirtualProtect")
	procGetModuleInformation = modPsapi.NewProc("GetModuleInformation")
)

func CloseHandleNu1r(hObject uintptr) {
	syscall.Syscall(procCloseHandle.Addr(), 1, uintptr(hObject), 0, 0)
	return
}

func CreateFileANu1r(lpFileName uintptr, dwDesiredAccess uintptr, dwShareMode uintptr, lpSecurityAttributes uintptr, dwCreationDisposition uintptr, dwFlagsAndAttributes uintptr, hTemplateFile uintptr) (ntdllFile uintptr) {
	r0, _, _ := syscall.Syscall9(procCreateFileA.Addr(), 7, uintptr(lpFileName), uintptr(dwDesiredAccess), uintptr(dwShareMode), uintptr(lpSecurityAttributes), uintptr(dwCreationDisposition), uintptr(dwFlagsAndAttributes), uintptr(hTemplateFile), 0, 0)
	ntdllFile = uintptr(r0)
	return
}

func CreateFileMappingWNu1r(hFile uintptr, lpFileMappingAttributes uintptr, flProtect uintptr, dwMaximumSizeHigh uintptr, dwMaximumSizeLow uintptr, lpName uintptr) (ntdllMapping uintptr) {
	r0, _, _ := syscall.Syscall6(procCreateFileMappingW.Addr(), 6, uintptr(hFile), uintptr(lpFileMappingAttributes), uintptr(flProtect), uintptr(dwMaximumSizeHigh), uintptr(dwMaximumSizeLow), uintptr(lpName))
	ntdllMapping = uintptr(r0)
	return
}

func GetCurrentProcessNu1r() (process uintptr) {
	r0, _, _ := syscall.Syscall(procGetCurrentProcess.Addr(), 0, 0, 0, 0)
	process = uintptr(r0)
	return
}

func GetModuleHandleANu1r(lpModuleName uintptr) (ntdllModule uintptr) {
	r0, _, _ := syscall.Syscall(procGetModuleHandleA.Addr(), 1, uintptr(lpModuleName), 0, 0)
	ntdllModule = uintptr(r0)
	return
}

func MapViewOfFileNu1r(hFileMappingObject uintptr, dwDesiredAccess uintptr, dwFileOffsetHigh uintptr, dwFileOffsetLow uintptr, dwNumberOfBytesToMap uintptr) (ntdllMappingAddress uintptr) {
	r0, _, _ := syscall.Syscall6(procMapViewOfFile.Addr(), 5, uintptr(hFileMappingObject), uintptr(dwDesiredAccess), uintptr(dwFileOffsetHigh), uintptr(dwFileOffsetLow), uintptr(dwNumberOfBytesToMap), 0)
	ntdllMappingAddress = uintptr(r0)
	return
}

func VirtualProtectNu1r(lpAddress uintptr, dwSize uintptr, flNewProtect uintptr, lpflOldProtect uintptr) (isProtected uintptr) {
	r0, _, _ := syscall.Syscall6(procVirtualProtect.Addr(), 4, uintptr(lpAddress), uintptr(dwSize), uintptr(flNewProtect), uintptr(lpflOldProtect), 0, 0)
	isProtected = uintptr(r0)
	return
}

func GetModuleInformationNu1r(hProcess uintptr, hModule uintptr, lpmodinfo uintptr, cb uintptr) {
	syscall.Syscall6(procGetModuleInformation.Addr(), 4, uintptr(hProcess), uintptr(hModule), uintptr(lpmodinfo), uintptr(cb), 0, 0)
	return
}
