package search

type Dictionary struct {
	size int
	data []*int
}

const defaultDictSize = 1046527

func NewDictionaryDefaultSize() Dictionary {
	return NewDictionary(defaultDictSize)
}

func NewDictionary(size int) Dictionary {
	return Dictionary{
		size: size,
		data: make([]*int, size),
	}
}

func h1(key, size int) int {
	return key % size
}

// 衝突した場合の第2のハッシュ関数
func h2(key, size int) int {
	return 1 + key%(size-1)
}

func h(key, i, size int) int {
	return (h1(key, size) + i*h2(key, size)) % size
}

func (d Dictionary) Insert(key int) {
	i := 0
	for {
		j := h(key, i, d.size)
		if d.data[j] == nil {
			d.data[j] = &key
			return
		} else {
			i++
		}
	}
}

func (d Dictionary) Search(key int) (int, bool) {
	i := 0
	for {
		j := h(key, i, d.size)
		if *d.data[j] == key {
			return j, true
		} else if d.data[j] == nil || i >= d.size {
			return 0, false
		} else {
			i++
		}
	}
}
