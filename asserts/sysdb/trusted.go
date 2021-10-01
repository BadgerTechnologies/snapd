// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016-2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package sysdb

import (
	"fmt"

	"github.com/snapcore/snapd/asserts"
	"github.com/snapcore/snapd/snapdenv"
)

const (
	encodedCanonicalAccount = `type: account
authority-id: canonical
account-id: canonical
display-name: Canonical
timestamp: 2016-04-01T00:00:00.0Z
username: canonical
validation: certified
sign-key-sha3-384: -CvQKAwRQ5h3Ffn10FILJoEZUXOv6km9FwA80-Rcj-f-6jadQ89VRswHNiEB9Lxk

AcLDXAQAAQoABgUCV7UYzwAKCRDUpVvql9g3IK7uH/4udqNOurx5WYVknzXdwekp0ovHCQJ0iBPw
TSFxEVr9faZSzb7eqJ1WicHsShf97PYS3ClRYAiluFsjRA8Y03kkSVJHjC+sIwGFubsnkmgflt6D
WEmYIl0UBmeaEDS8uY4Xvp9NsLTzNEj2kvzy/52gKaTc1ZSl5RDL9ppMav+0V9iBYpiDPBWH2rJ+
aDSD8Rkyygm0UscfAKyDKH4lrvZ0WkYyi1YVNPrjQ/AtBySh6Q4iJ3LifzKa9woIyAuJET/4/FPY
oirqHAfuvNod36yNQIyNqEc20AvTvZNH0PSsg4rq3DLjIPzv5KbJO9lhsasNJK1OdL6x8Yqrdsbk
ldZp4qkzfjV7VOMQKaadfcZPRaVVeJWOBnBiaukzkhoNlQi1sdCdkBB/AJHZF8QXw6c7vPDcfnCV
1lW7ddQ2p8IsJbT6LzpJu3GW/P4xhNgCjtCJ1AJm9a9RqLwQYgdLZwwDa9iCRtqTbRXBlfy3apps
1VjbQ3h5iCd0hNfwDBnGVm1rhLKHCD1DUdNE43oN2ZlE7XGyh0HFV6vKlpqoW3eoXCIxWu+HBY96
+LSl/jQgCkb0nxYyzEYK4Reb31D0mYw1Nji5W+MIF5E09+DYZoOT0UvR05YMwMEOeSdI/hLWg/5P
k+GDK+/KopMmpd4D1+jjtF7ZvqDpmAV98jJGB2F88RyVb4gcjmFFyTi4Kv6vzz/oLpbm0qrizC0W
HLGDN/ymGA5sHzEgEx7U540vz/q9VX60FKqL2YZr/DcyY9GKX5kCG4sNqIIHbcJneZ4frM99oVDu
7Jv+DIx/Di6D1ULXol2XjxbbJLKHFtHksR97ceaFvcZwTogC61IYUBJCvvMoqdXAWMhEXCr0QfQ5
Xbi31XW2d4/lF/zWlAkRnGTzufIXFni7+nEuOK0SQEzO3/WaRedK1SGOOtTDjB8/3OJeW96AUYK5
oTIynkYkEyHWMNCXALg+WQW6L4/YO7aUjZ97zOWIugd7Xy63aT3r/EHafqaY2nacOhLfkeKZ830b
o/ezjoZQAxbh6ce7JnXRgE9ELxjdAhBTpGjmmmN2sYrJ7zP9bOgly0BnEPXGSQfFA+NNNw1FADx1
MUY8q9DBjmVtgqY+1KGTV5X8KvQCBMODZIf/XJPHdCRAHxMd8COypcwgL2vDIIXpOFbi1J/B0GF+
eklxk9wzBA8AecBMCwCzIRHDNpD1oa2we38bVFrOug6e/VId1k1jYFJjiLyLCDmV8IMYwEllHSXp
LQAdm3xZ7t4WnxYC8YSCk9mXf3CZg59SpmnV5Q5Z6A5Pl7Nc3sj7hcsMBZEsOMPzNC9dPsBnZvjs
WpPUffJzEdhHBFhvYMuD4Vqj6ejUv9l3oTrjQWVC
`
	encodedBadgerAccount = `type: account
authority-id: canonical
account-id: 0UBOFtvmBvKkWOfwlOdWi6nYhDgxBnV3
display-name: Badger Technologies
timestamp: 2017-08-02T15:12:41.028285Z
username: badger-tech
validation: unproven
sign-key-sha3-384: BWDEoaqyr25nF5SNCvEv2v7QnM9QsfCc0PBMYD_i2NGSQ32EF2d4D0hqUel3m8ul

AcLBUgQAAQoABgUCWYHr6QAAGp8QAC42fKKbFsvNft4sDr18rh+1LcePvAezKUZc6qru+1D19Asj
7/WBBBHGK+GLg1s3RsWEwe42b+sCzmG3R5BFy6lDvXO/VaUqY/nuhG3eXAlujtpb5teXFFQHI/UE
TlhPsquA8+WDuS+J86rEbBVdvmp5gUM6DWqiPLOcdQSSdb46/+7fRpziMWmO9HXDcEa00BdGtaev
uDXQnJcR1eZu2RRucKGmk2m3CmD/RiVOEd4TXyc8K7GM4n5/yPHCZk6kMNvqrK/B5klAf7fvIQm7
fhVNvYx916DL/dOnh+sHrtYVeh72MdTHo3vrvpGUnMGG9q/mA9M3S8qT5+HTdMx4HXRTZkBa/Xsu
w8qPmz/W0ptjDyIA3MqkXgrDnMJj2o42ClGJLkuQhaUIcTq3nUXZokvlRIwQYyYAjvcVY06CPD9r
2mmqqvVX7Bxkqhl0/djL88d5zzk4xNZB5kcy6V20H+Az/jJXZv0TajWSzSYkWeHrUODDlwIjsBSi
GGC67xU6mwHupTBPu28eL4GJ24/Wasdl05yfIwTiXRk5KwO9PynpKQuq21IqKKLqXaRgcTiSClBf
Q93KY6PkYCRGui3+cd+2J5IIeEYFxZP3KckUqmxgzBxTdYzImE1BYB1/9eK+qXSRGO62FBNR3hsB
fQVYJQEt7IBndr8ul6q/5yyumzNH
`
	encodedBadgerAccountKey = `type: account-key
authority-id: canonical
public-key-sha3-384: E2EyRbU8HRPd3XhAzmnicBI2GRPpYfBo28LmhAY9lS22ItJOs2l8Doc92PWbiRtz
account-id: 0UBOFtvmBvKkWOfwlOdWi6nYhDgxBnV3
name: badger-devops
since: 2021-04-29T13:45:50Z
body-length: 717
sign-key-sha3-384: BWDEoaqyr25nF5SNCvEv2v7QnM9QsfCc0PBMYD_i2NGSQ32EF2d4D0hqUel3m8ul

AcbBTQRWhcGAARAAx9bDfoLUE3dy8ezeSvh556cCn1noL0szyxIehrPvgIB2sBcYHD7Zkqz8NPxX
zxkc7tyNzHNfVLNOQJl6btDP7YN4n5Ls2vdW3wzG0slieng5acWEo1cVIyqpnzAc4DTnAOk6VzUm
sjQLrSzmjnXj4g6jyq6b7bVMdcfUoZnmg7lRHVSWWrrSjwIFAz62K95cbf2C+XhwPWxM7Y+VAZA9
Z2ts4tfYJ80G9YKtihTzaLY0gDgY2RQg2oCWuAJDHuEkmOd7eb3qb0q0Z5MR/MddC8OsCYJDqdJA
24329TYnnYSJFzFFojP+B/nrsWNin7WTfziFv4xBlcuoM6TL4iIgCNxi/X1KMWEw78QL4swV6Kml
u58IPf+ZifMX9cIXWDv7E4I/YYKzZDaPg6ZSqvl+Iq4Sf47RJ66qqfde1UhLVpsCmn4o4BuVa7ui
C+Z6rpR/0TTtU3FWEyk6I+1FpGMIJGjMPasYl+WTnnszdRkUrhfjt9+8FVLG3Pbcl+RMtP4vDjMd
roiml72Zzknj05O0M93OlD626Iz7Uuervy46KXIdTXokpqeL+4OcDU6k2PfPe1bHZMuFYYUKYBRE
ubhVehkJMZpyvIcLA3bFpGKxpFQ4ZbY5kgzt5yhlEPLrxetWXkiE4UNgyVmJ/eCS+gTMoO+qPPwr
9CL7hmb9NTuSsgEAEQEAAQ==

AcLBUgQAAQoABgUCYIq4lAAAHJoQAHeYJ0ZFJbugAISK5V/V03y20Fy/7JyR9dyjb9A7rPGltsU1
SSXU/ld1Dqs8GHzKfOKAJElyrSJvjE1IvcguaFuSXKPGggzORLiyOO5FDUsJZ+P3ikZ7lnO+Ieqq
Q53Z9n37f8X+hsiknxpmqtBwAk5RovRT05k+W+BKULCYha++SPn3FXjN1rRU6qGP0B5diZ+eg97A
PT+gzZ7ug0+WMoyOHZvJHPPt/Y9/sGQzFOL5tn2WFN1PpIF9CcDmOqJCOfFxEU7VFVa01v2sMkwX
8GxOuyw5wyi482wVMyMwnI1+uyMn5R+bCRjEV2fj9izTdGnde1xlpEZd4Uca+FkKoUDtAXQWAX3a
ayGUKQHC5j8vAgbv1ZmlIcvunXUiO3Idam1mXFCOCUyZ5U3oBC36Nrxh5ef273xmFEesUEGRiRqD
ECWIhYu1zTUZ+aW0BxaZjacl96cyJH4fYc1vVSVItQlFg0nyg6p4zD7phkA8NzfPQ7psy9qocFUN
8usCBj7ufQIvVsl313AZBkgeJhvdxd2a9rehAg/+EWlkw90axa8dC++fZQH2F9PpclszO02zNXtZ
Op1ImOBcR4WI58Zj71PkS83FFNM9hX+U1QSoWnqOguOgdLssVQ0cIVznUtQLTUAu2wdHcoe4PxU+
i5ywh7ZfjXKbTKYfV+vHmsOH4Sch
`

	encodedCanonicalRootAccountKey = `type: account-key
authority-id: canonical
revision: 2
public-key-sha3-384: -CvQKAwRQ5h3Ffn10FILJoEZUXOv6km9FwA80-Rcj-f-6jadQ89VRswHNiEB9Lxk
account-id: canonical
name: root
since: 2016-04-01T00:00:00.0Z
body-length: 1406
sign-key-sha3-384: -CvQKAwRQ5h3Ffn10FILJoEZUXOv6km9FwA80-Rcj-f-6jadQ89VRswHNiEB9Lxk

AcbDTQRWhcGAASAA4Zdo3CVpKmTecjd3VDBiFbZTKKhcG0UV3FXxyGIe2UsdnJIks4NkVYO+qYk0
zW26Svpa5OIOJGO2NcgN9bpCYWZOufO1xTmC7jW/fEtqJpX8Kcq20+X5AarqJ5RBVnGLrlz+ZT99
aHdRZ4YQ2XUZvhbelzWTdK5+2eMSXNrFjO6WwGh9NRekE/NIBNwvULAtJ5nv1KwZaSpZ+klJrstU
EHPhs+NGGm1Aru01FFl3cWUm5Ao8i9y+pFcPoaRatgtpYU8mg9gP594lvyJqjFofXvHPwztmySqf
FVAp4gLLfLvRxbXkOfPUz8guidqvg6r4DUD+kCBjKYoT44PjK6l51MzEL2IEy6jdnFTgjHbaYML8
/5NpuPu8XiSjCpOTeNR+XKzXC2tHRU7j09Xd44vKRhPk0Hc4XsPNBWqfrcbdWmwsFhjfxFDJajOq
hzWVoiRc5opB5socbRjLf+gYtncxe99oC2FDA2FcftlFoyztho0bAzeFer1IHJIMYWxKMESjvJUE
pnMMKpIMYY0QfWEo5hXR0TaT+NxW2Z9Jqclgyw13y5iY72ZparHS66J+C7dxCEOswlw1ypNic6MM
/OzpafIQ10yAT3HeRCJQQOOSSTaold+WpWsQweYCywPcu9S+wCo6CrPzJCCIxOAnXjLYv2ykTJje
pNJ2+GZ1WH2UeJdJ5sR8fpxxRupqHuEKNRZ+2CqLmFC5kHNszoGolLEvGcK4BJciO4KihnKtxrdX
dUJIOPBLktA8XiiHSOmLzs2CFjcvlDuPSpe64HIL5yCxO1/GRux4A1Kht1+DqTrL7DjyIW+vIPro
A1PQwkcAJyScNRxT4bPpUj8geAXWd3n212W+7QVHuQEFezvXC5GbMyR+Xj47FOFcFcSZID1hTZEu
uMD+AxaBHQKwPfBx1arVKE1OhkuKHeSFtZRP8K8l3qj5W0sIxxIW19W8aziu8ZeDMT+nIEJrJvhx
zGEdxwCrp3k2/93oDV7g+nb1ZGfIhtmcrKziijghzPLaYaiM9LggqwTARelk3xSzd8+uk3LPXuVl
fP8/xHApss6sCE3xk4+F3OGbL7HbGuCnoulf795XKLRTy+xU/78piOMNJJQu+G0lMZIO3cZrP6io
MYDa+jDZw4V4fBRWce/FA3Ot1eIDxCq5v+vfKw+HfUlWcjm6VUQIFZYbK+Lzj6mpXn81BugG3d+M
0WNFObXIrUbhnKcYkus3TSJ9M1oMEIMp0WfFGAVTd61u36fdi2e+/xbLN0kbYcFRZwd9CmtEeDZ0
eYx/pvKKaNz/DfUr0piVCRwxuxQ0kVppklHPO4sOTFZUId8KLHg28LbszvupSsHP/nHlW8l5/VK6
4+KxRV2XofsUnwARAQAB

AcLDXAQAAQoABgUCV83kkgAKCRDUpVvql9g3IA9hIADAkn4VXnJIFblhMSBe6hbTy7z6AfOhZxXR
Ds/mHsiWfFT6ifGi9SpZowhRX+ff57YvFCjlBqMYLKYE0NsFQYEUc5uBWiFZwC0ENydNhO23DV1B
elTSs6mr9duPm1eJAozFrQETOD1kz5BIamqBUeaTczjM+9l5i485Ffknbc+EaGOrtMEap0GqjByQ
u+ykZGvryVQ447avgjvFsMtA0quFi+SoW9PT/9D26e5rD7RIICYWG8mzFRn5Isqs/X4W1uAiKQe9
pqHMbdNr/FCWX5ws0/nMaOq+b0z4EIIXIfT0JmIlFDQsAgFVnKwYw+zs32cTw4XuzvMhgMDtCowD
YodhiO/5AOMsMMV0qBsYxbIPJIEz7b6gwTYEJoTVkqTit6o3UgWrAy+p4Y7t0ickYIHgwiuKRS9E
fu0Ue+32NFp0XFqZElfXLK/U2yjto+fJXu6uAELsXesfFGIOp/nbRbNavUt9jAJeO7ftQczgf39T
YfA0OKerP5gAOd4+aO3gATPUjfWPsJ9908XC7QqK2BwS1kh/fMrd95mxcmXdF1bBElszKwaToBVQ
1m52EYp06kkPyOu+fGKFAoIMafcV/2Ztz1WMo/Vp0iP/r0WAtBDw6sDJyWOfRjUEvP7BBdEzraHV
VblbSrKzhYeEGdMDi6kFC+KEzfPDPFJX1l3saPBkz9VDuESbktyObQp9VfkFKYBgBnw3msQJk+6k
G4t0o3/DZ7qz/kTJXMogG26Z/FsMhPERsaLTbWRJ3WRyXX8COaTladSf8bG0Oib19outnjuvpjQ0
qEV9eeGRBlx9mbidSYH95cj0zD2DKpeSZ83M5K1pFg+8RKToGElGTTk8vtdTfDVbmi3+QntfLq+z
ZMgs2+SmCWrV/MPC04Dl00CXywdKPyf6toomqRP7A5fS7W8P9fdPn+a8JCblcleGj9nvJXBQjue7
97rofCEszhKhoE9fMCIUcSoTU9YAm5Jr+qclSEbV1pzwTvZ8auMIXtzEZV5n4aK4WPDV+lYCadrL
DlvJSJRuXRvIMbmvU9b8NxgG8AS88BkX3L9vlOpkMculwG1/iooQvxuFaJDargt370wAQo0lCpG3
MxnsSusymwnYegvvvr7Xp/KBLZK1+8Djzm3fwAryp4qNo29ciVw3O9lFKmmuiIcxSY0bauXaK6kv
pTnYkmx7XGPF7Ahb7Ov0/0FE2Lx3JZXSEKeW+VrCcpYQOY++t67b+jf0AV4rZExcLFJzP6MPMimP
ZCd383NzlzkXK+vAdvTi40HPiM9FYOp6g8JTs5TTdx2/qs/SWFC8AkahIQmH0IpFBJep2JKl2kyr
FZMvASkHA9bR/UuXDvbMzsUmT/xnERZosQaZgFEO
`
)

