package main

import (
	"testing"
	
)

func TestWallet( t *testing.T)  {
	t.Run("Wallet deposit", func(t *testing.T) {
		w:=Wallet{}
		w.Deposit(Bitcoin(10))
		got:=w.Balance()
		want:=Bitcoin(10)
		if want!=got{
			t.Errorf("wanted %d but got %d",want,got)
		}

	})
	t.Run("Wallet withdraw", func(t *testing.T){
		w:=Wallet{balance: 10}
		err:=w.Withdraw(Bitcoin(10));
		if err!=nil{
			 t.Errorf("Unexpected error: %s", err)
		}
		got:=w.Balance()
		want:=Bitcoin(0)
		if want!=got{
			t.Errorf("got %d but wanted %d",got,want)
		}
	})
	
}