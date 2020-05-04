# Kubernetes Walkthrough

Once your Kubernetes cluster has been created you will have a resource group containing:

1. One master accessible by SSH on port 22 or kubectl on port 443
2. A set of nodes in an availability set. The nodes can be accessed through a master. See [agent forwarding](https://developer.github.com/v3/guides/using-ssh-agent-forwarding/) for an example of how to do this.

The following image shows the architecture of a container service cluster with 1 master, and 2 agents:

![Image of Kubernetes cluster on azure](../static/img/kubernetes.png)

In the image above, you can see the following parts:

1. **Master Components** - The master runs the Kubernetes scheduler, api server, and controller manager.  Port 443 is exposed for remote management with the kubectl cli.
2. **Nodes** - the Kubernetes nodes run in an availability set.  Azure load balancers are dynamically added to the cluster depending on exposed services.
3. **Common Components** - All VMs run a kubelet, Docker, and a Proxy.
4. **Networking** - All VMs are assigned an ip address in the 10.240.0.0/16 network.  Each VM is assigned a /24 subnet for their pod CIDR enabling IP per pod.  The proxy running on each VM implements the service network 10.0.0.0/16.

All VMs are in the same private VNET and are fully accessible to each other.

## Create your First Kubernetes Service

After completing this walkthrough you will know how to:

* Access the Azure VM(s) running the Kubernetes control plane via SSH,
* Deploy a simple Docker application and expose to the world,
* The location of the kubeconfig file and how to access the Kubernetes cluster remotely,
* Use `kubectl exec` to run commands in a container,
* And finally access the Kubernetes Dashboard.

Follow these steps:

1. After successfully deploying the template write down the master FQDNs (Fully Qualified Domain Name).
   * If using Powershell or CLI, the output parameter is in the OutputsString section named 'masterFQDN'
   * If using Portal, to get the output you need to:
      1. Navigate to "resource group"
      2. Click on the resource group you just created
      3. Then click on "Succeeded" under *last deployment*
      4. Then click on the "Microsoft.Template"
      5. Now you can copy the output FQDNs and sample SSH commands
      ![Image of docker scaling](../static/img/portal-kubernetes-outputs.png)
1. SSH to the master FQDN obtained in step 1.
1. Explore your nodes and running pods:
   * To see a list of your nodes type `kubectl get nodes`.  If you want full detail of the nodes, add `-o yaml` to become `kubectl get nodes -o yaml`.
   * To see a list of running pods type `kubectl get pods --all-namespaces`.
1. Start your first Docker image by typing `kubectl run nginx --image nginx`.  This will start the nginx Docker container in a pod on one of the nodes.
1. Type `kubectl get pods -o yaml` to see the full details of the nginx deployment. You can see the host IP and the podIP.  The pod IP is assigned from the pod CIDR on the host.  Run curl to the pod IP to see the nginx output, eg. `curl 10.244.1.4`
  ![Image of curl to podIP](../static/img/kubernetes-nginx1.png)
1. The next step is to expose the nginx deployment as a Kubernetes service on the private service network 10.0.0.0/16:
   1. Expose the service with command `kubectl expose deployment nginx --port=80`.
   2. Get the service IP `kubectl get service`
   3. Run curl to the IP, eg. `curl 10.0.105.199`
   ![Image of curl to service IP](../static/img/kubernetes-nginx2.png)
1. The final step is to expose the service to the world.  This is done by changing the service type from `ClusterIP` to `LoadBalancer`:
   1. Edit the service: `kubectl edit svc/nginx`
   2. Change `type` from `ClusterIP` to `LoadBalancer` and save it. This will now cause Kubernetes to create an Azure Load Balancer with a public IP.
   3. The change will take about 2-3 minutes.  To watch the service change from "pending" to an external ip, type `watch 'kubectl get svc'`
   ![Image of watching the transition from pending to external ip](../static/img/kubernetes-nginx3.png)
   4. Once you see the external IP, you can navigate to it in your browser:
   ![Image of browsing to nginx](../static/img/kubernetes-nginx4.png)
1. The next step in this walkthrough is to show you how to remotely manage your Kubernetes cluster.  First download Kubectl to your machine and put it in your path:
   * [Windows Kubectl](https://storage.googleapis.com/kubernetes-release/release/v1.6.0/bin/windows/amd64/kubectl.exe)
   * [OSX Kubectl](https://storage.googleapis.com/kubernetes-release/release/v1.6.0/bin/darwin/amd64/kubectl)
   * [Linux](https://storage.googleapis.com/kubernetes-release/release/v1.6.0/bin/linux/amd64/kubectl)
1. The Kubernetes master contains the kube config file for remote access under the home directory ~/.kube/config.  Download this file to your machine, set the KUBECONFIG environment variable, and run kubectl to verify you can connect to cluster:
   * Windows to use pscp from [putty](http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html).  Ensure you have your certificate exposed through [pageant](https://developer.github.com/v3/guides/using-ssh-agent-forwarding/):
      ```sh
      # MASTERFQDN is obtained in step1
      pscp -P 22 azureuser@MASTERFQDN:.kube/config .
      SET KUBECONFIG=%CD%\config
      kubectl get nodes
      ```
   * OS X or Linux:
      ```sh
      # MASTERFQDN is obtained in step1
      scp azureuser@MASTERFQDN:.kube/config .
      export KUBECONFIG=`pwd`/config
      kubectl get nodes
      ```
1. The next step is to show you how to remotely run commands in a remote Docker container:
    1. Run `kubectl get pods` to show the name of your nginx pod
    2. Using your pod name, you can run a remote command on your pod.  eg. `kubectl exec nginx-701339712-retbj date`
    3. Try running a remote bash session. eg. `kubectl exec nginx-701339712-retbj -it bash`.
    The following screen shot shows these commands:
    ![Image of curl to podIP](../static/img/kubernetes-remote.png)
1. The final step of this tutorial is to show you the Kubernetes Dashboard:
    1. Run `kubectl proxy` to create a secure local proxy channel to your cluster.
    2. In your browser, access the Dashboard at `http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/`
    3. Browse around and explore your pods and services.
    ![Image of Kubernetes dashboard](../static/img/kubernetes-dashboard.png)
