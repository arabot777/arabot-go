package trace

// key必须要是string，方便查找，值因为是业务传入的，可以不一样。
// Expose extra KV pair for search
// Keys must be in form of string so that we can search with the specific keys, but value can be any form.
type Entry interface {
	Key() string
	Value() interface{}
	Entry()
}

type kv struct {
	k string
	v interface{}
}

func (m *kv) Key() string {
	return m.k
}

func (m *kv) Value() interface{} {
	return m.v
}

func (m *kv) Entry() {}

// 创建一个新的业务kv
// New a KV pair
func KV(key string, value interface{}) Entry {
	return &kv{k: key, v: value}
}
