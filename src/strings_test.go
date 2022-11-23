package ralphred

import (
	"testing"
)

func assertStringCommandResult(t *testing.T, input []string, expected string) {
	t.Helper()
	items, err := stringCommand(input)
	if err != nil {
		t.Fatalf("Got an err: %s", err)
	}

	if len(items) != 1 {
		t.Fatalf("Expected one result got %d", len(items))
	}

	result := items[0].Title
	if result != expected {
		t.Fatalf("Got %s expected %s", result, expected)
	}
}

func TestLength(t *testing.T) {
	t.Run("SingleWord", func(t *testing.T) {
		assertStringCommandResult(t, []string{"length", "word"}, "4")
	})
	t.Run("MultipleWords", func(t *testing.T) {
		assertStringCommandResult(t, []string{"length", "word", "word"}, "9")
	})
}

func TestWords(t *testing.T) {
	t.Run("SingleWord", func(t *testing.T) {
		assertStringCommandResult(t, []string{"words", "word"}, "1")
	})
	t.Run("MultipleWords", func(t *testing.T) {
		assertStringCommandResult(t, []string{"words", "word", "word"}, "2")
	})
}

func TestLower(t *testing.T) {
	t.Run("SingleWord", func(t *testing.T) {
		assertStringCommandResult(t, []string{"lower", "WoRd"}, "word")
	})
	t.Run("MultipleWords", func(t *testing.T) {
		assertStringCommandResult(t, []string{"lower", "WoRd", "wOrD"}, "word word")
	})
}

func TestTitle(t *testing.T) {
	t.Run("SingleWord", func(t *testing.T) {
		assertStringCommandResult(t, []string{"title", "wORD"}, "Word")
	})
	t.Run("MultipleWords", func(t *testing.T) {
		assertStringCommandResult(t, []string{"title", "WoRd", "wOrD"}, "Word Word")
	})
}

func TestUpper(t *testing.T) {
	t.Run("SingleWord", func(t *testing.T) {
		assertStringCommandResult(t, []string{"upper", "word"}, "WORD")
	})
	t.Run("MultipleWords", func(t *testing.T) {
		assertStringCommandResult(t, []string{"upper", "word", "WoRd"}, "WORD WORD")
	})
}

func TestPyMod(t *testing.T) {
	t.Run("SingleFile", func(t *testing.T) {
		assertStringCommandResult(t, []string{"pymod", "test.py"}, "test")
	})
	t.Run("FileInDict", func(t *testing.T) {
		assertStringCommandResult(t, []string{"pymod", "directory/test.py"}, "directory.test")
	})
}

func TestUnPyMod(t *testing.T) {
	t.Run("SingleFile", func(t *testing.T) {
		assertStringCommandResult(t, []string{"unpymod", "test"}, "test.py")
	})
	t.Run("FileInDict", func(t *testing.T) {
		assertStringCommandResult(t, []string{"unpymod", "directory.test"}, "directory/test.py")
	})
}

func TestHashes(t *testing.T) {
	t.Run("md5", func(t *testing.T) {
		assertStringCommandResult(t, []string{"md5", "word"}, "c47d187067c6cf953245f128b5fde62a")
	})
	t.Run("sha1", func(t *testing.T) {
		assertStringCommandResult(t, []string{"sha1", "word"}, "3cbcd90adc4b192a87a625850b7f231caddf0eb3")
	})
	t.Run("sha256", func(t *testing.T) {
		assertStringCommandResult(t, []string{"sha256", "word"}, "98c1eb4ee93476743763878fcb96a25fbc9a175074d64004779ecb5242f645e6")
	})
	t.Run("sha512", func(t *testing.T) {
		assertStringCommandResult(t, []string{"sha512", "word"}, "e1cc867e070565b17656702f48d54c483b3fb64fe4d2f0bb30b6c4ec84e4b8d51fe3cdebe2324e7dec3c82f6971d89b52a6c3beb8d5dda2b9b1a80ddc129d073")
	})
}
