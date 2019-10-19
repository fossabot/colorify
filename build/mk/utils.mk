utils: clean format help list nuke

clean: clean-bin clean-coverage clean-version clean-tests ## to clean up all generated files.

format: ## to format all Go files.
	@echo "$(INFO_CLR)$(MSG_PRFX) Formatting code$(MSG_SFX)$(NO_CLR)"
	@test -z "$$($(GO_FMT) -s -w $(GO_FLAGS) $(GO_FILES) 2>&1 | tee /dev/stderr)"

help: ## to get help about the targets.
	@$(call PROJECT_LOGO,$(OK_CLR)) 2>&1
	@echo "$(INFO_CLR)Please use \`make <target>\`, Available options for <target> are:$(NO_CLR)\n"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z._-]+:.*?## .*$$/ {sub("\\\\n", sprintf("\n%22c"," "), $$2); printf "  $(STAR) $(HELP_CLR)%-28s$(NO_CLR) %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort -u 2>&1
	@echo "\n$(INFO_CLR)Useful variables:$(NO_CLR)\n"
	@awk 'BEGIN { FS = "[:?]=" } /^## /{x = substr($$0, 4); getline; if (NF >= 2) printf "  $(PLUS) $(DISCLAIMER_CLR)%-29s$(NO_CLR) %s\n", $$1, x}' $(MAKEFILE_LIST) | sort -u 2>&1

list: ## to list all targets.
	@awk -F':' '/^[a-z0-9][^$#\/\t=]*:([^=]|$$)/ {split($$1,A,/ /);for(i in A)printf "$(LIST_CLR)%-30s$(NO_CLR)\n", A[i]}' $(MAKEFILE_LIST) | sort -u 2>&1

nuke: clean ## to do clean up and enforce removing the corresponding installed archive or binary.
	@echo "$(WARN_CLR)$(MSG_PRFX) 🧹 Cleaning up Go dependencies$(MSG_SFX)$(NO_CLR)"
	@$(GO) clean -i $(GO_FLAGS) ./... net 2>&1
