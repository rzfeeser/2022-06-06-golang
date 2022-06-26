/* Alta3 Research | RZFeeser
   Refer - no panic, refer or defer   */

package main

import(
    "os"
    "io"
)
    
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}

func main() {
             // new name of our file    // file we are copying
    CopyFile("/tmp/example-copy1.txt", "/tmp/example-copy.txt")
}

