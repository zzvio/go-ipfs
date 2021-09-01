package io.zzv.jni;

import java.nio.ByteBuffer;
import java.nio.ByteOrder;

public class Main {
    static native byte[] call(byte[] arr);

    static void callback(int token) {
        System.out.println("Token: " + token);
    }

    public static void main(String[] args) throws Exception {
        System.load(System.getProperty("user.home") + "/.java/packages/lib/" + System.mapLibraryName("goipfs"));

        // Commandline arguments passed to GoPlugin to Run GoIPFS with these
        // arguments.
        String goArgs = "add /home/talha/go/src/github.com/ipfs/go-ipfs/cmd/ipfs/jni_md.h";

        startPlugin(goArgs);
        for (;;)
            Thread.sleep(1000);
    }

    private static void startPlugin(String args) {
        // StartPlugin Message
        byte[] arr = { (byte) 0, (byte) 0 }; // MsgTypeStartPlugin
        int token = 123456; // Mock token, actual implementation in client code
        arr = append(arr, intToByteArray(token));
        arr = append(arr, args.getBytes());
        call(arr);
    }

    private static byte[] append(byte[] a, byte[] b) {
        byte[] arr = new byte[a.length + b.length];
        System.arraycopy(a, 0, arr, 0, a.length);
        System.arraycopy(b, 0, arr, a.length, b.length);
        return arr;
    }

    private static byte[] intToByteArray(int i) {
        return ByteBuffer.allocate(4).order(ByteOrder.BIG_ENDIAN).putInt(i).array();
    }

}
