// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (VirtualMachineFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "VirtualMachineFlavor resource contains common VirtualMachine configuration\nthat can be used by multiple VirtualMachine resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true\n+genclient",
		"spec": "VirtualMachineFlavorSpec for the flavor",
	}
}

func (VirtualMachineFlavorList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineFlavorList is a list of VirtualMachineFlavor resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineClusterFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "VirtualMachineClusterFlavor is a cluster scoped version of VirtualMachineFlavor resource.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true\n+genclient\n+genclient:nonNamespaced",
		"spec": "VirtualMachineFlavorSpec for the flavor",
	}
}

func (VirtualMachineClusterFlavorList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineClusterFlavorList is a list of VirtualMachineClusterFlavor resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachineFlavorSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "VirtualMachineFlavorSpec\n\n+k8s:openapi-gen=true1",
		"ioThreadsPolicy": "Optionally defines the IOThreadsPolicy to be used by the flavor.\n\n+optional",
		"launchSecurity":  "Optionally defines the LaunchSecurity to be used by the flavor.\n\n+optional",
	}
}

func (CPUFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "CPUFlavor\n\n+k8s:openapi-gen=true",
		"guest":                 "Number of vCPUs to expose to the guest.\nThe resulting CPU topology being derived from the optional PreferredCPUTopology attribute of CPUPreferences.",
		"model":                 "Model specifies the CPU model inside the VMI.\nList of available models https://github.com/libvirt/libvirt/tree/master/src/cpu_map.\nIt is possible to specify special cases like \"host-passthrough\" to get the same CPU as the node\nand \"host-model\" to get CPU closest to the node one.\nDefaults to host-model.\n+optional",
		"dedicatedCPUPlacement": "DedicatedCPUPlacement requests the scheduler to place the VirtualMachineInstance on a node\nwith enough dedicated pCPUs and pin the vCPUs to it.\n+optional",
		"numa":                  "NUMA allows specifying settings for the guest NUMA topology\n+optional",
		"isolateEmulatorThread": "IsolateEmulatorThread requests one more dedicated pCPU to be allocated for the VMI to place\nthe emulator thread on it.\n+optional",
		"realtime":              "Realtime instructs the virt-launcher to tune the VMI for lower latency, optional for real time workloads\n+optional",
	}
}

func (MemoryFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "FlavorMemory\n\n+k8s:openapi-gen=true",
		"guest":     "Guest allows to specifying the amount of memory which is visible inside the Guest OS.",
		"hugepages": "Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.\n+optional",
	}
}

func (VirtualMachinePreference) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachinePreference\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true\n+genclient",
	}
}

func (VirtualMachinePreferenceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VirtualMachinePreferenceList is a list of VirtualMachinePreference resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
		"items": "+listType=set",
	}
}

func (VirtualMachineClusterPreference) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineClusterPreference\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true\n+genclient\n+genclient:nonNamespaced",
	}
}

func (VirtualMachineClusterPreferenceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VirtualMachineClusterPreferenceList is a list of VirtualMachineClusterPreference resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
		"items": "+listType=set",
	}
}

func (VirtualMachinePreferenceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "VirtualMachinePreferenceSpec\n\n+k8s:openapi-gen=true",
		"cpu":      "+optional",
		"devices":  "+optional",
		"features": "+optional",
		"firmware": "+optional",
		"machine":  "+optional",
	}
}

func (CPUPreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "PreferencesCPU\n\n+k8s:openapi-gen=true",
		"preferredCPUTopology": "Defaults to\n+optional",
	}
}

