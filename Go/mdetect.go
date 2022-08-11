package mdetect

/* *******************************************
// Copyright 2010-2015, Anthony Hand
//
//
// File version 2015.05.13 (May 13, 2015)
// Updates:
//	- Moved MobileESP to GitHub. https://github.com/ahand/mobileesp
//	- Opera Mobile/Mini browser has the same UA string on multiple platforms and doesn't differentiate phone vs. tablet.
//		- Removed DetectOperaAndroidPhone(). This method is no longer reliable.
//		- Removed DetectOperaAndroidTablet(). This method is no longer reliable.
//	- Added support for Windows Phone 10: variable and DetectWindowsPhone10()
//	- Updated DetectWindowsPhone() to include WP10.
//	- Added support for Firefox OS.
//		- A variable plus DetectFirefoxOS(), DetectFirefoxOSPhone(), DetectFirefoxOSTablet()
//		- NOTE: Firefox doesn't add UA tokens to definitively identify Firefox OS vs. their browsers on other mobile platforms.
//	- Added support for Sailfish OS. Not enough info to add a tablet detection method at this time.
//		- A variable plus DetectSailfish(), DetectSailfishPhone()
//	- Added support for Ubuntu Mobile OS.
//		- DetectUbuntu(), DetectUbuntuPhone(), DetectUbuntuTablet()
//	- Added support for 2 smart TV OSes. They lack browsers but do have WebViews for use by HTML apps.
//		- One variable for Samsung Tizen TVs, plus DetectTizenTV()
//		- One variable for LG WebOS TVs, plus DetectWebOSTV()
//	- Updated DetectTizen(). Now tests for “mobile” to disambiguate from Samsung Smart TVs
//	- Removed variables for obsolete devices: deviceHtcFlyer, deviceXoom.
//	- Updated DetectAndroid(). No longer has a special test case for the HTC Flyer tablet.
//	- Updated DetectAndroidPhone().
//		- Updated internal detection code for Android.
//		- No longer has a special test case for the HTC Flyer tablet.
//		- Checks against DetectOperaMobile() on Android and reports here if relevant.
//	- Updated DetectAndroidTablet().
//		- No longer has a special test case for the HTC Flyer tablet.
//		- Checks against DetectOperaMobile() on Android to exclude it from here.
//	- DetectMeego(): Changed definition for this method. Now detects any Meego OS device, not just phones.
//	- DetectMeegoPhone(): NEW. For Meego phones. Ought to detect Opera browsers on Meego, as well.
//	- DetectTierIphone(): Added support for phones running Sailfish, Ubuntu and Firefox Mobile.
//	- DetectTierTablet(): Added support for tablets running Ubuntu and Firefox Mobile.
//	- DetectSmartphone(): Added support for Meego phones.
//	- Reorganized DetectMobileQuick(). Moved the following to DetectMobileLong():
//		- DetectDangerHiptop(), DetectMaemoTablet(), DetectSonyMylo(), DetectArchos()
//	- Removed the variable for Obigo, an embedded browser. The browser is on old devices.
//		- Couldn’t find info on current Obigo embedded browser user agent strings.
//
//
//
// LICENSE INFORMATION
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//        http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.
//
//
// ABOUT THIS PROJECT
//   Project Owner: Anthony Hand
//   Email: anthony.hand@gmail.com
//   Web Site: http://www.mobileesp.com
//   Source Files: https://github.com/ahand/mobileesp
//
//   Versions of this code are available for:
//      PHP, JavaScript, Java, ASP.NET (C#), Ruby, and Go
//
// *******************************************
*/

//**************************
// The UAgentInfo class encapsulates information about
//   a browser's connection to your web site.
//   You can use it to find out whether the browser asking for
//   your site's content is probably running on a mobile device.
//   The methods were written so you can be as granular as you want.
//   For example, enquiring whether it's as specific as an iPod Touch or
//   as general as a smartphone class device.
//   The object's methods return 1 for true, or 0 for false.

import (
	"net/http"
	"strings"
)

//standardized values for true and false.
const true = 1
const false = 0

//Initialize some initial smartphone string variables.
const engineWebKit = "webkit"
const deviceIphone = "iphone"
const deviceIpod = "ipod"
const deviceIpad = "ipad"
const deviceMacPpc = "macintosh" //Used for disambiguation

const deviceAndroid = "android"
const deviceGoogleTV = "googletv"

const deviceWinPhone7 = "windows phone os 7"
const deviceWinPhone8 = "windows phone 8"
const deviceWinPhone10 = "windows phone 10"
const deviceWinMob = "windows ce"
const deviceWindows = "windows"
const deviceIeMob = "iemobile"
const devicePpc = "ppc"     //Stands for PocketPC
const enginePie = "wm5 pie" //An old Windows Mobile

