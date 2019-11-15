// AUTOGENERATED CODE. DO NOT EDIT.

package gendoc

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var embeddedResources = map[string]string{
	"docbook.tmpl": "H4sIAAAAAAAA/+xZ30/bOhR+719h5fEiEq64SFeTW6TBqmkChIDt3U1OW2uOndlOAUX93yfHIb9/dLQUtvGCiM/nc+zj833HSfHpQ8jQCqSigo+df90jBwH3RUD5Yux8vZse/u+cTkaYSE19BpMRQlhTzWByLYUWvmDoXPhxCFwTTQXHnrWOEEoSSfgCkDulDNR6baYq8A3KmHNHSeJekRDW69JcMzsikiD3HJQvaWRmpS5Kfi9BKbLIXBfOEQ3GTpK405gx69ixLssRLwRftETti2tsdI7cz0RNKbBA5eNYkxkDNJckhLFDGMsD5iGxz4hSnISN6IUBWbe1BRkXCyniCPmCqbHzX8k5QtgMRuAb4z0N9HLs/ON4WyOO3JNh0HEdopdAgvIIQliK++oIQhi4lo+TdLfYsw/tkLvHCPoRF2QGrB9SOslWIPZqa8ReYyNYz0RQm1eq70o1DG68VPB968aM8u/I/AFeVLRJianorIrsI/YMbNLvz8ww2RqK28oAW/rnMCcx098Ii01Ug5tkYx9QktTtXgpIEuBBR9BG7k1aU3j1PKrZx55lRE5rLyVgQeGyh5y0nx40cCNzuyfuFSgNASoiDHD4ZB8cfhssz3OyLdM/EjWAuIrDGchXFoOWKhvM0SsJQtPfmeCaUE75oua5MPy66Nhj+atUB3uVi07ZVBQKj8P9XV12pXRpkofk7Xgf8ralLpm9/QZyYvO9cynZjpb7IlUHjbKHrrZefcnIy9u8hRwyWAHr7tOY8rmQIWG9bHnv5O+d/L2T/1GdvML7TcQnq5FbkCvqP+8TxMv38Jb+fQl6Kd7GN4YXFyy7VzTc6W/gRwxKo2HpugEVCa5gA+iL61N2lHsQpyw/NSXJRhsyYqmeWW+1BBJSvlivkUr/76Nz7xps5huLsMOdq7DmZy7jLd5+aoaSLrV+Xb31CSPS3tvTqq1yf/jG033fGeJ0i/1kCLBbe5NQjVPKTtqNpNCim9ZPFxOhTQK7AWcHB4NOvpAVGQRdP+ql4F2wWrk1pKMpHEXLSguiKhtdaXnqY+m3/hLrSs99ezCCYxI2iDqLoo28mcxtBLTZ64Q22Frnao2pVZ62XBdKlBxhL/+x5GcAAAD//x4vKjleGQAA",
	"html.tmpl": "H4sIAAAAAAAA/+xa32/juBF+918xp22Rdi+ybOfHpl5ZBZofKIpLurhki+tTQUu0RSxFqiKVS2r4fy9ISrIoU06cdYo+XPKwEof6+M3Mx+Ew2PCHq79fPvzzyzWkMqPRYBCaf8MUoyQaAACEkkiKoy8FlzzmFK54XGaYSSQJZ2FgrGZmhiWCOEWFwHLmfX248S+8ykQJ+wYFpjNPyGeKRYqx9EA+53jmSfwkg1gID9ICL2ZeKmU+DYIFZ1IMl5wvKUY5EcOYZ2ranxcoI/R59nVeMllOT0ej40+j0fHpaEQkoiT2gmpNvZJ5Vj9znjzDqnlVP7+SRKZTOB/h7LNlyFCxJGwKY5wBKiW3rTGnvJjCh8lkYhsUZd/Qm4JnCHrHIBATvsAFWdjTc5QkhC39OZeSZ1M4bdNYD5rHdNzhrdf5FZNlKqfAeJEhaiPPeZHgogEe508gOCUJfEAI7SYxGp7hJzeNSYfG4VaxYj48wxmM3BRO/i8igToslIL9BMe80LtCMWHYrZmz8094cuZElWhOsVuh49Ho9w6pCfIfPIWLrq3yOeaUolzgKdRP7mXVXt8V1k+jUQcfxd+WBS9Z4tduJbH6dePrjSeLKZOpH6eEJn/Aj5j9sSumbdDFXP26QbuMjc9WguM4dia4yixMerIrE8hdCSYswUzqouBWrluzCq7l+7jruIU9+gzBR7jjYAaAM1iQQkjIgTAF9TFwrRF8hAetHr6ABcE0Ee2JQz3kG33JpENGfXyjJrQ/gbb27LL0MuqkQn14zvHBQE8q0J/QHNMe1PN9QU8r0Css4oLkavP2QNunQE/o8ZPETBDO7PA3w7tScF1P2jdiO9G/JxU7get0/AWJwwLXKbkrszkueqDP3oJ8duBkszKDR0RLLIZ2ulmZ7cr0Hcr2D1kP5uQ10dob9eSwkRIxoqgwsdI9Xydgxu5ru6/tNamiUyvT6kg66emU2qvGnEmsWsj2Wh8kj31lQYThAkraWYASIX3dOGoirpO8bhMoXrgOA0oY9mue462z2XlWtNlBBJRAtNVjWEfynNOkLwA3hGJQZzphS0jIY+ckoIqdMb6y1UiIyCl6npoG5c3tVe35qeruXF2em2RP1+nKjE3UjzGlL+M7ezlEyZJNoVDR3msNS4kphqPbo2M4uj4CxBI4+uUI5ihZYqEP8xTDA7+00qOtPXkZnnel1qhs29TQJEwLck55/O3zYIdStzHakYgxk7hwtKBuRYLdoZ4rSTnb4Ys/zdHpxcvt5WIxii86GM1W0t2cuqqZJ39rTzqaw25vuZnmFyghpVBb+6krnzCoLpXm7Qffh68CFxCXQvIMLu/vwfffcOXdzBiq0SAahIG5gg9C1T7rm/kYSDLz9IXb672Pp2N1e59ETRm8rMpgGKSTaBCqgqBh2nWwvqWXtLZ5m3vzalUgtsQwVIVFrNctw+/UTvoXU+fZdAZDdbC17CElkRXfEFXuflitqsle1DyGAepML6k9YLG5xUKgpUWod2EngZuS0ppEKHLEIKZIiJmnN6EX3YaBGlUEf+Js2UPSqMK14GqFWeJg13hwzcrs/ehfvzv9ptl9qw8b+azX/qZ3dvvzS+WPUqFP8SOmm5ZZHNKve1w8kvgdhXW/ycxBshIG7Y1if9edrzzYkN1uvLzo3jRr/9DNmro26OBuUM1qYZCQx2gw2CoOfUVBF54qDu3zty496USXHndhSCct0lUBfOC5jlvFZLUiGVJEWj3ret0i2KoXYXpSL9bOVWenpCe7UMkChn9FQl+cNabpWBsPmwtn7aBs/0m1iEKZRPrjMJCJflOxbl70Lbd5a61uxgJZ1NmwcM1ZAXbVrhjWK2+0IJN2kGVim7Y0rQi24mRejTi636opygMnbq7MVkDBxPMKL1BJpZbeel29TUHPblsqUYdB3lp6E5O25sOgikkY6GyolBrrJodWIQvnRdRJZudSuSOhDZI7qequ3LyYG9tBUrxVif9Xaba+vTTdBGHLDsrG8ArZmLC8Ujcu2cA766aln845vk9ZybslpSO55mK+Q20K8T3kpKP1Fintm7732cBNFnrbi54WQmXmt53/287fIbjtFrF/0x9+w99imfIErH3/M/53iYUES24/Y5FzJrA9+j1CM0t/t8oqth2ZVKN7aa52cQvKDL+Idehq1Iik7mRf21vrW7mlhb4/h+5QxjAvuOR2uu+4VPjV2+WPP9rmv6FHZI98eZYpZ62xPfRhPGvpo94U+u8UdYK6mmkZ3SVBudBju8zzHV8q73aYjavdCa/NeRjUA/p/S/w3AAD//7koENtFIQAA",
	"markdown.tmpl": "H4sIAAAAAAAA/+RWz2+jOhC+81fMC+/wqor0XiU5vPZVT6u2qtpqL9Vq6yQTgmRsFptoK/D/vvIPsAnQVNrVXpYLnhk8zHzf54EYHkou+YZTuOabKkcmicw4ixYEGMlxOZO8mK0WF2QVRXEMz2RNEfgOrjiTyKSI6rokLEWY32QUhVJRXf+9yyh+1dvhcgnze5KjUgm81LVbf/kn7tZnEUCX4w6FIKlOA2A33HKWhptuKkrDjci2SoUp/mNV/lP7v0tkIuOsTaL7SigekIKPmWS+T6US7GITiZ+wPGSbsLfTdbX3BF6eNoSSEj4TWiE8vxWoaxDGmRy0M5HaeRZ9mBHPcVeLY3pRAKFZypazMkv3crZaENiXuFvOYiOHZ17o5xYXhVVFtz+q6/k1ik2ZFVpFSgXVeG57L/aNe5mZjJ650azZDub/E3GTId3qnA2YJTQGHGjglqyRQgPBTmiiBhJ9gb1D33QXNCGGOj8khsvGdwpNJy/9vpBGa5/Zp00V5vG6ZnxdQr8T28c17khFpeFWKXDmJZi+w5DTgymwrxCPSCjgqPGa9cj8S4S+3Vf5GssphIYodSBNo+XfPYpYDzBr6zlCMpax9Dhiy/s90NnQ4q8kAWRbyJ1UIUlWgYLdaPkF8m1Ah05x8AHATXOTYJ+CMABjCALqblsExrQVwDAxCUNIRsfon6vP14FAXycV6jnpUXCkTv99+ZhA3xHnHco937YafcRvFQrZsvOIouBMYGtPsnNMxLF5bDfht0IXMD12XUnH09e5gyFsSXD+J1kiyTOWKgXCrDvIXVbb2TCt9Y/ktYH3Ep86eAOOheOxpdedzziG4S+AJmte6D+4lo17LlFAA1fn563rEzmQdv3wJvecOWucrIAVGJ61kUNly/Jkheozf5cWsxlcrKDvcszqkrvzVhRhTNce2rb+1tOh8yMAAP//me02bccKAAA=",
	"scalars.json": "H4sIAAAAAAAA/9yXzW4aMRDH7zzFiFMqBZDSlEa9JZGQOOQEOUWp5GVnvW6NTewxzaqq1HfoG/ZJqt0F1gYvpChVk9zQfHg9v/nPWNx1AL53AAC6C6NJT4sFdj9BN9Uukdg9rV1KE9rSvDbMFot45MzG7XxzciY1o+H52vGFLVk8ZZEvgpyNvaBcq6jLuKRYO0aVowPw47SlxiB1X4lBYFNhYN4q8P1ZrMCwjn9dn1DeNTb13Vq0sGRGsERiT6LilAOqmU6F4n0YK8wyMROoCDJtNh5QyBmJJYJy8wSNhd8/f4HIoNDOQCZQpiAsSPEVZQGkIWdl7DppyaRDewrOItjqYiCUJWRpPwI8uHkDXKgI7iDWh+1Fe6iFIuRo4rC9FB/1leDKzUEbGInH8tcJs2DwwQmD6btDPWi0/sJ6MDw/0IPm5k0PpFY82oT4SPvhu10YWDLCC9huxsBP3+3IXvLuGPlHOLgWMbqoGt0zy/EAgaM06Y4RZQuZmERcXCPu/4jkKET2qN05EVxhCkJRPWt9mOZoEebaIGxGWhZ1Cu6OM+VMgUHuJDNQXcG+7fVoj9qPz855eN7G+VWvwEw8YhqR8aX8xgoLWflqJAWh7cNNgK4G5FbvdLZ6OYAZBJ0RKuAGGaGp484+n128vM35zEqtWEakumKJgud0CGb54B6G+WH4epft/mF/uh7f9tL7Cy29wa2UaC2f8q/Lj2vK9K08eqZfZWlHpiKFbnuCCrecfoFT4/BaMmsHIyZt/XN/twOSTa+hdsDcWQJWd36mFTGh4HY66l2sXq+01NjHXiIILifX4zEQPlJMF+GHGmKhnbdczKc2CZuft3wiZGbJDJwS5ZVj3Ooz4aQqbf98VMrfAXbDig0fpgpgJhFkmCnA4oNDNSvXafvUtNG5KggnLYTu7svjYoR2s55OaR+dqsO9i6vxtEbUue/8CQAA//8z+wC/ohEAAA==",
}

func fetchResource(name string) ([]byte, error) {
	raw, ok := embeddedResources[name]
	if !ok {
		return nil, fmt.Errorf("Could not find resource for '%s'", name)
	}

	compressed, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	buf := bytes.NewBuffer(compressed)
	
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
