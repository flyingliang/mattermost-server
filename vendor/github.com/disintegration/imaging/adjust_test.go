package imaging

import (
	"image"
	"testing"
)

func TestGrayscale(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Grayscale 3x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x3d, 0x3d, 0x3d, 0x01, 0x78, 0x78, 0x78, 0x02, 0x17, 0x17, 0x17, 0x03,
					0x1f, 0x1f, 0x1f, 0xff, 0x25, 0x25, 0x25, 0xff, 0x66, 0x66, 0x66, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Grayscale(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestInvert(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Invert 3x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x33, 0xff, 0xff, 0x01, 0xff, 0x33, 0xff, 0x02, 0xff, 0xff, 0x33, 0x03,
					0xee, 0xdd, 0xcc, 0xff, 0xcc, 0xdd, 0xee, 0xff, 0x55, 0xcc, 0x44, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xcc, 0xcc, 0xcc, 0xff, 0x00, 0x00, 0x00, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Invert(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestAdjustContrast(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		p    float64
		want *image.NRGBA
	}{
		{
			"AdjustContrast 3x3 10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			10,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xd5, 0x00, 0x00, 0x01, 0x00, 0xd5, 0x00, 0x02, 0x00, 0x00, 0xd5, 0x03,
					0x05, 0x18, 0x2b, 0xff, 0x2b, 0x18, 0x05, 0xff, 0xaf, 0x2b, 0xc2, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x2b, 0x2b, 0x2b, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustContrast 3x3 100",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			100,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0x01, 0x00, 0xff, 0x00, 0x02, 0x00, 0x00, 0xff, 0x03,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustContrast 3x3 -10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			-10,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xc4, 0x0d, 0x0d, 0x01, 0x0d, 0xc4, 0x0d, 0x02, 0x0d, 0x0d, 0xc4, 0x03,
					0x1c, 0x2b, 0x3b, 0xff, 0x3b, 0x2b, 0x1c, 0xff, 0xa6, 0x3b, 0xb5, 0xff,
					0x0d, 0x0d, 0x0d, 0xff, 0x3b, 0x3b, 0x3b, 0xff, 0xf2, 0xf2, 0xf2, 0xff,
				},
			},
		},
		{
			"AdjustContrast 3x3 -100",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			-100,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x01, 0x80, 0x80, 0x80, 0x02, 0x80, 0x80, 0x80, 0x03,
					0x80, 0x80, 0x80, 0xff, 0x80, 0x80, 0x80, 0xff, 0x80, 0x80, 0x80, 0xff,
					0x80, 0x80, 0x80, 0xff, 0x80, 0x80, 0x80, 0xff, 0x80, 0x80, 0x80, 0xff,
				},
			},
		},
		{
			"AdjustContrast 3x3 0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := AdjustContrast(d.src, d.p)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestAdjustContrastGolden(t *testing.T) {
	src, err := Open("testdata/lena_128.png")
	if err != nil {
		t.Errorf("Open: %v", err)
	}
	for name, p := range map[string]float64{
		"out_contrast_m10.png": -10,
		"out_contrast_p10.png": 10,
	} {
		got := AdjustContrast(src, p)
		want, err := Open("testdata/" + name)
		if err != nil {
			t.Errorf("Open: %v", err)
		}
		if !compareNRGBA(got, toNRGBA(want), 0) {
			t.Errorf("resulting image differs from golden: %s", name)
		}
	}
}

func TestAdjustBrightness(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		p    float64
		want *image.NRGBA
	}{
		{
			"AdjustBrightness 3x3 10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			10,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xe6, 0x1a, 0x1a, 0x01, 0x1a, 0xe6, 0x1a, 0x02, 0x1a, 0x1a, 0xe6, 0x03,
					0x2b, 0x3c, 0x4d, 0xff, 0x4d, 0x3c, 0x2b, 0xff, 0xc4, 0x4d, 0xd5, 0xff,
					0x1a, 0x1a, 0x1a, 0xff, 0x4d, 0x4d, 0x4d, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustBrightness 3x3 100",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			100,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0x01, 0xff, 0xff, 0xff, 0x02, 0xff, 0xff, 0xff, 0x03,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustBrightness 3x3 -10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			-10,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xb3, 0x00, 0x00, 0x01, 0x00, 0xb3, 0x00, 0x02, 0x00, 0x00, 0xb3, 0x03,
					0x00, 0x09, 0x1a, 0xff, 0x1a, 0x09, 0x00, 0xff, 0x91, 0x1a, 0xa2, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x1a, 0x1a, 0x1a, 0xff, 0xe6, 0xe6, 0xe6, 0xff,
				},
			},
		},
		{
			"AdjustBrightness 3x3 -100",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			-100,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"AdjustBrightness 3x3 0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := AdjustBrightness(d.src, d.p)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestAdjustBrightnessGolden(t *testing.T) {
	src, err := Open("testdata/lena_128.png")
	if err != nil {
		t.Errorf("Open: %v", err)
	}
	for name, p := range map[string]float64{
		"out_brightness_m10.png": -10,
		"out_brightness_p10.png": 10,
	} {
		got := AdjustBrightness(src, p)
		want, err := Open("testdata/" + name)
		if err != nil {
			t.Errorf("Open: %v", err)
		}
		if !compareNRGBA(got, toNRGBA(want), 0) {
			t.Errorf("resulting image differs from golden: %s", name)
		}
	}
}

