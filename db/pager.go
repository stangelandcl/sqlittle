package db

import (
	"fmt"
)

const (
	sqlitePendingByte  = 0x40000000
	sqliteReservedByte = sqlitePendingByte + 1
	sqliteSharedFirst  = sqlitePendingByte + 2
	sqliteSharedSize   = 510
)

type pager interface {
	// load a page from storage.
	page(n int, pagesize int) ([]byte, error)
	// as it says
	Close() error
	// read lock
	RLock() error
	// unlock read lock
	RUnlock() error
	// true if there is any 'RESERVED' lock on this file
	CheckReservedLock() (bool, error)
}

type bufferPager struct {
	buf []byte
}

func newBufferPager(buf []byte) *bufferPager {
	return &bufferPager{buf: buf}
}

func (p *bufferPager) page(id int, pagesize int) ([]byte, error) {
	pos := int64(id-1) * int64(pagesize)
	if pos < 0 || int64(len(p.buf)) < pos+int64(pagesize) {
		return nil, fmt.Errorf("pager: invalid ReadAt offset %d", pos)
	}
	buf := make([]byte, pagesize)
	copy(buf, p.buf[pos:])
	return buf, nil
}
func (p *bufferPager) Close() error {
	return nil
}
func (p *bufferPager) RLock() error {
	return nil
}
func (p *bufferPager) RUnlock() error {
	return nil
}
func (p *bufferPager) CheckReservedLock() (bool, error) {
	return false, nil
}
