package imaging

import (
	"image"
	"testing"
)

func TestBlur(t *testing.T) {
	td := []struct {
		desc  string
		src   image.Image
		sigma float64
		want  *image.NRGBA
	}{
		{
			"Blur 3x3 0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x66, 0xaa, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
			0.0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x66, 0xaa, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			"Blur 3x3 0.5",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x66, 0xaa, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
			0.5,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x66, 0xaa, 0xff, 0x04, 0x66, 0xaa, 0xff, 0x18, 0x66, 0xaa, 0xff, 0x04,
					0x66, 0xaa, 0xff, 0x18, 0x66, 0xaa, 0xff, 0x9e, 0x66, 0xaa, 0xff, 0x18,
					0x66, 0xaa, 0xff, 0x04, 0x66, 0xaa, 0xff, 0x18, 0x66, 0xaa, 0xff, 0x04,
				},
			},
		},
		{
			"Blur 3x3 10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x66, 0xaa, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
			10,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x66, 0xaa, 0xff, 0x1c, 0x66, 0xaa, 0xff, 0x1c, 0x66, 0xaa, 0xff, 0x1c,
					0x66, 0xaa, 0xff, 0x1c, 0x66, 0xaa, 0xff, 0x1c, 0x66, 0xaa, 0xff, 0x1c,
					0x66, 0xaa, 0xff, 0x1c, 0x66, 0xaa, 0xff, 0x1c, 0x66, 0xaa, 0xff, 0x1c,
				},
			},
		},
	}
	for _, d := range td {
		got := Blur(d.src, d.sigma)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestBlurGolden(t *testing.T) {
	src, err := Open("testdata/lena_128.png")
	if err != nil {
		t.Errorf("Open: %v", err)
	}
	for name, sigma := range map[string]float64{
		"out_blur_0.5.png": 0.5,
		"out_blur_1.5.png": 1.5,
	} {
		got := Blur(src, sigma)
		want, err := Open("testdata/" + name)
		if err != nil {
			t.Errorf("Open: %v", err)
		}
		if !compareNRGBA(got, toNRGBA(want), 0) {
			t.Errorf("resulting image differs from golden: %s", name)
		}
	}
}

func TestSharpen(t *testing.T) {
	td := []struct {
		desc  string
		src   image.Image
		sigma float64
		want  *image.NRGBA
	}{
		{
			"Sharpen 3x3 0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x77, 0x77, 0x77, 0x77, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				},
			},
			0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x77, 0x77, 0x77, 0x77, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				},
			},
		},
		{
			"Sharpen 3x3 0.5",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x77, 0x77, 0x77, 0x77, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				},
			},
			0.5,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x66, 0x66, 0x66, 0x66, 0x64, 0x64, 0x64, 0x64, 0x66, 0x66, 0x66, 0x66,
					0x64, 0x64, 0x64, 0x64, 0x7d, 0x7d, 0x7d, 0x7e, 0x64, 0x64, 0x64, 0x64,
					0x66, 0x66, 0x66, 0x66, 0x64, 0x64, 0x64, 0x64, 0x66, 0x66, 0x66, 0x66,
				},
			},
		},
		{
			"Sharpen 3x3 100",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x77, 0x77, 0x77, 0x77, 0x66, 0x66, 0x66, 0x66,
					0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				},
			},
			100,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64,
					0x64, 0x64, 0x64, 0x64, 0x86, 0x86, 0x86, 0x86, 0x64, 0x64, 0x64, 0x64,
					0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64, 0x64,
				},
			},
		},
		{
			"Sharpen 3x1 10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 0),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
				},
			},
			10,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 1),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Sharpen(d.src, d.sigma)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestSharpenGolden(t *testing.T) {
	src, err := Open("testdata/lena_128.png")
	if err != nil {
		t.Errorf("Open: %v", err)
	}
	for name, sigma := range map[string]float64{
		"out_sharpen_0.5.png": 0.5,
		"out_sharpen_1.5.png": 1.5,
	} {
		got := Sharpen(src, sigma)
		want, err := Open("testdata/" + name)
		if err != nil {
			t.Errorf("Open: %v", err)
		}
		if !compareNRGBA(got, toNRGBA(want), 0) {
			t.Errorf("resulting image differs from golden: %s", name)
		}
	}
}
