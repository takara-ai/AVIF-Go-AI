package avif

import (
	"encoding/binary"
	"io"
)

// AVIF box types
const (
	BoxTypeFtyp = "ftyp"
	BoxTypeMeta = "meta"
	BoxTypeMdat = "mdat"
	BoxTypeColr = "colr"
	BoxTypePixi = "pixi"
	BoxTypeIspe = "ispe"
	BoxTypeIpma = "ipma"
	BoxTypeIprp = "iprp"
	BoxTypeHdlr = "hdlr"
)

// BoxHeader represents the common header for all AVIF boxes
type BoxHeader struct {
	Size uint32
	Type string
}

// ReadBoxHeader reads a box header from the given reader
func ReadBoxHeader(r io.Reader) (BoxHeader, error) {
	var header BoxHeader
	if err := binary.Read(r, binary.BigEndian, &header.Size); err != nil {
		return header, err
	}
	typeBytes := make([]byte, 4)
	if _, err := io.ReadFull(r, typeBytes); err != nil {
		return header, err
	}
	header.Type = string(typeBytes)
	return header, nil
}

// WriteBoxHeader writes a box header to the given writer
func WriteBoxHeader(w io.Writer, size uint32, boxType string) error {
	if err := binary.Write(w, binary.BigEndian, size); err != nil {
		return err
	}
	_, err := w.Write([]byte(boxType))
	return err
}