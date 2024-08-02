package main

func main() {

	lru := LRUCache{
		keyMap: map[string]*Node{},
		list: DLL{
			maxSize: 3,
		},
	}

	//lru.Set("key1", "val1")
	//lru.Set("key2", "val2")
	//lru.PrintMap()
	//lru.PrintList()
	//_ = lru.Get("key1")
	//lru.PrintMap()
	//lru.PrintList()
	//lru.Set("key3", "val3")
	//lru.Set("key4", "val4")
	//lru.PrintMap()
	//lru.PrintList()
	//lru.Set("key2", "val21")
	//lru.PrintMap()
	//lru.PrintList()
	//lru.Set("key5", "val5")
	//lru.PrintMap()
	//lru.PrintList()
	lru.Set("key1", "val1")
	lru.Set("key2", "val2")
	//_ = lru.Get("key1")
	lru.Set("key3", "val3")
	lru.Set("key4", "val4")
	lru.Set("key5", "val5")
	lru.Set("key6", "val6")
	lru.Set("key7", "val7")
	lru.PrintMap()
	lru.PrintList()

	////fmt.Println(val)
	//lru.Set("key1", "val1")
	//lru.PrintMap()
	//lru.PrintList()

}
