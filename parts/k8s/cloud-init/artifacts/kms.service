[Unit]
Description=azurekms
Requires=docker.service
After=network-online.target

[Service]
Type=simple
Restart=always
TimeoutStartSec=0
ExecStart=/usr/bin/docker run \
  --net=host \
  --volume=/opt:/opt \
  --volume=/etc/kubernetes:/etc/kubernetes \
  --volume=/etc/ssl/certs/ca-certificates.crt:/etc/ssl/certs/ca-certificates.crt \
  --volume=/var/lib/waagent:/var/lib/waagent \
  mcr.microsoft.com/k8s/kms/keyvault:v0.0.9

[Install]
WantedBy=multi-user.target
