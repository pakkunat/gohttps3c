# go binary path
EXE = $(shell which go)

# binary name
BIN = https3c

# library path
#LIB = github.com/lib/pq

# certification files
#CRT = certificate.crt 
#KEY = private.key

# sources
SRC = https3c.go

# task
# compile
$(BIN): clean fmt preprocess
	@$(EXE) build

# format
fmt:
	@$(EXE)fmt -w $(SRC)

# preprocess
preprocess:
#	@$(EXE) get $(LIB) 

# clean
clean:
ifeq ("$(shell ls | grep -x $(BIN))", "$(BIN)")
	@rm $(BIN)
endif

#.PHONY
#.PHONY: preprocess
