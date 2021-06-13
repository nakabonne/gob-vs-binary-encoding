package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
)

type DataPoint struct {
	// The actual value. This field must be set.
	Value float64
	// Unix timestamp.
	Timestamp int64
}

func main() {
	point := &DataPoint{Value: 0.1, Timestamp: 1}

	out1, err := EncodeDecodeWithBinary(point)
	if err != nil {
		panic(err)
	}
	fmt.Printf("encoding/binary decoded point: %#v\n\n", out1)

	out2, err := EncodeDecodeWithGob(point)
	if err != nil {
		panic(err)
	}
	fmt.Printf("encoding/gob decoded point: %#v\n", out2)
}

func EncodeDecodeWithBinary(point *DataPoint) (*DataPoint, error) {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, point); err != nil {
		return nil, err
	}
	//fmt.Println("encoding/binary encoded binary length:", buf.Len())
	//fmt.Println("encoding/binary encoded string:", buf.String())
	dst := &DataPoint{}
	if err := binary.Read(buf, binary.LittleEndian, dst); err != nil {
		return nil, err
	}
	return dst, nil
}

func EncodeDecodeWithGob(point *DataPoint) (*DataPoint, error) {
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(point); err != nil {
		return nil, err
	}
	//fmt.Println("encoding/gob encoded binary length:", buf.Len())
	//fmt.Println("encoding/gob encoded string:", buf.String())
	dst := &DataPoint{}
	if err := gob.NewDecoder(buf).Decode(dst); err != nil {
		return nil, err
	}
	return dst, nil
}
