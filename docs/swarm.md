# Microsoft Azure Container Service Engine - Swarm Walkthrough

## Deployment

Here are the steps to deploy a simple Swarm cluster:

1. [Install aks-engine](acsengine.md#downloading-and-building-aks-engine)
2. [Generate your SSH key](ssh.md#ssh-key-generation)
3. [Edit the Swarm example](../examples/swarm.json) and fill in the blank strings
4. [Generate the template](acsengine.md#generate-templates)
5. [Deploy the output azuredeploy.json and azuredeploy.parameters.json](../docs/acsengine.md#deploy-templates)

## Walkthrough

Once your Swarm cluster has been deployed you will have a resource group containing:

1. a set of 1,3, or 5 masters in a master availability set.  Each master's SSH can be accessed via the public dns address at ports 2200..2204. First master's SSH can also be accessed via public dns address on port 22.

2. a set of agents in a VM scale set (VMSS).  The agent VMs can be accessed through a master.  See [agent forwarding](ssh.md#key-management-and-agent-forwarding-with-windows-pageant) for an example of how to do this.

The following image shows the architecture of a container service cluster with 3 masters, and 3 agents:

 ![Image of Swarm container service on azure](images/swarm.png)

 All VMs are in the same VNET where the masters are on private subnet 172.16.0.0/24 and the agents are on the private subnet, 10.0.0.0/16, and fully accessible to each other.

## Create your First Two Swarm Docker Applications: hello-world, and Docker web app

After completing this walkthrough you will know how to:
 * display information from Swarm,
 * deploy a simple Docker hello-world app using docker-compose,
 * and deploy a simple Docker web app publically available to the world.


 1. After successfully deploying the template write down the two output master and agent FQDNs (Fully Qualified Domain Name).
    1. If using Powershell or CLI, the output parameters are the last values printed.
    2. If using Portal, to get the output you need to:
        1. navigate to "resource group"
        2. click on the resource group you just created
        3. then click on "Succeeded" under *last deployment*
        4. then click on the "Microsoft.Template"
        5. now you can copy the output FQDNs and sample SSH commands

        ![Image of docker scaling](images/findingoutputs.png)

 2. SSH to port 2200 of the master FQDN. See [agent forwarding](ssh.md#key-management-and-agent-forwarding-with-windows-pageant) for an example of how to do this.

 3. Set the DOCKER_HOST environment variable to `:2375`: e.g. ```export DOCKER_HOST=:2375```

 4. Type `docker info` to see the status of the agent nodes.
 ![Image of docker info](images/dockerinfo.png)

 5. Type `docker run -it hello-world` to see the hello-world test app run on one of the agents (the '-it' switches ensure output is displayed on your client)

 6. Now let's create a simple web app and expose to the world.  Start by using your favorite linux file editor to create a file named `docker-compose.yml` with the following contents:
    ```
    web:
      image: "yeasy/simple-web"
      ports:
        - "80:80"
      restart: "always"
    ```
 7. type `docker-compose up -d` to create the simple web server.  This will take a few minutes to pull the image

 8. once completed, type `docker ps` to see the running image.

 ![Image of docker ps](images/dockerps.png)

 9. in your web browser hit the AGENTFQDN endpoint (**not the master FQDN**) you recorded in step #1 and you should see the following page, with a counter that increases on each refresh.

 ![Image of the web page](images/swarmbrowser.png)

 10. You can now scale the web application.  For example, if you have 3 agents, you can type `docker-compose scale web=3`, and this will scale to the rest of your agents.  Note that in this example you can only scale up to the number of agents that you have since each container requires port 80, so if you deployed a single agent, you won't be able to scale up.  The Azure load balancer will automatically pick up the new containers.

 ![Image of docker scaling](images/dockercomposescale.png)

# Learning More

Here are recommended links to learn more about Swarm, Docker, and Docker Compose:

1. [Docker](https://docs.docker.com/) - learn more through Docker documentation.

2. [Docker Swarm](https://docs.docker.com/swarm/overview/) - learn more about Docker Swarm.

3. [Docker Compose](https://docs.docker.com/compose/overview/) - Learn more about Docker Compose.
