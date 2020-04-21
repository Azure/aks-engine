// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package common

type ContainerdConfig struct {
	OomScore int     `toml:"oom_score,omitempty"`
	Root     string  `toml:"root,omitempty"`
	Version  int     `toml:"version,omitempty"`
	Plugins  Plugins `toml:"plugins,omitempty"`
}

type ContainerdCNIPlugin struct {
	ConfTemplate string `toml:"conf_template,omitempty"`
}

type ContainerdRuntime struct {
	RuntimeType string `toml:"runtime_type,omitempty"`
}

type ContainerdPlugin struct {
	DefaultRuntimeName string                       `toml:"default_runtime_name,omitempty"`
	Runtimes           map[string]ContainerdRuntime `toml:"runtimes,omitempty"`
}

type IoContainerdGrpcV1Cri struct {
	SandboxImage string              `toml:"sandbox_image,omitempty"`
	CNI          ContainerdCNIPlugin `toml:"cni,omitempty"`
	Containerd   ContainerdPlugin    `toml:"containerd,omitempty"`
}

type Plugins struct {
	IoContainerdGrpcV1Cri IoContainerdGrpcV1Cri `toml:"io.containerd.grpc.v1.cri,omitempty"`
}

type DockerConfig struct {
	ExecOpts             []string                       `json:"exec-opts,omitempty"`
	DataRoot             string                         `json:"data-root,omitempty"`
	LiveRestore          bool                           `json:"live-restore,omitempty"`
	LogDriver            string                         `json:"log-driver,omitempty"`
	LogOpts              LogOpts                        `json:"log-opts,omitempty"`
	DefaultRuntime       string                         `json:"default-runtime,omitempty"`
	DockerDaemonRuntimes map[string]DockerDaemonRuntime `json:"runtimes,omitempty"`
}

type LogOpts struct {
	MaxSize string `json:"max-size,omitempty"`
	MaxFile string `json:"max-file,omitempty"`
}

type DockerDaemonRuntime struct {
	Path        string   `json:"path,omitempty"`
	RuntimeArgs []string `json:"runtimeArgs"`
}
