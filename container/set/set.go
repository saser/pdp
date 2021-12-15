package set

type Set[T comparable] map[T]struct{}

func Union[T comparable](ss ...Set[T]) Set[T] {
	u := New[T]()
	for _, s := range ss {
		for t := range s {
			u[t] = struct{}{}
		}
	}
	return u
}

func New[T comparable](ts ...T) Set[T] {
	s := make(Set[T])
	for _, t := range ts {
		s[t] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(ts ...T) {
	for _, t := range ts {
		s[t] = struct{}{}
	}
}

func (s Set[T]) Remove(ts ...T) {
	for _, t := range ts {
		delete(s, t)
	}
}

func (s Set[T]) Contains(ts ...T) bool {
	for _, t := range ts {
		if _, ok := s[t]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) ContainsAny(ts ...T) bool {
	for _, t := range ts {
		if _, ok := s[t]; ok {
			return true
		}
	}
	return false
}
