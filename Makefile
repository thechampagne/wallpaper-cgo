CC := gcc
CFLAGS := -Wall -Wextra -std=c99 -pthread
CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/wallpaper.a build/wallpaper.so

.PHONY: all
all: $(LIBS)

build/wallpaper.a: wallpaper.go
	$(CGO) $(STATIC) -o build/wallpaper.a $<

build/wallpaper.so: wallpaper.go
	$(CGO) $(SHARED) -o build/wallpaper.so $<

test: all
	$(CC) $(CFLAGS) test/main.c -o test/main -Ibuild -Lbuild -l:wallpaper.a
	test/main

clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete