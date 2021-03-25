export GOOS=windows
export GOARCH=386
export CGO_ENABLED=1
export CC="i686-w64-mingw32-gcc -static-libgcc -static-libstdc++ -Wl,-Bstatic -lstdc++ -lpthread -Wl,-Bdynamic"
export CXX="i686-w64-mingw32-g++ -static-libgcc -static-libstdc++ -Wl,-Bstatic -lstdc++ -lpthread -Wl,-Bdynamic"
go build

# brew install mingw-w64 for installing C pipeline to compile on macos for windows
# CC --> Toolchain for C code
# CXX --> Toolchain for C++ Code
# -shared -static -static-libgcc -static-libstdc++ -lwinpthread