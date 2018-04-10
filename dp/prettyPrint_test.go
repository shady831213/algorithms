package dp

import (
	"fmt"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	words := "I can hear the drummers drumming And the trumpets, someone's tryna summon someone, I know something 's coming But I 'm running from it to be standing at the summit And plummet, how come it wasn 't what I thought it was Was it, too good to be true? Have nothing, get it all but too much of it Then lose it again, then I swallow hallucinogens Cause if not, where the hell did it go? Cause here I sit in Lucifer’s den by the dutch oven Just choosing to sin Even if it means I 'm selling my soul, just to be the undisputed again Do whatever I gotta do just to win Cause I got this motherfucking cloud over my head Crown around it, thorns on it Cracks in it, bet you morons didn 't Think I 'd be back, did ya? How 'b out that I’m somehow now back to the underdog But no matter how loud that I bark, this sport is something I never bow - wow 'd at I complain about the game, I shout and I pout, it 's a love - hate But I found out that I can move a mountain of doubt Even when you bitches are counting me out, and I appear to be down for the count Only time I ever been out and about is driving around town with my fucking whereabouts, in a doubt cause I been lost tryna think of what I did to get here but I 'm not a quitter Gotta get up, give it all I got or give up Spit or shit or stepped on but kept going I 'm tryna be headstrong but it feels like I slept on my neck wrong Cause you 'r e moving onto the next, but is the respect gone? Cause someone told me that(Kings never die)"
	alignedIdx, alignedWords := prettyPrint(words, 150)
	if alignedIdx != 48442 {
		t.Log(fmt.Sprintf("alighedIdx expect 48442, but get %d", alignedIdx))
		t.Fail()
	}
	if alignedWords != "I can hear the drummers drumming And the trumpets, someone's tryna summon someone, I know something 's coming But I 'm running from\n"+
		"it to be standing at the summit And plummet, how come it wasn 't what I thought it was Was it, too good to be true? Have nothing, get\n"+
		"it all but too much of it Then lose it again, then I swallow hallucinogens Cause if not, where the hell did it go? Cause here I sit in\n"+
		"Lucifer’s den by the dutch oven Just choosing to sin Even if it means I 'm selling my soul, just to be the undisputed again Do whatever\n"+
		"I gotta do just to win Cause I got this motherfucking cloud over my head Crown around it, thorns on it Cracks in it, bet you morons didn\n"+
		"'t Think I 'd be back, did ya? How 'b out that I’m somehow now back to the underdog But no matter how loud that I bark, this sport is\n"+
		"something I never bow - wow 'd at I complain about the game, I shout and I pout, it 's a love - hate But I found out that I can move a\n"+
		"mountain of doubt Even when you bitches are counting me out, and I appear to be down for the count Only time I ever been out and about\n"+
		"is driving around town with my fucking whereabouts, in a doubt cause I been lost tryna think of what I did to get here but I 'm not\n"+
		"a quitter Gotta get up, give it all I got or give up Spit or shit or stepped on but kept going I 'm tryna be headstrong but it feels\n"+
		"like I slept on my neck wrong Cause you 'r e moving onto the next, but is the respect gone? Cause someone told me that(Kings never die) " {
		t.Log("output not aligned\n")
		t.Log(alignedWords)
		t.Fail()
	}
}
