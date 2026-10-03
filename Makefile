SHELL = /bin/sh
target = public
HUGO ?= $(shell hvm status --printExecPathCached 2>/dev/null || command -v hugo)

all : $(target) ;

$(target) :
	$(HUGO) --gc -d $@

.PHONY : clean
clean :
	-rm -r $(target)

.PHONY : serve
serve :
	$(HUGO) serve --buildDrafts --navigateToChanged --watch
