// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate jsonenums -type=ShirtSize -type-prefix Size -lowercase

type ShirtSize byte

const (
	SizeNA ShirtSize = iota
	SizeXS
	SizeS
	SizeM
	SizeL
	SizeXL
)

//go:generate jsonenums -type=WeekDay -type-prefix Day -lowercase

type WeekDay int

const (
	DayMonday WeekDay = iota
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)

// func main() {
//     v := struct {
//         Size ShirtSize
//         Day  WeekDay
//     }{SizeM, DayFriday}
//     if err := json.NewEncoder(os.Stdout).Encode(v); err != nil {
//         log.Fatal(err)
//     }
//
//     input := `{"Size":"xl", "Day":"Dimarts"}`
//     if err := json.NewDecoder(strings.NewReader(input)).Decode(&v); err != nil {
//         log.Fatal(err)
//     }
//     fmt.Printf("decoded %s as %+v\n", input, v)
// }
