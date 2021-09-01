go-build:
	mkdir -p ~/.java/packages/lib/
	cd cmd/ipfs; CGO_CFLAGS="-I${PWD}" go build -o libgoipfs.so -buildmode=c-shared;
	cp cmd/ipfs/libgoipfs.so ~/.java/packages/lib/libgoipfs.so
	
java-build:
	cd gojni/src/main/java; javac io/zzv/jni/Main.java
run:
	cd gojni/src/main/java; java io.zzv.jni.Main
	