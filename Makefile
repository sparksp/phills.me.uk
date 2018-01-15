SHELL = /bin/sh
target = public

all : $(target) ;

$(target) :
	hugo -d $@

.PHONY : clean
clean :
	-rm -r $(target)

.PHONY : serve
serve :
	HUGO_GOOGLE_ANALYTICS=dev hugo serve --buildDrafts --navigateToChanged --watch
