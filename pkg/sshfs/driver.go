package sshfs

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/glog"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
)

type driver struct {
	csiDriver *csicommon.CSIDriver
	endpoint  string

	//ids *identityServer
	ns    *nodeServer
	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

const (
	driverName = "csi-sshfs"
)

var (
	Version   = "latest"
	BuildTime = "1970-01-01 00:00:00"
)

func NewDriver(nodeID, endpoint string) *driver {
	glog.Infof("Starting new %s driver in version %s built %s", driverName, Version, BuildTime)

	d := &driver{}

	d.endpoint = endpoint

	csiDriver := csicommon.NewCSIDriver(driverName, Version, nodeID)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})
	// SSHFS plugin does not support ControllerServiceCapability now.
	// If support is added, it should set to appropriate
	// ControllerServiceCapability RPC types.
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{csi.ControllerServiceCapability_RPC_UNKNOWN})

	d.csiDriver = csiDriver

	return d
}

func NewNodeServer(d *driver) *nodeServer {
	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.csiDriver),
		mounts:            map[string]*mountPoint{},
	}
}

func (d *driver) Run() {
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(d.endpoint,
		csicommon.NewDefaultIdentityServer(d.csiDriver),
		// SSHFS plugin has not implemented ControllerServer
		// using default controllerserver.
		csicommon.NewDefaultControllerServer(d.csiDriver),
		NewNodeServer(d))
	s.Wait()
}
