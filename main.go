package main

import (
    "fmt"
    "os"
    "github.com/abadojack/whatlanggo"
)


func main() {
    // Copied from https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
    ISO_639_3_TO_639_1 := make(map[string]string)
    ISO_639_3_TO_639_1["abk"] = "ab"
    ISO_639_3_TO_639_1["aar"] = "aa"
    ISO_639_3_TO_639_1["afr"] = "af"
    ISO_639_3_TO_639_1["aka"] = "ak"
    ISO_639_3_TO_639_1["sqi"] = "sq"
    ISO_639_3_TO_639_1["amh"] = "am"
    ISO_639_3_TO_639_1["ara"] = "ar"
    ISO_639_3_TO_639_1["arg"] = "an"
    ISO_639_3_TO_639_1["hye"] = "hy"
    ISO_639_3_TO_639_1["asm"] = "as"
    ISO_639_3_TO_639_1["ava"] = "av"
    ISO_639_3_TO_639_1["ave"] = "ae"
    ISO_639_3_TO_639_1["aym"] = "ay"
    ISO_639_3_TO_639_1["aze"] = "az"
    ISO_639_3_TO_639_1["bam"] = "bm"
    ISO_639_3_TO_639_1["bak"] = "ba"
    ISO_639_3_TO_639_1["eus"] = "eu"
    ISO_639_3_TO_639_1["bel"] = "be"
    ISO_639_3_TO_639_1["ben"] = "bn"
    ISO_639_3_TO_639_1["bis"] = "bi"
    ISO_639_3_TO_639_1["bos"] = "bs"
    ISO_639_3_TO_639_1["bre"] = "br"
    ISO_639_3_TO_639_1["bul"] = "bg"
    ISO_639_3_TO_639_1["mya"] = "my"
    ISO_639_3_TO_639_1["cat"] = "ca"
    ISO_639_3_TO_639_1["cha"] = "ch"
    ISO_639_3_TO_639_1["che"] = "ce"
    ISO_639_3_TO_639_1["nya"] = "ny"
    ISO_639_3_TO_639_1["zho"] = "zh"
    ISO_639_3_TO_639_1["chv"] = "cv"
    ISO_639_3_TO_639_1["cor"] = "kw"
    ISO_639_3_TO_639_1["cos"] = "co"
    ISO_639_3_TO_639_1["cre"] = "cr"
    ISO_639_3_TO_639_1["hrv"] = "hr"
    ISO_639_3_TO_639_1["ces"] = "cs"
    ISO_639_3_TO_639_1["dan"] = "da"
    ISO_639_3_TO_639_1["div"] = "dv"
    ISO_639_3_TO_639_1["nld"] = "nl"
    ISO_639_3_TO_639_1["dzo"] = "dz"
    ISO_639_3_TO_639_1["eng"] = "en"
    ISO_639_3_TO_639_1["epo"] = "eo"
    ISO_639_3_TO_639_1["est"] = "et"
    ISO_639_3_TO_639_1["ewe"] = "ee"
    ISO_639_3_TO_639_1["fao"] = "fo"
    ISO_639_3_TO_639_1["fij"] = "fj"
    ISO_639_3_TO_639_1["fin"] = "fi"
    ISO_639_3_TO_639_1["fra"] = "fr"
    ISO_639_3_TO_639_1["ful"] = "ff"
    ISO_639_3_TO_639_1["glg"] = "gl"
    ISO_639_3_TO_639_1["kat"] = "ka"
    ISO_639_3_TO_639_1["deu"] = "de"
    ISO_639_3_TO_639_1["ell"] = "el"
    ISO_639_3_TO_639_1["grn"] = "gn"
    ISO_639_3_TO_639_1["guj"] = "gu"
    ISO_639_3_TO_639_1["hat"] = "ht"
    ISO_639_3_TO_639_1["hau"] = "ha"
    ISO_639_3_TO_639_1["heb"] = "he"
    ISO_639_3_TO_639_1["her"] = "hz"
    ISO_639_3_TO_639_1["hin"] = "hi"
    ISO_639_3_TO_639_1["hmo"] = "ho"
    ISO_639_3_TO_639_1["hun"] = "hu"
    ISO_639_3_TO_639_1["ina"] = "ia"
    ISO_639_3_TO_639_1["ind"] = "id"
    ISO_639_3_TO_639_1["ile"] = "ie"
    ISO_639_3_TO_639_1["gle"] = "ga"
    ISO_639_3_TO_639_1["ibo"] = "ig"
    ISO_639_3_TO_639_1["ipk"] = "ik"
    ISO_639_3_TO_639_1["ido"] = "io"
    ISO_639_3_TO_639_1["isl"] = "is"
    ISO_639_3_TO_639_1["ita"] = "it"
    ISO_639_3_TO_639_1["iku"] = "iu"
    ISO_639_3_TO_639_1["jpn"] = "ja"
    ISO_639_3_TO_639_1["jav"] = "jv"
    ISO_639_3_TO_639_1["kal"] = "kl"
    ISO_639_3_TO_639_1["kan"] = "kn"
    ISO_639_3_TO_639_1["kau"] = "kr"
    ISO_639_3_TO_639_1["kas"] = "ks"
    ISO_639_3_TO_639_1["kaz"] = "kk"
    ISO_639_3_TO_639_1["khm"] = "km"
    ISO_639_3_TO_639_1["kik"] = "ki"
    ISO_639_3_TO_639_1["kin"] = "rw"
    ISO_639_3_TO_639_1["kir"] = "ky"
    ISO_639_3_TO_639_1["kom"] = "kv"
    ISO_639_3_TO_639_1["kon"] = "kg"
    ISO_639_3_TO_639_1["kor"] = "ko"
    ISO_639_3_TO_639_1["kur"] = "ku"
    ISO_639_3_TO_639_1["kua"] = "kj"
    ISO_639_3_TO_639_1["lat"] = "la"
    ISO_639_3_TO_639_1["ltz"] = "lb"
    ISO_639_3_TO_639_1["lug"] = "lg"
    ISO_639_3_TO_639_1["lim"] = "li"
    ISO_639_3_TO_639_1["lin"] = "ln"
    ISO_639_3_TO_639_1["lao"] = "lo"
    ISO_639_3_TO_639_1["lit"] = "lt"
    ISO_639_3_TO_639_1["lub"] = "lu"
    ISO_639_3_TO_639_1["lav"] = "lv"
    ISO_639_3_TO_639_1["glv"] = "gv"
    ISO_639_3_TO_639_1["mkd"] = "mk"
    ISO_639_3_TO_639_1["mlg"] = "mg"
    ISO_639_3_TO_639_1["msa"] = "ms"
    ISO_639_3_TO_639_1["mal"] = "ml"
    ISO_639_3_TO_639_1["mlt"] = "mt"
    ISO_639_3_TO_639_1["mri"] = "mi"
    ISO_639_3_TO_639_1["mar"] = "mr"
    ISO_639_3_TO_639_1["mah"] = "mh"
    ISO_639_3_TO_639_1["mon"] = "mn"
    ISO_639_3_TO_639_1["nau"] = "na"
    ISO_639_3_TO_639_1["nav"] = "nv"
    ISO_639_3_TO_639_1["nde"] = "nd"
    ISO_639_3_TO_639_1["nep"] = "ne"
    ISO_639_3_TO_639_1["ndo"] = "ng"
    ISO_639_3_TO_639_1["nob"] = "nb"
    ISO_639_3_TO_639_1["nno"] = "nn"
    ISO_639_3_TO_639_1["nor"] = "no"
    ISO_639_3_TO_639_1["iii"] = "ii"
    ISO_639_3_TO_639_1["nbl"] = "nr"
    ISO_639_3_TO_639_1["oci"] = "oc"
    ISO_639_3_TO_639_1["oji"] = "oj"
    ISO_639_3_TO_639_1["chu"] = "cu"
    ISO_639_3_TO_639_1["orm"] = "om"
    ISO_639_3_TO_639_1["ori"] = "or"
    ISO_639_3_TO_639_1["oss"] = "os"
    ISO_639_3_TO_639_1["pan"] = "pa"
    ISO_639_3_TO_639_1["pli"] = "pi"
    ISO_639_3_TO_639_1["fas"] = "fa"
    ISO_639_3_TO_639_1["pol"] = "pl"
    ISO_639_3_TO_639_1["pus"] = "ps"
    ISO_639_3_TO_639_1["por"] = "pt"
    ISO_639_3_TO_639_1["que"] = "qu"
    ISO_639_3_TO_639_1["roh"] = "rm"
    ISO_639_3_TO_639_1["run"] = "rn"
    ISO_639_3_TO_639_1["ron"] = "ro"
    ISO_639_3_TO_639_1["rus"] = "ru"
    ISO_639_3_TO_639_1["san"] = "sa"
    ISO_639_3_TO_639_1["srd"] = "sc"
    ISO_639_3_TO_639_1["snd"] = "sd"
    ISO_639_3_TO_639_1["sme"] = "se"
    ISO_639_3_TO_639_1["smo"] = "sm"
    ISO_639_3_TO_639_1["sag"] = "sg"
    ISO_639_3_TO_639_1["srp"] = "sr"
    ISO_639_3_TO_639_1["gla"] = "gd"
    ISO_639_3_TO_639_1["sna"] = "sn"
    ISO_639_3_TO_639_1["sin"] = "si"
    ISO_639_3_TO_639_1["slk"] = "sk"
    ISO_639_3_TO_639_1["slv"] = "sl"
    ISO_639_3_TO_639_1["som"] = "so"
    ISO_639_3_TO_639_1["sot"] = "st"
    ISO_639_3_TO_639_1["spa"] = "es"
    ISO_639_3_TO_639_1["sun"] = "su"
    ISO_639_3_TO_639_1["swa"] = "sw"
    ISO_639_3_TO_639_1["ssw"] = "ss"
    ISO_639_3_TO_639_1["swe"] = "sv"
    ISO_639_3_TO_639_1["tam"] = "ta"
    ISO_639_3_TO_639_1["tel"] = "te"
    ISO_639_3_TO_639_1["tgk"] = "tg"
    ISO_639_3_TO_639_1["tha"] = "th"
    ISO_639_3_TO_639_1["tir"] = "ti"
    ISO_639_3_TO_639_1["bod"] = "bo"
    ISO_639_3_TO_639_1["tuk"] = "tk"
    ISO_639_3_TO_639_1["tgl"] = "tl"
    ISO_639_3_TO_639_1["tsn"] = "tn"
    ISO_639_3_TO_639_1["ton"] = "to"
    ISO_639_3_TO_639_1["tur"] = "tr"
    ISO_639_3_TO_639_1["tso"] = "ts"
    ISO_639_3_TO_639_1["tat"] = "tt"
    ISO_639_3_TO_639_1["twi"] = "tw"
    ISO_639_3_TO_639_1["tah"] = "ty"
    ISO_639_3_TO_639_1["uig"] = "ug"
    ISO_639_3_TO_639_1["ukr"] = "uk"
    ISO_639_3_TO_639_1["urd"] = "ur"
    ISO_639_3_TO_639_1["uzb"] = "uz"
    ISO_639_3_TO_639_1["ven"] = "ve"
    ISO_639_3_TO_639_1["vie"] = "vi"
    ISO_639_3_TO_639_1["vol"] = "vo"
    ISO_639_3_TO_639_1["wln"] = "wa"
    ISO_639_3_TO_639_1["cym"] = "cy"
    ISO_639_3_TO_639_1["wol"] = "wo"
    ISO_639_3_TO_639_1["fry"] = "fy"
    ISO_639_3_TO_639_1["xho"] = "xh"
    ISO_639_3_TO_639_1["yid"] = "yi"
    ISO_639_3_TO_639_1["yor"] = "yo"
    ISO_639_3_TO_639_1["zha"] = "za"
    ISO_639_3_TO_639_1["zul"] = "zu"

    args := os.Args[1:]
    info := whatlanggo.Detect(args[0])
    // fmt.Println(args[0])
    lang_639_3 := whatlanggo.LangToString(info.Lang)
    // fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])
    fmt.Println(ISO_639_3_TO_639_1[lang_639_3])
}
