// This file is automatically generated by generate-charset-data.
// Do not hand-edit.

package data

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/triggit/go-charset/charset"
)

func init() {
	charset.RegisterDataFile("jisx0201kana.dat", func() (io.ReadCloser, error) {
		r := strings.NewReader("｡｢｣､･ｦｧｨｩｪｫｬｭｮｯｰｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜﾝﾞﾟ")
		return ioutil.NopCloser(r), nil
	})
}
