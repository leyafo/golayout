package eth

import (
	"testing"
)

func TestWallet(t *testing.T) {

	w := InitWallet("f4054ec6b71dda8fff8ea1c6713443f38079cb281658eff54bef28e41796946f")
	t.Log(w.PublicKey)

}
