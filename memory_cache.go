package main

import (
	"time"
)

type Cache struct {
	dict       map[string]string
	expireDict map[string]int64
}

func (c *Cache) Expire(key string, second int) bool {
	if _, ok := c.dict[key]; ok {
		e := time.Now().Unix() + int64(second)
		c.expireDict[key] = e
		return true
	}
	return false
}

func (c *Cache) Put(key, value string) {
	c.dict[key] = value
}

func (c *Cache) Get(key string) string {
	if v, ok := c.expireDict[key]; ok {
		//如果设置了过期时间
		now := time.Now().Unix()
		if now < v {
			return c.dict[key]
		} else {
			//过期处理
			delete(c.dict, key)
			return ""
		}
	}
	//没有过期时间限制
	return c.dict[key]
}

func (c *Cache) Delete(key string) {
	delete(c.dict, key)
	delete(c.expireDict, key)
}

func (c *Cache) Set(key, value string)  {
	c.dict[key] = value
}

func main() {
	cache := Cache{
		dict:       make(map[string]string),
		expireDict: make(map[string]int64),
	}

	//存取测试
	cache.Put("key1", "value1")
	val1 := cache.Get("key1")
	println("value1" == val1)

	//过期测试
	cache.Expire("key1", 2)
	time.Sleep(time.Second * 3)
	val1 = cache.Get("key1")
	println(val1 != "value1")
	println(val1 == "")

	//此时过期的key1已经被自动删除
	_, ok := cache.dict["key1"]
	println(ok == false)

	//更改测试
	cache.Put("key2", "value2")
	cache.Set("key2", "value")
	val2 := cache.Get("key2")
	println(val2 == "value")

}
