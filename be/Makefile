BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=main

export GOPATH=$(CURDIR)

myname:
	@echo "I am makefile"

makedir:
	@echo "start building tree..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

build:
	@echo "start building..."
	$(GOINSTALL) $(EXENAME)
	@echo "all DONE"

get:
	$(GOGET) github.com/gorilla/mux
	$(GOGET) github.com/go-sql-driver/mysql
	$(GOGET) github.com/denisenkom/go-mssqldb



clean:
	@echo "cleanning..."
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg

all: makedir build
