module github.com/projectdiscovery/nuclei/v2

go 1.17

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/alecthomas/jsonschema v0.0.0-20211022214203-8b29eab41725
	github.com/andygrunwald/go-jira v1.15.1
	github.com/antchfx/htmlquery v1.2.4
	github.com/apex/log v1.9.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/bluele/gcache v0.0.2
	github.com/corpix/uarand v0.1.1
	github.com/go-playground/validator/v10 v10.10.1
	github.com/go-rod/rod v0.104.4
	github.com/gobwas/ws v1.1.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/itchyny/gojq v0.12.7
	github.com/json-iterator/go v1.1.12
	github.com/julienschmidt/httprouter v1.3.0
	github.com/karlseguin/ccache v2.0.3+incompatible
	github.com/karrick/godirwalk v1.16.1
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/miekg/dns v1.1.47
	github.com/olekukonko/tablewriter v0.0.5
	github.com/owenrumney/go-sarif v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/projectdiscovery/clistats v0.0.8
	github.com/projectdiscovery/cryptoutil v1.0.0
	github.com/projectdiscovery/fastdialer v0.0.15-0.20220127193345-f06b0fd54d47
	github.com/projectdiscovery/filekv v0.0.0-20210915124239-3467ef45dd08
	github.com/projectdiscovery/fileutil v0.0.0-20210928100737-cab279c5d4b5
	github.com/projectdiscovery/goflags v0.0.8-0.20220223132346-fcffa47ead36
	github.com/projectdiscovery/gologger v1.1.4
	github.com/projectdiscovery/hmap v0.0.2-0.20210917080408-0fd7bd286bfa
	github.com/projectdiscovery/interactsh v1.0.2-0.20220309105842-7aa59db95941
	github.com/projectdiscovery/nuclei-updatecheck-api v0.0.0-20211006155443-c0a8d610a4df
	github.com/projectdiscovery/rawhttp v0.0.8-0.20220321180300-366b511e8bfd
	github.com/projectdiscovery/retryabledns v1.0.13-0.20211109182249-43d38df59660
	github.com/projectdiscovery/retryablehttp-go v1.0.2
	github.com/projectdiscovery/stringsutil v0.0.0-20220208075244-7c05502ca8e9
	github.com/projectdiscovery/yamldoc-go v1.0.3-0.20211126104922-00d2c6bb43b6
	github.com/remeh/sizedwaitgroup v1.0.0
	github.com/rs/xid v1.4.0
	github.com/segmentio/ksuid v1.0.4
	github.com/shirou/gopsutil/v3 v3.22.2
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cast v1.4.1
	github.com/syndtr/goleveldb v1.0.0
	github.com/tj/go-update v2.2.5-0.20200519121640-62b4b798fd68+incompatible
	github.com/valyala/fasttemplate v1.2.1
	github.com/weppos/publicsuffix-go v0.15.1-0.20210928183822-5ee35905bd95
	github.com/xanzy/go-gitlab v0.77.0
	github.com/ysmood/gson v0.7.0 // indirect
	github.com/ysmood/leakless v0.7.0 // indirect
	go.uber.org/atomic v1.9.0
	go.uber.org/multierr v1.8.0
	go.uber.org/ratelimit v0.2.0
	golang.org/x/net v0.0.0-20220805013720-a33c5aa5df48
	golang.org/x/oauth2 v0.0.0-20220722155238-128564f6959c
	golang.org/x/text v0.3.7
	gopkg.in/yaml.v2 v2.4.0
	moul.io/http2curl v1.0.0
)

require github.com/aws/aws-sdk-go v1.43.29

require github.com/projectdiscovery/folderutil v0.0.0-20211206150108-b4e7ea80f36e

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/docker/go-units v0.4.0
	github.com/h2non/filetype v1.1.3
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/openrdap/rdap v0.9.1-0.20191017185644-af93e7ef17b7
	github.com/projectdiscovery/iputil v0.0.0-20210804143329-3a30fcde43f3
	github.com/projectdiscovery/nvd v1.0.9-0.20220314070650-d4a214c1f87d
	github.com/projectdiscovery/sliceutil v0.0.0-20220225084130-8392ac12fa6d
	github.com/projectdiscovery/wappalyzergo v0.0.34
	github.com/stretchr/testify v1.8.0
	github.com/zmap/zcrypto v0.0.0-20211005224000-2d0ffdec8a9b
)

require (
	git.mills.io/prologic/smtpd v0.0.0-20210710122116-a525b76c287a // indirect
	github.com/Mzack9999/ldapserver v1.0.2-0.20211229000134-b44a0d6ad0dd // indirect
	github.com/PuerkitoBio/goquery v1.6.0 // indirect
	github.com/akrylysov/pogreb v0.10.1 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/andybalholm/cascadia v1.1.0 // indirect
	github.com/antchfx/xpath v1.2.0 // indirect
	github.com/bits-and-blooms/bitset v1.2.0 // indirect
	github.com/bits-and-blooms/bloom/v3 v3.0.1 // indirect
	github.com/c4milo/unpackit v0.1.0 // indirect
	github.com/caddyserver/certmagic v0.15.3 // indirect
	github.com/cnf/structhash v0.0.0-20201127153200-e1b16c1ebc08 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/frankban/quicktest v1.14.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/goburrow/cache v0.1.4 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gosuri/uilive v0.0.4 // indirect
	github.com/gosuri/uiprogress v0.0.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0 // indirect
	github.com/itchyny/timefmt-go v0.1.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/klauspost/pgzip v1.2.5 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/lor00x/goldap v0.0.0-20180618054307-a546dffdd1a3 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mholt/acmez v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nwaples/rardecode v1.1.2 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/projectdiscovery/blackrock v0.0.0-20210415162320-b38689ae3a2e // indirect
	github.com/projectdiscovery/mapcidr v0.0.8 // indirect
	github.com/projectdiscovery/networkpolicy v0.0.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/trivago/tgo v1.0.7 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/ulule/deepcopier v0.0.0-20200430083143-45decc6639b6 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/yl2chen/cidranger v1.0.2 // indirect
	github.com/ysmood/goob v0.3.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	github.com/zmap/rc2 v0.0.0-20131011165748-24b9757f5521 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.uber.org/zap v1.21.0 // indirect
	goftp.io/server/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
	golang.org/x/tools v0.1.6-0.20210726203631-07bc1bf47fb2 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/corvus-ch/zbase32.v1 v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
