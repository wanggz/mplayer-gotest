package library

import "errors"

type MusicManager struct {
	muiscs []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.muiscs)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.muiscs) {
		return nil, errors.New("Index out of range.")
	}
	return &m.muiscs[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.muiscs) == 0 {
		return nil
	}
	for _, m := range m.muiscs {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.muiscs = append(m.muiscs, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.muiscs) {
		return nil
	}

	removedMuisc := &m.muiscs[index]

	if index < len(m.muiscs)-1 {
		m.muiscs = append(m.muiscs[:index-1], m.muiscs[index+1:]...)
	} else if index == 0 {
		m.muiscs = make([]MusicEntry, 0)
	} else {
		m.muiscs = m.muiscs[:index-1]
	}

	return removedMuisc
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.muiscs) == 0 {
		return nil
	}
	for i, v := range m.muiscs {
		if v.Name == name {
			return m.Remove(i)
		}
	}
	return nil
}
