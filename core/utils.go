package core

import (
	"io/ioutil"
	"os"
)

type Vec2 struct {
    X int32
    Y int32
}

func Vec2Zeros() Vec2 {
    return Vec2{ X: 0, Y: 0 }
}

func SaveFileData(path string, data []byte) error {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }

    if _, err := f.Write(data); err != nil {
        return err
    }

    if err := f.Close(); err != nil {
        return err
    }

    return nil
}

func LoadFileData(path string) ([]byte, error) {
    fileBuffer, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    return fileBuffer, nil
}
