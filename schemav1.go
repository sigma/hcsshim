package hcsshim

type Layer struct {
	ID   string
	Path string
}

type MappedDir struct {
	HostPath          string
	ContainerPath     string
	ReadOnly          bool
	BandwidthMaximum  uint64
	IOPSMaximum       uint64
	CreateInUtilityVM bool
}

type MappedPipe struct {
	HostPath          string
	ContainerPipeName string
}

type HvRuntime struct {
	ImagePath           string `json:",omitempty"`
	SkipTemplate        bool   `json:",omitempty"`
	LinuxInitrdFile     string `json:",omitempty"` // File under ImagePath on host containing an initrd image for starting a Linux utility VM
	LinuxKernelFile     string `json:",omitempty"` // File under ImagePath on host containing a kernel for starting a Linux utility VM
	LinuxBootParameters string `json:",omitempty"` // Additional boot parameters for starting a Linux Utility VM in initrd mode
	BootSource          string `json:",omitempty"` // "Vhd" for Linux Utility VM booting from VHD
	WritableBootSource  bool   `json:",omitempty"` // Linux Utility VM booting from VHD
}

type MappedVirtualDisk struct {
	HostPath          string `json:",omitempty"` // Path to VHD on the host
	ContainerPath     string // Platform-specific mount point path in the container
	CreateInUtilityVM bool   `json:",omitempty"`
	ReadOnly          bool   `json:",omitempty"`
	Cache             string `json:",omitempty"` // "" (Unspecified); "Disabled"; "Enabled"; "Private"; "PrivateAllowSharing"
	AttachOnly        bool   `json:",omitempty:`
}

// ContainerConfig is used as both the input of CreateContainer
// and to convert the parameters to JSON for passing onto the HCS
type ContainerConfig struct {
	SystemType                  string              // HCS requires this to be hard-coded to "Container"
	Name                        string              // Name of the container. We use the docker ID.
	Owner                       string              `json:",omitempty"` // The management platform that created this container
	VolumePath                  string              `json:",omitempty"` // Windows volume path for scratch space. Used by Windows Server Containers only. Format \\?\\Volume{GUID}
	IgnoreFlushesDuringBoot     bool                `json:",omitempty"` // Optimization hint for container startup in Windows
	LayerFolderPath             string              `json:",omitempty"` // Where the layer folders are located. Used by Windows Server Containers only. Format  %root%\windowsfilter\containerID
	Layers                      []Layer             // List of storage layers. Required for Windows Server and Hyper-V Containers. Format ID=GUID;Path=%root%\windowsfilter\layerID
	Credentials                 string              `json:",omitempty"` // Credentials information
	ProcessorCount              uint32              `json:",omitempty"` // Number of processors to assign to the container.
	ProcessorWeight             uint64              `json:",omitempty"` // CPU shares (relative weight to other containers with cpu shares). Range is from 1 to 10000. A value of 0 results in default shares.
	ProcessorMaximum            int64               `json:",omitempty"` // Specifies the portion of processor cycles that this container can use as a percentage times 100. Range is from 1 to 10000. A value of 0 results in no limit.
	StorageIOPSMaximum          uint64              `json:",omitempty"` // Maximum Storage IOPS
	StorageBandwidthMaximum     uint64              `json:",omitempty"` // Maximum Storage Bandwidth in bytes per second
	StorageSandboxSize          uint64              `json:",omitempty"` // Size in bytes that the container system drive should be expanded to if smaller
	MemoryMaximumInMB           int64               `json:",omitempty"` // Maximum memory available to the container in Megabytes
	HostName                    string              `json:",omitempty"` // Hostname
	MappedDirectories           []MappedDir         `json:",omitempty"` // List of mapped directories (volumes/mounts)
	MappedPipes                 []MappedPipe        `json:",omitempty"` // List of mapped Windows named pipes
	HvPartition                 bool                // True if it a Hyper-V Container
	NetworkSharedContainerName  string              `json:",omitempty"` // Name (ID) of the container that we will share the network stack with.
	EndpointList                []string            `json:",omitempty"` // List of networking endpoints to be attached to container
	HvRuntime                   *HvRuntime          `json:",omitempty"` // Hyper-V container settings. Used by Hyper-V containers only. Format ImagePath=%root%\BaseLayerID\UtilityVM
	Servicing                   bool                `json:",omitempty"` // True if this container is for servicing
	AllowUnqualifiedDNSQuery    bool                `json:",omitempty"` // True to allow unqualified DNS name resolution
	DNSSearchList               string              `json:",omitempty"` // Comma seperated list of DNS suffixes to use for name resolution
	ContainerType               string              `json:",omitempty"` // "Linux" for Linux containers on Windows. Omitted otherwise.
	TerminateOnLastHandleClosed bool                `json:",omitempty"` // Should HCS terminate the container once all handles have been closed
	MappedVirtualDisks          []MappedVirtualDisk `json:",omitempty"` // Array of virtual disks to mount at start
}
