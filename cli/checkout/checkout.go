package checkout

import (
	"math"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "variant-1",
	Short: "Run variant 1",
	RunE:  run,
}

type Inventory struct {
	SKU map[string]Item
}

type Item struct {
	Name              string
	Price             float64
	InventoryQuantity int
}

type Cart struct {
	SKU      string
	Quantity int
}

const jsonFile = "https://mghifariyusuf.github.io/fita/items.json"

func run(cmd *cobra.Command, args []string) error {
	// i dont know about graphql, so i just write down the function that will be called 🙏
	return nil
}

func (i *Inventory) Checkout(cart []Cart) (total float64) {
	if len(cart) == 0 {
		return total
	}

	return i.validatePromo(cart)
}

func (i *Inventory) validatePromo(cart []Cart) (total float64) {
	var discount float64
	scanned := make(map[string]int)

	for _, c := range cart {
		// skip item that already scanned before
		if scanned[c.SKU] == 1 {
			continue
		}

		// promo 1, if buy MacBook Pro, get free Raspberry Pi B
		if c.SKU == "43N23P" {
			discount = discount + i.SKU["234234"].Price
		}

		// promo 2, if buy 3 Google Home, discount as 1 price of Google Home
		if c.SKU == "120P90" && c.Quantity == 3 {
			discount = discount + i.SKU[c.SKU].Price
		}

		// promo 3, if buy 3 Alexa Speaker, all speakers discount 10%
		if c.SKU == "A304SD" && c.Quantity >= 3 {
			discount = discount + (i.SKU[c.SKU].Price*10/100)*float64(c.Quantity)
		}

		// flag item already scanned
		scanned[c.SKU] = 1

		// calculate total with items quantity in cart
		total = total + i.SKU[c.SKU].Price*float64(c.Quantity)
	}

	// add discount and round it
	total -= discount
	total = math.Round((total)*100) / 100

	return total
}
