//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

func (s *TStrings) AddStrings(list IStrings) {
	s.AddStrings3(list, false)
}

func (s *TStrings) AddStrings2(sArr []string) {
	s.BeginUpdate()
	defer s.EndUpdate()
	for _, v := range sArr {
		s.Add(v)
	}
}

func (s *TStrings) AddStrings3(list IStrings, clearFirst bool) {
	if list == nil {
		return
	}
	s.BeginUpdate()
	defer s.EndUpdate()
	if clearFirst {
		s.Clear()
	}
	if s.Count()+list.Count() > s.Capacity() {
		s.SetCapacity(s.Count() + list.Count())
		for i := int32(0); i < list.Count(); i++ {
			s.AddObject(list.S(i), list.Objects(i))
		}
	}
}

func (s *TStrings) AddPair(name, value string) *TStrings {
	return s.AddPair2(name, value, nil)
}

func (s *TStrings) AddPair2(name, value string, object IObject) *TStrings {
	s.AddObject(name+string(s.NameValueSeparator())+value, object)
	return s
}

// 文件流加载。
func (s *TStrings) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStreamFromBytes(data)
	defer mem.Free()
	mem.SetPosition(0)
	s.LoadFromStream(mem)
}
