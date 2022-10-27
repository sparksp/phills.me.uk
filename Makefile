SHELL = /bin/sh
target = public

all : $(target) ;

$(target) :
	hugo --gc -d $@

.PHONY : clean
clean :
	-rm -r $(target)

.PHONY : serve
serve :
	hugo serve --buildDrafts --navigateToChanged --watch
