package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution(in))
}

func solution(data string) int {
	var sum [2]int
	var lines = strings.Split(data, "\n")

	for _, line := range lines {
		res := checkLine(line)

		switch res {
		case 1:
			sum[0]++
		case 2:
			sum[1]++
		case 3:
			sum[0]++
			sum[1]++
		}
	}

	return sum[0] * sum[1]
}

func checkLine(line string) int {
	var chars = map[byte]int{}
	var has2, has3 bool

	for _, char := range []byte(line) {
		chars[char]++
	}

	for _, char := range chars {
		switch char {
		case 2:
			has2 = true
		case 3:
			has3 = true
		}

		if has2 && has3 {
			return 3
		}
	}

	switch {
	case has2:
		return 1
	case has3:
		return 2
	}

	return 0
}

const in = `tqzvwnogbarflkpcbdewsmjhxi
tqyvunogzarfckpcbdewsmzhci
tqyvunojzarflkpcbdcwsmyhxi
tqyvunogzarclkpcbdewmmjrxi
hqyvunogzarflkpcbczwsmjhxi
tqyvunogzarflppcudewsmjhei
tqrvunogzarflkpcbpewsmjhdi
aqyvunogzarflkpcbdewsjjsxi
tqyvtnogzarflkkcbdewymjhxi
tcyeunogzarfkkpcbdewsmjhxi
tqyvunozzarfvkpcbdewsqjhxi
tkyvuwjgzarflkpcbdewsmjhxi
tqevunogzarflkpnbdkwsmjhxi
tqyvunogqarflkpcppewsmjhxi
tqyvunsgzarflkpcbrewsmjhxk
tqyvunogzarffkpdbnewsmjhxi
tqyvunogvarflkpjbdewsojhxi
tqyvgnogzarflkpfbdswsmjhxi
tqyvunogzarfxkpcbtersmjhxi
tqyvukhgzarflupcbdewsmjhxi
tqyvdnoozyrflkpcbdewsmjhxi
tqyvunogzorflkpcbdewsvjhxd
tqyvunzqzarflkpcbdewxmjhxi
tqykunogzarulkpcbdhwsmjhxi
tqycdnogzarflkpcbdewsmjhxz
eqyvunogzarflkpcbdhwqmjhxi
cqyvunogzaralkpcbdewsmjvxi
vqyvunogzarflklcbeewsmjhxi
tqyvunogzarffkpcqdewlmjhxi
eqyvunogzarflkpcbdejsmahxi
tqyvunogjarflkocsdewsmjhxi
tqzvunogzanflkpcbdewsmjhxg
tqycunogzarflkpabdewsmkhxi
tqyvunogzarnlkpcbdewimjaxi
tqyvunogzarfzkpcbdepsyjhxi
gqykunogzarelkpcbdewsmjhxi
tqyvuwogzarflkpcbdrwsmjmxi
tqdqunogzarflkpcbdewsvjhxi
tqmvunbgzarflkpcbdewsvjhxi
tqyvunogzarflkpcbheesmjhdi
tqxvunogzarfxkpcbdewsmhhxi
tqyvunogzarflkpabdjosmjhxi
tqyvugogztrflkpybdewsmjhxi
tqyvundgzarflkxcbdewsmjhmi
tqyvunogzurfgkpcbdeksmjhxi
tqyvunokzarfckpcbdewsojhxi
tqyvufobzarflkpcldewsmjhxi
tqyvunogznrflkncbdeusmjhxi
tqyvuncgzarfxkpcmdewsmjhxi
oqyvunogzarflkpybdewbmjhxi
tqyvjnogqarfmkpcbdewsmjhxi
tqyvunogzacflkpcidewsmjhwi
tqyvunogzarflkpcbqlwxmjhxi
tqyvunogzarflkpnbhewsvjhxi
vqyiunogzarflkpcbdewsmjhqi
tbyvuncgzarfljpcbdewsmjhxi
tqylunogzarflkpcldexsmjhxi
tqyvulogzarflktcbdewsxjhxi
tqyvmnlgzarflkpcbxewsmjhxi
tqyvunogzartlkpcbdewsmihxp
nqyvunogzarflkpcbdewsmnwxi
tqyvunogzarflkpczdewsmjcxj
tzyvunogzariwkpcbdewsmjhxi
tqyiufogzarflipcbdewsmjhxi
oqyvunogzasflkpcbdewsmjhxv
tqyvunoqmarflkpcwdewsmjhxi
tqrvunogzarflkpqbdewnmjhxi
tqyvunogzarmlkocbdewsmjhri
tqyvunogzakflkpcbveasmjhxi
tqyvunorearflkpcbdewsmfhxi
tqynrnogzarflkpcbdewsmjhxp
tiyvueogzaeflkpcbdewsmjhxi
tqyrunogzarflkpccdewbmjhxi
tqtvunogzarflkpcbdewbnjhxi
tqyvfnogzarflhpcbdewsmjqxi
tqyjunoazarflkpcbdewssjhxi
tqyvunxizarflkpcbdewsmjnxi
tqyvunogzarfhkpobdewsmjhai
tqyaunogzanflkpcbxewsmjhxi
tqyvunogzarflkpsbuewsmjmxi
tqyvunogzarzlkwybdewsmjhxi
tqyvunogzarflkpibdawsmhhxi
tqyvunogzarflkycbdewamjwxi
tqyvunourarflkpcbdewsujhxi
tqyvnnogzirflkpcbdewsdjhxi
tayvunogzauflkpcqdewsmjhxi
tqyvunobzfrflkscbdewsmjhxi
tqygvnogzarflkpcbdewsmjnxi
oqyvunogzarflkpcbdewsmjsgi
tqyvunokzarflkpcbdewshjhii
tqyvunobzarflkvcbdewskjhxi
aqyvunogvarflkpcbdlwsmjhxi
tqyvunogzmrrlbpcbdewsmjhxi
tqyvunggzaqolkpcbdewsmjhxi
tqyvunogqarflkpybdewsmjaxi
tqyvunogzxrflkpcxsewsmjhxi
zqyvunogzarflppcbdewsmjhni
tqvvunogzakslkpcbdewsmjhxi
tqyrunogzarzlkpcbdewsmjtxi
tqyvhnogzarfxkpcbdewqmjhxi
tqyvunogzarflkecbdewgdjhxi
tqyvuwogzerfhkpcbdewsmjhxi
tqmvunogzarflkpcbddwsmdhxi
tqyvunogzarflcqcbdewsmihxi
tqyvunogzarvlkpcbdewsmjmxd
tqyvknogzarfllncbdewsmjhxi
tqyvunogzarflbpcbdrwsajhxi
tqyvunogzarfukpcbddwsmjhii
tqyvuvojzahflkpcbdewsmjhxi
tqyvunogzarfchpcbdeqsmjhxi
wqivueogzarflkpcbdewsmjhxi
tqyvunogzwrflkpcbdewstjhxz
tqyvunogzarfloscbdewsmjhxf
tqfbuiogzarflkpcbdewsmjhxi
tqyvfuogzarflkpcxdewsmjhxi
tqyvunogzarflkpcpdewsmgqxi
tqyvunogzdrflkpcbdewsmqhci
tqyvunogzartlkpcbdewsmjpxj
tqyfunogzarfwkpcbdewsmwhxi
tqyvuntgzarflkjcmdewsmjhxi
tqyvunqfzarflkpckdewsmjhxi
nqyvunogznrflkpcbiewsmjhxi
tqymunobzarflkccbdewsmjhxi
tqyvunogzaoflkprbdewzmjhxi
tqyvunogzaiflkpcvdewbmjhxi
tqwvunogzarfzkpcbdewsmjhxx
txhvunogzarflkpcbdewsijhxi
tqyeunogkarflkicbdewsmjhxi
tqylunogzarylkpcbdewsmkhxi
tqyvriogzarflkpcbdewsmohxi
tqyvunogzsrflkpcbdeasijhxi
tqyvunogzarflkpcbfewcmjhxh
tqyvunoozvrflkpcbdewimjhxi
tqyvunogqayflkpcidewsmjhxi
tqyounogzarflkpccdewsmjhxg
tqgvunogsarflkpcbdewqmjhxi
tqevunogzarflkpcbmewsmjhpi
tqivunogzarflkgcbdewqmjhxi
tqyuunogzlrflkgcbdewsmjhxi
xqyvbnogznrflkpcbdewsmjhxi
tqyvunogzarfjkpebdewsmnhxi
tqyvvnogzarfskpcxdewsmjhxi
thovunogzarflkpcbdewsmjhvi
tqyvunugzarflkpcpdewsmjrxi
tcyvvnogzarflkdcbdewsmjhxi
tqdfunogzarflkpbbdewsmjhxi
tqyvunogzarflkpcbdnwsejzxi
tqyvunivkarflkpcbdewsmjhxi
tqyvunogzawflopcbdewsmjhmi
tqyvunogzarflkpcbdkwsdjqxi
tqyvunodzarflkpcbdewlmjhei
oqyvunoozarflkpcbdemsmjhxi
tiyvunogzarjlkpcbdewfmjhxi
tqrvunogearflkpcbdewsojhxi
tqyvunkgzarflkpcbdcwtmjhxi
tqmvunogzarflkpcbdpwsmjtxi
tqyvunogzarflkpcydeptmjhxi
tqyvunkglarflkpcbdfwsmjhxi
tqyaunogzarflkpcbzeesmjhxi
tqyvunogzarrlkpcbdkwsmjhui
tpyvunogzarflkqcbdewsmjhxd
tvyvunogzarfkkpcbdewsajhxi
gqynunogzarflepcbdewsmjhxi
zqvvunogzarflkpcbdexsmjhxi
tqyyunogzawflkpcbdewsmjhxw
tfyvunogzarflkpcbdewomjrxi
tqyeunogzvrflrpcbdewsmjhxi
nqyvunogzarftkpabdewsmjhxi
tzyvunogzariwkpcbdewmmjhxi
tiyvunogzarflkpcbbewsmjhxa
tqydujoyzarflkpcbdewsmjhxi
tqyvunpgzarflkpcbdeysmjhwi
tqyvunogvarllkpcbdewsmshxi
tqyvunogzbrvlkpcbdewsmjhxp
tcyvueogzarflkacbdewsmjhxi
tqyvunogzrhflkpcbdetsmjhxi
tqavunogzrdflkpcbdewsmjhxi
tqyvunogzjrflkpcbdewstjhki
tqyqunolzarflkpcbdewbmjhxi
tqyvunogzarflkqczdgwsmjhxi
tqyvunogzqrfrkpcbrewsmjhxi
tqyvcnogzhrflkacbdewsmjhxi
tqyvunogzarflkpcbdewsmdhtk
tqyvunoggarftkpcbjewsmjhxi
tvyvunogzarfkkpcpdewsmjhxi
tqyvunogzawflkpcndedsmjhxi
tqyvunogzrrflkpcbdemsmmhxi
tqyvunogzarclkpcbrpwsmjhxi
tryvunogztrflkpcbbewsmjhxi
cqyvundgzarflkpcbdewgmjhxi
tqyvunogzarflktcbkewsmjqxi
tqyvjuogzarflkpcadewsmjhxi
tqyvunogzyrflkpcbbxwsmjhxi
ttyvunogzarflkpcbdewsnmhxi
tqyvunogzarflkpcbdeqsmlhki
fqyvugogzarflapcbdewsmjhxi
tqyvunogzartlkppbdewszjhxi
tqyvunfgztrflzpcbdewsmjhxi
tqyvunogmazflkpcbdewsmjhki
tqyvunzdzarflkpcbdewsmjhvi
tqyvunogzarflkpqbzewsujhxi
tqyvunogzarzlkpcbeewymjhxi
tqyounogzarflkpcbdewsnwhxi
tqysunogsaiflkpcbdewsmjhxi
tqdvunogdarflkpcboewsmjhxi
teyvunogzarflkscbdfwsmjhxi
tqyvunoazarflkpcbdvwsmjhpi
tqyvunooearflkpcbdewrmjhxi
tqyvunoszarflnrcbdewsmjhxi
tqyvunogzalflkpcblewsjjhxi
uqlvunkgzarflkpcbdewsmjhxi
tqyvunobzarflkpcidewsmjhvi
tnyvunogzarflkpcbdnwamjhxi
tqyoudogzarflkpcbdgwsmjhxi
tqyvunoggarflkpcbmewsmwhxi
tqykunogzazflkpcbddwsmjhxi
tfyvunogzarflkpcbdewsmjhgo
tqyvunogztrflkpcbdewoojhxi
tqyvunogzarflkpcbdewbmjoni
tmyvunogzalflkpabdewsmjhxi
tqyvunogzarflkpbbvewqmjhxi
tqyvunofzarflkpcwdexsmjhxi
tayvunogzasflkpcbdewsmhhxi
tqyvlnogzarflkpcbdewsmjwxd
tvyvunogzarflkpcbdewhmjrxi
tqyvunogzarplkpcsdewsmihxi
tqyvunogzarplklcbdewsmjtxi
rqbvunogzarhlkpcbdewsmjhxi
tqyvuxogzarftkpcbdewsmthxi
tqtvunogzarfikpczdewsmjhxi
tqyvunwgzarflkpcbdewsmjxxk
tqygunogzarzlkpnbdewsmjhxi
tqyvunogzarjlkpcbdbwswjhxi
tqyvunogzalfnkpcbdewsmjhxf
tqyucuogzarflkpcbdewsmjhxi
tzyvunogzvrflkpcbdewsmjaxi
tjyvunlgzarflkpcbdewgmjhxi
tqyvcnogzarklkpcbdewsmjhfi
tqyvunogzaaflkpsbaewsmjhxi
tsyvunogzarflkpctdewsmbhxi
tqyeunbgzarflkpcbdewrmjhxi
tqyvunogzcrflqpcbdeismjhxi
eqylunogzarflkpcbdewsmjhxy
tqyvundgzarflkpcbdewmmnhxi
tzyvunogzarflkpcndewsmjhxb
tkyvunogzxrflkpcbdewfmjhxi
tqyvunogzarflkxcbdzwsmjfxi
kqavunogzarflkycbdewsmjhxi
tqyvunogzarjlkpcbdxwkmjhxi
tqyvinogzarfqkpcbdewsmjpxi`