func TestAdjustGamma(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		p    float64
		want *image.NRGBA
	}{
		{
			"AdjustGamma 3x3 0.75",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			0.75,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xbd, 0x00, 0x00, 0x01, 0x00, 0xbd, 0x00, 0x02, 0x00, 0x00, 0xbd, 0x03,
					0x07, 0x11, 0x1e, 0xff, 0x1e, 0x11, 0x07, 0xff, 0x95, 0x1e, 0xa9, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x1e, 0x1e, 0x1e, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustGamma 3x3 1.5",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1.5,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xdc, 0x00, 0x00, 0x01, 0x00, 0xdc, 0x00, 0x02, 0x00, 0x00, 0xdc, 0x03,
					0x2a, 0x43, 0x57, 0xff, 0x57, 0x43, 0x2a, 0xff, 0xc3, 0x57, 0xcf, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x57, 0x57, 0x57, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustGamma 3x3 1.0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1.0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := AdjustGamma(d.src, d.p)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestAdjustGammaGolden(t *testing.T) {
	src, err := Open("testdata/lena_128.png")
	if err != nil {
		t.Errorf("Open: %v", err)
	}
	for name, g := range map[string]float64{
		"out_gamma_0.75.png": 0.75,
		"out_gamma_1.25.png": 1.25,
	} {
		got := AdjustGamma(src, g)
		want, err := Open("testdata/" + name)
		if err != nil {
			t.Errorf("Open: %v", err)
		}
		if !compareNRGBA(got, toNRGBA(want), 0) {
			t.Errorf("resulting image differs from golden: %s", name)
		}
	}
}

func TestAdjustSigmoid(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		m    float64
		p    float64
		want *image.NRGBA
	}{
		{
			"AdjustSigmoid 3x3 0.5 3.0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			0.5,
			3.0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xd4, 0x00, 0x00, 0x01, 0x00, 0xd4, 0x00, 0x02, 0x00, 0x00, 0xd4, 0x03,
					0x0d, 0x1b, 0x2b, 0xff, 0x2b, 0x1b, 0x0d, 0xff, 0xb1, 0x2b, 0xc3, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x2b, 0x2b, 0x2b, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustSigmoid 3x3 0.5 -3.0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			0.5,
			-3.0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xc4, 0x00, 0x00, 0x01, 0x00, 0xc4, 0x00, 0x02, 0x00, 0x00, 0xc4, 0x03,
					0x16, 0x2a, 0x3b, 0xff, 0x3b, 0x2a, 0x16, 0xff, 0xa4, 0x3b, 0xb3, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x3b, 0x3b, 0x3b, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"AdjustSigmoid 3x3 0.5 0.0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			0.5,
			0.0,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0x00, 0x00, 0x01, 0x00, 0xcc, 0x00, 0x02, 0x00, 0x00, 0xcc, 0x03,
					0x11, 0x22, 0x33, 0xff, 0x33, 0x22, 0x11, 0xff, 0xaa, 0x33, 0xbb, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x33, 0x33, 0x33, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := AdjustSigmoid(d.src, d.m, d.p)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}
