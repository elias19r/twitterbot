package twitter

import (
	"context"
	"net/url"
	"regexp"
	"strings"

	"github.com/elias19r/twitterbot/pkg/int64slice"
)

// Mute mutes a user.
func Mute(user User) error {
	_, err := api.MuteUser(user.Username, url.Values{})
	if err != nil {
		return err
	}
	logger.Printf("muted %s (%s)\n", user.Username, user.Name)
	return nil
}

// Unmute unmutes a user.
func Unmute(user User) error {
	_, err := api.UnmuteUser(user.Username, url.Values{})
	if err != nil {
		return err
	}
	logger.Printf("unmuted %s (%s)\n", user.Username, user.Name)
	return nil
}

// GetMutedIDs returns a slice with all muted user IDs.
func GetMutedIDs(ctx context.Context) int64slice.Int64Slice {
	logger.Println("building muted list...")
	muted := int64slice.Int64Slice{}
	v := url.Values{}
	nextCursor := "-1"
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			v.Set("cursor", nextCursor)
			c, err := api.GetMutedUsersIds(v)
			if err != nil {
				break
			}
			muted = append(muted, c.Ids...)
			nextCursor = c.Next_cursor_str
			if nextCursor == "0" {
				break loop
			}
		}
	}
	logger.Println("muted list built")
	return muted
}

// WillMute decides whether to mute a user.
func WillMute(lu Lookup) error {
	content := lu.User.Username + " " + lu.User.Name + " " + lu.User.Description + " " + lu.User.Location

	if !lu.Muting && (muteRegexp.MatchString(content) || lu.User.Description == "") {
		return Mute(lu.User)
	}

	return ErrNoUnMutePerformed
}

var muteRegexp = regexp.MustCompile(`(?i)` + strings.Join(mutePatterns, "|"))

var mutePatterns = []string{
	// Follow-backs.
	`i?fb`,
	`f4f`,
	`st?dv`,
	`(un)?fo?llow`,
	`texasgain`,
	`(ifb|1d)drive`,
	`1first`,
	`mgwv`,
	`notifica`,
	`unf ?[=:/] ?unf`,
	`segue`,
	`sigam`,
	`sigo( ?t[uo]dos?)? de? volta`,
	`[0-9]{1,3}[,.]?[0-9]? ?k`,
	`[#@]team`,
	`indica[çc][ãa]?[oõ]e?s?`,

	// Language specific.
	`\p{Arabic}+`,
	`\p{Hiragana}+`,
	`\p{Katakana}+`,
	`\p{Han}+`,
	`\p{Greek}+`,
	`\p{Devanagari}+`,
	`\p{Cyrillic}+`,
	`\p{Bengali}+`,
	`[şöüğıûî¿¡ñ]`,
	`🇹🇷`,
	`🇳🇬`,
	`turkey`,
	`istanbul`,
	`madrid`,
	`united ?states`,
	`pakistan`,
	`stockholm`,
	`argentina`,
	`india`,

	// Politics.
	`influencer`,
	`politics`,
	`pol[íi]tica`,
	`jou?rnalista?`,
	`censorship`,
	`lula`,
	`bolsonaro`,
	`dilma`,
	`temer`,
	`trump`,
	`nazismo?`,
	`feminismo?`,
	`ac?tivista?`,
	`united`,
	`insafian`,
	`preconceito`,
	`discrimina`,
	`democrac(ia|y)`,
	`che ?guevara`,
	`marx`,
	`justi[cç]`,
	`rep[uú]blic`,

	// Religion.
	`jesus`,
	`crist[aã]?o`,
	`Christ(ian)?`,
	`deus`,
	`god`,
	`satan[aá]s`,
	`de(vil|mon)`,
	`b[íi]blia`,
	`bible`,
	`hell`,
	`heaven`,
	`muslim`,
	`muçulman[oa]`,
	`islam(ismo?)?`,
	`m[aou]h[ae]mm?[ae]d`,
	`all[ae]hu`,
	`[ae]kb[ae]r`,
	`ahmm?[ae]d`,
	`fé`,
	`faith`,
	`dalal`,
	`mustafa`,
	`abbad`,

	// Porn.
	`porno?`,
	`sexo?`,
	`\+(18|21)`,
	`(18|21)\+`,
	`🔞`,
	`adulto?`,
	`genital`,
	`tes[ãa]o`,
	`er[oó]tico?`,
	`relationship`,
	`anal`,
	`vagina`,

	// Social networks.
	`youtube`,
	`twitch`,
	`whatsapp`,
	`insta`,

	// Cryptocurrency.
	`bitcoi?n`,
	`bloc?k?chain`,
	`block`,
	`cr[yi]pto`,
	`#btc`,
	`#ico`,
	`altcoi?n`,

	// Other
	`offers?`,
	`discounts?`,
	`business?`,
	`shape`,
	`fitness?`,
	`workout`,
	`warrior`,
	`gamer?`,
	`zoei?r`,
	`maconha`,
	`weed`,
	`4[:h]?20`,
	`life ?style`,
	`parody`,
	`sport`,
	`goss?ip`,
	`market`,
	`compra`,
	`vegan`,
	`student`,
	`minecraft`,
	`[❥🔔📢💯✡🕉🕇🔮☪ღ★™®]`,
	`conspiracy`,
	`conspira[cç][aã]o`,
	`hoax`,
	`black +friday`,
	`philosopher`,
	`fil[óo]sofo`,
	`fashion`,
}
