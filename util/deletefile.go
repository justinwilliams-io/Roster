package util

import (
	"fmt"
	"os"
)

func DeleteFile(filename string) error {
    file, err := os.Open("./files/" + filename)
    if err != nil {
        fmt.Println(err)
        return err
    }

    fmt.Println("Deleting File")

    file.Close()
    os.Remove("./files/" + filename)

    return nil
}
