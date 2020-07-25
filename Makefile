all:
	go generate
	go build

gen-impl:
	@[[ -z "$$IMPL_TYPE" ]] && echo "IMPL_TYPE is empty!" && exit 1; exit 0;
	@echo "> Generating $$GOFILE..";
	@echo "$$(head -n $$GOLINE $$GOFILE)" > $$GOFILE;
	cat ./flowimpl-tpl.go | sed "s#FlowImplBase#$$IMPL_TYPE#g" | \
		grep -v "^\(//\|package \|import \)" >> "$$GOFILE";