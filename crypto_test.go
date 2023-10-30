package toolkit

import (
	"fmt"
	"testing"
)

func TestCrypto(t *testing.T) {
	type sample struct {
		Key       string
		Iv        string
		PlainText string
		Excepted  string
	}

	samples := []sample{
		//正常加解密测试
		{
			Key:       "12345678901234567890123456789012",
			Iv:        "1234567890123456",
			PlainText: "任何一种技艺达到完美~都会令人无法抗拒~",
			Excepted:  "iN0ElJauMiSjgvG3GlPa2PxdlnmZ/nkP55r0Ti9eQxG019//ViLasmSNsHrwmMwbpwqCmqdhgaJTc91qj8TxDQ==",
		},
		{
			Key:       "12345678901234567890123456789012",
			Iv:        "1234567890123456",
			PlainText: "",
			Excepted:  "tVIs1T89WnKMXFMUdO/BUA==",
		},
		//报错测试
		{
			Key:       "1234567890123456789AA0123456789012",
			Iv:        "1234567890123456",
			PlainText: "任何一种技艺达到完美~都会令人无法抗拒~",
		},
		{
			Key:       "",
			Iv:        "1234567890123456",
			PlainText: "任何一种技艺达到完美~都会令人无法抗拒~",
		},
		{
			Key:       "12345678901234567890123456789012",
			Iv:        "123456789012?3456",
			PlainText: "任何一种技艺达到完美~都会令人无法抗拒~",
		},
		{
			Key:       "12345678901234567890123456789012",
			Iv:        "",
			PlainText: "任何一种技艺达到完美~都会令人无法抗拒~",
		},
	}

	for i, s := range samples {
		fmt.Println("-> sample", i)
		fmt.Println("PlainText:     ", s.PlainText)
		result, err := Encrypt(s.PlainText, s.Key, s.Iv)
		if err != nil {
			fmt.Println("Encrypt error: ", err)
		} else {
			if result != s.Excepted {
				t.Errorf("Encrypt error: result is \"%s\" , which excepted to be \"%s\"\n", result, s.Excepted)
			}
		}
		fmt.Println("result:        ", result)

		raw, err := Decrypt(result, s.Key, s.Iv)
		if err != nil {
			fmt.Println("Decrypt error: ", err)
		}

		fmt.Println("raw:           ", raw)
	}

}