var (
	trustedAssertions        []asserts.Assertion
	trustedStagingAssertions []asserts.Assertion
	trustedExtraAssertions   []asserts.Assertion
)

func init() {
	canonicalAccount, err := asserts.Decode([]byte(encodedCanonicalAccount))
	if err != nil {
		panic(fmt.Sprintf("cannot decode trusted assertion: %v", err))
	}
	canonicalRootAccountKey, err := asserts.Decode([]byte(encodedCanonicalRootAccountKey))
	if err != nil {
		panic(fmt.Sprintf("cannot decode trusted assertion: %v", err))
	}
	badgerAccount, err := asserts.Decode([]byte(encodedBadgerAccount))
	if err != nil {
		panic(fmt.Sprintf("cannot decode trusted assertion: %v", err))
	}
	badgerAccountKey, err := asserts.Decode([]byte(encodedBadgerAccountKey))
	if err != nil {
		panic(fmt.Sprintf("cannot decode trusted assertion: %v", err))
	}
	trustedAssertions = []asserts.Assertion{canonicalAccount, canonicalRootAccountKey, badgerAccount, badgerAccountKey}
}

// Trusted returns a copy of the current set of trusted assertions as used by Open.
func Trusted() []asserts.Assertion {
	trusted := []asserts.Assertion(nil)
	if !snapdenv.UseStagingStore() {
		trusted = append(trusted, trustedAssertions...)
	} else {
		if len(trustedStagingAssertions) == 0 {
			panic("cannot work with the staging store without a testing build with compiled-in staging keys")
		}
		trusted = append(trusted, trustedStagingAssertions...)
	}
	trusted = append(trusted, trustedExtraAssertions...)
	return trusted
}

// InjectTrusted injects further assertions into the trusted set for Open.
// Returns a restore function to reinstate the previous set. Useful
// for tests or called globally without worrying about restoring.
func InjectTrusted(extra []asserts.Assertion) (restore func()) {
	prev := trustedExtraAssertions
	trustedExtraAssertions = make([]asserts.Assertion, len(prev)+len(extra))
	copy(trustedExtraAssertions, prev)
	copy(trustedExtraAssertions[len(prev):], extra)
	return func() {
		trustedExtraAssertions = prev
	}
}
