prow: prow-config prow-secrets prow-services
.PHONY: prow

prow-config:
	kubectl create cm config --from-file=config.yaml=config.yaml
	kubectl create cm plugins --from-file=plugins.yaml=plugins.yaml
.PHONY: prow-config

prow-config-update:
	kubectl create cm config --from-file=config.yaml=config.yaml -o yaml --dry-run | kubectl replace -f -
	kubectl create cm plugins --from-file=plugins.yaml=plugins.yaml -o yaml --dry-run | kubectl replace -f -
.PHONY: prow-config-update

prow-secrets:
	# hmac is used for encrypting Github webhook payloads.
	kubectl create secret generic hmac-token --from-file=hmac
	# oauth is used for merging PRs, adding/removing labels and comments.
	kubectl create secret generic oauth-token --from-file=oauth
.PHONY: prow-secrets

prow-services:
	kubectl create -f deck.yaml
	kubectl create -f hook.yaml
	kubectl create -f tide.yaml
	kubectl create -f ingress.yaml
.PHONY: prow-services
