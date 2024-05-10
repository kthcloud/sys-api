// Code generated by tygo. DO NOT EDIT.

//////////
// source: capacities.go

export interface TimestampedCapacities {
  capacities: Capacities;
  timestamp: string;
}
export interface Capacities {
  cpuCore: CpuCoreCapacities;
  ram: RamCapacities;
  gpu: GpuCapacities;
  hosts: HostCapacities[];
}
export interface ClusterCapacities {
  cluster: string;
  RAM: RamCapacities;
  CpuCore: CpuCoreCapacities;
}
export interface HostGpuCapacities {
  count: number /* int */;
}
export interface HostRamCapacities {
  total: number /* int */;
}
export interface HostCapacities {
  HostBase: HostBase;
  Capacities: any /* host_api.Capacities */;
}
export interface RamCapacities {
  total: number /* int */;
}
export interface CpuCoreCapacities {
  total: number /* int */;
}
export interface GpuCapacities {
  total: number /* int */;
}

//////////
// source: discovery.go

export interface HostRegisterParams {
  /**
   * Name is the host name of the node
   */
  name: string;
  /**
   * DisplayName is the human readable name of the node
   * This is optional, and is set to Name if not provided
   */
  displayName: string;
  ip: string;
  /**
   * Port is the port the node is listening on for API requests
   */
  port: number /* int */;
  zone: string;
  /**
   * Token is the discovery token validated against the config
   */
  token: string;
}
export interface ClusterRegisterParams {
}

//////////
// source: error.go

export interface BindingError {
  validationErrors: { [key: string]: string[]};
}

//////////
// source: gpu_info.go

export interface GpuInfo {
  hosts: HostGpuInfo[];
}
export interface TimestampedGpuInfo {
  gpuInfo: GpuInfo;
  timestamp: string;
}
export interface HostGpuInfo {
  HostBase: HostBase;
  gpus: any /* host_api.GpuInfo */[];
}

//////////
// source: host.go

export interface HostBase {
  name: string;
  displayName: string;
  /**
   * Zone is the name of the zone where the host is located.
   * This field might not yet be present in all responses, in which case ZoneID should be used.
   */
  zone?: string;
}
export interface HostInfo {
  HostBase: HostBase;
}

//////////
// source: stats.go

export interface Stats {
  k8s: K8sStats;
}
export interface TimestampedStats {
  stats: Stats;
  timestamp: string;
}
export interface K8sStats {
  podCount: number /* int */;
  clusters: ClusterStats[];
}
export interface ClusterStats {
  cluster: string;
  podCount: number /* int */;
}

//////////
// source: status.go

export interface Status {
  hosts: HostStatus[];
}
export interface TimestampedStatus {
  status: Status;
  timestamp: string;
}
export interface HostStatus {
  HostBase: HostBase;
  Status: any /* host_api.Status */;
}
