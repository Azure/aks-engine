# Support
## How to file issues and get help

This project uses GitHub Issues to track bugs and feature requests. Please search the [existing issues][github-issues] before filing new issues to avoid duplicates. For new issues, file your bug or feature request as a new Issue.

The AKS Engine project maintainers will respond to the best of their abilities.

## Support FAQ

### How often are new AKS Engine releases published?

There is no definite release cadence. Project maintainers use the following as guidance to prioritize testing, validating, and publishing a new release:

- Are there new, validated versions of Kubernetes that haven't yet been included in an AKS Engine release?
- Are there improvements to the "default" Kubernetes configuration surface area?
- Is there a remediated Kubernetes CVE?
- Has a significant amount of time (say, more than 2 months) transpired since the last released VHD (Ubuntu OS image)?
- Has a bug with broad impact been fixed?
- Is there any new Kubernetes + Azure integration present and validated in the main branch, but not yet released?
- Is the community asking for a new release, for any of the above reasons?

If the answer to one or more of the above is true, then we may consider a release. In practice, expect a new release with a recently built VHD including pre-downloaded Kubernetes artifacts every 3-5 weeks.

### Does AKS Engine support OS patching for security vulnerabilities?

No. In general, the most recent release of AKS Engine will include a VHD derived from a build of Canonical Ubuntu about a week before the release is published. In other words, the maintainers will build a VHD from the "latest" tag of the supported Ubuntu releases, and then run about a week's worth of testing and validation prior to release. In practice, this means that on the day of a new AKS Engine release, OS security patches will be roughly a week behind of "latest", which may or may not have any practical effect (the Canonical Ubuntu "latest" tag isn't necessarily changed every week). As time progresses beyond that initial release interval, the most recent version of AKS Engine will eventually be several weeks behind "latest" when building a new cluster. If this does not meet your OS security requirements, we recommend curating your own Shared Image Gallery with OS images running one of the AKS Engine supported Ubuntu releases, with your desired OS patches already applied. See [the `imageReference` documentation here for more information](docs/topics/clusterdefinitions.md).

## More information

For help and questions about using this project, please:

  - Check out [Frequently Asked Questions][faq] about the AKS Engine project and tool
  - Read the [documentation][docs]
  - Join us on the [#aks-engine-users][aks-engine-users-slack] Slack channel
## Microsoft Support Policy

Support for this project is limited to the resources listed above.

[aks-engine-users-slack]: https://kubernetes.slack.com/archives/CU3N85WJK
[docs]: https://github.com/Azure/aks-engine/tree/master/docs
[faq]: https://github.com/Azure/aks-engine/blob/master/docs/faq.md
[github-issues]: https://github.com/Azure/aks-engine/issues
