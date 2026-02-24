package safeio

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadFile читает файл целиком и ВСЕГДА его закрывает.
// При ошибке возвращает err вида "read <path>: <исходная ошибка>".
func ReadFile(path string) ([]byte, error) {
	// читаем содержимое целиком
	// ReadFile автоматически сам закрывает файл
	// под капотом:
	// file, _ := os.Open("file.txt")
	// defer file.Close()
	// data, _ := io.ReadAll(file)
	data, err := os.ReadFile(path)

	if err != nil {
		// Оборачиваем ошибку в нужный формат
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return data, nil

	/*
		// Открываем файл
			file, err := os.Open(path)
			if err != nil {
				// Оборачиваем ошибку в нужный формат
				return nil, fmt.Errorf("read %s: %w", path, err)
			}

			// Гарантируем закрытие файла
			defer file.Close()

			// Читаем всё содержимое файла
			data, err := io.ReadAll(file)
	*/
}

// WriteFileAtomic пишет data во временный файл рядом с target,
// затем атомарно переименовывает его в <path>.
// • при любой ошибке временный файл удаляется
// • права файла — как в аргументе perm
// • возвращаемый err имеет вид "write <path>: <исходная ошибка>"
func WriteFileAtomic(path string, data []byte, perm os.FileMode) error {
	// создаем временный файл
	tmpFile, err := os.CreateTemp(filepath.Dir(path), ".tmp-*")

	if err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	// сохраняем имя временного файла, чтобы в дальнейшем использовать
	tmpName := tmpFile.Name()

	// флаг для контроля успешности
	success := false

	// удаление файла при любой ошибке
	defer func() {
		// Закрываем файл (игнорируем ошибку здесь)
		err := tmpFile.Close()
		if err != nil {
			return
		}

		// операция не успешна -> удаляем файл
		if !success {
			err := os.Remove(tmpName)
			if err != nil {
				return
			}
		}
	}()

	// права доступа
	if err := tmpFile.Chmod(perm); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	// запись
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	// Важно: синхронизируем данные с диском
	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	// закрытие файла перед переименованием
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	// атомарное переименование
	if err := os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	// меняем флаг на успешно
	success = true
	return nil
}

// WithFile открывает файл, передаёт *os.File в fn и ГАРАНТИРОВАННО его закрывает.
//   - если fn вернул ошибку — она выходит наружу без изменений
//   - если внутри fn возникла паника — она перехватывается, файл закрывается,
//     наружу возвращается error: "panic while <path>: <panic-value>".
func WithFile(
	path string,
	flag int,
	perm os.FileMode,
	fn func(*os.File) error,
) (err error) {
	// Открываем файл с указанными флагами и правами
	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return err
	}

	// гарантированное закрытие файла
	defer func() {
		// если паника
		if r := recover(); r != nil {
			err = fmt.Errorf("panic while %s: %v", path, r)
		}

		// Гарантированно закрываем
		closeErr := file.Close()

		// Если fn не вернул ошибку и panic не было,
		// но Close() вернул ошибку — возвращаем её
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	return fn(file)
}