func (DevicePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                                    "DevicePreferences contains various optional defaults for Devices.\n\n+k8s:openapi-gen=true",
		"preferredAutoattachGraphicsDevice":   "PreferredAutoattachGraphicsDevice optionally defines the preferred value of AutoattachGraphicsDevice\n\n+optional",
		"preferredAutoattachMemBalloon":       "PreferredAutoattachMemBalloon optionally defines the preferred value of AutoattachMemBalloon\n\n+optional",
		"preferredAutoattachPodInterface":     "PreferredAutoattachPodInterface optionally defines the preferred value of AutoattachPodInterface\n\n+optional",
		"preferredAutoattachSerialConsole":    "PreferredAutoattachSerialConsole optionally defines the preferred value of AutoattachSerialConsole\n\n+optional",
		"preferredDisableHotplug":             "PreferredDisableHotplug optionally defines the preferred value of DisableHotplug\n\n+optional",
		"preferredVirtualGPUOptions":          "PreferredVirtualGPUOptions optionally defines the preferred value of VirtualGPUOptions\n\n+optional",
		"preferredSoundModel":                 "PreferredSoundModel optionally defines the preferred model for Sound devices.\n\n+optional",
		"preferredUseVirtioTransitional":      "PreferredUseVirtioTransitional optionally defines the preferred value of UseVirtioTransitional\n\n+optional",
		"preferredInputBus":                   "PreferredInputBus optionally defines the preferred bus for Input devices.\n\n+optional",
		"preferredInputType":                  "PreferredInputType optionally defines the preferred type for Input devices.\n\n+optional",
		"preferredDiskBus":                    "PreferredDiskBus optionally defines the preferred bus for Disk Disk devices.\n\n+optional",
		"preferredLunBus":                     "PreferredLunBus optionally defines the preferred bus for Lun Disk devices.\n\n+optional",
		"preferredCdromBus":                   "PreferredCdromBus optionally defines the preferred bus for Cdrom Disk devices.\n\n+optional",
		"preferredDiskDedicatedIoThread":      "PreferredDedicatedIoThread optionally enables dedicated IO threads for Disk devices.\n\n+optional",
		"preferredDiskCache":                  "PreferredCache optionally defines the DriverCache to be used by Disk devices.\n\n+optional",
		"preferredDiskIO":                     "PreferredIo optionally defines the QEMU disk IO mode to be used by Disk devices.\n\n+optional",
		"preferredDiskBlockSize":              "PreferredBlockSize optionally defines the block size of Disk devices.\n\n+optional",
		"preferredInterfaceModel":             "PreferredInterfaceModel optionally defines the preferred model to be used by Interface devices.\n\n+optional",
		"preferredRng":                        "PreferredRng optionally defines the preferred rng device to be used.\n\n+optional",
		"preferredBlockMultiQueue":            "PreferredBlockMultiQueue optionally enables the vhost multiqueue feature for virtio disks.\n\n+optional",
		"preferredNetworkInterfaceMultiQueue": "PreferredNetworkInterfaceMultiQueue optionally enables the vhost multiqueue feature for virtio interfaces.\n\n+optional",
	}
}

func (FeaturePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "FeaturePreferences contains various optional defaults for Features.\n\n+k8s:openapi-gen=true",
		"preferredAcpi":       "PreferredAcpi optionally enables the ACPI feature\n\n+optional",
		"preferredApic":       "PreferredApic optionally enables and configures the APIC feature\n\n+optional",
		"preferredHyperv":     "PreferredHyperv optionally enables and configures HyperV features\n\n+optional",
		"preferredKvm":        "PreferredKvm optionally enables and configures KVM features\n\n+optional",
		"preferredPvspinlock": "PreferredPvspinlock optionally enables the Pvspinlock feature\n\n+optional",
		"preferredSmm":        "PreferredSmm optionally enables the SMM feature\n\n+optional",
	}
}

func (FirmwarePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                       "FirmwarePreferences contains various optional defaults for Firmware.\n\n+k8s:openapi-gen=true",
		"preferredUseBios":       "PreferredUseBios optionally enables BIOS\n\n+optional",
		"preferredUseBiosSerial": "PreferredUseBiosSerial optionally transmitts BIOS output over the serial.\n\nRequires PreferredUseBios to be enabled.\n\n+optional",
		"preferredUseEfi":        "PreferredUseEfi optionally enables EFI\n\n+optional",
		"preferredUseSecureBoot": "PreferredUseSecureBoot optionally enables SecureBoot and the OVMF roms will be swapped for SecureBoot-enabled ones.\n\nRequires PreferredUseEfi and PreferredSmm to be enabled.\n\n+optional",
	}
}

func (MachinePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "MachinePreferences contains various optional defaults for Machine.\n\n+k8s:openapi-gen=true",
		"preferredMachineType": "PreferredMachineType optionally defines the preferred machine type to use.\n\n+optional",
	}
}
