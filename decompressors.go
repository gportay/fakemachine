package fakemachine

import (
	"io"
	"compress/gzip"

	// BSD 3-clause licenses
	"github.com/klauspost/compress/zstd"
	"github.com/ulikunitz/xz"
)

func ZstdDecompressor(dst io.Writer, src io.Reader) error {
	decompressor, err := zstd.NewReader(src)
	if err != nil {
		return err
	}
	defer decompressor.Close()

	_, err = io.Copy(dst, decompressor)
	return err
}

func XzDecompressor(dst io.Writer, src io.Reader) error {
	decompressor, err := xz.NewReader(src)
	if err != nil {
		return err
	}
	// There is not Close() API just yet, see
	// https://github.com/ulikunitz/xz/issues/45
	//defer decompressor.Close()

	_, err = io.Copy(dst, decompressor)
	return err
}

func GzipDecompressor(dst io.Writer, src io.Reader) error {
	decompressor, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer decompressor.Close()

	_, err = io.Copy(dst, decompressor)
	return err
}
