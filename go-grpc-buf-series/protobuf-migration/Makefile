.PHONY: generate generate-user generate-billing generate-common lint breaking

generate:
	buf generate

generate-user:
	buf generate --path proto/user/v1 --path proto/user/v2

generate-billing:
	buf generate --path proto/billing/v1

generate-common:
	buf generate --path proto/common/v1

lint:
	buf lint

breaking:
	buf breaking --against '.git#branch=main'
