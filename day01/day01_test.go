package day01

import (
	"testing"
)

func TestGetCalibrationNumberEnds(t *testing.T) {
	line := "1abc2"
	got := GetCalibrationNumber(line)
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberInside(t *testing.T) {
	line := "pqr3stu8vwx"
	got := GetCalibrationNumber(line)
	want := 38

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberMultiple(t *testing.T) {
	line := "a1b2c3d4e5f"
	got := GetCalibrationNumber(line)
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetCalibrationNumberSingle(t *testing.T) {
	line := "treb7uchet"
	got := GetCalibrationNumber(line)
	want := 77

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumCalibrationValuesExample(t *testing.T) {
	lines := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	got := SumCalibrationValues(lines)
	want := 142

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSumCalibrationValuesInput(t *testing.T) {
	lines := `29lzrxseven
	9nnqljsixkzphvtmtr
	five73kskpfgbkcltwoccvf
	fiveq7
	npfdbnjlthree6cmdbsrsqc9five
	fourthree5threenineq
	vdxpfourone1393mtvtlzmj
	onetwo6
	62one
	vttq7sixninekgxtpspjgkninesnpktchz
	2954two
	395fiveqqjqrkxxrbseven6ncnnpzffj
	one92seven3pfbhlqzt
	5sixtwo75four
	eightcvcprdttmrbtttznine9
	two4eight
	ndndrfmoneone3
	85one7szmffjstpbdssixpfqbbcljn6
	16nine9eight
	nrshscpxfivepqxccgpzfeight3zcvgvkvlcdmrmqmrqrj
	eight3c
	sixfive3nine57sevenvqbtsh3
	38pbqgtwofmhjnjbkeight
	1vvtcclvnineeight
	fjeightwo9xtxrmfive5three
	threeqndg6prxnnrndpmktbsccpfktzfqvrdlnvvq
	1fivefive3pvthree
	eighttwo257djtdp5two
	lgqhbcsfmlg3onepfsvsxbvnnqeight
	threesix1dgbmcfmcrl1eightnrvgbmrvb
	qgrgqjlszpcnpq82
	3fourlbgkgvgpqml
	bdrdrt38
	three3six7two4fivecvxkpnt
	rdtvngcd4twotwoone4sevenbdxjfive
	9foureightonemd
	ninevbpvvhlzncthree5f
	nineeightjknsdxtx1fivefive3ccscvtp
	66one4pzctwo
	dcqgxpczqlsrh61txjmpscxqbhshtz16
	1fourkvjqdd
	75eight88
	11nine
	three9two5two135
	cq23jgpmpgqqj52
	xbdqfpdnfour7
	3ninejnjfjmfsjntxeight
	one3ninejvhnxphl82seven9nine
	sixhlhxseven6gnppph
	78qknxhztlvtgbrtkcrxmfldjbbrdgxeight
	nine8n9385oneone
	94tvmmthree7jgfourrkhxdszrgc
	ftwonevbhlhlqhsix5dzfqgsone
	vjppsvjrh8nrzlvhxnrxdninedrqmkld44
	9sevenfivebxhbvqhblcqpznkonepmsbvfmtjzbn
	47nine46six7
	9qvmrcjd8lpskqjgsdtxz
	6zsninesevenxzddcfninefivetwo
	97jsrrspbhl4sixs4
	1nine9eightfive12five
	6gjx3tqdjstonenndkdsvrtthree
	tvdd7htfive225four1
	six7733eight88
	vsseven6
	eightftwoone5twombmjfbnkk
	2lbpgtwokrphdqzeightctbc9frdtpgrzf
	39one8fourfive57
	one6xzxmmjpone7ngdlvvhljmhztlbsgxthree
	gcjcdhmp74nine
	39vrbnz3s3sevenone
	gvzqtjslsmbzqzfgc2
	4tppjrkprbkqkbtgrqfivecpztkprhhheight
	599
	pctwone5one
	trrzzszlsq7onenine47lvprfqdpbkqgvvgrqg
	2eightkrqftb
	vgrrthreeone9twosix
	2bb
	onemlpmcvhvzqssxzptd63five48mjnfsmq
	8nmeightsixhmdjphreightsixseven
	f6seven4pxmglflmqckcsskfkqrmzhbrnvlqgr
	3jmxzhttjc5four32
	cqdfninekrldlfzdnhvcnl8
	572qntvlqjsfzrsnpsgsgdfpzmdcdlhsp
	6bngnt52
	ceightwo7hnqnvfx8nine3eight6
	three93fvkgjsr73
	seven7dgdpxqncgv1fpdxvtrmg8
	53fiveseveneightninekk
	flzmhmbdcb5sevenjcnshhcz1
	ntwonefoursixkpmjlzfpjdkdjfgvkfj2jpxt
	kzbnbqfcknhnine5lqfgcnl
	zsqsdrxxtnlcmx1pjtkzkbqpvzxtvvbrhf
	stbqnrhdqnjcvjgthtmht8xndfgprq3eightwol
	three7eightsix8seven7
	knttsixseven73ninebsjkdxxlp5
	sevensxkltql27
	onelxsixbmznqnfqtmdscjthree5dltxeightwox
	6onegqvfdtpbsixonesixthreeqfdvpncjfsgv
	m2slphnnccvbsqfivegtwo81
	sixseven5bqlfourfmtpvzjknrlppnbffn
	59ghxvhtdsgeightoneoneseven
	4jzeightoneeightwobvh
	2two1bnprdmxhgkvfour
	s2onepkqseven9ltmqtjj7pdhtrrnqsd
	dpkfdlhcrjfivedntgkqclhzzntwodtsmfourseven1
	fbnbj5two3twoneg
	twofour3six9one26nvd
	twonine1threeonetwosevenvgsdxxhpl
	onefourhx6deightfiveb
	two718
	sevenhhrslnjn6onemhg7two
	62cfgvscmxxfgthhjgsb
	8ptwooneone2hfvcfgb
	5two5
	eightseven86sjsqsmbone
	lgtxmqhdfp9
	zdmskllfour27four
	5fourfive2tdnfctxvzmfrnfive2
	five34fourncpzvdnmhrrg7
	one48eightjxjfmxljvs
	three2sixprmhc
	rx5six
	fourdb13lfjjbpx74three
	five42
	8sixthreemfgl7vqhtr783
	8two1msvkqrr6klgsdmrtcqv
	85rjrhtk4xqbeightthree1
	79three18kjvmnbnjnrvrgh
	8cfpseven97eightkslkrlxtrssdh
	trrmhcmtwo1rpkb27fh
	threerjoneonepmdjgcrjlmlqvqbpzg7three
	leightwothreeninesixtsljvdl1nflg249
	ph59fzlxzfjdxpqmzp24
	hhjxhjthcnkhxrgmcrlonevzrxjqs87
	91psxvvmnqqsixxthreetwo
	9jbb73kdxhxxcvlhlcdknine
	ninesixntrngtp7jzkgbgxjp
	12sevenqnchlxzh1mhthddbl
	vqmdttbkpkbhzcf9341
	bgvqpsxhtwoseven94lrkxkrjxq6
	92one5hkmdbmgkqdqnsk
	2jdkfjttwo1
	q9
	pzqjsvvjnsfzxrtmqdxv6
	kxcxtnqsixseven2rnjncvbcbbrxhbgtwonine
	jxpvsbmxsglfqs4pjbffive4ninesix2
	eightlnv4
	qxdzrcjmcbx9rchjmvkt
	571z6
	vqfqtsn7oneonethreeeight
	lnnmzone216nscppjvtlqztkbcxsix
	six33lfvgmjx
	kjpscbx41bvgjcpdkfive48sfpgdchclhlzmn
	rfqhqqrxqtrnmkqsrkponeltl1
	three7nine45xqspsfour1jdqcnjz
	96rt49fvhzbknvxpfive
	eightxpgqg9
	onetqcrxzzmt6cbjh
	ninerfbvpbzbxqseven2
	9sixkhqrrk6four
	9jff17
	4five3ldtbmsdbcvttrkvprseven
	7sixvgbttfdtcjcxvflnklpsnjpp54three
	eightmpndj74nine
	hmbncvctfive72h
	fxnm7twoseven87
	oneeight11ninejbhcprshx74
	78eight29fvqtf
	2ffdz
	meightwoeightfournineqjjxfourvkxsvcbn58
	7cq5smpmgntnhvgrslb16one
	33six
	one5grbxtprmbtlninefive
	three7qfnzmbqmnh563six55eightwot
	kvtpmkrhvzfourxdknp6fivegqcshz4one
	sixznzftmbf9ccflfiveone
	6two6rmsjnmdmtbxmtjhcpsqkvppjfvp
	lxqnslxqneighttwo7
	3ninelmzqrrkvmxpkpsdtghtkxrfcfive
	tgknhvqz5cdjfourpzssxgxxrx5
	6223
	1threesmvhfrkxdnxhfthree1nqcrfjhnt2
	four7bflnvpfqhfclhnsc
	msmd6
	xcqx5ninecc4onefour
	sqrdkpzeight936oneighth
	three9pgm
	zzskhnglgvthree83xs4
	8xccmplxb5
	j8171cqjmonetwo
	7xxlxkheight7eightjpxdql56
	oneseven9nine
	19fivezcltxsklvx8nine2nine
	lcjvnn991six6six
	zcjcnnvfive852hrqlktzskpqgtchjb
	sevennine59
	9fdbstknkd3
	8nineninekcrrshrgpttjsskvbfnhgrmbkzfd
	njtonexpzhdszdcthreedsdqsfvp5
	4rhfk
	68six9ninembksbtzk
	57eightthreetwothree
	5eight8three5lrhdbsnj3lncs
	lthree68seven
	fivetwo2vmndlgtbkg4
	2rfjdshttkkvl523
	b7mrzsfour
	nlzsg13
	jpfcvrfkqbszvht93
	9jqts83lxprbsbpqx5seven
	kbnf78hktfn
	seven43rvjldsfvzpssjlsevenfcpjfxtls4
	sixsix12djn
	fouroneeighthxtpppjrhqqmlzvvsrqk9gmkqpqkplpg
	8twoone577
	zcn6sixfour5three
	gb8three
	nine7rnzsqshxrxnineone
	bgrnone7vqsxmjneightpfour9nine
	fbznfkxhvd9nhvpvqhlmzvr66five
	bvkglbdzmvnine3eight1t9
	365
	fxvfqkclm8oneonesix8
	45four2twombjhhssnonefour
	sixgkhcf92
	7thtrjnlxj3sevenone
	fourthreeeight9hhs49gqrxcmk
	ftbzpdgljnknktwosevenvlk4
	two6tmjfllfdhthreeseven8
	eightzfphtm6ct39five
	7sevenfivethreegnmpteight
	1one3rfjfhmzqqz522fourhmhvqrqmx
	383nineeight9eightjfdhmjfrj
	sqkmgfxzdgp7sevenqlvcrtlqkjdrbhnrvhrxjlkrfmgcfour
	twoshslmzcjshcnxvhnh9twomscx1
	hrdtwosix4threehrqqczzpflseven5nqcct
	6ninezltlrxrgsnjxpkjvnmlzvcpjxkkmbld
	984nqngbcdnscqhbdzfmpjs
	5sixfivethreesixfivemrsxjglmzbtsneight
	gsg8xsgtwofive13one
	cgdhztzxg57rmj
	8sixldnzrrrthreeonesixsixfive
	47twofd
	gs4ntjbjnfmn2kdxxsix
	1gbfj86xntxgnine55gsmz
	5gtsgxbpxkblvfh4six6
	2fzzsskzthreethreetwogqldvtfmzcsz
	jbjggvkf51fourrcltfrfive
	foureight5
	rnrt5fjkjjfkgdnqdseven9tbrrpc
	cbftwonesfvlgqq939
	kjzqzdv75eightt
	4mlsninejknfgvdhpgtwokphtgvqsg53
	6d2fivefgfoneeight2
	onejrnjpkbcsjkfeightxvbzjtdmb3qz6
	8fmmthreeeight6five
	t19182foureight
	seven4nine4five16eight
	4bklkqnfht6tcnxjxrtmn1zglmtrfive
	qbzszxdv42r8
	19kgnmmf3txszjcjrqgxrfive
	7sixqlvhrkrn7fivesix
	8fzqkmfrvhqn6brzxx4
	2sevenfour6sixfourrrvptzrhtv
	k47one13qxzpzstnm
	3fourfourcgj4
	4lhh
	6nmjbcbqqjkttvcxjhjrplpcsix4
	8rntsllnpd
	j5eightfhxzhjmjcthrrtxf3
	6ndkkbtljlkr5
	sbs5nine
	prjsqdx11
	7onesevenseventwo3one1twones
	691
	twonhmbtxd8981
	b5
	4mgzfvgcbztqc
	qlgkqzbtclsixdcsrjmldfqcjnk29
	sixngv4
	kqqjpnnmfourchkcphjq1xn
	four87one9twoxzcskqbgc7
	seveneightone3hxsxjrninepzqzfjc
	2four35five
	ffftthhnv4jpddcphseven795
	six9sdlljbmxtnnkhfourzn5zhr
	7four24twoltngrkhtwo
	573sdvhrhfhxfxdsqhjrz
	dvcfqmzltkone2mrtcc4eight
	1jrgxblmdqz4
	fiveeightninethree4bsevenfour
	crpz9three3791
	onethreedfrxkvqs53three1
	dnrmthvknineoneeight8lgvdnjt
	bjbvskthreeseven29mrmt
	lbbjdvlqfb5rnine
	zfloneight2phbzdvjhmtwoninepgjlxpbjtjmgdbeight
	seven3xhfgkoneseven
	jlgmxqrbtwo4
	one7sffjzqzsxq
	6xncthreefqvhkmjdd
	7twooneone5
	7dhzht44pmknlmsqdhnine
	fxbcxbnkdninesqsixtwo2twoksdcqfqccnltk
	chhl58
	9threefourfour
	54eight3eight
	79sbsqj62
	blncmcqnsxfive5threesixqkzv9three
	nnrxqdtth7threesixsix
	jklzjg5fouronelsfnqtpsix
	eightlmvqn8rhvlm
	eight3dv32twofour
	cxndxfcsix5rgvtmlbpvcftphq73
	fourhcqrdsfcxjdxj6crnzrclbggvsvq
	ninesix9znineninetgvr5
	nine5rsqbrrlmnztdgspx
	6nine7lthcjbbqskktvnmnz6four8nine
	pg4fxgvfhtg32one1
	pceightwosixfourthreemkqtncvztwopqdvmtqlleightrjnmm9
	3eightsvqdbgcntv37twonine
	4twoktwoone42zsgcxlgkjl
	4hlfvmpmssf7
	zhcvpsevenseven3
	7tsdxhfrfdkeighteightnine3
	tzeight8sixeightmqjglrcpc
	kponeightcmhqfzc9seven6lxrxh5
	84dpccxdjvxkmngdlxzct
	6scjxxctt3txvnlm4n2nine
	2fivexjjf551pnkgdz
	34dclnf633four
	7vvkqbhscconefive7sixfour
	fourfiveoneone4dh
	fivesevenlzbbzr7one
	jxsxcvp5sixnine6pleightdp
	5xbcnhrkgzxsonethreefoursix
	6ld4
	vl38nine2
	335
	svreight93fkksv
	26hfsgjrhmftqgljglbfgrvzjvl5
	sixtgngfsmq148
	four9qzndrcxkr72241
	3xhznspqsix1
	7ninesevensix6xtlzkhgrjxzngqvg
	nine51
	one92mrxhfxkrv
	2ptjpbk42six
	4cnrclct
	zxpfnqt48m4
	qmcbkgrzf4xkjfhclfr76xgvzvhhvsml
	fourqf4354
	8tgfspgpbnine49fivendfhr5four
	2fvdx
	7vksevensix2pjqnqgsqtjoneone
	rxnneight9four
	fivetwoeight9fivefbmlthrneightq3
	2cfpzkbvncvnine7lzvfiveeightkvrdqeight
	8sixvsvnnm29
	gbmppktltvzzcprsm825
	dkthree7cjhfvpfpfq
	tltssfqbjmtrnmjnntsfzcqmhfkncgsixl5
	xbskrbrszxbmbxskcvrrvfzd8
	4vd
	8fourthree
	b6
	onefourcrtxdhkz6
	z4hhfive5three5
	rklffnxntvgprtrxztwoeight31
	8b
	eightnqnzthreerfpr7
	rkjlggxft4cvnzncrrlseven
	lvoneeight62533ltbq
	6eightninetwoztphnkhdthreehdldhvkfvp2
	fpnqfqone4sixmcjnqbj
	twoeighttwo8kmzcsvxhqt
	5bstjtjbgfsixfour2vhsblxnvptx
	5pdg6
	9xztgbqc91
	7nrjxssvnxr99five24nfn
	79456eightsix1
	onembcqhcszk2nineone45hkvxtkdfive
	69slhrknjnvsevenone1ctwo5
	22fiveninemn6pxqhgvhqkjonethree
	1five21five
	fivezjvzfive4
	five9sshrnpxxxjjlkt
	fdznqdjdcmvnhjrhpgdnjmsljdone8
	two8threemmeight
	bdzrdgpmd1
	1cnmvtknine2oneeight
	one4zklszzvdd
	twokvppczbhdkdnpkmkcp7
	mrnsdlc465seveneight9nmlpfive
	lzgrvsflljtkcs3fiveonenxnrljrbdvs
	1sevenfour
	ninehzrfourseventhree91
	3vlsix
	8bpgkrcznine4
	eight33threeeight3twonepr
	26threekfnklxrsc1eight
	seven92b8nlvtrrdkksone
	9fivenine5dqdbdnnnbd2xpnnnlxszzeight
	four5fiveqvjxtbvvpn
	htwonine6
	tgvhnmphninepsrxjqrsktnj4three6
	seven9lpgxhqbhfrvjkptfjkxgslc2two5bc
	1dcbhbvvl29threezthggl9six
	tjt9fivehseven8one9
	cjzgdrb56twoljpdhhnfive
	9onefourclmsrzkcxkkxjbjbnine6
	ftgzppdr2
	dfdtddxjgxhthvzpnv6two63fivetb
	4bczfllzdztqv
	fpgppvsl2fivesevenslhbjqsf5one2
	tfpzhjbmltxkc1rptpdnlgm7
	ctrfnfxvbjkr9five1threeeightwontc
	vstwone2one
	zssoneight9
	four1twooneightkvm
	9five2
	9six3nine5921seven
	3one1
	7pxlf
	g12xtfdmsix8six14
	cmfourcgjthreesixfivenine6
	2pjpbpnljp
	tcqgcdlfjmm4two7one
	gx3nrvmrvfrbd
	onefourmxrcbmf1six
	five58f5dsbxfvgztwo
	24four62one
	48mnfcrkgbvxc
	five2krvfsg
	2eight7
	67sevenkgsvlthreekhsix
	tone36tnbbgftvzffeight
	kpc87
	eightsqm4hbfndbfnseventrsrsrvhqssbgvq
	nine7139tgdxlklkdcvqdvm
	qzlts53dvrseven7two
	167ggvsvgrdgs2zgjmxlnthreetrgx
	gdplsdjf379eight
	blkfjx4
	558rtsqsgpgjknine
	81six
	2ninenvp
	62sm68one
	twobhrtzvmfive8msxvl3nbdpn
	4jgeightbjxzsqtnvd3nineeight1
	8564oneddbcqkxgfqjcl8
	6two68six61nine
	cztwone2thdmkgftvmcccfsxrfpxxm
	four8fcdpdcghzzmnconerjmkznbkv17
	jfourfive4v53sfnine
	4543fiveptzlmzpblflvqklseven
	58qglltjtwosixclpjbmq
	3gmsphjqjxn9dthhn6eight
	fourpgqstrrc6xdfnrmfprn59
	5threecxxjnk
	five6jhbkbbc
	nvvtslxbggvx23foursixkld
	srcrvmcczfvgfqqh79seven7four
	jmvlfqhxtfvn12hd
	hv5seven
	nxrtnggfnineqhmhmrhp8ninetwo
	mrtkcqvgsixsixfive54
	h5sevenmrsk
	qftone1tmgsrdsix
	vqvfbxfsix13sdhjznmttnine1
	3threetworcpsxtbpvsgtfmnine
	81nq9
	two19ninetwo
	dtvrjxzlthree4
	four1853dqone56
	pvcnbrjpsznfjv1jfhbhsmvtwo1two
	7seventwo3eight
	8threetwofour464gdqrzphgkxnc
	64fpvcznnsrbfour
	9twofive
	9339five2jvcsdfivemzrgjfmnnl
	four5cfstbksix7sevenpsix4
	jctwonecfqgpqkhhmninenine3ninejhqsbfsgjj
	4twosix4gshj7lpkxrtdfb
	js7
	77rxmvrlqkcskjnqxmqxbg
	ct2one
	glcbv6eightseven3eight3sevenqcg
	6sixonedxeightone8
	4oneeighttwonine
	three9nljrrxqtxhnpsrlvbtjdznhjgmhkncrpmcpx
	vvnclvjj18
	eightfivetbqc9seven93
	dmsh1onethreeeightsixxqdcxfive
	mxkfpmmvhtwo8rbxrvgbsgzmcrfsfhmqcxkr5jxzzb
	onepqxxnxknfourthree2
	four9185nftjmvfour
	one8ninerrssdsfhkjltdsndjgpfour
	tttqlggtnine33tzp
	7eight99mcztone3vhm
	vbvpzvt9three8
	95jmbnlfvk1four7
	2tgzdneightfive4
	8seven3two3kprxhpcxssevenmd
	twoeight89mngzqlrx2
	v78jnbrxpqpsixeighttwo5ljpvphjld
	djmmbtgpchninekhpgrzskdddmmg61
	j13six78zdsmllspj
	3qbqqpkljnqsix6xrldfr9
	93fourfxnb553fhpdv3
	rvnlmxfbfour3
	6t51
	69oneznlvmmpbfourbvkpxg
	2dffthreetwothree
	sevenonekvqrfv748pzkhxgddh7jtcbhdrcmj
	onexhzzdqmtb6dmqmkvzpjftcm
	four2onethreerk
	threesevenvxzdlpfmq8
	seventhree172three9
	fourspjcvlfkm9
	ninefoureight4six
	nine6nstfivesix7eight
	18mrfktwohnslcdxl
	ldxjnvjpq7kstghdnkf
	dbdvfoursixonefourghplcnlqrd998
	1xhhrqzx24
	5fgfjvgps241fqzzf
	4rjhdv
	five8qfnkqrrtbninefour3vqb
	6jbtmcfrjq
	rnspzzhqtnine7lfllf
	tkgxsmjhprnnd1pbgftd7
	hcjf4four1fourseven141
	3eightfiveplfrzmbnclsevenfive1
	mx1hshhprppqnl64
	bvpqd6kmlntxv95ninekgdtks4
	kbcnqxc4eighths4dvkkvndpzhsix
	1sixfivefivesix7f7twonel
	ztqxklqbcnineeightthreesix5fivetwotwo
	4sixkmgpq68five7
	fpccz628p3
	sixfiveseven8lsqtwohkreight2
	5ninevjncfour
	sevenjxbflccqdxbkqonethreenineqr2
	lcseven3
	eight8threeone
	6four8eightwokt
	8jrckfmrjct
	sevendxcrjj1
	1gmngdnkhntwofourvfdtzhdnrzdqprrnf7two
	5cmsfmdqtjc368five
	mqxkkhxdhvd6zjlzqzqz
	2qphczzbsixnine5fourlsgh
	twofourtwo6one3pgd
	jt8four9threeffvnnk9
	xkcskxvlrsix4onesixfive7
	jqbzkfjlnk8eightnine79nine5four
	drppqxhzkxtbzrsbh6seven
	5fivepnjdfgmd8cvtgjvfrslcjll
	jr69
	rlftnx6six
	38three3
	27eightthreefour3pvtwo
	29fourqlkslrsq
	79g2one3sixqfbqhgbtdt
	97sevenseven
	fivefive5fivebdnmh
	4nine8nine4seven
	918pcggpvmhgs
	five37three
	htn4tztkrq
	23onefourthreefour5
	mbzmtbsvone7xbxvprdcjdonefourbtnkkgnmcr
	2kksixsix
	44three
	7nine5fiveoneninevpbnvrqxlsix2
	jsthree4p8427
	jmthree4
	5six1fiveseven4
	fbqrfn94twosix9kfgdcx
	threethreevdcqqlrt2fourns7fourseven
	rrrxrxtpstwo9mtmlxrjdfour35eightsix
	5four2
	cbeightwo3
	9zrfx
	vgxdtlr3rsninekzgvsgvt
	bzvbxprzvnfivetwobclh881
	4nzztwo5five7hkx
	rblhc87
	cljrtkmldjonemsqchdcjc6
	xhninesevenkonedjsfhtznkkszm7three
	63gvpflbxrvlsevenlqdsx2
	ninefour6fiveclqnnq7pmg
	6fourjvljvgsevenk3
	twofive5nine
	tworfnjt8fivesdmlnm72tpglhrone
	14twotwo
	84qsoneightdxt
	six9pcthree
	five9935oneptphsctwo6
	seventhreenine63nnszbxddjslqsgjdj
	zpmclmrmql8dmfkfour8three
	6onethreefivekpmv
	1jdcmhfour73eight9six6
	397threesixvlk8one
	sixmqntxqlpdh54
	hcleightwo8fourtwotwotwo2fournfour
	two7four6five7threepdhpfive
	onegdxlkddhkndtgbfgh95three
	skjqgfnq55sixsrxzprdqfour4
	113bsixsxtwo
	9tssbdp92453onesxfpkg
	sevensix1mpc
	one8sixeightseven
	threevsbtbtjpdnbtbjjkseven1
	ns5twoneg
	fourfive6five6six
	7twokcdcpb7
	43nine1lgcdskpztsixthree1
	fourzlqhcpn8eight
	3btgcksvlzzqpghllq
	82threebrnxvgqtlvxnxzpk84
	three872
	7three3
	94fivedktkggcmfour
	fourgvkdzseven63two
	mtbrh6cv36five2gm7
	8ngzftpfvjzthsv
	bmhbq5
	4threekprhxjbj6
	l7ninenineone9two36
	threefpzq4tlzmrjthsthree6ninesix
	4bbjkdrc89fvdqsrcl
	62lpmdbmjbfourglxxh
	1mtnpqdlctt4
	qf9gblvrhkqv9rknxrgf9two
	12fjdjrcfmx
	nine1six
	22two3rkpxjvkhflqfjgf
	2488four
	8vlgcjqqstz1qs1two
	vsoneighteightcqtv7fphhbbj1sixzrhhfgj3
	xeightworlggcdsdk18five8one47five
	4eight8
	4ldcmctd4j
	xmoneightxkq512pdlpknhrvhonesevensix
	lptfv5lhhkdvlktzttjh7
	dtflmvvczninesevenfive8
	pmbhzsqpxtvninetwoctktrmtcvv7
	3rv
	two1mvhttdnpd94seven
	gqpbznlqsdsh6zvgseven
	twohsixgqgcvhkjrz4four3
	eight7nine1
	ninempfnhgdfpkxrzl3hzcgxhtcfctwosixkphztwo
	69five
	nine3rfmm
	76mmxrr72five
	cshknkjeightnine9
	hqhzptzcpfsn78cdfdzmtxjrvqvrkpc9
	ninenine6four
	four27
	3djdrbffcqfourfivethree92six
	vjdnineszrzqcbd3594
	bk56nine24
	eightqtvd563
	rhcfbpssl5tvlbjd3one
	fourzpl3
	qjxdnjx6lhmmcsix7threergqb
	3six67threehq6ldh
	nine5three1
	seven9fivehqkdngpeightlbplfpgnjlvlmjmttone1
	snlrgqgfninegxdvnbkjjs84six3qtqhghh
	tvchztbkninenineseven83
	3kgbsvmhlgtvhscl
	svzz1four4eightxtfsvt34
	nfks2vzvsgbqf
	8eightfivethree
	hltdcxmfxbnskvhqsevenmgb1
	2threesix6zv4two5seven
	dttqcxj2phftphqnqczcgbzsevenfour9nxjf
	mj689mddc1three
	four11sixrmtjtdsevenseven1
	2nlszmgkthreesix
	15sixkdzcndm6three2gpftjlr
	8dxkhtpfqm
	7jgnhvb
	4tscbjh
	nineseven6slg6jfdjvtwoknxmxsb2
	mcdfnbzbtl3sixfgpsmhnqxdldz4
	three3seventmbpxmfksrh3
	4vtbkjlzhdvninetwo152k
	threeclmlfbz4sixeightvkrtlsnlkczm7pnrrxrdh
	onef5eightsevenzlrlvj
	eight99five
	onekrtmtsthreefourninell64
	eight5three6six
	29ht
	qbsjkkkjrv7seven
	gsxfpgreightzvqfbmh4
	onetwornine2nine
	87xhgxb4kbmqds
	4qhjghqqzfour8
	213five5
	crqvxvfsktwo9
	64qfpfpbxhzpdlthree8five
	sixfhdr7
	hzdnzbc4fourthree
	lprmfrxxzf2
	onekkdnpczfj7
	xl2one8cfxvgpgxzp41fourh
	848jqvfxqmdfktwofournlqsqqsfmb
	81six291
	xnxnqtlzqcttwo9
	lvfq254one
	seventhree3
	fiveonenineqcldqlhlqjbdmtbntqs9xbshtltxg
	5fiveljvgksvfq6krfljkqdqnmvmtg6
	8fournine
	4mxpfdq7zlmvslnj99
	6txrss9threejhxjps
	fourfive1qrkbjppbnqeight
	bb2
	3nine34652tthctwonebp
	vpdz7eight8fouronefive5
	ckpqmllrqrcbgdcmzpxcpqxtx66
	nzloneightoneninefivecrxtd5
	4dpdvdqdqq6mtnl
	nine9eightbkdpnkpmfour
	tvtlcqjqdp1vzgr82pgsbffjpbeightjfzdzmzvm
	gqoneightsevensix94five
	2fivesevenoneeightqmpcslzr
	8xvsbxpfkfdthreeqsksgbhh4onefivejhjfpbs6
	5three87
	29eightsevenfour9
	7cfivefiverqzbhdftrggvnhnf
	xjtwonefivevtjvkz61nrsnmfr
	1hnrjdgtvz6two7
	kfdqhfml55
	vbphsbsfz8twonine6
	eight9two
	eightkvrbdl3one
	2csixnine
	tsqmlqrt9zonetwo
	mqhhlf3eightonen
	kj5nslbsn15
	2966threeone
	3six415vnine
	9khvgnllhpq967nine99
	qzrvbnckh3
	458sevenonenine
	8twofivexk5chqhvntdlp4zkbtr8twoneqh
	fivezcprsxseven88sevencplfgznnmvhoneightstr
	3ffztxgzqfivetjzqdkr1three
	nine3zsjtr2seven
	grlh3bbrlsevenqqhnzclxvlrgrtthcxjqpdbxkbsrr
	4four2eightdkgxp728
	9rsjhxllvthtlcsevensix
	qhnmhzzfxfmftghvbv6eighthljb2
	nbmczbmcdlmg6seventkbvzkcdspcmtslcp
	bninefivecone83
	2dfhkq
	3bfsbpvxkjdhhpjskfd
	9xskpfbm
	two3vgr8eight
	pfprtzxpxks3fourmkjkrcgcrstwo9seven
	dtgmvxskmfdfzxmfour3two3
	fourrlksqss3
	hbxbsixrkbrcnz5bbnmqfzgsbs
	sdhtrdsdtjdsqvzsix1jnclhjsd
	fourtwo4
	xrsjrgffb2gksvsxfldjzlgkg
	6gzgrzgplqfkfjkrvtwo5
	five8fjzttsdsreightxxshcbxfour2
	pcvx2mmnhsgkqtcppb9twoq6
	qtwone1bssfvxsdd28qhone
	fiveninethfrpzlxkj8fdvsixtwosgd
	sixthmsjkp9twohvndmpv93
	1sixfour
	bznkdpmrprrdfxb6fpnpr4
	lshcpgffvxccd1fourninednm
	onesix8xnhpcfkrgnbbxzmq6
	25sixfour1lbbdqmsxzfnkgvtftgsgdrg
	2mdqzqqc
	hjjsjp7cgtqdfxhqvkzzdc
	2snrf71one
	threesevenfour516
	jmfltlkhv7rktq4two
	5eighteightr3
	vkdscldtfourhmzzthree12ninexg
	nineninetwo9nine3
	8nineeightwof
	fivefivesixbr7xfkhrqzhdeight
	four75csrpchl
	76smqztbhqsixfourthree1zvn
	66nzdsevendpnzcqlcone
	6cnfrldrkonefive
	517837
	nineeight36
	four7ssnxm
	lrcgxfgd2nxzsrqncdm
	37
	sixnineskxb3fourthree6nine8
	9foursix
	jpnmjvl26threejjbdmptsix
	4933eightqvjsrsfzeightkfvc
	2sixtskzrsptnnonexrkv6eight5zvcsnbkp
	sqmfbdzj3q
	nine37mrksjgjlltwo
	cfgcfone6
	six92one
	4one6
	fivefivefourphdnrvvpcvrqnsxpthreeh76
	dsjhfdjkninedpkxdktf6onejddfour8
	vkckpclbthree14jmf8dxf
	4njxbptmg6four5qdmkjfpfourfive9
	35pfnrthreegkh
	two1onekdkfmz885
	txrjqbqbn27rmgxmkqbv
	8fourkoneeight9three
	mhdgskhtlzb6threefoursevenpft
	9slcxvjz8seveneightnine
	cqsnvpc7
	three72
	djsxgdfxkjvbxnnlbseven85twoseven
	fmzrnvsone8ninejk6
	97one
	2sevennine
	65znshqqpkqjktr
	sixfourfivefour9xfvrdzfgjnjhzjkzcbpfpqkzfg
	four4pzqtvgxhqxdhxfnkeight4
	4xgjrseven
	ninetwolcjfmsixone67
	9jptgxvhk19bnfcnmpgdxxxrvmrfnsg
	8fourlvvtgdjvzxxxtwonine9one
	oneeightb17foureighttwo
	eightjfcl2one4two6eightjsfmn
	7seven5
	pktvqjkkjdhlrx7zn
	onellmtcqmnjfourvnvgmshonesixthree8
	3xhvcrnlsixeight
	4vfcqdfxjplnksthreenineplvhvqhxfive2
	six8fourcbxdvnvfjzhjdzmt
	cgrktgfour8vfoursevenjtcknp
	vdnxzc53g
	hfmjzlpvbkxpnf1bkeightwolsx
	xgglvqfhp1ccxmhzeightfive9hl
	7xcgmcjfivelgj
	sixmmmfrkqvdhkg7gxfourtwonekh
	cfivekng6
	1six1mpqmlmkbngdvrtwothree
	6hglld
	5sevengjbdhffvb
	rlfmsrgndnhqnhlbp11
	mxsbhvkvbklzvdfour7fkkrnxb
	24nine
	zgxjtcflbzmxpkvthree1nttnnvdtg9
	nsmxzcnmvnineglkxcsq596qzmj7one
	twothreexgbvc7vhjzcfgrleightzgqsfq
	dnjfbpjd9
	three2qblhgxkltwo
	tjbhkqvvcnbjd91
	qhrmsix8
	jggjzfgmnd3
	1onethreenine
	twojxgpxtfv1threeeightfiveddszdds8three
	hldvlq4321szlg3
	fivefour2
	twovfsrf7seven4two
	fvjlchqlfivetwo1
	five1lkgh
	1pgtgsevensevenbjqbpsxzqf
	lgkpn8hbggmcvql13nine1lrdzone
	fourjltwoseven6drrxqssxk1
	8threepgqzlbvkfiveonehkzqlpbssix
	qptsc9
	ctrznpqeightzrsqfour63
	cqhhjlbpbxthpnfnzsevenqgjhxssvdstffive9seven
	5twoonezgvgsslkgpqlff6xmdznqx
	xhmpflgfpplpnxzgleightvdchndpqkphqqltkg7two
	sixtwofkh85clksveightwob
	58pccnnlzfpprdcbhtnrhgbjfour2
	1sixnpr
	lbxxjkhn5one7
	threeonepzltnzdninefive45two
	2dcf
	oneonezrpqdxtmmhbvlt3sixone2
	248
	3dhbfxlf4three
	sevensix4
	7glsgonegl8
	seven3tfxbpmmr
	onemgqtzzfgrgdffive6fourpvz1tdxtrtx
	five6pvdsfiveone2threeddglqxqzqdxrnjkneightwox
	2cljsix1twoqvqbzg
	6nhjfk
	mchqvjtrtlfour2
	85two32xsgjpznmdz
	xspzlcjzvvc5mq3kthf2
	threefourrjshnlzm9nineeightdzllhjhl8gdtqmb
	9cqhrtwo1five12m
	pb89hzbnxgbrrd
	jreight1five
	four18pmcrdbkcdcjcsix2fhssbzjjjvzhnjjk
	rn4sjbjtn
	vbt3kmkskgn5l
	threexsrfcqdlll9bfjk
	zfourninekjztvzphtc9eight6
	75three1six67
	eight73seven41fsxmkdqd
	threeeightlccnztnrfx2mtqvfcqtjx9lfnkmldpm
	14seven92sevensixxzrmhtchqk
	1nhjllhkglzseven6kfqfszkfgb6lhfljnspj
	onetwofive182
	vnzjcptm83one9
	bqdhljbknreightthreeeight2
	3txkkxsrxbzsixj
	2cjthreemhtd4jvzpgbkngsxbnrtgj
	one2pjpjsfpdznnxgkhdqnine1fivetwo
	qhcdv73vn
	66fourrz79seven
	five7threepxrfbrqz
	prlsxjkhm6kkhrpdltqznineseven7
	89fouroneq15
	4twofmsjsbqls25
	4smqsdrpdqqcdlxjfvrnsevenseventwo
	pnqqpctwollkbbpmfmnzlmzsz2z
	eightxcjlsbtzzbrbkmseveneight47fivelrp
	hmqlmzhvfl1lv6fourfiverfspsnvtwo
	7sevencvsxqmkthree
	273ninefour6eighthnhdzm
	four3vdcsfnineseven5
	mvpctnvzcnpdqbxbd2knlrzvfpgqdsdjztlgvf
	sixfive7
	3ninenpxqtj49four
	1lplgcsz5seven
	lvvmmvn37
	2sevenqbppzsldfourseven
	bmrs1rbfsfdhconeonesix
	3sixrrxrbbzsixsixfrhghcvkrfnggmhqx95
	4one4nine
	2fourkrpsmnkcslnszhdvjfour
	mqnmstznbnbrgzrvgdtzbxbnfrvcm2lbpprxqc
	dzlbvpj5hxmtzfzb3fdhxvtb8rhtlnhpjvtf
	qcptwone1fiveqxfrtmbpdtdtfceightstfivenine
	42seventwo
	four84mlspvcfnineeight84
	1bsix5rbdqnc2fiveseven
	mzbgh9vhksjone
	4eightfivefivetwo1oneightvhr
	one7bpqltzf9ninesix
	9sevenrc
	threethree3four64
	krpzm9vgzzcgjsxhnkckqv98lmjthbdqsixsix
	seven4six9gxrqzbmnone9
	xjzfxlzzb5rdstbqnjt
	bfpmf5seven6eightxdxsctrcd
	fivekhcfqgffntsmxxxseven655
	2dbcfmvsevenxfkxfjfkfive
	twofrgktmh9nskvhxchnine
	43slhfrqsfp9
	6fiveeight79995vkkncvsqhk
	eight6chnqqvbng9threedlsvdnqkpnxszkshtqppgdjn
	onepkgxfvgngt8six
	88gdtttxsjkpvpqbeighthpxlsdsbl6one
	eight6xkm8
	threepsmmfbdqx4
	seven5639
	six3qlgbddfn6eightnine
	2lrlzvdpklf4five8two4
	sixvhbz7
	ksfhbrfnhdqbfkgcblthreefour55
	one6958mmhvjj
	oneqmfmnjg7fmrnine
	6ninefourgvztqcztk4sevensix1
	5gsjgpcrvq
	ktss45
	kbtjtdcljbjcghxcbjlv3foureight
	fivefive6vcgxjjthreenine
	qns2fgklcgzzbczjdnvgcxsix
	74gfpsfjfour81zjxqfbb
	scvqlthree1
	398ggxg
	mpvoneightsixbdhbmhklveightthree182
	7bvjnine2fivethreefour
	52pfivenineqpshjdbb
	59threexjjbczsix5
	oneeight69dkhvvplxjfourseven
	oneeight9threevctrd4
	5threesix7
	8five7kltvbjtbsvnkhlzrtntwo
	33four8cd
	48three5fiveseven2hmb
	one5seven
	7onetwovxmtvrjxb1
	45five7fivegpzhcfbbfiveqhnhhzdqtnltgnkrxz
	kfjzhdtgkjtqthree7pvmxs
	7two5one8
	3eighteightllkbxkbs9zgznxtj8lfflcst`

	got := SumCalibrationValues(lines)
	want := 55834

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
