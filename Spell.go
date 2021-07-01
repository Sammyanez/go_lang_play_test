package main

type Spell struct {
	spellName        string
	manaCost, damage int
}

func (s *Spell) SpellName() string {
	return s.spellName
}

func (s *Spell) SetSpellName(spellName string) {
	s.spellName = spellName
}

func (s *Spell) ManaCost() int {
	return s.manaCost
}

func (s *Spell) SetManaCost(manaCost int) {
	s.manaCost = manaCost
}

func (s *Spell) Damage() int {
	return s.damage
}

func (s *Spell) SetDamage(damage int) {
	s.damage = damage
}
