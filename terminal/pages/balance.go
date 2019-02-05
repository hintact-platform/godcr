package pages

import (
	"fmt"
	"context"

	"github.com/rivo/tview"
	//"github.com/decred/dcrd/dcrutil"
	"github.com/raedahgroup/godcr/app/walletcore"
)

func BalancePage(ctx context.Context, wallet walletcore.Wallet) tview.Primitive {
	body := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(fmt.Sprintf("\n\n\nBalance : %s", "0 GODCR"))

	getBal := body.SetText
	getBalance(getBal)

	getBalance(ctx, wallet, getBal)

	return body
}


func getBalance(ctx context.Context, wallet walletcore.Wallet, getBal func(text string) *tview.TextView) {
	_, err := wallet.AccountsOverview(walletcore.DefaultRequiredConfirmations)
	if err != nil {
		fmt.Println("err")
	}
}