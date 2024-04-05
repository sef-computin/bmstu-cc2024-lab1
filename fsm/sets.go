package fsm

type SymbolSet struct {
	set []rune
}

func (set *SymbolSet) Add(sym rune) {
	if !set.Contains(sym) {
		set.set = append(set.set, sym)
	}
}

func (set *SymbolSet) Remove(sym rune) {
	for ind, val := range set.set {
		if val == sym {
			set.set = append(set.set[:ind], set.set[ind+1:]...)
		}
	}
}

func (set *SymbolSet) Contains(sym rune) bool {
	for _, val := range set.set {
		if val == sym {
			return true
		}
	}
	return false
}

func (set *SymbolSet) Len() int {
	return len(set.set)
}

func (set *SymbolSet) GetAll() []rune {
	return set.set
}

func NewSymbolSet() *SymbolSet {
	return &SymbolSet{set: []rune{}}
}

func (set *SymbolSet) Equals(set2 SymbolSet) bool {
	if set.Len() == set2.Len() {

		for _, val := range set.set {
			if !set2.Contains(val) {
				return false
			}
		}
		return true
	}
	return false
}


type StateSet struct {
	set []State
}

func (set *StateSet) Add(state State) {
	if !set.Contains(state) {
		set.set = append(set.set, state)
	}
}

func (set *StateSet) Remove(state State) {
	for ind, val := range set.set {
		if val.Equals(state) {
			set.set = append(set.set[:ind], set.set[ind+1:]...)
		}
	}
}

func (set *StateSet) Contains(state State) bool {
	for _, val := range set.set {
		if val.Equals(state) {
			return true
		}
	}
	return false
}

func (set *StateSet) Len() int {
	return len(set.set)
}

func (set *StateSet) GetAll() []State {
	return set.set
}

func NewStateSet() *StateSet {
	return &StateSet{set: []State{}}
}

func (set *StateSet) Equals(set2 StateSet) bool {
	if set.Len() == set2.Len() {

		for _, val := range set.set {
			if !set2.Contains(val) {
				return false
			}
		}
		return true
	}
	return false
}

func (set *StateSet) Union(set2 StateSet) *StateSet{
  ret := NewStateSet()
  for _, elem := range set2.GetAll(){
    ret.Add(elem)
  } 
  for _, elem := range set.GetAll(){
    ret.Add(elem)
  }
  return ret
}

func (set *StateSet) Intersect(set2 StateSet) *StateSet{
  ret := NewStateSet()

  for _, elem := range set.GetAll(){
    if set2.Contains(elem){
      ret.Add(elem)
    }
  }
  return ret
}
