package matematika

import "testing"

func TestTambah(t *testing.T) {
	testCases := []struct {
		bil1, bil2, ekspektasi int
	}{
		{1, 2, 3},
		{10, 20, 30},
		{23, 3, 26},
	}
	for _, c := range testCases {
		hasil := Tambah(c.bil1, c.bil2)

		if hasil != c.ekspektasi {
			t.Errorf("Tambah(%q) == %q, want %q", c.bil1, c.bil2, hasil, c.ekspektasi)
		}
	}
}

func TestKurang(t *testing.T) {
	testCases := []struct {
		bil1, bil2, ekspektasi int
	}{
		{10, 2, 8},
		{100, 20, 80},
		{23, 3, 20},
	}
	for _, c := range testCases {
		hasil := Kurang(c.bil1, c.bil2)

		if hasil != c.ekspektasi {
			t.Errorf("Kurang(%q) == %q, want %q", c.bil1, c.bil2, hasil, c.ekspektasi)
		}
	}
}

func TestKali(t *testing.T) {
	testCases := []struct {
		bil1, bil2, ekspektasi int
	}{
		{1, 2, 2},
		{10, 20, 200},
		{23, 3, 69},
	}
	for _, c := range testCases {
		hasil := Kali(c.bil1, c.bil2)

		if hasil != c.ekspektasi {
			t.Errorf("Kali(%q) == %q, want %q", c.bil1, c.bil2, hasil, c.ekspektasi)
		}
	}
}
