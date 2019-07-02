
REM Stop the docker & kubelet service
powershell "Stop-Service docker -Force"
powershell "Stop-Service kubelet -Force"