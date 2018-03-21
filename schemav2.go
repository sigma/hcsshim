package hcsshim

// NOTE: The v2 schema is in flux and under development as at March 2018.
// Requires RS5+

type GuestOsV2 struct {
	HostName string `json:"HostName,omitempty"`
}

type ContainersResourcesLayerV2 struct {
	Id    string `json:"Id,omitempty"`
	Path  string `json:"Path,omitempty"`
	Cache string `json:"Cache,omitempty"` //  Unspecified defaults to Enabled
}

type ContainersResourcesStorageQoSV2 struct {
	IOPSMaximum      uint64 `json:"IOPSMaximum,omitempty"`
	BandwidthMaximum uint64 `json:"BandwidthMaximum,omitempty"`
}

type ContainersResourcesStorageV2 struct {
	// List of layers that describe the parent hierarchy for a container's
	// storage. These layers combined together, presented as a disposable
	// and/or committable working storage, are used by the container to
	// record all changes done to the parent layers.
	Layers []ContainersResourcesLayerV2 `json:"Layers,omitempty"`

	// Path that points to the scratch space of a container, where parent
	// layers are combined together to present a new disposable and/or
	// committable layer with the changes done during its runtime.
	Path string `json:"Path,omitempty"`

	StorageQoS *ContainersResourcesStorageQoSV2 `json:"StorageQoS,omitempty"`
}

type ContainersResourcesMappedDirectoryV2 struct {
	HostPath      string `json:"HostPath,omitempty"`
	ContainerPath string `json:"ContainerPath,omitempty"`
	ReadOnly      bool   `json:"ReadOnly,omitempty"`
}

type ContainersResourcesMappedPipeV2 struct {
	ContainerPipeName string `json:"ContainerPipeName,omitempty"`
	HostPath          string `json:"HostPath,omitempty"`
}

type ContainersResourcesMemoryV2 struct {
	Maximum uint64 `json:"Maximum,omitempty"`
}

type ContainersResourcesProcessorV2 struct {
	Count   uint32 `json:"Count,omitempty"`
	Maximum uint64 `json:"Maximum,omitempty"`
	Weight  uint64 `json:"Weight,omitempty"`
}

type ContainersResourcesNetworkingV2 struct {
	AllowUnqualifiedDnsQuery   bool     `json:"AllowUnqualifiedDnsQuery,omitempty"`
	DNSSearchList              string   `json:"DNSSearchList,omitempty"`
	NetworkSharedContainerName string   `json:"NetworkSharedContainerName,omitempty"`
	Namespace                  string   `json:"Namespace,omitempty"`       //  Guid in windows; string in linux
	NetworkAdapters            []string `json:"NetworkAdapters,omitempty"` // JJH Query. Guid in schema.containers.resources.mars
}

type HvSocketServiceConfigV2 struct {
	//  SDDL string that HvSocket will check before allowing a host process to bind  to this specific service.
	// If not specified, defaults to the system DefaultBindSecurityDescriptor.
	BindSecurityDescriptor string `json:"BindSecurityDescriptor,omitempty"`

	//  SDDL string that HvSocket will check before allowing a host process to connect
	// to this specific service.  If not specified, defaults to the system DefaultConnectSecurityDescriptor.
	ConnectSecurityDescriptor string `json:"ConnectSecurityDescriptor,omitempty"`

	//  If true, HvSocket will process wildcard binds for this service/system combination.
	// Wildcard binds are secured in the registry at  SOFTWARE/Microsoft/Windows NT/CurrentVersion/Virtualization/HvSocket/WildcardDescriptors
	AllowWildcardBinds bool `json:"AllowWildcardBinds,omitempty"`

	Disabled bool `json:"Disabled,omitempty"`
}

type HvSocketSystemConfigV2 struct {
	//  SDDL string that HvSocket will check before allowing a host process to bind  to an unlisted service for this specific container/VM (not wildcard binds).
	DefaultBindSecurityDescriptor string `json:"DefaultBindSecurityDescriptor,omitempty"`

	//  SDDL string that HvSocket will check before allowing a host process to connect  to an unlisted service in the VM/container.
	DefaultConnectSecurityDescriptor string `json:"DefaultConnectSecurityDescriptor,omitempty"`

	ServiceTable map[string]HvSocketServiceConfigV2 `json:"ServiceTable,omitempty"`
}

type ContainersResourcesHvSocketV2 struct {
	Config                 *HvSocketSystemConfigV2 `json:"Config,omitempty"`
	EnablePowerShellDirect bool                    `json:"EnablePowerShellDirect,omitempty"`
	EnableUtcRelay         bool                    `json:"EnableUtcRelay,omitempty"`
	EnableAuditing         bool                    `json:"EnableAuditing,omitempty"`
}

type RegistryKeyV2 struct {
	Hive     string `json:"Hive,omitempty"`
	Name     string `json:"Name,omitempty"`
	Volatile bool   `json:"Volatile,omitempty"`
}

type RegistryValueV2 struct {
	// JJH Check the types in this structure
	Key         *RegistryKeyV2 `json:"Key,omitempty"`
	Name        string         `json:"Name,omitempty"`
	Type        string         `json:"Type,omitempty"`
	StringValue string         `json:"StringValue,omitempty"` //  One and only one value type must be set.
	BinaryValue string         `json:"BinaryValue,omitempty"`
	DWordValue  int32          `json:"DWordValue,omitempty"`
	QWordValue  int32          `json:"QWordValue,omitempty"`
	CustomType  int32          `json:"CustomType,omitempty"` //  Only used if RegistryValueType is CustomType  The data is in BinaryValue
}

type RegistryChangesV2 struct {
	AddValues  []RegistryValueV2 `json:"AddValues,omitempty"`
	DeleteKeys []RegistryKeyV2   `json:"DeleteKeys,omitempty"`
}

type ContainerConfigV2 struct {
	GuestOS           *GuestOsV2                             `json:"GuestOS,omitempty"`
	Storage           *ContainersResourcesStorageV2          `json:"Storage,omitempty"`
	MappedDirectories []ContainersResourcesMappedDirectoryV2 `json:"MappedDirectories,omitempty"`
	MappedPipes       []ContainersResourcesMappedPipeV2      `json:"MappedPipes,omitempty"`
	Memory            *ContainersResourcesMemoryV2           `json:"Memory,omitempty"`
	Processor         *ContainersResourcesProcessorV2        `json:"Processor,omitempty"`
	Networking        *ContainersResourcesNetworkingV2       `json:"Networking,omitempty"`
	HvSocket          *ContainersResourcesHvSocketV2         `json:"HvSocket,omitempty"`
	RegistryChanges   *RegistryChangesV2                     `json:"RegistryChanges,omitempty"`
}
