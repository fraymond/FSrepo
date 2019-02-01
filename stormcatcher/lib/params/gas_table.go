// Copyright 2016 The MOAC-core Authors
// This file is part of the MOAC-core library.
//
// The MOAC-core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The MOAC-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the MOAC-core library. If not, see <http://www.gnu.org/licenses/>.

package params

type GasTable struct {
	ExtcodeSize uint64
	ExtcodeCopy uint64
	Balance     uint64
	SLoad       uint64
	Calls       uint64
	Suicide     uint64

	ExpByte uint64

	// CreateBySuicide occurs when the
	// refunded account is one that does
	// not exist. This logic is similar
	// to call. May be left nil. Nil means
	// not charged.
	CreateBySuicide uint64
}

var (
	// GasTablePangu contain the gas prices for
	// the pangu phase.
	//homestead
	// GasTablePangu = GasTable{
	// 	ExtcodeSize: 20,
	// 	ExtcodeCopy: 20,
	// 	Balance:     20,
	// 	SLoad:       50,
	// 	Calls:       40,
	// 	Suicide:     0,
	// 	ExpByte:     10,
	// }

	// GasTablePangu contain the gas re-prices for
	// the pangu phase.
	// GasTableEIP150
	// TODO rename to GasTableEIP150
	// GasTablePanguGasRepriceFork = GasTable{
	// 	ExtcodeSize: 700,
	// 	ExtcodeCopy: 700,
	// 	Balance:     400,
	// 	SLoad:       200,
	// 	Calls:       700,
	// 	Suicide:     5000,
	// 	ExpByte:     10,

	// 	CreateBySuicide: 25000,
	// }

	//	GasTableEIP158 = GasTable{
	//With high ExpBytes gas
	GasTablePangu = GasTable{
		ExtcodeSize: 700,
		ExtcodeCopy: 700,
		Balance:     400,
		SLoad:       200,
		Calls:       700,
		Suicide:     5000,
		ExpByte:     50,

		CreateBySuicide: 25000,
	}
)
