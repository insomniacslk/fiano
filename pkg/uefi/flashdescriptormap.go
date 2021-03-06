// Copyright 2018 the LinuxBoot Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uefi

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	// FlashDescriptorMapMaxBase is the maximum base address for a flash descriptor
	// region
	FlashDescriptorMapMaxBase = 0xe0
)

// FlashDescriptorMap represent an Intel flash descriptor. This object provides
// accessors to the various descriptor fields.
type FlashDescriptorMap struct {
	// FLMAP0
	ComponentBase      uint8
	NumberOfFlashChips uint8
	RegionBase         uint8
	NumberOfRegions    uint8
	// FLMAP1
	MasterBase        uint8
	NumberOfMasters   uint8
	PchStrapsBase     uint8
	NumberOfPchStraps uint8
	// FLMAP2
	ProcStrapsBase          uint8
	NumberOfProcStraps      uint8
	IccTableBase            uint8
	NumberOfIccTableEntries uint8
	// FLMAP3
	DmiTableBase            uint8
	NumberOfDmiTableEntries uint8
	Reserved0               uint8
	Reserved1               uint8
}

// NewFlashDescriptorMap initializes a FlashDescriptor from a slice of bytes.
func NewFlashDescriptorMap(buf []byte) (*FlashDescriptorMap, error) {
	r := bytes.NewReader(buf)
	var descriptor FlashDescriptorMap
	if err := binary.Read(r, binary.LittleEndian, &descriptor); err != nil {
		return nil, err
	}
	return &descriptor, nil
}

func (d *FlashDescriptorMap) String() string {
	return fmt.Sprintf("FlashDescriptorMap{NumberOfRegions=%v, NumberOfFlashChips=%v, NumberOfMasters=%v, NumberOfPCHStraps=%v, NumberOfProcessorStraps=%v, NumberOfICCTableEntries=%v, DMITableEntries=%v}",
		d.NumberOfRegions,
		d.NumberOfFlashChips,
		d.NumberOfMasters,
		d.NumberOfPchStraps,
		d.NumberOfProcStraps,
		d.NumberOfIccTableEntries,
		d.NumberOfDmiTableEntries,
	)
}
