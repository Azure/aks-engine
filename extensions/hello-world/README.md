# hello-world Extension

Sample hello-world extension.  Calls the following on the node:

```
 echo "hello"
```

You can validate that the extension was run by running (make sure you have tunneled into the master):
```
ls -l /var/log
```

# Configuration
|Name|Required|Acceptable Value|
|---|---|---|
|name|yes|hello-world|
|version|yes|v1|
|extensionParameters|no||
|rootURL|optional||
|script|required|hello.sh|

# Example
``` javascript
    "masterProfile": {
      ...
      "extensions": [
        {
          "name": "hello-world",
          "singleOrAll": "single"
        }
     ]
    },
    ...
    "extensionProfiles": [
      {
        "name": "hello-world",
        "version": "v1",
        "script": "hello.sh"
      }
    ]


```

# Supported Orchestrators
All
