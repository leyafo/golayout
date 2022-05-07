package crypt

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestMain(m *testing.M) {
	err := SetKey([]byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDQWALs5+ykkwAKM1QVJHjmQe0FynaDZOfmOiCqVwVJK8hfLOde
+PSb2V2US/Q3+gHhiZWggGSdOs0qYwshRDZcf6CXYyxpyDI8aSdsPxhEwIFWpyWN
AdWzNyCClfP2D0itsLcS3+UY4Zsd2t8VNOt9rMlBP8hsH6M+ornusrGjfwIDAQAB
AoGASPSYyaaJGjQTjn7c0a5823x4aE+2YlpiTh9KsvtX8YBwYMuTlZEt7qkV+MkE
Etnr8LNB/vsWwGwHzfDyw8pkEiu3/gmEkGK1lMQSkHT+NDKfIrMK9O4dPAw8OoaB
YoBAYd70oP6lHL5MbMTRhLU9ldMGVBFd9b2gc6eegwm9GlkCQQD49dxRb5DOsJJB
HOuhtuk5OoJbAE0lN4hi7ZVyrBx+9qX3uNErcu2uYOVXq39aAtsOgfst8KkYE73a
KO2cdnAtAkEA1jwkH3w85kw0LArDo9TvIrWHlNNFOgSrnC+MiXlj0FdcaK55tkB9
jdx+h+IrSAOhSLSVCmE0F3/IQ9chJN2B2wJBAJ2RxpLIQOeAe+C4NC6S/QOak3yT
MUB36FtssaT1Z8e3xg2GrOSKBgLTEvSs95p5qjmBbP+DjRJPFF8qflED6TUCQD9l
uMLgfx0fu+i0nsSixMmesqqmArxymV406//avmDvGVeZGkeGuiD6+S65DVnYSSg8
2EYkEchKdjctOI+yRTUCQQD0UO2nZMUED1bwiTpFfqb/LStXrp/MzyWRqt1Q91f3
sy2Bt/IQV/tPKmGvMQBGxY8WYFczylOvk3VSQgvYFHRv
-----END RSA PRIVATE KEY-----
	`))

	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestDeEncrypt(t *testing.T) {
	randBytes := make([]byte, 16)
	for i := 0; i < 100; i++ {
		rand.Read(randBytes)
		m, err := RsaEncrypt(randBytes)
		if err != nil {
			panic(err)
		}
		d, err := RsaDecrypt(m)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(randBytes, d) {
			panic("something wrong....")
		}
	}
}

func TestBase64DeEncrypt(t *testing.T) {
	randBytes := make([]byte, 16)
	for i := 0; i < 100; i++ {
		rand.Read(randBytes)
		m, err := RsaEncryptBase64(randBytes)
		if err != nil {
			panic(err)
		}
		d, err := RsaEncryptBase64(m)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(randBytes, d) {
			panic("something wrong....")
		}
	}
}

func TestSplitEncrypt(t *testing.T) {
	randBytes := make([]byte, 200)
	for i := 0; i < 100; i++ {
		rand.Read(randBytes)
		m, err := SplitEncrypt(randBytes)
		if err != nil {
			panic(err)
		}
		cryptTexts := bytes.Split(m, []byte(";"))
		buf := new(bytes.Buffer)
		for i := 0; i < len(cryptTexts); i++ {
			if len(cryptTexts[i]) == 0 {
				continue
			}
			p, err := DecryptBase64(cryptTexts[i])
			if err != nil {
				t.Error(err)
				return
			}
			_, err = buf.Write(p)
			if err != nil {
				t.Error(err)
				return
			}
		}

		if !bytes.Equal(randBytes, buf.Bytes()) {
			t.Logf("exepect: %v, actual: %v", randBytes, buf.Bytes())
		}
	}
}

func TestBase64strDecrypt(t *testing.T) {
	data := "i3QBzxS3o0512kaui2SBX/j/wgAzCS0psSAH/wpuLlWgt9HNmZd3u+HKANezdmlhcVreNoqPfpJZ7MOwkyB8VU8ZiMGDHx/xfVogIPDw7dRjHIPXch3Y+6pq43K5BduAz0EMslk499/D3G8FPZynkrQM77SxC9rFuEdj0ipXiYA="

	plainText, err := DecryptBase64Str(data)
	if err != nil {
		t.Error(err)
	}

	t.Logf(string(plainText))
}
