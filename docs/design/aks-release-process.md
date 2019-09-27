# AKS Release Process Proposal

## Problem Statement

Today, AKS regularly updates its dependency upon AKS Engine by tracking the public release velocity of AKS Engine, in order to rapidly get the changes it needs that reside in the AKS Engine surface area. Because AKS Engine introduces non-AKS-required changes at a high velocity, in practice this approach brings more change, more frequently, than AKS actually desires, and thus increases the risk to AKS that some part of the changeset not strictly required by AKS may introduce undesired changes or regressions. So long as AKS strives to match AKS Engine release velocity, it will be required to regularly accept and validate that *all* AKS Engine changes work in the context of an AKS downstream library dependency.

## Proposal

Enable AKS to regularly update its downstream AKS Engine dependencies according to its own specific requirements.

Rather than requiring AKS to pull down *all* changes as a result of matching AKS Engine release velocity, we are proposing that:

1. AKS maintain its own release channel for rapid (aka weekly) iterations against its AKS Engine surface area, and
2. AKS regularly (aka monthly) re-align with the "mainline" AKS Engine release channel.

#1 above addresses the Problem Statement for most release iterations. #2 ensures that the AKS <-- AKS Engine relationship doesn't deviate too far from mainline, and that AKS isn't effectively maintaining a fork of AKS Engine.

## How this will work

Every month, we will cut a new AKS release @ n-1 compared to the "mainline", public AKS Engine release channel. For example:

- For an "October" release: we derive a new AKS release branch from AKS Engine n-1:
  - e.g., AKS Engine is at v0.42._n_
  - e.g., we cut a new AKS release branch from AKS Engine the latest v0.41 release
- Any additional commits in AKS Engine ahead of this new AKS release branch will be cherry-picked by AKS.
- For the duration of the "October" release cycle, AKS continues to cherry-pick changes from `aks-engine@master` into its "October" release branch.
- When "November" arrives, AKS cuts a new release branch from AKS Engine according to the above n-1 calculation.

At a glance, that's the process.

## How this will work successfully

- Additional testing/validation will be required for the monthly "catch-up" AKS releases. This additional validation overhead will be scheduled so as not to slow down feature and bugfix velocity during these intervals.
  - I.e., these releases will include _all_ AKS Engine changes at the AKS Engine minor release that AKS derives its new release from, not just AKS-specific changes
- All AKS-targeted changes to AKS Engine will be introduced via the normal PR-against-master workflow. In other words, the AKS release channel will not be a first class commit destination for novel additions/deletions to code; all changes will be cherry-picked from master.
  - This will prevent AKS-targeted changes that benefit AKS Engine generally from (1) not making it to AKS Engine, and (2) more importantly, this will prevent AKS-required changes to the AKS Engine surface area from being inadvertently overwritten during the monthly AKS Engine n-1 reconciliation.
