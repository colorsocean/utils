package utils

/****************************************************************
**
********/

type HashSet struct {
	data map[interface{}]bool
}

func (this *HashSet) Init() *HashSet {
	this.data = make(map[interface{}]bool)
	return this
}

/****************************************************************
**
********/

func (this *HashSet) Add(value interface{}) {
	this.data[value] = true
}

func (this *HashSet) Contains(value interface{}) (exists bool) {
	_, exists = this.data[value]
	return
}

func (this *HashSet) Remove(value interface{}) {
	if this.Contains(value) {
		delete(this.data, value)
	}
}

func (this *HashSet) Length() int {
	return len(this.data)
}