const deviceBB = "blackberry"
const deviceBB10 = "bb10"                   //For the new BB 10 OS
const vndRIM = "vnd.rim"                    //Detectable when BB devices emulate IE or Firefox
const deviceBBStorm = "blackberry95"        //Storm 1 and 2
const deviceBBBold = "blackberry97"         //Bold 97x0 (non-touch)
const deviceBBBoldTouch = "blackberry 99"   //Bold 99x0 (touchscreen)
const deviceBBTour = "blackberry96"         //Tour
const deviceBBCurve = "blackberry89"        //Curve2
const deviceBBCurveTouch = "blackberry 938" //Curve Touch
const deviceBBTorch = "blackberry 98"       //Torch
const deviceBBPlaybook = "playbook"         //PlayBook tablet

const deviceSymbian = "symbian"
const deviceS60 = "series60"
const deviceS70 = "series70"
const deviceS80 = "series80"
const deviceS90 = "series90"

const devicePalm = "palm"
const deviceWebOS = "webos"   //For Palm devices
const deviceWebOStv = "web0s" //For LG TVs
const deviceWebOShp = "hpwos" //For HP"s line of WebOS devices

const deviceNuvifone = "nuvifone" //Garmin Nuvifone
const deviceBada = "bada"         //Samsung"s Bada OS
const deviceTizen = "tizen"       //Tizen OS
const deviceMeego = "meego"       //Meego OS
const deviceSailfish = "sailfish" //Sailfish OS
const deviceUbuntu = "ubuntu"     //Ubuntu Mobile OS

const deviceKindle = "kindle"         //Amazon Kindle, eInk one
const engineSilk = "silk-accelerated" //Amazon"s accelerated Silk browser for Kindle Fire

const engineBlazer = "blazer" //Old Palm browser
const engineXiino = "xiino"   //Another old Palm

//Initialize variables for mobile-specific content.
const vndwap = "vnd.wap"
const wml = "wml"

//Initialize variables for other random devices and mobile browsers.
const deviceTablet = "tablet" //Generic term for slate and tablet devices
const deviceBrew = "brew"
const deviceDanger = "danger"
const deviceHiptop = "hiptop"
const devicePlaystation = "playstation"
const devicePlaystationVita = "vita"
const deviceNintendoDs = "nitro"
const deviceNintendo = "nintendo"
const deviceWii = "wii"
const deviceXbox = "xbox"
const deviceArchos = "archos"

const engineFirefox = "firefox"      //For Firefox OS
const engineOpera = "opera"          //Popular browser
const engineNetfront = "netfront"    //Common embedded OS browser
const engineUpBrowser = "up.browser" //common on some phones
const engineOpenWeb = "openweb"      //Transcoding by OpenWave server
const deviceMidp = "midp"            //a mobile Java technology
const uplink = "up.link"
const engineTelecaQ = "teleca q" //a modern feature phone browser

const devicePda = "pda" //some devices report themselves as PDAs
const mini = "mini"     //Some mobile browsers put "mini" in their names.
const mobile = "mobile" //Some mobile browsers put "mobile" in their user agent strings.
const mobi = "mobi"     //Some mobile browsers put "mobi" in their user agent strings.

//Smart TV strings
const smartTV1 = "smart-tv" //Samsung Tizen smart TVs
const smartTV2 = "smarttv"  //LG WebOS smart TVs

//Use Maemo, Tablet, and Linux to test for Nokia"s Internet Tablets.
const maemo = "maemo"
const linux = "linux"
const qtembedded = "qt embedded" //for Sony Mylo and others
const mylocom2 = "com2"          //for Sony Mylo also

//In some UserAgents, the only clue is the manufacturer.
const manuSonyEricsson = "sonyericsson"
const manuericsson = "ericsson"
const manuSamsung1 = "sec-sgh"
const manuSony = "sony"
const manuHtc = "htc" //Popular Android and WinMo manufacturer

//In some UserAgents, the only clue is the operator.
const svcDocomo = "docomo"
const svcKddi = "kddi"
const svcVodafone = "vodafone"

//Disambiguation strings.
const disUpdate = "update" //pda vs. update

type headers struct {
	userAgentHeader  string
	httpAcceptHeader string
}

type devices struct {
	initCompleted       int //Stores whether we're currently initializing the most popular func (base *UAgentInfo)s.
	IsWebkit            int //Stores the result of DetectWebkit()
	IsMobilePhone       int //Stores the result of DetectMobileQuick()
	IsIphone            int //Stores the result of DetectIphone()
	IsAndroid           int //Stores the result of DetectAndroid()
	IsAndroidPhone      int //Stores the result of DetectAndroidPhone()
	IsTierTablet        int //Stores the result of DetectTierTablet()
	IsTierIphone        int //Stores the result of DetectTierIphone()
	IsTierRichCss       int //Stores the result of DetectTierRichCss()
	IsTierGenericMobile int //Stores the result of DetectTierOtherPhones()
}

