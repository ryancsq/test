package gogb2312

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func hexprint_utf8(u8 int) {
	fmt.Printf("%x %x %x\n", u8>>16, (u8>>8)&0xff, u8&0xff)
}

func hexprint_utf8string(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func hexprint_bytes(s []byte) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func test_unicode2utf8(t *testing.T) {
	s := unicode2utf8(0x90ed)
	hexprint_utf8(s)
	b := make([]byte, 3)
	b[0] = byte(s >> 16)
	b[1] = byte((s >> 8) & 0xff)
	b[2] = byte(s & 0xff)
	fmt.Println(string(b))
}

func test_RE(t *testing.T) {
	var ss = []string{
		"who  \t are u \t\t  ? ",
		"i \t \t am\t\t\t   boy.  ",
		"i \t \t am\t\t\t 90  boy.  ",
	}
	var gb = []byte("\x90\xae")
	var b = byte(98)

	for _, s := range ss {
		ret := re_space.ReplaceAllString(s, " ")
		_ = ret
		//fmt.Println(s, " - ", ret)
	}
	_ = gb
	_ = b
}

func Test_MakeData(t *testing.T) {
	//readcp936("./CP936.TXT")
}

func convert_file(t *testing.T, fp string) {
	buf, err := ioutil.ReadFile("./test/test1.txt")
	if err != nil {
		t.Error(err.Error())
	}
	//hexprint_bytes(buf)
	output, cerr, il, ol := ConvertGB2312(buf)
	if cerr != nil {
		t.Error(cerr.Error())
	}
	_ = il
	_ = ol
	_ = output
}

func test_Convert(t *testing.T) {
	bn := []byte("\xbf\xc6\xd1\xa7\xC3\xF1\xD6\xF7\xCF\xDC\xD5\xFE")
	sn := string(bn)

	cbn, err1, _, _ := ConvertGB2312(bn)
	if err1 != nil {
		t.Error("convert failed!")
	}
	fmt.Printf("%s\n", cbn)
	csn, err2, _, _ := ConvertGB2312String(sn)
	if err2 != nil {
		t.Error("convert failed!")
	}
	fmt.Printf("%s\n", csn)
}

func _Test_isutf8(t *testing.T) {
	bn0 := []byte("\xbf\xc6\xd1\xa7\xC3\xF1\xD6\xF7\xCF\xDC\xD5\xFE")
	bn1 := []byte("\xbf\xc6\xe2\x80\x94\xd1\xa7\xC3\xF1\xD6\xF7\xCF\xDC\xD5\xFE")

	test_isutf8(bn0)
	test_isutf8(bn1)
	ss, err, _, _ := ConvertHybirdString(string(bn1))
	if err == nil {
		fmt.Println(ss)
	} else {
		t.Error(err.Error())
	}
}

func test_isutf8(bs []byte) {
	tlen := len(bs)
	for i := 0; i < tlen; i++ {
		res := isutf8(bs[i:], tlen-i)
		if res > 0 {
			fmt.Printf("found utf-8 at %d, length %d. %q\n", i, res, bs[i:])
		}
	}
	fmt.Println()
}

func saveToFile(buf []byte) {
	err := ioutil.WriteFile("d:\\f1.txt", buf, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	obuf, err, _, _ := ConvertGB2312(buf)
	if err != nil {
		panic(err.Error())
	}
	err = ioutil.WriteFile("d:\\f2.txt", obuf, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
}

// 两个问题：
// 英文名字中间的点convert不正确，温度上标convert不正确
// 有些页面整体convert错误，例如：http://sports.sina.com.cn/n/2014-02-24/04137038088.shtml
//                           该页面是由于其本身进行了混淆处理。
func TestUrl(t *testing.T) {
	resp, err := http.Get("http://ent.sina.com.cn/zl/bagua/blog/2014-02-25/15341218/2036655340/7964e4ec0101sigp.shtml")
	if err != nil {
		t.Fatal("get page failed!\n")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("read response failed!\n")
	}

	saveToFile(body)
}
