package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "numencode <numerical value> -dtype <dtype>\n")
		flag.PrintDefaults()
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		flag.Usage()
		os.Exit(1)
	}
}

func run() error {
	dtype := flag.String("dtype", "", "data type of argument. accept: [short, float, double]")
	flag.Parse()

	if flag.NArg() != 1 || *dtype == "" {
		return errors.New("invalid arguments")
	}
	var input string
	b, err := encode(input, *dtype)
	if err != nil {
		return err
	}
	fmt.Println(b)
	return nil
}

func encode(s string, dtype string) ([]byte, error) {
	switch dtype {
	case "short":
		v, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16ToBytes(int16(v))
	case "float":
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return nil, err
		}
		return float32ToBytes(float32(v))
	case "double":
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		return float64ToBytes(v)
	default:
		return nil, errors.New("unknown dtype")
	}
}

func decode(b []byte, dtype string) (string, error) {
	switch dtype {
	case "short":
		v, err := bytesToInt16(b)
		if err != nil {
			return "", nil
		}
		return fmt.Sprint(v), nil
	case "float":
		v, err := bytesToFloat32(b)
		if err != nil {
			return "", nil
		}
		return fmt.Sprint(v), nil
	case "double":
		v, err := bytesToFloat64(b)
		if err != nil {
			return "", nil
		}
		return fmt.Sprint(v), nil
	default:
		return "", errors.New("unknown dtype")
	}
}

func bytesToFloat64(b []byte) (float64, error) {
	var v float64
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &v)
	return v, err
}

func float64ToBytes(v float64) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, v)
	return buf.Bytes(), err
}

func bytesToFloat32(b []byte) (float32, error) {
	var v float32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &v)
	return v, err
}

func float32ToBytes(v float32) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, v)
	return buf.Bytes(), err
}

func bytesToInt16(b []byte) (int16, error) {
	var v int16
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &v)
	return v, err
}

func int16ToBytes(v int16) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, v)
	return buf.Bytes(), err
}
