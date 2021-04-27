// +build linux

package util

import (
	"fmt"
	"github.com/Mellanox/sriovnet"
	"net"
)

// TODO:(Adrianc) use consts from sriovnet lib
const (
	PORT_FLAVOUR_PHYSICAL = iota
	_
	_
	PORT_FLAVOUR_PCI_PF
	PORT_FLAVOUR_PCI_VF
	PORT_FLAVOUR_UNKNOWN = 0xffff
)

type PortFlavour int

type SriovnetOps interface {
	GetNetDevicesFromPci(pciAddress string) ([]string, error)
	GetUplinkRepresentor(vfPciAddress string) (string, error)
	GetVfIndexByPciAddress(vfPciAddress string) (int, error)
	GetVfRepresentor(uplink string, vfIndex int) (string, error)
	GetPfPciFromVfPci(vfPciAddress string) (string, error)
	GetVfRepresentorSmartNIC(pfID, vfIndex string) (string, error)
	GetRepresentorMacAddress(netdev string) (net.HardwareAddr, error)
	GetRepresentorPortFlavour(netdev string) (PortFlavour, error)
}

type defaultSriovnetOps struct {
}

var sriovnetOps SriovnetOps = &defaultSriovnetOps{}

// SetSriovnetOpsInst method would be used by unit tests in other packages
func SetSriovnetOpsInst(mockInst SriovnetOps) {
	sriovnetOps = mockInst
}

// GetSriovnetOps will be invoked by functions in other packages that would need access to the sriovnet library methods.
func GetSriovnetOps() SriovnetOps {
	return sriovnetOps
}

func (defaultSriovnetOps) GetNetDevicesFromPci(pciAddress string) ([]string, error) {
	return sriovnet.GetNetDevicesFromPci(pciAddress)
}

func (defaultSriovnetOps) GetUplinkRepresentor(vfPciAddress string) (string, error) {
	return sriovnet.GetUplinkRepresentor(vfPciAddress)
}

func (defaultSriovnetOps) GetVfIndexByPciAddress(vfPciAddress string) (int, error) {
	return sriovnet.GetVfIndexByPciAddress(vfPciAddress)
}

func (defaultSriovnetOps) GetVfRepresentor(uplink string, vfIndex int) (string, error) {
	return sriovnet.GetVfRepresentor(uplink, vfIndex)
}

func (defaultSriovnetOps) GetPfPciFromVfPci(vfPciAddress string) (string, error) {
	return sriovnet.GetPfPciFromVfPci(vfPciAddress)
}

func (defaultSriovnetOps) GetVfRepresentorSmartNIC(pfID, vfIndex string) (string, error) {
	return sriovnet.GetVfRepresentorSmartNIC(pfID, vfIndex)
}

// TODO:(adrianc) replace with sriovnet impl once merged and dependency updated.
func (defaultSriovnetOps) GetRepresentorMacAddress(netdev string) (net.HardwareAddr, error) {
	if netdev == "pf0hpf" {
		return net.ParseMAC("0c:42:a1:c6:cf:7c")
	}
	return nil, fmt.Errorf("unexpected netdev")
}

func (defaultSriovnetOps) GetRepresentorPortFlavour(netdev string) (PortFlavour, error) {
	if netdev == "pf0hpf" {
		return PORT_FLAVOUR_PCI_PF, nil
	}
	return PORT_FLAVOUR_UNKNOWN, nil
}
