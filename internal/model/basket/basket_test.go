package basket

//func TestAddItem(t *testing.T) {
//	//1.
//
//	testCases := []struct {
//		name       string
//		before     Basket
//		insertItem Item
//		expected   Basket
//	}{
//		{name: "test1",
//			before: Basket{make(map[string]CardItem)},
//
//			insertItem: Item{
//				name:     "Milk",
//				weight:   900,
//				priceRUB: 70,
//			},
//
//			expected: Basket{map[string]CardItem{
//				"Milk": {
//					item:      &Item{"Milk", 900, 70},
//					weightSum: 900,
//					priceSum:  70,
//					countItem: 1,
//				},
//			},
//			},
//		},
//
//		{name: "test2",
//			before: Basket{map[string]CardItem{
//				"Bread": {
//					item:      &Item{"Bread", 200, 40},
//					weightSum: 200,
//					priceSum:  40,
//					countItem: 1,
//				},
//			},
//			},
//
//			insertItem: Item{
//				name:     "Bread",
//				weight:   200,
//				priceRUB: 40,
//			},
//
//			expected: Basket{map[string]CardItem{
//				"Bread": {
//					item:      &Item{"Bread", 200, 40},
//					weightSum: 400,
//					priceSum:  80,
//					countItem: 2,
//				},
//			},
//			},
//		},
//	}
//
//	for _, testCase := range testCases {
//		testCase.before.AddItem(&testCase.insertItem)
//
//		assert.Equal(t, testCase.expected, testCase.before)
//	}
//}
//
//func TestRemoveItem(t *testing.T) {
//	testCases := []struct {
//		name       string
//		before     Basket
//		removeItem Item
//		expected   Basket
//	}{
//		{name: "test1",
//
//			before: Basket{make(map[string]CardItem)},
//
//			removeItem: Item{
//				name:     "Milk",
//				weight:   900,
//				priceRUB: 70,
//			},
//
//			expected: Basket{make(map[string]CardItem)},
//		},
//		{name: "test2",
//
//			before: Basket{map[string]CardItem{
//
//				"Milk": {
//					item:      &Item{"Milk", 900, 70},
//					weightSum: 900,
//					priceSum:  70,
//					countItem: 1,
//				},
//			},
//			},
//
//			removeItem: Item{
//				name:     "Milk",
//				weight:   900,
//				priceRUB: 70,
//			},
//
//			expected: Basket{make(map[string]CardItem)},
//		},
//
//		{name: "test3",
//
//			before: Basket{map[string]CardItem{
//
//				"Milk": {
//					item:      &Item{"Milk", 900, 70},
//					weightSum: 1800,
//					priceSum:  140,
//					countItem: 2,
//				},
//			},
//			},
//
//			removeItem: Item{
//				name:     "Milk",
//				weight:   900,
//				priceRUB: 70,
//			},
//
//			expected: Basket{map[string]CardItem{
//
//				"Milk": {
//					item:      &Item{"Milk", 900, 70},
//					weightSum: 900,
//					priceSum:  70,
//					countItem: 1,
//				},
//			},
//			},
//		},
//	}
//
//	for _, testCase := range testCases {
//		testCase.before.RemoveItem(&testCase.removeItem)
//
//		assert.Equal(t, testCase.expected, testCase.before)
//	}
//}
