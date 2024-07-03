package avif

// Metadata struct to hold AVIF image information
type Metadata struct {
	Width          uint32
	Height         uint32
	ColorProfile   ColorProfile
	PixelInfo      PixelInfo
	ItemProperties map[uint16]ItemProperty
	MdatOffset     int64
	MdatSize       uint32
}

type ColorProfile struct {
	ColorPrimaries          uint16
	TransferCharacteristics uint16
	MatrixCoefficients      uint16
	FullRangeFlag           bool
}

type PixelInfo struct {
	NumChannels    uint8
	BitsPerChannel uint8
}

type ItemProperty struct {
	ID   uint16
	Type string
	// Add other properties as needed
}