type UAgentInfo struct {
	headers
	devices
}

//**************************
//The constructor. Allows the latest PHP (5.0+) to locate a constructor object and initialize the object.
func NewMDetect(request *http.Request) *UAgentInfo {
	uAgent, httpAccept := uAgentInfo(request)

	base := UAgentInfo{}
	base.httpAcceptHeader = httpAccept
	base.userAgentHeader = uAgent

	base.initDeviceScan()
	return &base
}

//**************************
//The object initializer. Initializes several default variables.
func uAgentInfo(request *http.Request) (string, string) {
	userAgentHeader := request.Header.Get("HTTP_USER_AGENT")
	httpAcceptHeader := request.Header.Get("HTTP_ACCEPT")

	if userAgentHeader != "" {
		userAgentHeader = strings.ToLower(userAgentHeader)
	}

	if httpAcceptHeader != "" {
		httpAcceptHeader = strings.ToLower(httpAcceptHeader)
	}

	return userAgentHeader, httpAcceptHeader
}

//**************************
// Initialize Key Stored Values.
func (base *UAgentInfo) initDeviceScan() {
	base.IsWebkit = base.DetectWebkit()
	base.IsIphone = base.DetectIphone()
	base.IsAndroid = base.DetectAndroid()
	base.IsAndroidPhone = base.DetectAndroidPhone()

	//These tiers are the most useful for web development
	base.IsMobilePhone = base.DetectMobileQuick()
	base.IsTierIphone = base.DetectTierIphone()
	base.IsTierTablet = base.DetectTierTablet()

	//Optional: Comment these out if you NEVER use them.
	base.IsTierRichCss = base.DetectTierRichCss()
	base.IsTierGenericMobile = base.DetectTierOtherPhones()

	base.initCompleted = true
}

//**************************
//Returns the contents of the User Agent value, in lower case.
func (base *UAgentInfo) GetUserAgent() string {
	return base.userAgentHeader
}

//**************************
//Returns the contents of the HTTP Accept value, in lower case.
func (base *UAgentInfo) GetHttpAccept() string {
	return base.httpAcceptHeader
}

//*****************************
// Start device detection
//*****************************

//**************************
// Detects if the current device is an iPhone.
func (base *UAgentInfo) DetectIphone() int {
	if base.initCompleted == true || base.IsIphone == true {
		return base.IsIphone
	}

	if strings.Index(base.userAgentHeader, deviceIphone) > -1 {
		//The iPad and iPod Touch say they're an iPhone. So let's disambiguate. 
		if base.DetectIpad() == true || base.DetectIpod() == true {
			return false
		} else {
			//Yay! It's an iPhone!
			return true
		}
	} else {
		return false
	}
}

