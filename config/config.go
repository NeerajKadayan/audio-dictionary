package config

var (
	LastFmApiKey       string // 1b6fe22b7645f5005d4037fe84629796
	LastFmSharedSecret string // e33a625d95c7bbb9e09a7cc8bf6f7d70
	MusixMatchApiKey   string // fcee7240c5a3ce2115f1c7d904067e51 -- eg authentication track.get?apikey=xxxx
)

func init() {
	LastFmApiKey = "1b6fe22b7645f5005d4037fe84629796"
	LastFmSharedSecret = "e33a625d95c7bbb9e09a7cc8bf6f7d70"
	MusixMatchApiKey = "fcee7240c5a3ce2115f1c7d904067e51"
}
