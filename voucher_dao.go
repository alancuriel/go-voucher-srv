package main

type VoucherDao interface {
	CreateVoucherEntry(entry VoucherEntry) error
	ListAllVoucherEntry() ([]*VoucherEntry, error)
}