//**************************
// Detects if the current device is an iPod Touch.
func (base *UAgentInfo) DetectIpod() int {
	if strings.Index(base.userAgentHeader, deviceIpod) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is an iPad tablet.
func (base *UAgentInfo) DetectIpad() int {
	if strings.Index(base.userAgentHeader, deviceIpad) > -1 && base.DetectWebkit() == true {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is an iPhone or iPod Touch.
func (base *UAgentInfo) DetectIphoneOrIpod() int {
	//We repeat the searches here because some iPods may report themselves as an iPhone, which would be okay.
	if base.DetectIphone() == true || base.DetectIpod() == true {
		return true
	} else {
		return false
	}
}

//**************************
// Detects *any* iOS device: iPhone, iPod Touch, iPad.
func (base *UAgentInfo) DetectIos() int {
	if (base.DetectIphoneOrIpod() == true) || (base.DetectIpad() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects *any* Android OS-based device: phone, tablet, and multi-media player.
// Also detects Google TV.
func (base *UAgentInfo) DetectAndroid() int {
	if base.initCompleted == true || base.IsAndroid == true {
		return base.IsAndroid
	}

	if (strings.Index(base.userAgentHeader, deviceAndroid) > -1) || (base.DetectGoogleTV() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is a (small-ish) Android OS-based device
// used for calling and/or multi-media (like a Samsung Galaxy Player).
// Google says these devices will have 'Android' AND 'mobile' in user agent.
// Ignores tablets (Honeycomb and later).
func (base *UAgentInfo) DetectAndroidPhone() int {
	if base.initCompleted == true || base.IsAndroidPhone == true {
		return base.IsAndroidPhone
	}

	//First, let's make sure we're on an Android device.
	if base.DetectAndroid() == false {
		return false
	}

	//If it's Android and has 'mobile' in it, Google says it's a phone.
	if strings.Index(base.userAgentHeader, mobile) > -1 {
		return true
	}
	//Special check for Android devices with Opera Mobile/Mini. They should report here.
	if base.DetectOperaMobile() == true {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is a (self-reported) Android tablet.
// Google says these devices will have 'Android' and NOT 'mobile' in their user agent.
func (base *UAgentInfo) DetectAndroidTablet() int {
	//First, let's make sure we're on an Android device.
	if base.DetectAndroid() == false {
		return false
	}

	//Special check for Android devices with Opera Mobile/Mini. They should NOT report here.
	if base.DetectOperaMobile() == true {
		return false
	}

	//Otherwise, if it's Android and does NOT have 'mobile' in it, Google says it's a tablet.
	if strings.Index(base.userAgentHeader, mobile) > -1 {
		return false
	} else {
		return true
	}
}

//**************************
// Detects if the current device is an Android OS-based device and
//   the browser is based on WebKit.
func (base *UAgentInfo) DetectAndroidWebKit() int {
	if (base.DetectAndroid() == true) && (base.DetectWebkit() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is a GoogleTV.
func (base *UAgentInfo) DetectGoogleTV() int {
	if strings.Index(base.userAgentHeader, deviceGoogleTV) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is based on WebKit.
func (base *UAgentInfo) DetectWebkit() int {
	if base.initCompleted == true || base.IsWebkit == true {
		return base.IsWebkit
	}

	if strings.Index(base.userAgentHeader, engineWebKit) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a 
// Windows Phone 7, 8, or 10 device.
func (base *UAgentInfo) DetectWindowsPhone() int {
	if (base.DetectWindowsPhone7() == true) || (base.DetectWindowsPhone8() == true) || (base.DetectWindowsPhone10() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a Windows Phone 7 device (in mobile browsing mode).
func (base *UAgentInfo) DetectWindowsPhone7() int {
	if strings.Index(base.userAgentHeader, deviceWinPhone7) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a Windows Phone 8 device (in mobile browsing mode).
func (base *UAgentInfo) DetectWindowsPhone8() int {
	if strings.Index(base.userAgentHeader, deviceWinPhone8) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a Windows Phone 10 device (in mobile browsing mode).
func (base *UAgentInfo) DetectWindowsPhone10() int {
	if strings.Index(base.userAgentHeader, deviceWinPhone10) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a Windows Mobile device.
// Excludes Windows Phone 7 and later devices.
// Focuses on Windows Mobile 6.xx and earlier.
func (base *UAgentInfo) DetectWindowsMobile() int {
	if base.DetectWindowsPhone() == true {
		return false
	}

	//Most devices use 'Windows CE', but some report 'iemobile'
	//  and some older ones report as 'PIE' for Pocket IE.
	if strings.Index(base.userAgentHeader, deviceWinMob) > -1 || strings.Index(base.userAgentHeader, deviceIeMob) > -1 ||
		strings.Index(base.userAgentHeader, enginePie) > -1 {
		return true
	} //Test for Windows Mobile PPC but not old Macintosh PowerPC.
	if strings.Index(base.userAgentHeader, devicePpc) > -1 && !(strings.Index(base.userAgentHeader, deviceMacPpc) > 1) {
		return true
	} //Test for certain Windwos Mobile-based HTC devices.
	if strings.Index(base.userAgentHeader, manuHtc) > -1 && strings.Index(base.userAgentHeader, deviceWindows) > -1 {
		return true
	}
	if base.DetectWapWml() == true && strings.Index(base.userAgentHeader, deviceWindows) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is any BlackBerry device.
// Includes BB10 OS, but excludes the PlayBook.
func (base *UAgentInfo) DetectBlackBerry() int {
	if (strings.Index(base.userAgentHeader, deviceBB) > -1) || (strings.Index(base.httpAcceptHeader, vndRIM) > -1) {
		return true
	}
	if base.DetectBlackBerry10Phone() == true {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a BlackBerry 10 OS phone.
// Excludes tablets.
func (base *UAgentInfo) DetectBlackBerry10Phone() int {
	if (strings.Index(base.userAgentHeader, deviceBB10) > -1) && (strings.Index(base.userAgentHeader, mobile) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is on a BlackBerry tablet device.
//    Examples: PlayBook
func (base *UAgentInfo) DetectBlackBerryTablet() int {
	if strings.Index(base.userAgentHeader, deviceBBPlaybook) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a BlackBerry phone device AND uses a
//    WebKit-based browser. These are signatures for the new BlackBerry OS 6.
//    Examples: Torch. Includes the Playbook.
func (base *UAgentInfo) DetectBlackBerryWebKit() int {
	if (base.DetectBlackBerry() == true) && (base.DetectWebkit() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a BlackBerry Touch phone device with
//    a large screen, such as the Storm, Torch, and Bold Touch. Excludes the Playbook.
func (base *UAgentInfo) DetectBlackBerryTouch() int {
	if (strings.Index(base.userAgentHeader, deviceBBStorm) > -1) || (strings.Index(base.userAgentHeader, deviceBBTorch) > -1) ||
		(strings.Index(base.userAgentHeader, deviceBBBoldTouch) > -1) || (strings.Index(base.userAgentHeader, deviceBBCurveTouch) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a BlackBerry OS 5 device AND
//    has a more capable recent browser. Excludes the Playbook.
//    Examples, Storm, Bold, Tour, Curve2
//    Excludes the new BlackBerry OS 6 and 7 browser!!
func (base *UAgentInfo) DetectBlackBerryHigh() int {
	//Disambiguate for BlackBerry OS 6 or 7 (WebKit) browser
	if base.DetectBlackBerryWebKit() == true {
		return false
	}
	if base.DetectBlackBerry() == true {
		if (base.DetectBlackBerryTouch() == true) || strings.Index(base.userAgentHeader, deviceBBBold) > -1 ||
			strings.Index(base.userAgentHeader, deviceBBTour) > -1 || strings.Index(base.userAgentHeader, deviceBBCurve) > -1 {
			{
				return true
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a BlackBerry device AND
//    has an older, less capable browser. 
//    Examples: Pearl, 8800, Curve1.
func (base *UAgentInfo) DetectBlackBerryLow() int {
	if base.DetectBlackBerry() == true {
		//Assume that if it's not in the High tier, then it's Low.
		if (base.DetectBlackBerryHigh() == true) || (base.DetectBlackBerryWebKit() == true) {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is the Nokia S60 Open Source Browser.
func (base *UAgentInfo) DetectS60OssBrowser() int {
	//First, test for WebKit, then make sure it's either Symbian or S60.
	if base.DetectWebkit() == true {
		if strings.Index(base.userAgentHeader, deviceSymbian) > -1 || strings.Index(base.userAgentHeader, deviceS60) > -1 {
			{
				return true
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

//**************************
// Detects if the current device is any Symbian OS-based device,
//   including older S60, Series 70, Series 80, Series 90, and UIQ, 
//   or other browsers running on these devices.
func (base *UAgentInfo) DetectSymbianOS() int {
	if strings.Index(base.userAgentHeader, deviceSymbian) > -1 || strings.Index(base.userAgentHeader, deviceS60) > -1 ||
		strings.Index(base.userAgentHeader, deviceS70) > -1 || strings.Index(base.userAgentHeader, deviceS80) > -1 ||
		strings.Index(base.userAgentHeader, deviceS90) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is on a PalmOS device.
func (base *UAgentInfo) DetectPalmOS() int {
	//Most devices nowadays report as 'Palm', but some older ones reported as Blazer or Xiino.
	if strings.Index(base.userAgentHeader, devicePalm) > -1 ||
		strings.Index(base.userAgentHeader, engineBlazer) > -1 ||
		strings.Index(base.userAgentHeader, engineXiino) > -1 {
		//Make sure it's not WebOS first
		if base.DetectPalmWebOS() == true {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is on a Palm device
//   running the new WebOS.
func (base *UAgentInfo) DetectPalmWebOS() int {
	if strings.Index(base.userAgentHeader, deviceWebOS) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is on an HP tablet running WebOS.
func (base *UAgentInfo) DetectWebOSTablet() int {
	if (strings.Index(base.userAgentHeader, deviceWebOShp) > -1) && (strings.Index(base.userAgentHeader, deviceTablet) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is on a WebOS smart TV.
func (base *UAgentInfo) DetectWebOSTV() int {
	if (strings.Index(base.userAgentHeader, deviceWebOStv) > -1) && (strings.Index(base.userAgentHeader, smartTV2) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is Opera Mobile or Mini.
func (base *UAgentInfo) DetectOperaMobile() int {
	if (strings.Index(base.userAgentHeader, engineOpera) > -1) &&
		((strings.Index(base.userAgentHeader, mini) > -1) ||
			(strings.Index(base.userAgentHeader, mobi) > -1)) {
		return true
	}
	return false
}

//**************************
// Detects if the current device is an Amazon Kindle (eInk devices only).
// Note: For the Kindle Fire, use the normal Android methods. 
func (base *UAgentInfo) DetectKindle() int {
	if strings.Index(base.userAgentHeader, deviceKindle) > -1 &&
		base.DetectAndroid() == false {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current Amazon device has turned on the Silk accelerated browsing feature.
// Note: Typically used by the the Kindle Fire.
func (base *UAgentInfo) DetectAmazonSilk() int {
	if strings.Index(base.userAgentHeader, engineSilk) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if a Garmin Nuvifone device.
func (base *UAgentInfo) DetectGarminNuvifone() int {
	if strings.Index(base.userAgentHeader, deviceNuvifone) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a device running the Bada OS from Samsung.
func (base *UAgentInfo) DetectBada() int {
	if strings.Index(base.userAgentHeader, deviceBada) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a device running the Tizen smartphone OS.
func (base *UAgentInfo) DetectTizen() int {
	if strings.Index(base.userAgentHeader, deviceTizen) > -1 && (strings.Index(base.userAgentHeader, mobile) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is on a Tizen smart TV.
func (base *UAgentInfo) DetectTizenTV() int {
	if (strings.Index(base.userAgentHeader, deviceTizen) > -1) && (strings.Index(base.userAgentHeader, smartTV1) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a device running the Meego OS.
func (base *UAgentInfo) DetectMeego() int {
	if strings.Index(base.userAgentHeader, deviceMeego) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a phone running the Meego OS.
func (base *UAgentInfo) DetectMeegoPhone() int {
	if (strings.Index(base.userAgentHeader, deviceMeego) > -1) && (strings.Index(base.userAgentHeader, mobi) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a mobile device (probably) running the Firefox OS.
func (base *UAgentInfo) DetectFirefoxOS() int {
	if (base.DetectFirefoxOSPhone() == true) || (base.DetectFirefoxOSTablet() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a phone (probably) running the Firefox OS.
func (base *UAgentInfo) DetectFirefoxOSPhone() int {
	//First, let's make sure we're NOT on another major mobile OS.
	if base.DetectIos() == true || base.DetectAndroid() == true || base.DetectSailfish() == true {
		return false
	}

	if (strings.Index(base.userAgentHeader, engineFirefox) > -1) &&
		(strings.Index(base.userAgentHeader, mobile) > -1) {
		return true
	}
	return false
}

//**************************
// Detects a tablet (probably) running the Firefox OS.
func (base *UAgentInfo) DetectFirefoxOSTablet() int {
	//First, let's make sure we're NOT on another major mobile OS.
	if base.DetectIos() == true || base.DetectAndroid() == true || base.DetectSailfish() == true {
		return false
	}

	if (strings.Index(base.userAgentHeader, engineFirefox) > -1) &&
		(strings.Index(base.userAgentHeader, deviceTablet) > -1) {
		return true
	}
	return false
}

//**************************
// Detects a device running the Sailfish OS.
func (base *UAgentInfo) DetectSailfish() int {
	if strings.Index(base.userAgentHeader, deviceSailfish) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a phone running the Sailfish OS.
func (base *UAgentInfo) DetectSailfishPhone() int {
	if (base.DetectSailfish() == true) &&
		(strings.Index(base.userAgentHeader, mobile) > -1) {
		return true
	}
	return false
}

//**************************
// Detects a mobile device running the Ubuntu Mobile OS.
func (base *UAgentInfo) DetectUbuntu() int {
	if (base.DetectUbuntuPhone() == true) || (base.DetectUbuntuTablet() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects a phone running the Ubuntu Mobile OS.
func (base *UAgentInfo) DetectUbuntuPhone() int {
	if (strings.Index(base.userAgentHeader, deviceUbuntu) > -1) && (strings.Index(base.userAgentHeader, mobile) > -1) {
		return true
	}
	return false
}

//**************************
// Detects a tablet running the Ubuntu Mobile OS.
func (base *UAgentInfo) DetectUbuntuTablet() int {
	if (strings.Index(base.userAgentHeader, deviceUbuntu) > -1) &&
		(strings.Index(base.userAgentHeader, deviceTablet) > -1) {
		return true
	}
	return false
}

//**************************
// Detects the Danger Hiptop device.
func (base *UAgentInfo) DetectDangerHiptop() int {
	if strings.Index(base.userAgentHeader, deviceDanger) > -1 ||
		strings.Index(base.userAgentHeader, deviceHiptop) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current browser is a Sony Mylo device.
func (base *UAgentInfo) DetectSonyMylo() int {
	if (strings.Index(base.userAgentHeader, manuSony) > -1) &&
		((strings.Index(base.userAgentHeader, qtembedded) > -1) ||
			(strings.Index(base.userAgentHeader, mylocom2) > -1)) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is on one of the Maemo-based Nokia Internet Tablets.
func (base *UAgentInfo) DetectMaemoTablet() int {
	if strings.Index(base.userAgentHeader, maemo) > -1 {
		return true
	} //For Nokia N810, must be Linux + Tablet, or else it could be something else.
	if (strings.Index(base.userAgentHeader, linux) > -1) && (strings.Index(base.userAgentHeader, deviceTablet) > -1) && (base.DetectWebOSTablet() == false) && (base.DetectAndroid() == false) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is an Archos media player/Internet tablet.
func (base *UAgentInfo) DetectArchos() int {
	if strings.Index(base.userAgentHeader, deviceArchos) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is an Internet-capable game console.
// Includes many handheld consoles.
func (base *UAgentInfo) DetectGameConsole() int {
	if base.DetectSonyPlaystation() == true {
		return true
	} else if base.DetectNintendo() == true {
		return true
	} else if base.DetectXbox() == true {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is a Sony Playstation.
func (base *UAgentInfo) DetectSonyPlaystation() int {
	if strings.Index(base.userAgentHeader, devicePlaystation) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is a handheld gaming device with
// a touchscreen and modern iPhone-class browser. Includes the Playstation Vita.
func (base *UAgentInfo) DetectGamingHandheld() int {
	if (strings.Index(base.userAgentHeader, devicePlaystation) > -1) &&
		(strings.Index(base.userAgentHeader, devicePlaystationVita) > -1) {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is a Nintendo game device.
func (base *UAgentInfo) DetectNintendo() int {
	if strings.Index(base.userAgentHeader, deviceNintendo) > -1 ||
		strings.Index(base.userAgentHeader, deviceWii) > -1 ||
		strings.Index(base.userAgentHeader, deviceNintendoDs) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device is a Microsoft Xbox.
func (base *UAgentInfo) DetectXbox() int {
	if strings.Index(base.userAgentHeader, deviceXbox) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects whether the device is a Brew-powered device.
func (base *UAgentInfo) DetectBrewDevice() int {
	if strings.Index(base.userAgentHeader, deviceBrew) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects whether the device supports WAP or WML.
func (base *UAgentInfo) DetectWapWml() int {
	if strings.Index(base.httpAcceptHeader, vndwap) > -1 ||
		strings.Index(base.httpAcceptHeader, wml) > -1 {
		return true
	} else {
		return false
	}
}

//**************************
// Detects if the current device supports MIDP, a mobile Java technology.
func (base *UAgentInfo) DetectMidpCapable() int {
	if strings.Index(base.userAgentHeader, deviceMidp) > -1 ||
		strings.Index(base.httpAcceptHeader, deviceMidp) > -1 {
		return true
	} else {
		return false
	}
}

//*****************************
// Device Classes
//*****************************

//**************************
// Check to see whether the device is *any* 'smartphone'.
//   Note: It's better to use DetectTierIphone() for modern touchscreen devices.
func (base *UAgentInfo) DetectSmartphone() int {
	//Exclude duplicates from TierIphone
	if (base.DetectTierIphone() == true) || (base.DetectS60OssBrowser() == true) || (base.DetectSymbianOS() == true) ||
		(base.DetectWindowsMobile() == true) || (base.DetectBlackBerry() == true) || (base.DetectMeegoPhone() == true) ||
		(base.DetectPalmWebOS() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// The quick way to detect for a mobile device.
//   Will probably detect most recent/current mid-tier Feature Phones
//   as well as smartphone-class devices. Excludes Apple iPads and other modern tablets.
func (base *UAgentInfo) DetectMobileQuick() int {
	if base.initCompleted == true || base.IsMobilePhone == true {
		return base.IsMobilePhone
	}

	//Let's exclude tablets
	if base.IsTierTablet == true {
		return false
	}

	//Most mobile browsing is done on smartphones
	if base.DetectSmartphone() == true {
		return true
	}
	//Catch-all for many mobile devices
	if strings.Index(base.userAgentHeader, mobile) > -1 {
		return true
	}
	if base.DetectOperaMobile() == true {
		return true
	}
	//We also look for Kindle devices
	if base.DetectKindle() == true ||
		base.DetectAmazonSilk() == true {
		return true
	}
	if (base.DetectWapWml() == true) || (base.DetectMidpCapable() == true) || (base.DetectBrewDevice() == true) {
		return true
	}
	if (strings.Index(base.userAgentHeader, engineNetfront) > -1) || (strings.Index(base.userAgentHeader, engineUpBrowser) > -1) {
		return true
	}
	return false
}

//**************************
// The longer and more thorough way to detect for a mobile device.
//   Will probably detect most feature phones,
//   smartphone-class devices, Internet Tablets,
//   Internet-enabled game consoles, etc.
//   This ought to catch a lot of the more obscure and older devices, also --
//   but no promises on thoroughness!
func (base *UAgentInfo) DetectMobileLong() int {
	if base.DetectMobileQuick() == true {
		return true
	}
	if base.DetectGameConsole() == true {
		return true
	}
	if (base.DetectDangerHiptop() == true) || (base.DetectMaemoTablet() == true) || (base.DetectSonyMylo() == true) ||
		(base.DetectArchos() == true) {
		return true
	}
	if (strings.Index(base.userAgentHeader, devicePda) > -1) && !(strings.Index(base.userAgentHeader, disUpdate) > -1) {
		return true
	}
	//Detect older phones from certain manufacturers and operators.
	if (strings.Index(base.userAgentHeader, uplink) > -1) || (strings.Index(base.userAgentHeader, engineOpenWeb) > -1) ||
		(strings.Index(base.userAgentHeader, manuSamsung1) > -1) || (strings.Index(base.userAgentHeader, manuSonyEricsson) > -1) ||
		(strings.Index(base.userAgentHeader, manuericsson) > -1) || (strings.Index(base.userAgentHeader, svcDocomo) > -1) ||
		(strings.Index(base.userAgentHeader, svcKddi) > -1) || (strings.Index(base.userAgentHeader, svcVodafone) > -1) {
		return true
	}
	return false
}

//*****************************
// For Mobile Web Site Design
//*****************************

//**************************
// The quick way to detect for a tier of devices.
//   This method detects for the new generation of
//   HTML 5 capable, larger screen tablets.
//   Includes iPad, Android (e.g., Xoom), BB Playbook, WebOS, etc.
func (base *UAgentInfo) DetectTierTablet() int {
	if base.initCompleted == true || base.IsTierTablet == true {
		return base.IsTierTablet
	}

	if (base.DetectIpad() == true) || (base.DetectAndroidTablet() == true) || (base.DetectBlackBerryTablet() == true) ||
		(base.DetectFirefoxOSTablet() == true) || (base.DetectUbuntuTablet() == true) || (base.DetectWebOSTablet() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// The quick way to detect for a tier of devices.
//   This method detects for devices which can
//   display iPhone-optimized web content.
//   Includes iPhone, iPod Touch, Android, Windows Phone, BB10, Playstation Vita, etc.
func (base *UAgentInfo) DetectTierIphone() int {
	if base.initCompleted == true || base.IsTierIphone == true {
		return base.IsTierIphone
	}

	if (base.DetectIphoneOrIpod() == true) || (base.DetectAndroidPhone() == true) || (base.DetectWindowsPhone() == true) ||
		(base.DetectBlackBerry10Phone() == true) || (base.DetectPalmWebOS() == true) || (base.DetectBada() == true) ||
		(base.DetectTizen() == true) || (base.DetectFirefoxOSPhone() == true) || (base.DetectSailfishPhone() == true) ||
		(base.DetectUbuntuPhone() == true) || (base.DetectGamingHandheld() == true) {
		return true
	}
	//Note: BB10 phone is in the previous paragraph
	if (base.DetectBlackBerryWebKit() == true) &&
		(base.DetectBlackBerryTouch() == true) {
		return true
	} else {
		return false
	}
}

//**************************
// The quick way to detect for a tier of devices.
//   This method detects for devices which are likely to be capable
//   of viewing CSS content optimized for the iPhone, 
//   but may not necessarily support JavaScript.
//   Excludes all iPhone Tier devices.
func (base *UAgentInfo) DetectTierRichCss() int {
	if base.initCompleted == true || base.IsTierRichCss == true {
		return base.IsTierRichCss
	}

	if base.DetectMobileQuick() == true {
		//Exclude iPhone Tier and e-Ink Kindle devices
		if (base.DetectTierIphone() == true) || (base.DetectKindle() == true) {
			return false
		}

		//The following devices are explicitly ok.
		if base.DetectWebkit() == true {
			//Any WebKit
			return true
		}
		if base.DetectS60OssBrowser() == true {
			return true
		}
		//Note: 'High' BlackBerry devices ONLY
		if base.DetectBlackBerryHigh() == true {
			return true
		}
		//Older Windows 'Mobile' isn't good enough for iPhone Tier.
		if base.DetectWindowsMobile() == true {
			return true
		}
		if strings.Index(base.userAgentHeader, engineTelecaQ) > -1 {
			return true
		} else {
			//default
			return false
		}
	} else {
		return false
	}
}

//**************************
// The quick way to detect for a tier of devices.
//   This method detects for all other types of phones,
//   but excludes the iPhone and RichCSS Tier devices.
func (base *UAgentInfo) DetectTierOtherPhones() int {
	if base.initCompleted == true ||
		base.IsTierGenericMobile == true {
		return base.IsTierGenericMobile
	}

	//Exclude devices in the other 2 categories
	if (base.DetectMobileLong() == true) && (base.DetectTierIphone() == false) && (base.DetectTierRichCss() == false) {
		return true
	} else {
		return false
	}
}
