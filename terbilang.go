package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ToTerbilang ...
// display from any number (e.g. 109209) to
// seratus sembilan ribu dua ratus sembilan
//
// Example:
//  fmt.Println(terbilang.ToTerbilang(102209))
//  // Output: seratus sembilan ribu dua ratus sembilan
func ToTerbilang(num int) string {
	var s string
	satuan := [12]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}
	if num < 12 {
		s = satuan[num]
	} else if num < 20 {
		s = fmt.Sprintf("%s belas", ToTerbilang(num-10))
	} else if num < 100 {
		s = fmt.Sprintf("%s puluh %s", ToTerbilang(num/10), ToTerbilang(num%10))
	} else if num < 200 { // ratus
		s = fmt.Sprintf("seratus %s", ToTerbilang(num-100))
	} else if num < 1000 {
		s = fmt.Sprintf("%s ratus %s", ToTerbilang(num/100), ToTerbilang(num%100))
	} else if num < 2000 { // ribu
		s = fmt.Sprintf("seribu %s", ToTerbilang(num-1000))
	} else if num < 1000000 {
		s = fmt.Sprintf("%s ribu %s", ToTerbilang(num/1000), ToTerbilang(num%1000))
	} else if num < 2000000 { // juta
		s = fmt.Sprintf("satu juta %s", ToTerbilang(num-1000000))
	} else if num < 1000000000 {
		s = fmt.Sprintf("%s juta %s", ToTerbilang(num/1000000), ToTerbilang(num%1000000))
	} else if num < 2000000000 { // milyar
		s = fmt.Sprintf("satu milyar %s", ToTerbilang(num-1000000000))
	} else if num < 1000000000000 {
		s = fmt.Sprintf("%s milyar %s", ToTerbilang(num/1000000000), ToTerbilang(num%1000000000))
	} else if num < 2000000000000 { // triliun
		s = fmt.Sprintf("satu triliun %s", ToTerbilang(num-1000000000000))
	} else if num < 1000000000000000 {
		s = fmt.Sprintf("%s triliun %s", ToTerbilang(num/1000000000000), ToTerbilang(num%1000000000000))
	}
	return strings.TrimSpace(s)
}

// ToTerbilangRp ...
// like ToTerbilang, but add "rupiah" as suffix.
//
// Example:
//  fmt.Println(terbilang.ToTerbilangRp(102209))
//  // Output: seratus sembilan ribu dua ratus sembilan rupiah
func ToTerbilangRp(num int) string {
	return fmt.Sprintf("%s rupiah", ToTerbilang(num))
}

func ToTerbilangStringRp(num string) string {
	numSplit := strings.Split(num, ".")
	prefix := "0"
	if len(numSplit) > 0 {
		numf, _ := strconv.Atoi(numSplit[0])
		if numf > 0 {
			prefix = ToTerbilang(numf)
		}
	}
	suffix := ""
	if len(numSplit) > 1 {
		numf, _ := strconv.Atoi(numSplit[1])
		if numf > 0 {
			suffix = ToTerbilang(numf)
		}
	}

	koma := ""
	if suffix != "" {
		koma = " koma "
	}
	return fmt.Sprintf("%s rupiah", prefix+koma+suffix)
}

func ToTerbilangString(num, curr string) string {
	numSplit := strings.Split(num, ".")
	prefix := "0"
	if len(numSplit) > 0 {
		numf, _ := strconv.Atoi(numSplit[0])
		if numf > 0 {
			prefix = ToTerbilang(numf)
		}
	}

	suffix := ""
	if len(numSplit) > 1 {
		numf, _ := strconv.Atoi(numSplit[1])
		fmt.Println(numf)
		if numf > 0 {
			suffix = ToTerbilang(numf)
		}
	}

	koma := ""
	if suffix != "" {
		koma = " koma "
	}
	return fmt.Sprintf("%s", BuilderStringWithoutSpace(prefix, koma, suffix, " ", curr))
}

// ToTerbilangSuffix ...
// like ToTerbilangRp, but you can add any word as suffix.
//
// Example:
//  fmt.Println(terbilang.ToTerbilangSuffix("ekor",102209))
//  // Output: seratus sembilan ribu dua ratus sembilan ekor
func ToTerbilangSuffix(suff string, num int) string {
	return fmt.Sprintf("%s %s", ToTerbilang(num), suff)
}

func BuilderStringWithoutSpace(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
