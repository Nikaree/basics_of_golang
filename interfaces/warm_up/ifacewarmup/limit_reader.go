package ifacewarmup

import "io"

type limitedReader struct {
	r io.Reader
	n int64
}

// Read читает не более оставшегося лимита.
// Когда лимит исчерпан — возвращает (0, io.EOF).
func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}

	// Если запрашиваем больше, чем осталось — уменьшаем срез
	if int64(len(p)) > l.n {
		p = p[:l.n]
	}

	n, err = l.r.Read(p)
	l.n -= int64(n)

	// Если лимит закончился — сигнализируем EOF
	if l.n <= 0 && err == nil {
		return n, io.EOF
	}
	return n, err
}

// LimitReader возвращает io.Reader, который считывает из r не более n байт.
// Когда лимит исчерпан, последующие Read возвращают (0, io.EOF).
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{
		r: r,
		n: n,
	}
}
