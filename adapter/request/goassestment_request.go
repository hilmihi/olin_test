package request

type Soal1Request struct {
	Nums   []int `validate:"required" json:"nums"`
	Target int   `validate:"required" json:"target"`
}

type Soal2Request struct {
	Nums []int `validate:"required" json:"nums"`
}

type Soal3Request struct {
	Str   string   `validate:"required" json:"str"`
	Words []string `validate:"required" json:"words"`
}
