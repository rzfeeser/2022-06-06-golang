/* Alta3 Research | RZFeeser
   Refer - using the defer statement  */

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
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

func main() {
    CopyFile("/tmp/example-copy2.txt", "/tmp/example-copy.txt")
}

