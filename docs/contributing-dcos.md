# Porting a new DC/OS version to AKS Engine

## 1. Locate the official ARM Template

Go to `https://dcos.io/docs/X.X/administration/installing/cloud/azure/`, where `X.X` should be replaced by the version you are looking to port.
In the documentation, you will the link to the ARM templates you are looking for.
The latest stable templates should be at `https://downloads.dcos.io/dcos/stable/azure.html`
Early Access at: `https://downloads.dcos.io/dcos/EarlyAccess/azure.html`
Etc.

## 2. Find the package GUIDs

Following the previous step, you should now have 3 ARM templates (1, 3 and 5 masters variants).
We now need to find the package GUID of each variant.
In each template you should find a string that looks like: `dcos-config--setup_<Some GUID>`, this GUID is what we are looking for.
Extract the GUIDs from the 3 differents templates, and them in `engine.go/getPackageGUID` for your specific DC/OS version.

In DC/OS 1.11 and greater, you should find a the contents for `/etc/mesosphere/setup-flags/repository-url` in each 3 ARM templates (1, 3 and 5 masters variants).
Extract the GUIDs from the 3 differents templates, and them in `dcos/dcoscustomdataXXX.t` for your specific DC/OS version.


## 3. Extract the cloud-config data from the template

In one of the template (no matter which one), grab the data from the MasterVM.osProfile.customData.
If you remove the concat operation, you should end up which a big string of unescaped JSON.
Unescape it (for example using this [online tool](http://www.freeformatter.com/javascript-escape.html#ad-output)), and convert it to yaml (you can use [json2yaml](https://www.json2yaml.com/)).
You should now have a clean yaml.

## 4. Create and customize the custom data file.

under the `parts` directory, create a new file called `dcoscustomdataXXX.t` replacing `XXX` by the correct version number.
Paste the yaml from the previous step inside.

In the new file, under the `runcmd` section you should find 4 sucessive `curl` calls downloading some `.deb` packages followed by a bash script installing each one of them. This is handled by `parts\dcos\dcosprovision.sh` in AKS Engine, so make sure the dependencies didn't change and replace the `curl` and `bash` calls by a link to the script.

For example, in DC/OS 1.9:
```yaml
- curl -fLsSv --retry 20 -Y 100000 -y 60 -o /var/tmp/1.deb https://az837203.vo.msecnd.net/dcos-deps/libipset3_6.29-1_amd64.deb
- curl -fLsSv --retry 20 -Y 100000 -y 60 -o /var/tmp/2.deb https://az837203.vo.msecnd.net/dcos-deps/ipset_6.29-1_amd64.deb
- curl -fLsSv --retry 20 -Y 100000 -y 60 -o /var/tmp/3.deb https://az837203.vo.msecnd.net/dcos-deps/unzip_6.0-20ubuntu1_amd64.deb
- curl -fLsSv --retry 20 -Y 100000 -y 60 -o /var/tmp/4.deb https://az837203.vo.msecnd.net/dcos-deps/libltdl7_2.4.6-0.1_amd64.deb
- sed -i "s/^Port 22$/Port 22\nPort 2222/1" /etc/ssh/sshd_config
- service ssh restart
- bash -c "try=1;until dpkg -i /var/tmp/{1,2,3,4}.deb || ((try>9));do echo retry \$((try++));sleep
  \$((try*try));done"
```

becomes

```yaml
- /opt/azure/containers/provision.sh
```

Additional modifications under `runcmd`:
* Replace every occurence of the Package GUID (that we found in step 2) by `DCOSGUID`.
* the `content` of the cmd with path `/etc/mesosphere/setup-flags/late-config.yaml` should be modified to accept AKS Engine bindings instead of variable where needed (look at a previous custom data file for reference).
* At the very end of the file, replace
```yaml
- content: ''
  path: "/etc/mesosphere/roles/master"
- content: ''
  path: "/etc/mesosphere/roles/azure_master"
- content: ''
  path: "/etc/mesosphere/roles/azure"
```

by

```yaml
- content: ''
  path: /etc/mesosphere/roles/azure
- content: 'PROVISION_STR'
  path: "/opt/azure/containers/provision.sh"
  permissions: "0744"
  owner: "root"
```

## 5. Adding the support of the new version inside to .go files

### pkg/engine/defaults.go

- Around line 30, add your `DCOSXXXBootstrapDownloadURL` variable (replace XXX with the version number), inside the `fmt.Sprintf()` function replace the second and third parameters with the version `EA, Stable, Beta, ...` and the commit hash.

> You can find the commit hash from the https://downloads.dcos.io/dcos/stable/X.XX.X/azure.html page.

Example for version 1.10
[https://downloads.dcos.io/dcos/stable/1.10.0/azure.html](https://downloads.dcos.io/dcos/stable/1.10.0/azure.html)

```
DCOS110BootstrapDownloadURL: fmt.Sprintf(AzureEdgeDCOSBootstrapDownloadURL, "stable", "e38ab2aa282077c8eb7bf103c6fff7b0f08db1a4"),
```

### pkg/engine/engine.go

- Around line 39, add `dcosCustomDataXXX    = "dcos/dcoscustomdataXXX.t"` variable

Example for version 1.10:
```
dcosCustomData110    = "dcos/dcoscustomdata110.t"
```

- Around line  578, add the code case block for your version.

Example for version 1.10:
```
case api.DCOSRelease1Dot10:
  dcosBootstrapURL = cloudSpecConfig.DCOSSpecConfig.DCOS110BootstrapDownloadURL
 ```

- Around line 1170, add your api case version.

 Example for version 1.10:
 ```
 case api.DCOSRelease1Dot10:
			switch masterCount {
			case 1:
				return "c4ec6210f396b8e435177b82e3280a2cef0ce721"
			case 3:
				return "08197947cb57d479eddb077a429fa15c139d7d20"
			case 5:
				return "f286ad9d3641da5abb622e4a8781f73ecd8492fa"
			}
 ```

 > In the return function, paste the package GUID from the step 2 for each cases.

- Around line 1558, add your api case version.

Example for version 1.10:
```
		case api.DCOSRelease1Dot10:
 			yamlFilename = dcosCustomData110
```

### pkg/engine/types.go

- Around line 40, add your the type for your new version.

Example for version 1.10 :
```
DCOS110BootstrapDownloadURL     string
```

### pkg/api/common/const.go

- Around line 59, declare a new const with your `DCOSRelease`

Example for version 1.10 :
```
// DCOSRelease1Dot10 is the major.minor string prefix for 1.9 versions of DCOS
DCOSRelease1Dot10 string = "1.10"
```

- Around line 72, add your `DCOSReleaseToVersion` in the map

Example for version 1.10 :
```
DCOSRelease1Dot10: "1.10.0",
```

### pkg/api/const.go

- Around line 76, add the const for your DCOS release

Example for version 1.10 :
```
// DCOSRelease1Dot10 is the major.minor string prefix for 1.10 versions of DCOS
DCOSRelease1Dot10 string = "1.10"
```

### pkg/api/convertertoapi.go

- Around line 572 and 601 (two places) add the case for your release

Example for version 1.10 :
```
case DCOSRelease1Dot10, DCOSRelease1Dot9, DCOSRelease1Dot8:
```
```
case DCOSRelease1Dot10, DCOSRelease1Dot9, DCOSRelease1Dot8, DCOSRelease1Dot7:
```

### pkg/api/v20170701/validate.go

- Around line 33, add the case for your release

Example for version 1.10 :
```
case common.DCOSRelease1Dot10:
```

### pkg/api/vlabs/validate.go

- Around line 37, add the case for your release

Example for version 1.10 :
```
case common.DCOSRelease1Dot10:
```


## Conclusion

We encourage you to look at previous PR as example, listed bellow :

- [Adding DC/OS 1.10 stable version support #1439](https://github.com/Azure/aks-engine/pull/1439/files)
- [setting dcos test to 1.9 (current default)](https://github.com/Azure/aks-engine/pull/1443)
- [[DC/OS] Set 1.9 as default DCOS version and upgrade Packages](https://github.com/Azure/aks-engine/pull/457)
- [[DC/OS] Add support for DCOS 1.9 EA](https://github.com/Azure/aks-engine/pull/360)
- [DCOS 1.8.8 Support](https://github.com/Azure/aks-engine/pull/278)
