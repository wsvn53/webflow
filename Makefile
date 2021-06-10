VERSION = 0.4.4
FLOWS = headless false; \
	userdata "~/.webflow"; 		\
	open "https://git.wsen.me/utils/webflow/releases/new"; 	\
	wait "\#tag-name";	\
	setvalue "\#tag-name" "v$(VERSION)";	\
	setvalue `input[name="title"]` "webflow v$(VERSION)"; \
	click "\#dropzone"; \
	wait "div.dz-success-mark"; 	\
	wait 5000;	\
	click "button.ui.primary.button";	\
	wait 3000; 	\
	wait "div.content.active";	\
	eval "document.querySelector(\"div.content.active a span[title=webflow]\").parentNode.parentNode.href";  \
	wait 1000;

all:
	go generate
	go build

linux:
	GOOS=linux GOARCH=amd64 go build -o webflow-linux

gen-impl:
	@[[ -z "$$IMPL_TYPE" ]] && echo "IMPL_TYPE is empty!" && exit 1; exit 0;
	@echo "> Generating $$GOFILE..";
	@echo "$$(head -n $$GOLINE $$GOFILE)" > $$GOFILE;
	cat ./flowimpl-tpl.go | sed "s#FlowImplBase#$$IMPL_TYPE#g" | \
		grep -v "^\(//\|package \|import \)" >> "$$GOFILE";

brew:
	webflow -c '$(FLOWS)';