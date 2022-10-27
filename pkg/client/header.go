package client

import (
	"net/http"
	"net/textproto"
)

var (
	DefHeader = http.Header{
		textproto.CanonicalMIMEHeaderKey(`accept`):             []string{`application/json`, `text/plain`, `*/*`},
		textproto.CanonicalMIMEHeaderKey(`Content-Type`):       []string{`application/json;charset=UTF-8`},
		textproto.CanonicalMIMEHeaderKey(`accept-encoding`):    []string{`gzip`, `deflate`, `br`},
		textproto.CanonicalMIMEHeaderKey(`accept-language`):    []string{`zh`, `zh-CN;q=0.9`, `en-US;q=0.8`, `en;q=0.7`},
		textproto.CanonicalMIMEHeaderKey(`cache-control`):      []string{`no-cache`},
		textproto.CanonicalMIMEHeaderKey(`cookie`):             []string{`fingerprint=de73a8e14afc457e972e97e4eb29f25c83; low_login_enable=1; CheckKey=fced66fc601fc0cd85a24b9f19rdn1; RK=2cPZdOphff; ptcz=3d8b354d9d71f4a2e947e18159938a30ba5784074f34231f1d446db2a50ba312; luin=o1016017516; lskey=00010000d6843fa41f96eb01c724ced4accce85c46fc6ed25a45f19c3199d2c398cfcdff21dded00a95333a1; p_luin=o1016017516; p_lskey=00040000ded573d6c5df94ef6c09037878a071d0c45a4d781bef9c4916178b26ae8c57aca6709e76493a8330; has_been_login=1; DOC_SID=8d18bc0cb6fe40b7834ecfe9b3a9f4214df51d2cf92048e693af24c4d28a89cf; SID=8d18bc0cb6fe40b7834ecfe9b3a9f4214df51d2cf92048e693af24c4d28a89cf; loginTime=1649672318331; backup_cdn_domain=docs.gtimg.com; optimal_cdn_domain=docs2.gtimg.com; traceid=927d688e40; TOK=927d688e40d8b727; hashkey=927d688e`},
		textproto.CanonicalMIMEHeaderKey(`pragma`):             []string{`no-cache`},
		textproto.CanonicalMIMEHeaderKey(`referer`):            []string{`https://docs.qq.com/sheet/DQkRORVFKc2dTeXlB?u=161031ceb2e74df8a0b51bda76bde179&tab=BB08J2&_t=1649672315303`},
		textproto.CanonicalMIMEHeaderKey(`sec-ch-ua`):          []string{`" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`},
		textproto.CanonicalMIMEHeaderKey(`sec-ch-ua-mobile`):   []string{`?0`},
		textproto.CanonicalMIMEHeaderKey(`sec-ch-ua-platform`): []string{`"macOS"`},
		textproto.CanonicalMIMEHeaderKey(`sec-fetch-dest`):     []string{`empty`},
		textproto.CanonicalMIMEHeaderKey(`sec-fetch-mode`):     []string{`cors`},
		textproto.CanonicalMIMEHeaderKey(`sec-fetch-site`):     []string{`same-origin`},
		textproto.CanonicalMIMEHeaderKey(`traceparent`):        []string{`00-b62ed7fb49dc5482684c90c1a4a45407-4bcc539a429b6960-01`},
		textproto.CanonicalMIMEHeaderKey(`user-agent`):         []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36`},
	}
)
