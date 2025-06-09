package binance

import "testing"

func TestSign(t *testing.T) {
	in := "hello world"
	secret := "123456"

	sign := Sign(in, secret)
	t.Logf("%+v", sign)
}

func BenchmarkSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sign("hello world", "123456")
	}
}
