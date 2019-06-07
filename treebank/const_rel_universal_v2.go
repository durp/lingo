// +build universalrelv2

package treebank

import "github.com/chewxy/lingo"

var dependencyTable = map[string]lingo.DependencyType{
	"acl":          lingo.Acl,         // clausal modifier of noun (adjectival clause)
	"acl:relcl":    lingo.AclRelcl,    // relative clause modifier
	"advcl":        lingo.AdvCl,       // adverbial clause modifier
	"advmod":       lingo.AdvMod,      // adverbial modifier
	"amod":         lingo.AMod,        // adjectival modifier
	"appos":        lingo.Appos,       // appositional modifier
	"aux":          lingo.Aux,         // auxiliary
	"aux:pass":     lingo.AuxPass,     // passive auxiliary
	"case":         lingo.Case,        // case marking
	"cc":           lingo.CC,          // coordinating conjunction
	"cc:preconj":   lingo.CCPreconj,   // preconjunct
	"ccomp":        lingo.CComp,       // clausal complement
	"clf":          lingo.Clf,         // classifier
	"compound":     lingo.Compound,    // compound
	"compound:prt": lingo.CompoundPrt, // separable verb particle
	"conj":         lingo.Conj,        // conjunct
	"cop":          lingo.Cop,         // copula
	"csubj":        lingo.CSubj,       // clausal subject
	"csubj:pass":   lingo.CSubjPass,   //clausal passive subject
	"dep":          lingo.Dep,         // unspecified dependency
	"det":          lingo.Det,         // determiner
	"det:predet":   lingo.DetPredet,   // predeterminer
	"discourse":    lingo.Discourse,   // discourse element
	"dislocated":   lingo.Dislocated,  // dislocated elements
	"expl":         lingo.Expl,        // expletive
	"fixed":        lingo.Fixed,       // fixed multiword expression
	"flat":         lingo.Flat,        // flat multiword expression
	"flat:foreign": lingo.FlatForeign, // foreign words
	"goeswith":     lingo.GoesWith,    // goes with
	"iobj":         lingo.IObj,        // indirect object
	"list":         lingo.List,        // list
	"mark":         lingo.Mark,        // marker
	"nmod":         lingo.NMod,        // nominal modifier
	"nmod:npmod":   lingo.NModNPMod,   // noun phrase as adverbial modifier
	"nmod:tmod":    lingo.NModTMod,    // temporal nominal modifier
	"nmod:poss":    lingo.NModPoss,    // possessive nominal modifier
	"nsubj":        lingo.NSubj,       // nominal subject
	"nsubj:pass":   lingo.NSubjPass,   // passive nominal subject
	"nummod":       lingo.NumMod,      // numeric modifier
	"obj":          lingo.Obj,         // object
	"obl":          lingo.Obl,         // oblique nominal
	"obl:npmod":    lingo.OblNPMod,    // oblique noun phrase as adverbial modifier
	"obl:tmod":     lingo.OblTMod,     // temporal modifier
	"orphan":       lingo.Orphan,      // orphan
	"parataxis":    lingo.Parataxis,   // parataxis
	"punct":        lingo.Punct,       // punctuation
	"reparandum":   lingo.Reparandum,  // overridden disfluency
	"root":         lingo.Root,        // root
	"vocative":     lingo.Vocative,    // vocative
	"xcomp":        lingo.XComp,       // open clausal complement
	"-NULL-":       lingo.NoDepType,
}
