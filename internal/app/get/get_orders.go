package get

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"test_orders/internal/service"
	"test_orders/internal/storage"

	_ "github.com/lib/pq"
)

type Getter interface {
	GetOrders([]int64) ([]storage.Product, error)
}

func GetOrders(getter Getter) {
	ordersStr := os.Args[1:]
	if len(ordersStr) == 0 {
		log.Fatal("get 0 orders")
	}
	orders := make([]int64, 0, len(ordersStr))
	for _, orderStr := range ordersStr {
		order, err := strconv.Atoi(orderStr)
		if err != nil {
			log.Fatal("invalid order")
		}
		orders = append(orders, int64(order))
	}

	getOrders(getter, orders)
}

func getOrders(getter Getter, orders []int64) {
	result, err := getter.GetOrders(orders)
	if err != nil {
		if errors.Is(err, storage.ErrNoResult) {
			fmt.Println("can't get result")
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("Страница сборки заказов %s\n", service.GetStringOrders(orders))

	prevShelf := ""
	for _, product := range result {
		if product.MainShelf != prevShelf || prevShelf == "" {
			prevShelf = product.MainShelf
			fmt.Printf("===Стеллаж %s\n", product.MainShelf)
		}

		fmt.Printf("%s (id=%d)\n", product.NameProduct, product.IdProduct)
		fmt.Printf("заказ %d, %d шт\n", product.Order, product.CountProduct)
		if len(product.AdditionalSheves) > 0 {
			fmt.Printf("доп стеллаж: %s\n", service.GetStringAdditionalSheves(product.AdditionalSheves))
		}

		fmt.Println()
	}
}
