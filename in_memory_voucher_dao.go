package main

import "errors"

type InMemoryVoucherDao struct {
	Vouchers []*VoucherEntry
}

func (d *InMemoryVoucherDao) CreateVoucherEntry(entry VoucherEntry) error {
	if d == nil {
		return errors.New("No dao present, nil")
	}

	if d.Vouchers == nil {
		d.Vouchers = []*VoucherEntry{}
	}

	d.Vouchers = append(d.Vouchers, &entry)

	return nil
}

func (d *InMemoryVoucherDao) ListAllVoucherEntry() ([]*VoucherEntry, error) {
	if d == nil {
		return nil, errors.New("No dao present, nil")
	}

	if d.Vouchers == nil || len(d.Vouchers) < 0 {
		return []*VoucherEntry{}, nil
	}

	return d.Vouchers, nil
}
