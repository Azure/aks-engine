# WinDSR

## Configuration

WinDSR is currently **disabled** by default.

WinDSR can be enabled by setting the `enableWinDSR` feature flag to `true` in the 
apimodel.json. When `enableWinDSR` is set to `true`, `'WinDSR=true','WinOverlay=false'`
will be appened to `--feature-gates` in Kubeproxy's arguments and `enable-dsr` will be
set to `true` in kubeproxy's arguments.

``` javascript
{
  "properties": {
    "featureFlags": {
      "enableWinDSR": true
    }
  }
}
```

Please note that WinDSR requires,
  1. Windows Server 2019 with 10C patch (KB4580390). It means that all Windows VHD older than 17763.1554.201109 does not support it.
  1. Kubernetes >= 1.20. `enable-dsr` is introduced into kubeproxy since 1.20.