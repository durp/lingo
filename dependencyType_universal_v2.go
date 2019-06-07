// +build universalrelv2

package lingo

const BUILD_RELSET = "universalrelv2"

//go:generate go run golang.org/x/tools/cmd/stringer -tags universalrelv2 -type=DependencyType -output=dependencyType_universal_v2_string.go

// https://universaldependencies.org/u/dep/index.html
const (
	NoDepType   DependencyType = iota
	Acl                        // clausal modifier of noun (adjectival clause)
	AclRelcl                   // relative clause modifier
	AdvCl                      // adverbial clause modifier
	AdvMod                     // adverbial modifier
	AMod                       // adjectival modifier
	Appos                      // appositional modifier
	Aux                        // auxiliary
	AuxPass                    // passive auxiliary
	Case                       // case marking
	CC                         // coordinating conjunction
	CCPreconj                  // preconjunct
	CComp                      // clausal complement
	Clf                        // classifier,
	Compound                   // compound
	CompoundPrt                // separable verb particle
	Conj                       // conjunct
	Cop                        // copula
	CSubj                      // clausal subject
	CSubjPass                  // clausal passive subject
	Dep                        // unspecified dependency
	Det                        // determiner
	DetPredet                  // predeterminer
	Discourse                  // discourse element
	Dislocated                 // dislocated elements
	Expl                       // expletive
	Fixed                      // fixed multiword expression
	Flat                       // flat multiword expression
	FlatForeign                // foreign words
	GoesWith                   // goes with
	IObj                       // indirect object
	List                       // list
	Mark                       // marker
	NMod                       // nominal modifier
	NModNPMod                  // noun phrase as adverbial modifier
	NModPoss                   // possessive nominal modifier
	NModTMod                   // temporal nominal modifier
	NSubj                      // nominal subject
	NSubjPass                  //passive nominal subject
	NumMod                     // numeric modifier
	Obj                        // object
	Obl                        // oblique nominal
	OblNPMod                   // oblique noun phrase as adverbial modifier
	OblTMod                    // temporal modifier
	Orphan                     // orphan
	Parataxis                  // parataxis
	Punct                      // punctuation
	Reparandum                 // overridden disfluency
	Root                       // root
	Vocative                   // vocative
	XComp                      // open clausal complement
	MAXDEPTYPE
)

var Modifiers = []DependencyType{AdvMod, AMod, Discourse}
var Compounds = []DependencyType{Compound}
var DeterminerRels = []DependencyType{Det}
var MultiWord = []DependencyType{Compound, Fixed, Flat}
var QuantifyingMods = []DependencyType{NumMod}
