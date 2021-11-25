VERSION = 0.5.1
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

build:
	go generate
	go build -ldflags '-X "main.VERSION=v$(VERSION)"'

release:
	go generate
	GOOS=darwin GOARCH=amd64 go build -ldflags '-X "main.VERSION=v$(VERSION)"' -o webflow-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags '-X "main.VERSION=v$(VERSION)"' -o webflow-darwin-arm64
	GOOS=linux GOARCH=amd64 go build -ldflags '-X "main.VERSION=v$(VERSION)"' -o webflow-linux-amd64

gen-impl:
	@[[ -z "$$IMPL_TYPE" ]] && echo "IMPL_TYPE is empty!" && exit 1; exit 0;
	@echo "> Generating $$GOFILE..";
	@echo "$$(head -n $$GOLINE $$GOFILE)" > $$GOFILE;
	cat ./flowimpl-tpl.go | sed "s#FlowImplBase#$$IMPL_TYPE#g" | \
		grep -v "^\(//\|package \|import \)" >> "$$GOFILE";

brew: build
	url=$$(webflow -c '$(FLOWS)'); \
		vim -c 'let @q="/url\<Esc>f\"va\"c\"'$$url'\"\<Esc>" | argdo normal @q | ZZ' ./webflow.rb;
	vim -c 'let @q="/version\<Esc>f\"va\"c\"v$(VERSION)\"\<Esc>" | argdo normal @q | ZZ' ./webflow.rb;
	sha256=$$(openssl sha256 ./webflow | cut -d= -f2 | cut -d' ' -f2);  \
		vim -c 'let @q="/sha256\<Esc>f\"va\"c\"'$$sha256'\"\<Esc>" | argdo normal @q | ZZ' ./webflow.rb;
