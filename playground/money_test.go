// reference: https://golangprojectstructure.com/representing-money-and-currency-go-code/
package playground_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"
)

type currency uint8

const (
	currencyIDR currency = iota // Rupiah Indonesia
	currencyVND                 // Dong Vietnam
)

// return IDR -> https://api.exchangerate.host/latest?base=IDR
func (c currency) String() string {
	switch c {
	case currencyIDR:
		return "IDR"
	case currencyVND:
		return "VND"
	}

	return ""
}

type money struct {
	value uint64
	curr  currency
	mutex *sync.RWMutex
}

// return *money
// because of function Add
func NewMoney(dollars, cents uint64, curr currency) *money {
	return &money{
		value: (dollars * 100) + cents,
		curr:  curr,
		mutex: &sync.RWMutex{},
	}
}

// function to move money
func (m *money) Add(m2 *money, move bool) {
	m2.mutex.Lock()
	defer m2.mutex.Unlock()
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.value += m2.value

	if move {
		m2.value = 0
	}
}

func (m money) String() string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var builder strings.Builder

	switch m.curr {
	case currencyIDR:
		builder.WriteString("Rp")
	case currencyVND:
		builder.WriteString("₫")
	}

	builder.WriteString(strconv.FormatUint(uint64(m.value)/100, 10))
	builder.WriteByte('.')

	cents := strconv.FormatUint(uint64(m.value)%100, 10)
	if len(cents) == 1 {
		builder.WriteByte('0')
	}
	builder.WriteString(cents)

	return builder.String()
}

type convertCurrencyResult struct {
	Rates map[string]float64
}

func (m *money) ConvertCurrency(curr currency) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	oldCurr := m.curr

	if curr == oldCurr {
		return nil
	}

	res, err := http.Get("https://api.exchangerate.host/latest?base=" + oldCurr.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("status code %d: %s", res.StatusCode, res.Status)
	}

	var result convertCurrencyResult

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return err
	}

	rate, foundRate := result.Rates[curr.String()]
	if !foundRate {
		return fmt.Errorf(
			"cannot load currency data for converting from %s to %s currency",
			oldCurr.String(),
			curr.String(),
		)
	}

	m.value = uint64(float64(m.value) * rate)
	m.curr = curr

	return nil
}

func TestAdd(t *testing.T) {
	m := NewMoney(1000, 0, currencyIDR)
	m2 := NewMoney(2000, 100, currencyIDR)

	fmt.Printf("I used to have %s.\n", m)
	fmt.Printf("You used to have %s.\n", m2)

	m.Add(m2, true)
	fmt.Printf("\nI now have %s.\n", m)
	fmt.Printf("You now have %s.\n", m2)
}

// using uint64
// 1000 IDR = 1515 VND = 999 IDR
func TestConvertCurrencyWithUint64(t *testing.T) {
	valueInIDR := uint64(1000)
	IDRtoVNDrate := 1.515609

	valueIDRtoVND := uint64(float64(valueInIDR) * IDRtoVNDrate)
	fmt.Printf("%d IDR = %d VND\n", valueInIDR, valueIDRtoVND)

	VNDtoIDRrate := 0.659801
	valueVNDtoIDR := uint64(float64(valueIDRtoVND) * VNDtoIDRrate)

	fmt.Printf("%d VND = %d IDR\n", valueIDRtoVND, valueVNDtoIDR)
}

// using float64
// 1000 IDR = 1515.609 VND = 1000,000334 IDR
func TestConvertCurrencyWithFloat64(t *testing.T) {
	valueInIDR := float64(1000)
	IDRtoVNDrate := 1.515609

	valueIDRtoVND := valueInIDR * IDRtoVNDrate
	fmt.Printf("%f IDR = %f VND\n", valueInIDR, valueIDRtoVND)

	VNDtoIDRrate := 0.659801
	valueVNDtoIDR := valueIDRtoVND * VNDtoIDRrate

	fmt.Printf("%f VND = %f IDR\n", valueIDRtoVND, valueVNDtoIDR)
}

func TestConvertCurrency(t *testing.T) {
	m := NewMoney(1000, 0, currencyIDR)

	fmt.Printf("I used to have %s. value %d\n", m, m.value)

	if err := m.ConvertCurrency(currencyVND); err != nil {
		panic(err)
	}

	fmt.Printf("I then had %s. value %d\n", m, m.value)

	if err := m.ConvertCurrency(currencyIDR); err != nil {
		panic(err)
	}

	fmt.Printf("I now have %s. value %d\n", m, m.value)
}
